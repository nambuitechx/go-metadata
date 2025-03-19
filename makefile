POSTGRES_IMAGE = go-metadata/postgres
BACKEND_IMAGE = go-metadata/backend
METADATA_INGESTION_IMAGE = go-metadata/metadata-ingestion

build-postgres:
	docker build -t $(POSTGRES_IMAGE) ./postgres

build-backend:
	docker build -t ${BACKEND_IMAGE} ./backend

build-metadata-ingestion:
	docker build -t ${METADATA_INGESTION_IMAGE} ./metadata-ingestion


build: build-postgres build-backend build-metadata-ingestion

run:
	docker compose up -V

up_postgres:
	docker compose up postgres -d -V

up_metadata_ingestion:
	docker compose up metadata-ingestion -d -V

up: up_postgres up_metadata_ingestion

down:
	docker compose down

# Đăng nhập vào ECR
login:
	aws ecr get-login-password --region $(REGION) | docker login --username AWS --password-stdin $(ECR_REGISTRY)

# Đăng xuất khỏi ECR
logout:
	docker logout $(ECR_REGISTRY)
