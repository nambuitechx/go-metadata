from datetime import datetime
from fastapi import APIRouter, BackgroundTasks

from configs.logger import get_logger
from entities.base import DefaultResponsePayload
from entities.ingestion import TestConnectionPayload
from services.ingestion import init_test_connection_result_db, save_test_connection_result, get_test_connection_result, test_connection_via_om

logger = get_logger(name=__name__)
init_test_connection_result_db()

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
    return {"message": "Start test connection successfully"}

@router.get("/test-connection/{id}")
def get_test_connection_handler(id: str):
    return get_test_connection_result(id)
