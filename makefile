IMAGE_POSTGRES = go-metadata/postgres
IMAGE_BACKEND = go-metadata/backend

build-postgres:
	docker build -t $(IMAGE_POSTGRES) ./postgres

build-backend:
	docker build -t ${IMAGE_BACKEND} ./backend


build: build-postgres build-backend

run:
	docker compose up -d -V

up_postgres:
	docker compose up postgres -d -V

up: up_postgres

down:
	docker compose down

# Đăng nhập vào ECR
login:
	aws ecr get-login-password --region $(REGION) | docker login --username AWS --password-stdin $(ECR_REGISTRY)

# Đăng xuất khỏi ECR
logout:
	docker logout $(ECR_REGISTRY)
