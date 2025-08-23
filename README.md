# Hospital-Microservice
A hospital management systems utilizing microservices

### Run with Docker
docker-compose up --build

### Run with MakeFile (not recomended)
make all

### Reset Swagger
swag init --generalInfo cmd/app/main.go --output ./docs --parseDependency

