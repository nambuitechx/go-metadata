from fastapi import FastAPI
from fastapi.exceptions import HTTPException, RequestValidationError
from fastapi.middleware.cors import CORSMiddleware

from services.exceptions import (
    SystemException,
    default_exception_handler,
    system_exception_handler,
    http_exception_handler,
    validation_exception_handler,   
)
from apis.ingestion import router as ingestion_router

app = FastAPI()

@app.get("/health")
def health_check():
    return { "message": "healthy" }

# CORS Middleware configuration
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Allows all origins, or you can specify specific origins
    allow_credentials=True,
    allow_methods=["*"],  # Allows all HTTP methods (GET, POST, etc.)
    allow_headers=["*"],  # Allows all headers
)

# Exception Handler
app.add_exception_handler(Exception, default_exception_handler)
app.add_exception_handler(SystemException, system_exception_handler)
app.add_exception_handler(HTTPException, http_exception_handler)
app.add_exception_handler(RequestValidationError, validation_exception_handler)

app.include_router(prefix="/api/v1", router=ingestion_router)

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8080)
