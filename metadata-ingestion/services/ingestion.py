import json
import traceback
import psycopg2

from fastapi import HTTPException
from typing import List, Optional
from datetime import datetime

from metadata.generated.schema.entity.services.connections.metadata.openMetadataConnection import OpenMetadataConnection
from metadata.generated.schema.entity.services.databaseService import (
    DatabaseServiceType,
    DatabaseService,
)
from metadata.generated.schema.metadataIngestion.workflow import (
    OpenMetadataWorkflowConfig,
    SourceConfig,
    Source,
)
from metadata.generated.schema.metadataIngestion.databaseServiceMetadataPipeline import DatabaseServiceMetadataPipeline
from metadata.generated.schema.type.filterPattern import FilterPattern
from metadata.generated.schema.entity.automations.workflow import (
    Workflow as AutomationWorkflow,
)

import metadata.ingestion.connections.test_connections
from metadata.ingestion.source.connections import get_connection, get_test_connection_fn
from metadata.ingestion.ometa.ometa_api import OpenMetadata
from metadata.ingestion.connections.test_connections import (
    TestConnectionStep,
    # TestConnectionResult,
    TestConnectionStepResult,
    StatusType,
    _test_connection_steps_during_ingestion,
)
from metadata.workflow.metadata import MetadataWorkflow

from configs.settings import settings
from configs.logger import get_logger
from configs.constants import INTERNAL_SCHEMA
from entities.ingestion import TestConnectionPayload, TestConnectionResult, IngestMetadataPayload, IngestionRun, IngestionRunStatus

logger = get_logger(__name__)


def get_db_connection():
    return psycopg2.connect(
        host=settings.DATABASE_HOST,
        port=settings.DATABASE_PORT,
        dbname=settings.DATABASE_NAME,
        user=settings.DATABASE_USER,
        password=settings.DATABASE_PASSWORD,
    )
    
    
def init_db():
    init_test_connection_result_db()
    init_ingestion_run_db()


def init_test_connection_result_db():
    conn = get_db_connection()
    
    with conn.cursor() as cursor:
        cursor.execute(
            """
            CREATE TABLE IF NOT EXISTS z_test_connection_result (
                id TEXT PRIMARY KEY,
                result TEXT NOT NULL
            );"""
        )
        conn.commit()
    
    conn.close()


def init_ingestion_run_db():
    conn = get_db_connection()
    
    with conn.cursor() as cursor:
        cursor.execute(
            """
            CREATE TABLE IF NOT EXISTS z_ingestion_run (
                id TEXT PRIMARY KEY,
                connection_name TEXT NOT NULL,
                connection_type TEXT NOT NULL,
                status TEXT NOT NULL,
                start_time BIGINT,
                end_time BIGINT,
                error_message TEXT
            );"""
        )
        conn.commit()
    
    conn.close()


def save_test_connection_result(item: dict):
    conn = get_db_connection()
    
    with conn.cursor() as cursor:
        cursor.execute("DELETE FROM z_test_connection_result WHERE id = %s", (item["id"],))
        cursor.execute(
            "INSERT INTO z_test_connection_result (id, result) VALUES (%s, %s) ON CONFLICT (id) DO UPDATE SET result = %s",
            (item["id"], json.dumps(item), json.dumps(item),)
        )
        conn.commit()
    
    conn.close()


def get_test_connection_result(id: str) -> dict:
    conn = get_db_connection()
    
    with conn.cursor() as cursor:
        cursor.execute(
            "SELECT result FROM z_test_connection_result WHERE id = %s", (id,)
        )
        data = cursor.fetchone()
        conn.commit()
        conn.close()
        
        if data == None:
            raise HTTPException(status_code=404, detail="Test connection result not found")
        
        return json.loads(data[0])


def _test_connection_steps(
    metadata: OpenMetadata,
    steps: List[TestConnectionStep],
    automation_workflow: Optional[AutomationWorkflow] = None,
) -> TestConnectionResult:
    """
    Run all the function steps and raise any errors
    """
    if automation_workflow:
        return _test_connection_steps_automation_workflow(
            metadata=metadata, steps=steps, automation_workflow=automation_workflow
        )

    test_connection_result = _test_connection_steps_during_ingestion(steps=steps)
    test_connection_result.lastUpdatedAt = int(datetime.now().timestamp() * 1000)
    test_connection_result.status = (
        StatusType.Failed.value
        if any(
            step
            for step in test_connection_result.steps
            if (not step.passed) and step.mandatory
        )
        else StatusType.Successful.value
    )
    return test_connection_result


def _test_connection_steps_automation_workflow(
    metadata: OpenMetadata,
    steps: List[TestConnectionStep],
    automation_workflow: dict,
) -> TestConnectionResult:
    """
    Run the test connection as part of the automation workflow
    We need to update the automation workflow in each step
    """
    # /api/v1/services/databaseServices/{id}/testConnectionResult
    connectionId = automation_workflow["connectionId"]

    def updateConnectionTestStatus(data: TestConnectionResult):
        # logger.info(f"Data type: {type(data)}")
        # logger.info(f"Data content: {data}")
        automation_workflow.update(data.model_dump())
        save_test_connection_result(automation_workflow)
        if connectionId != None:
            metadata.client.put(
                "/services/databaseServices/{id}/testConnectionResult".format(
                    id=connectionId
                ),
                data.model_dump_json(),
            )
    
    test_connection_result = TestConnectionResult(
        status=StatusType.Running.value,
        steps=[],
    )
    
    try:
        for step in steps:
            test_connection_result.lastUpdatedAt = int(
                datetime.now().timestamp() * 1000
            )
            test_connection_result.status = StatusType.Running.value
            updateConnectionTestStatus(test_connection_result)
            try:
                logger.info(f"========== Running {step.name} step...")
                step.function()
                test_connection_result.steps.append(
                    TestConnectionStepResult(
                        name=step.name,
                        mandatory=step.mandatory,
                        passed=True,
                    )
                )
            except Exception as err:
                logger.error(
                    f"========== Wild error happened while appending test connection result at steps - {err}"
                )
                logger.debug(traceback.format_exc())
                logger.warning(f"{step.name}-{err}")
                test_connection_result.steps.append(
                    TestConnectionStepResult(
                        name=step.name,
                        mandatory=step.mandatory,
                        passed=False,
                        message=step.error_message,
                        errorLog=str(err),
                    )
                )
                if step.short_circuit:
                    # break the workflow if the step is a short circuit step
                    break
        test_connection_result.lastUpdatedAt = int(datetime.now().timestamp() * 1000)
        test_connection_result.status = (
            StatusType.Failed.value
            if any(
                step
                for step in test_connection_result.steps
                if (not step.passed) and step.mandatory
            )
            else StatusType.Successful.value
        )
        updateConnectionTestStatus(test_connection_result)
        return test_connection_result

    except Exception as err:
        logger.error(
            f"========== Wild error happened while testing the connection in the workflow - {err}"
        )
        test_connection_result.lastUpdatedAt = int(datetime.now().timestamp() * 1000)
        test_connection_result.status = StatusType.Failed.value
        updateConnectionTestStatus(test_connection_result)
        return test_connection_result


metadata.ingestion.connections.test_connections._test_connection_steps = (
    _test_connection_steps
)


def get_metadata_client():
    metadata = OpenMetadata(
        config=OpenMetadataConnection.model_validate(
            {
                "hostPort": settings.BACKEND_URL,
                "authProvider": "openmetadata",
                "securityConfig": {"jwtToken": "token"},
            }
        )
    )
    return metadata


def test_connection_via_om(payload: TestConnectionPayload):
    connection_config = payload.connectionInfo.connection.config
    test_connection_fn = get_test_connection_fn(connection_config)
    
    metadata_client = get_metadata_client()
    
    if payload.connectionId != None:
        db_service: DatabaseService = metadata_client.get_by_id(DatabaseService, payload.connectionId)
        connection_config = db_service.connection.config
    
    result = {"id": payload.requestId, "connectionId": payload.connectionId}
    connection = get_connection(connection_config)
    test_connection_fn(metadata_client, connection, connection_config, result)


def run_metadata_ingestion_for_om(payload: IngestMetadataPayload):
    connection_name = payload.connectionName
    client = get_metadata_client()
    connection: DatabaseService = client.get_by_name(
        DatabaseService,
        fqn=connection_name,
        fields=["tags"],
    )
    
    if connection == None:
        raise HTTPException(status_code=404, detail="Connection not found")
    
    source = Source(
        type=connection.serviceType.value,
        serviceName=connection_name,
        # serviceConnection={
        #     "config": {
        #         "hostPort": "techx-kalliope-dev-openmetadatadb.coadbyfuowjn.ap-southeast-1.rds.amazonaws.com:5432",
        #         "database": "openmetadatadb",
        #         "username": "openmetadata_db_user",
        #         "authType": {
        #             "password": "ML1EsY81dBUevxuo"
        #         },
        #         "sslMode": "require"
        #     }
        # },
        sourceConfig=SourceConfig(
            config=DatabaseServiceMetadataPipeline(
                useFqnForFiltering=True,
                markDeletedTables=True,
                schemaFilterPattern=FilterPattern(
                    excludes=get_schema_filter_pattern(connection=connection)
                ),
            )
        ),
    )
    
    logger.info(f"========== Running metadata ingestion for {connection_name}")
    runWorkFlow(source, connection.id.root.__str__())


def get_schema_filter_pattern(connection: DatabaseService):
    schema_fqn_prefix = connection.name.root.__str__() + "." + connection.name.root.__str__() + "."
    if connection.serviceType == DatabaseServiceType.Postgres:
        schema_fqn_prefix = connection.name.root.__str__() + "." + connection.connection.config.database + "."
    if connection.serviceType in INTERNAL_SCHEMA:
        return [schema_fqn_prefix + e for e in INTERNAL_SCHEMA[connection.serviceType]]
    else:
        return []


def runWorkFlow(
    metadata_config: Source,
    entity_id: str,
):
    ingestion_run = IngestionRun(
        id=entity_id,
        connection_name=metadata_config.serviceName,
        connection_type=metadata_config.type,
        status=IngestionRunStatus.RUNNING.value,
        start_time=get_now(),
    )
    insert_or_update_ingestion_run_status(ingestion_run)
    
    try:
        workflow = MetadataWorkflow(
            OpenMetadataWorkflowConfig.model_validate(
                {
                    "source": metadata_config.model_dump(),
                    "sink": {"type": "metadata-rest", "config": {}},
                    "workflowConfig": {
                        "loggerLevel": "INFO",
                        "openMetadataServerConfig": {
                            "hostPort": settings.BACKEND_URL,
                            "authProvider": "openmetadata",
                            "securityConfig": {"jwtToken": "token"},
                        },
                    },
                }
            ),
        )
        workflow.execute()
        workflow.print_status()
        workflow.raise_from_status()
        workflow.stop()
        ingestion_run.status = "SUCCEEDED"
        ingestion_run.end_time = get_now()
        insert_or_update_ingestion_run_status(ingestion_run)
    except Exception as e:
        ingestion_run.status = "FAILED"
        ingestion_run.end_time = get_now()
        ingestion_run.error_message = str(e)
        insert_or_update_ingestion_run_status(ingestion_run)
        raise e


def get_now():
    return int(datetime.now().timestamp())


def insert_or_update_ingestion_run_status(ingestion_run: IngestionRun):
    conn = get_db_connection()
    
    with conn.cursor() as cursor:
        cursor.execute(
            # "UPDATE z_ingestion_run SET status = %s, error_message = %s WHERE id = %s",
            # (ingestion_run.status, ingestion_run.error_message, ingestion_run.id,)
            """
            INSERT INTO z_ingestion_run(id, connection_name, connection_type, status, start_time, end_time, error_message)
            VALUES (%s, %s, %s, %s, %s, %s, %s)
            ON CONFLICT (id) DO UPDATE SET status = %s, start_time = %s, end_time = %s, error_message = %s
            """,
            (
                ingestion_run.id,
                ingestion_run.connection_name,
                ingestion_run.connection_type,
                ingestion_run.status,
                ingestion_run.start_time,
                ingestion_run.end_time,
                ingestion_run.error_message,
                ingestion_run.status,
                ingestion_run.start_time,
                ingestion_run.end_time,
                ingestion_run.error_message,
            ),
        )
        conn.commit()
    
    conn.close()
