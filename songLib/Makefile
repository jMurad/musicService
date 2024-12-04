# run:
# 	@docker-compose up -d
# 	@sleep 4
# 	@go run ./cmd/main.go
# migrate-up:
# 	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost user=postgres dbname=test password=12345678 port=5432 sslmode=disable" goose -dir ./migrations/ up 
# migrate-down:
# 	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost user=postgres dbname=test password=12345678 port=5432 sslmode=disable" goose -dir ./migrations/ down 


migrate-up:
	migrate -database "postgres://user:123456@localhost:5432/songslib?sslmode=disable" -path migrations up
migrate-down:
	migrate -database "postgres://user:123456@localhost:5432/songslib?sslmode=disable" -path migrations down