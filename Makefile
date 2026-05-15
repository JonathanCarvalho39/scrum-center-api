.PHONY: help run build build-static test clean migrate-up migrate-down swagger docker-up docker-down docker-build docker-push

help: ## Mostra este menu de ajuda
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

run: ## Executa a aplicação em modo desenvolvimento
	@echo "🚀 Starting server..."
	@go run cmd/api/main.go

build: ## Compila a aplicação
	@echo "🔨 Building..."
	@go build -o bin/api cmd/api/main.go
	@echo "✅ Build complete: bin/api"

build-static: ## Compila binário estático (para Docker)
	@echo "🔨 Building static binary..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o bin/api ./cmd/api
	@echo "✅ Static build complete: bin/api"

test: ## Executa todos os testes
	@echo "🧪 Running tests..."
	@go test -v ./...

test-coverage: ## Executa testes com coverage
	@echo "📊 Running tests with coverage..."
	@go test -cover ./...
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report: coverage.html"

test-race: ## Executa testes com detecção de race condition
	@echo "🏁 Running tests with race detector..."
	@go test -race -v ./...

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

docker-build: ## Build imagem Docker local
	@echo "🐳 Building Docker image..."
	@docker build -t scrum-center-api:latest .
	@echo "✅ Docker image built: scrum-center-api:latest"

docker-build-ghcr: ## Build e tag para GitHub Container Registry
	@echo "🐳 Building Docker image for GHCR..."
	@docker build -t ghcr.io/jonathancarvalho39/scrum-center-api:latest .
	@echo "✅ Docker image built and tagged for GHCR"

docker-push: docker-build-ghcr ## Push imagem para GitHub Container Registry
	@echo "📤 Pushing to GitHub Container Registry..."
	@docker push ghcr.io/jonathancarvalho39/scrum-center-api:latest
	@echo "✅ Image pushed successfully"

docker-up: ## Sobe containers Docker
	@echo "🐳 Starting Docker containers..."
	@docker-compose up -d

docker-down: ## Para containers Docker
	@echo "🛑 Stopping Docker containers..."
	@docker-compose down

docker-logs: ## Mostra logs dos containers
	@docker-compose logs -f

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

lint-fix: ## Executa linter com auto-fix
	@echo "🔧 Linting with auto-fix..."
	@golangci-lint run --fix

ci-local: lint test build ## Simula pipeline CI localmente
	@echo "✅ Local CI pipeline completed successfully"

.DEFAULT_GOAL := help
