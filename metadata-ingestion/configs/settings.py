from pydantic import ConfigDict
from pydantic_settings import BaseSettings

##
# Mapping environment variable to Settings config
##
class Settings(BaseSettings):
    model_config = ConfigDict(extra="ignore")
    
    # AWS region
    AWS_REGION: str = "ap-southeast-1"
    
    
    # Database connection information
    DB_NAME: str = ""
    DB_HOST: str = ""
    DB_PORT: str = ""
    DB_USER: str = ""
    DB_PASSWORD: str = ""
    
    BACKEND_URL: str = "http://backend:8585/api"
    

settings = Settings(_env_file=".env")
