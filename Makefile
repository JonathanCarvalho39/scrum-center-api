.PHONY: help run build test clean migrate-up migrate-down swagger docker-up docker-down

help: ## Mostra este menu de ajuda
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

run: ## Executa a aplicação em modo desenvolvimento
	@echo "🚀 Starting server..."
	@go run cmd/api/main.go

build: ## Compila a aplicação
	@echo "🔨 Building..."
	@go build -o bin/api cmd/api/main.go
	@echo "✅ Build complete: bin/api"

test: ## Executa todos os testes
	@echo "🧪 Running tests..."
	@go test -v ./...

test-coverage: ## Executa testes com coverage
	@echo "📊 Running tests with coverage..."
	@go test -cover ./...
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report: coverage.html"

clean: ## Remove arquivos de build e cache
	@echo "🧹 Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@go clean

swagger: ## Gera documentação OpenAPI/Swagger
	@echo "📚 Generating Swagger docs..."
	@swag init -g cmd/api/main.go -o docs/
	@echo "✅ Swagger docs generated"

migrate-up: ## Executa migrations (up)
	@echo "📈 Running migrations..."
	@echo "Migration tool not configured yet"

migrate-down: ## Reverte última migration
	@echo "📉 Reverting migration..."
	@echo "Migration tool not configured yet"

docker-up: ## Sobe containers Docker
	@echo "🐳 Starting Docker containers..."
	@docker-compose up -d

docker-down: ## Para containers Docker
	@echo "🛑 Stopping Docker containers..."
	@docker-compose down

install: ## Instala dependências
	@echo "📦 Installing dependencies..."
	@go mod download
	@go mod tidy

fmt: ## Formata código
	@echo "✨ Formatting code..."
	@go fmt ./...

lint: ## Executa linter
	@echo "🔍 Linting..."
	@golangci-lint run

.DEFAULT_GOAL := help
