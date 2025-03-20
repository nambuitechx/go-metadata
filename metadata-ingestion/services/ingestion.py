import json
import traceback
import psycopg2

from fastapi import HTTPException
from typing import List
from datetime import datetime

from metadata.ingestion.source.connections import get_connection, get_test_connection_fn
from metadata.ingestion.ometa.ometa_api import OpenMetadata
from metadata.generated.schema.entity.services.connections.metadata.openMetadataConnection import OpenMetadataConnection
from metadata.generated.schema.entity.services.databaseService import DatabaseService
import metadata.ingestion.connections.test_connections
from metadata.ingestion.connections.test_connections import (
    TestConnectionStep,
    # TestConnectionResult,
    TestConnectionStepResult,
    StatusType,
    _test_connection_steps_during_ingestion,
)

from configs.settings import settings
from configs.logger import get_logger
from entities.ingestion import TestConnectionPayload, TestConnectionResult

logger = get_logger(__name__)


def get_db_connection():
    return psycopg2.connect(
        host=settings.DATABASE_HOST,
        port=settings.DATABASE_PORT,
        dbname=settings.DATABASE_NAME,
        user=settings.DATABASE_USER,
        password=settings.DATABASE_PASSWORD,
    )


def init_test_connection_result_db():
    conn = get_db_connection()
    
    with conn.cursor() as cursor:
        cursor.execute(
            """
            CREATE TABLE IF NOT EXISTS test_connection_result (
                id TEXT PRIMARY KEY,
                result TEXT NOT NULL
            );"""
        )
        conn.commit()
    
    conn.close()


def save_test_connection_result(item: dict):
    conn = get_db_connection()
    
    with conn.cursor() as cursor:
        cursor.execute("DELETE FROM test_connection_result WHERE id = %s", (item["id"],))
        cursor.execute(
            "INSERT INTO test_connection_result (id, result) VALUES (%s, %s) ON CONFLICT (id) DO UPDATE SET result = %s",
            (item["id"], json.dumps(item), json.dumps(item),)
        )
        conn.commit()
    
    conn.close()


def get_test_connection_result(id: str) -> dict:
    conn = get_db_connection()
    
    with conn.cursor() as cursor:
        cursor.execute(
            "SELECT result FROM test_connection_result WHERE id = %s", (id,)
        )
        data = cursor.fetchone()
        conn.commit()
        conn.close()
        
        if data == None:
            raise HTTPException(status_code=404, detail="request not found")
        
        return json.loads(data[0])


def _test_connection_steps(
    metadata: OpenMetadata,
    steps: List[TestConnectionStep],
    automation_workflow: dict,
) -> None:
    """
    Run all the function steps and raise any errors
    """
    if automation_workflow:
        _test_connection_steps_automation_workflow(
            metadata=metadata, steps=steps, automation_workflow=automation_workflow
        )

    else:
        _test_connection_steps_during_ingestion(steps=steps)


def _test_connection_steps_automation_workflow(
    metadata: OpenMetadata,
    steps: List[TestConnectionStep],
    automation_workflow: dict,
) -> None:
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

    except Exception as err:
        logger.error(
            f"========== Wild error happened while testing the connection in the workflow - {err}"
        )
        test_connection_result.lastUpdatedAt = int(datetime.now().timestamp() * 1000)
        test_connection_result.status = StatusType.Failed.value
        updateConnectionTestStatus(test_connection_result)


metadata.ingestion.connections.test_connections._test_connection_steps = (
    _test_connection_steps
)


def test_connection_via_om(payload: TestConnectionPayload):
    connection_config = payload.connectionInfo.connection.config
    test_connection_fn = get_test_connection_fn(connection_config)
    
    metadata = OpenMetadata(
        config=OpenMetadataConnection.model_validate(
            {
                "hostPort": settings.BACKEND_URL,
                "authProvider": "openmetadata",
                "securityConfig": {"jwtToken": "token"},
            }
        )
    )
    
    if payload.connectionId != None:
        db_service: DatabaseService = metadata.get_by_id(DatabaseService, payload.connectionId)
        connection_config = db_service.connection.config
    
    result = {"id": payload.requestId, "connectionId": payload.connectionId}
    connection = get_connection(connection_config)
    test_connection_fn(metadata, connection, connection_config, result)
