from datetime import datetime
from fastapi import APIRouter, BackgroundTasks

from configs.logger import get_logger
from entities.base import DefaultResponsePayload
from entities.ingestion import TestConnectionPayload, IngestMetadataPayload
from services.ingestion import init_db, save_test_connection_result, get_test_connection_result, test_connection_via_om, run_metadata_ingestion_for_om

init_db()

logger = get_logger(name=__name__)

router = APIRouter(prefix="/ingestion")

@router.get("/health")
async def health_check():
    """
    Health check endpoint to monitor API availability.
    """
    return {"message": "healthy"}

@router.post("/test-connection", response_model=DefaultResponsePayload, tags=["ingestion"])
async def create_test_connection_handler(payload: TestConnectionPayload, background_tasks: BackgroundTasks):
    result = {"id": payload.requestId, "connectionId": payload.connectionId, "status": "Running", "lastUpdatedAt": int(datetime.now().timestamp() * 1000)}
    save_test_connection_result(result)
    background_tasks.add_task(
        test_connection_via_om,
        payload
    )
    return {"message": "Start testing connection successfully"}

@router.get("/test-connection/{id}")
def get_test_connection_handler(id: str):
    return get_test_connection_result(id)

@router.post("/ingest-metadata", response_model=DefaultResponsePayload, tags=["ingestion"])
async def ingest_metadata_handler(payload: IngestMetadataPayload, background_tasks: BackgroundTasks):
    background_tasks.add_task(
        run_metadata_ingestion_for_om,
        payload
    )
    return {"message": "Start ingesting metadata successfully"}
