# 🚀 Quick Reference - CI/CD Commands

## 📝 Comandos Make

```bash
# Desenvolvimento
make run              # Executa em modo dev
make test             # Roda todos os testes
make test-coverage    # Testes com coverage HTML
make test-race        # Testes com race detector
make lint             # Executa linter
make lint-fix         # Linter com auto-fix
make fmt              # Formata código
make build            # Compila binário

# Docker
make docker-build         # Build imagem local
make docker-build-ghcr    # Build e tag para GHCR
make docker-push          # Push para GHCR
make docker-up            # Sobe containers
make docker-down          # Para containers
make docker-logs          # Ver logs

# CI Local
make ci-local         # Simula pipeline (lint + test + build)

# Utilidades
make clean            # Limpa artifacts
make install          # Instala dependências
make help             # Lista comandos
```

---

## 🐳 Docker Commands

```bash
# Local Development
docker-compose up                    # Sobe API + PostgreSQL
docker-compose up -d                 # Modo background
docker-compose down                  # Para tudo
docker-compose logs -f api           # Logs em tempo real

# GHCR (GitHub Container Registry)
# Login
echo $GITHUB_TOKEN | docker login ghcr.io -u USERNAME --password-stdin

# Pull
docker pull ghcr.io/jonathancarvalho39/scrum-center-api:latest
docker pull ghcr.io/jonathancarvalho39/scrum-center-api:v1.0.0
docker pull ghcr.io/jonathancarvalho39/scrum-center-api:develop

# Run
docker run -d \
  --name scrum-api \
  -p 8080:8080 \
  -e DB_HOST=host.docker.internal \
  -e DB_PASSWORD=postgres \
  ghcr.io/jonathancarvalho39/scrum-center-api:latest

# Inspect
docker inspect ghcr.io/jonathancarvalho39/scrum-center-api:latest
docker history ghcr.io/jonathancarvalho39/scrum-center-api:latest
```

---

## 🔀 Git Workflows

```bash
# Feature Development
git checkout -b feature/minha-feature
# ... desenvolver com TDD
git add .
git commit -m "feat: adiciona nova funcionalidade"
git push origin feature/minha-feature
# → Abre PR → CI roda automaticamente

# Hotfix
git checkout -b hotfix/correcao-critica
# ... corrigir
git commit -m "fix: corrige bug crítico"
git push origin hotfix/correcao-critica
# → PR → CI

# Release
git checkout main
git pull origin main
git tag v1.0.0
git push origin v1.0.0
# → Release workflow: binários + Docker + changelog

# Ver tags
git tag -l
git tag -l "v*"

# Deletar tag local e remota
git tag -d v1.0.0
git push origin :refs/tags/v1.0.0
```

---

## 📦 Go Commands

```bash
# Testes
go test ./...                        # Todos os testes
go test -v ./...                     # Verbose
go test -race ./...                  # Race detector
go test -cover ./...                 # Coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out     # Coverage HTML

# Build
go build -o bin/api ./cmd/api                    # Build normal
CGO_ENABLED=0 go build -o bin/api ./cmd/api      # Build estático
go build -ldflags="-s -w" -o bin/api ./cmd/api   # Build otimizado

# Dependências
go mod download                      # Download deps
go mod tidy                          # Limpa deps
go mod verify                        # Verifica integridade
go get -u ./...                      # Atualiza deps

# Ferramentas
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install golang.org/x/vuln/cmd/govulncheck@latest

# Lint
golangci-lint run                    # Lint
golangci-lint run --fix              # Lint + fix
govulncheck ./...                    # CVE check
```

---

## 🔍 Verificar Status CI/CD

```bash
# Via GitHub CLI (gh)
gh run list                          # Lista workflows recentes
gh run view                          # Ver último workflow
gh run watch                         # Assistir workflow em tempo real
gh workflow list                     # Lista workflows disponíveis
gh workflow run ci.yml               # Trigger manual

# Via Browser
# https://github.com/JonathanCarvalho39/scrum-center-api/actions

# Ver Packages
# https://github.com/JonathanCarvalho39?tab=packages
```

---

## 📊 Verificar Coverage

```bash
# Local
make test-coverage
# Abre coverage.html no browser

# CI (após workflow rodar)
# Ver em: GitHub Actions → CI workflow → Artifacts
# Ou Codecov: https://codecov.io/gh/JonathanCarvalho39/scrum-center-api
```

---

## 🔐 Security Scans

```bash
# Local
golangci-lint run --enable=gosec     # Security linting
govulncheck ./...                    # CVE em dependências
docker scan scrum-center-api:latest  # Docker scan (Snyk)

# Trivy (container scan)
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock \
  aquasec/trivy image ghcr.io/jonathancarvalho39/scrum-center-api:latest

# CI/CD
# Todos rodam automaticamente em cada push
# Ver resultados em: GitHub → Security tab → Code scanning alerts
```

---

## 🌐 Health Checks

```bash
# Local
curl http://localhost:8080/health

# Docker container
docker exec scrum-api wget --quiet --tries=1 --spider http://localhost:8080/health

# Produção (exemplo)
curl https://api.seudominio.com/health

# Logs do container
docker logs scrum-api
docker logs -f scrum-api              # Follow mode
docker logs --tail 100 scrum-api      # Últimas 100 linhas
```

---

## 🧹 Cleanup

```bash
# Docker local
docker system prune -a                # Remove tudo não usado
docker image prune -a                 # Remove imagens não usadas
docker volume prune                   # Remove volumes não usados
docker builder prune                  # Limpa cache de build

# Go
go clean -cache                       # Limpa cache de build
go clean -modcache                    # Limpa cache de módulos

# Projeto
make clean                            # Remove bin/ coverage.*
rm -rf bin/ coverage.* *.test
```

---

## 🔧 Troubleshooting

```bash
# Verificar portas em uso
# Windows
netstat -ano | findstr :8080
# Linux/Mac
lsof -i :8080

# Matar processo
# Windows
taskkill /PID <PID> /F
# Linux/Mac
kill -9 <PID>

# Docker debug
docker ps -a                          # Todos containers
docker logs <container-id>            # Logs
docker exec -it <container-id> sh     # Entrar no container
docker inspect <container-id>         # Detalhes completos

# Rebuild do zero
docker-compose down -v                # Para e remove volumes
docker-compose build --no-cache       # Build sem cache
docker-compose up                     # Sobe tudo novamente

# Ver espaço em disco
docker system df                      # Uso de disco Docker
du -sh .                              # Tamanho do projeto
```

---

## 📝 Conventional Commits

```bash
# Estrutura
<type>(<scope>): <subject>

# Types
feat:      # Nova funcionalidade
fix:       # Correção de bug
docs:      # Documentação
style:     # Formatação
refactor:  # Refatoração
test:      # Testes
chore:     # Manutenção
ci:        # CI/CD
perf:      # Performance

# Exemplos
git commit -m "feat(team): adiciona endpoint de criação de time"
git commit -m "fix(invite): corrige validação de código"
git commit -m "test(member): adiciona testes de join team"
git commit -m "ci: atualiza workflow de release"
git commit -m "docs: atualiza README com instruções de deploy"
```

---

## 📚 Links Úteis

```bash
# Repositório
https://github.com/JonathanCarvalho39/scrum-center-api

# Actions
https://github.com/JonathanCarvalho39/scrum-center-api/actions

# Packages
https://github.com/JonathanCarvalho39?tab=packages

# Security
https://github.com/JonathanCarvalho39/scrum-center-api/security

# Releases
https://github.com/JonathanCarvalho39/scrum-center-api/releases

# GHCR Package
https://github.com/JonathanCarvalho39/scrum-center-api/pkgs/container/scrum-center-api
```

---

**💡 Dica:** Adicione este arquivo aos favoritos do browser ou IDE para referência rápida!

