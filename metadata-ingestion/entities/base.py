from pydantic import BaseModel
from typing import Any

class DefaultResponsePayload(BaseModel):
    is_success: bool = True
    status_code: int = 200
    message: str = "success"
    data: Any = {}


class PagingResponsePayload(BaseModel):
    limit: int
    offset: int
    total: int
