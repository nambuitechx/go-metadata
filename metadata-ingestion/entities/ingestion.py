from pydantic import BaseModel
from typing import Optional

from metadata.generated.schema.entity.automations.testServiceConnection import TestServiceConnectionRequest

class TestConnectionPayload(BaseModel):
    requestId: str
    connectionId: Optional[str] = None
    connectionInfo: TestServiceConnectionRequest

class TestConnectionResult(BaseModel):
    lastUpdatedAt: Optional[int] = None
    status: Optional[str] = None
    steps: list = []
