from entities.base import DefaultResponsePayload
from fastapi.exceptions import RequestValidationError, HTTPException
from fastapi.responses import JSONResponse


class SystemException(Exception):
    def __init__(self, msg: str = "Don't have specific message yet.", code: str = 50001, http_code: int = 500) -> None:
        self.msg = msg
        self.code = code
        self.http_code = http_code
        super().__init__(self.msg)


async def default_exception_handler(request, exc: Exception):
    return JSONResponse(
        status_code=500,
        content=DefaultResponsePayload(
            is_success=False, status_code=500, message=f"Something went wrong in system."
        ).model_dump(),
    )


async def system_exception_handler(request, exc: SystemException):
    return JSONResponse(
        status_code=exc.http_code,
        content=DefaultResponsePayload(
            is_success=False, status_code=exc.http_code, message=exc.msg
        ).model_dump(),
    )


async def http_exception_handler(request, exc: HTTPException):
    return JSONResponse(
        status_code=exc.status_code,
        content=DefaultResponsePayload(
            is_success=False,
            status_code=exc.status_code,
            message=exc.detail,
        ).model_dump(),
    )


async def validation_exception_handler(request, exc: RequestValidationError):
    return JSONResponse(
        status_code=400,
        content=DefaultResponsePayload(
            is_success=False,
            status_code=400,
            message=str(exc),
        ).model_dump(),
    )
