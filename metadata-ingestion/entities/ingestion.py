from pydantic import BaseModel
from typing import Optional
from enum import Enum

from metadata.generated.schema.entity.automations.testServiceConnection import TestServiceConnectionRequest

class TestConnectionPayload(BaseModel):
    requestId: str
    connectionId: Optional[str] = None
    connectionInfo: TestServiceConnectionRequest

class TestConnectionResult(BaseModel):
    lastUpdatedAt: Optional[int] = None
    status: Optional[str] = None
    steps: list = []

class IngestMetadataPayload(BaseModel):
    connectionName: str

class IngestionRun(BaseModel):
    id: str
    connection_name: str
    connection_type: str
    status: str
    start_time: Optional[int] = None
    end_time: Optional[int] = None
    error_message: Optional[str] = None


class IngestionRunStatus(Enum):
    PENDING = "PENDING"
    RUNNING = "RUNNING"
    SUCCESS = "SUCCESS"
    FAILED = "FAILED"
