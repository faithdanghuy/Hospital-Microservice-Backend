all:
	docker-compose up --build

user:
	docker compose build user-service
	docker compose up -d user-service
appointment:
	docker compose build appointment-service
	docker compose up -d appointment-service 
prescription:
	docker compose build prescription-service
	docker compose up -d prescription-service
report:
	docker compose build report-service
	docker compose up -d report-service
notify:
	docker compose build notify-service
	docker compose up -d notify-service