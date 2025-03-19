from fastapi import APIRouter, BackgroundTasks

from configs.logger import get_logger
from entities.base import DefaultResponsePayload
from entities.ingestion import TestConnectionPayload
from services.ingestion import test_connection_via_om

logger = get_logger(name=__name__)

router = APIRouter(prefix="/ingestion")

@router.get("/health")
async def health_check():
    """
    Health check endpoint to monitor API availability.
    """
    return {"message": "healthy"}

@router.post("/test-connection", response_model=DefaultResponsePayload, tags=["ingestion"])
async def test_connection_handler(payload: TestConnectionPayload, background_tasks: BackgroundTasks):
    
    logger.info(f"========== payload: {payload}")
    
    background_tasks.add_task(
        test_connection_via_om,
        payload
    )
    
    return {"message": "Start test connection successfully"}
