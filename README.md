# Hospital-Microservice
A hospital management systems utilizing microservices

### Run with Docker
docker-compose up --build
### Individual Docker
docker compose build user-service
docker compose up -d user-service

### Run with MakeFile (not recomended)
make all

### Reset Swagger
swag init --generalInfo cmd/app/main.go --output ./docs --parseDependency

### Start ngrok
ngrok http 3080

