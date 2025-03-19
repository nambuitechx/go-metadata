from pydantic import BaseModel
from typing import Optional

from metadata.generated.schema.entity.automations.testServiceConnection import TestServiceConnectionRequest

class TestConnectionPayload(BaseModel):
    requestId: str
    connectionId: Optional[str] = None
    connectionInfo: TestServiceConnectionRequest
