.PHONY: seed
seed:
	@go run cmd/migrate/seed/main.go
