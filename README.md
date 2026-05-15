# Scrum Center API

API REST para gerenciamento de projetos Scrum construída com Go e Clean Architecture.

## 🏗️ Arquitetura

Este projeto segue os princípios da **Clean Architecture** com separação clara entre camadas:

```
/cmd/api/          # Entrypoint da aplicação
/internal/
  domain/          # Entidades e interfaces de repositório
  usecase/         # Casos de uso (regras de negócio)
  infrastructure/  # Implementações concretas (postgres, http)
  handler/         # HTTP handlers (adapters de entrada)
/pkg/              # Utilitários reutilizáveis
/migrations/       # SQL migrations versionadas
/docs/             # Documentação OpenAPI gerada
```

### Camadas

- **Handler** → **UseCase** → **Repository**
- Interfaces definidas em `domain/`, implementadas em `infrastructure/`
- Inversão de dependência: domínio não conhece infraestrutura

## 🚀 Stack Tecnológica

- **Go 1.22+**
- **Gin** - HTTP web framework
- **pgx/v5** - PostgreSQL driver
- **JWT** - Autenticação
- **Swag** - OpenAPI/Swagger documentation
- **Validator** - Input validation
- **godotenv** - Environment variables

## 📋 Pré-requisitos

- Go 1.22 ou superior
- PostgreSQL 15+
- Make (opcional)

## ⚙️ Configuração

1. Clone o repositório:
```bash
git clone https://github.com/JonathanCarvalho39/scrum-center-api.git
cd scrum-center-api
```

2. Copie o arquivo de ambiente:
```bash
cp .env.example .env
```

3. Edite o `.env` com suas configurações

4. Instale as dependências:
```bash
go mod download
```

5. Execute as migrations (quando disponíveis):
```bash
make migrate-up
```

## 🏃 Executando

### Modo desenvolvimento
```bash
go run cmd/api/main.go
```

### Build e execução
```bash
go build -o bin/api cmd/api/main.go
./bin/api
```

## 📝 Endpoints

### Health Check
```
GET /health
```

### API v1
```
GET /api/v1/ping
```

## 🧪 Testes

### Executar todos os testes
```bash
go test ./...
```

### Testes com coverage
```bash
go test -cover ./...
```

### Estratégia de TDD

| Camada | Tipo de Teste | Ferramenta |
|--------|--------------|------------|
| Use cases | Unitário | Mocks das interfaces de repositório |
| Handlers | Unitário/Integração | `httptest` do Go stdlib |
| Repositórios | **Integração com PostgreSQL real** | `testcontainers-go` |

**Política obrigatória**: Repositórios testados com PostgreSQL real (sem mocks).

## 🐳 Docker

### Build local
```bash
docker build -t scrum-center-api .
```

### Executar com Docker Compose
```bash
docker-compose up
```

### Pull da imagem do GitHub Container Registry
```bash
docker pull ghcr.io/jonathancarvalho39/scrum-center-api:latest
```

## 🔄 CI/CD

GitHub Actions configurado com:
- **CI:** Lint, testes, build e security scan
- **Docker:** Build multi-arch e push para GitHub Container Registry

## 📚 Documentação API

A documentação OpenAPI/Swagger será gerada automaticamente com `swag`:

```bash
swag init -g cmd/api/main.go -o docs/
```

Acesse: `http://localhost:8080/swagger/index.html`

## 🔒 Segurança

- ✅ Validação de input em todos os handlers
- ✅ JWT para autenticação
- ✅ Middleware de autorização por rota
- ✅ Secrets via variáveis de ambiente
- ✅ Queries parametrizadas (proteção contra SQL Injection)
- ✅ CORS configurado

## 🤝 Contribuindo

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT.
