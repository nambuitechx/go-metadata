from metadata.ingestion.source.connections import get_connection, get_test_connection_fn
from metadata.ingestion.ometa.ometa_api import OpenMetadata
from metadata.generated.schema.entity.services.connections.metadata.openMetadataConnection import OpenMetadataConnection
from metadata.generated.schema.entity.services.databaseService import DatabaseService

from configs.settings import settings
from configs.logger import get_logger
from entities.ingestion import TestConnectionPayload

logger = get_logger(__name__)

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
    
    logger.info(f"========== result: {result}")