# ✅ Validation Checklist - CI/CD Setup

## 📦 Arquivos Criados - Checklist

### GitHub Actions Workflows
- [x] `.github/workflows/ci.yml` - Pipeline de CI completo
- [x] `.github/workflows/docker.yml` - Build e push Docker para GHCR
- [x] `.github/workflows/release.yml` - Release automático
- [x] `.github/workflows/cleanup.yml` - Limpeza de imagens antigas

### Configurações
- [x] `.golangci.yml` - Configuração do linter
- [x] `Dockerfile` - Otimizado com distroless
- [x] `docker-compose.yml` - Ambiente dev local
- [x] `docker-compose.prod.yml` - Backup da config de produção
- [x] `Makefile` - Comandos estendidos
- [x] `.gitignore` - Atualizado
- [x] `.env.example` - Já existia, mantido

### Documentação
- [x] `README.md` - Atualizado com CI/CD
- [x] `SETUP_SUMMARY.md` - Resumo completo do setup
- [x] `CI_CD_ARCHITECTURE.md` - Arquitetura detalhada
- [x] `GHCR_GUIDE.md` - Guia do GitHub Container Registry
- [x] `QUICK_REFERENCE.md` - Referência rápida de comandos
- [x] `VALIDATION_CHECKLIST.md` - Este arquivo

---

## 🧪 Validação Local

### 1. Verificar sintaxe YAML
```bash
# Instalar yamllint (opcional)
pip install yamllint

# Validar workflows
yamllint .github/workflows/*.yml
```

### 2. Testar build local
```bash
# Build normal
make build
# Verificar: bin/api existe

# Build estático (para Docker)
make build-static
# Verificar: bin/api existe e é estático
file bin/api  # Deve mostrar "statically linked"
```

### 3. Testar Docker build
```bash
# Build local
make docker-build
# Verificar: imagem criada

# Ver tamanho da imagem
docker images scrum-center-api:latest
# Esperado: ~15-20MB

# Testar imagem
docker run --rm -p 8080:8080 scrum-center-api:latest &
sleep 5
curl http://localhost:8080/health
# Esperado: resposta de health
```

### 4. Testar Linter
```bash
# Rodar linter
make lint
# Deve passar sem erros críticos

# Auto-fix
make lint-fix
```

### 5. Testar Tests
```bash
# Testes normais
make test

# Testes com race detector
make test-race

# Coverage
make test-coverage
# Verificar: coverage.html criado
```

### 6. Simular CI local
```bash
# Pipeline completo
make ci-local
# Deve executar: lint + test + build com sucesso
```

---

## 🔍 Validação GitHub

### Antes do Push

- [ ] Todos os testes passando localmente
- [ ] Linter sem erros
- [ ] Docker build funciona
- [ ] Arquivos sensíveis não commitados (.env)
- [ ] go.mod e go.sum estão tidy

### Primeiro Push

```bash
# Commit inicial
git add .github/ *.md Dockerfile docker-compose* Makefile .golangci.yml .gitignore
git commit -m "ci: setup GitHub Actions with GHCR integration

- Add CI workflow with lint, test, security
- Add Docker workflow for GHCR multi-arch builds
- Add release workflow for automated releases
- Add cleanup workflow for old images
- Optimize Dockerfile with distroless
- Add comprehensive documentation
- Update Makefile with Docker commands"

git push origin main
```

### Verificar no GitHub

1. **Actions Tab**
   - [ ] Workflow CI iniciou automaticamente
   - [ ] Jobs rodando em paralelo (lint, test, security)
   - [ ] Todos os jobs passaram (checkmarks verdes)

2. **Packages Tab**
   - [ ] Imagem Docker foi publicada
   - [ ] Tag `latest` existe
   - [ ] Tag `main` existe
   - [ ] Tag com SHA existe

3. **Security Tab**
   - [ ] Code scanning alerts habilitado
   - [ ] SARIF uploads aparecendo (gosec, trivy)

---

## 🚀 Validação Release

### Criar Primeiro Release

```bash
# Tag
git tag v0.1.0
git push origin v0.1.0
```

### Verificar Release Workflow

- [ ] Release workflow iniciou
- [ ] GitHub Release criado
- [ ] Binários anexados:
  - [ ] Linux amd64
  - [ ] Linux arm64
  - [ ] macOS amd64
  - [ ] macOS arm64
  - [ ] Windows amd64
- [ ] Checksums SHA256 anexados
- [ ] Docker images com tags:
  - [ ] `v0.1.0`
  - [ ] `0.1`
  - [ ] `0`
  - [ ] `latest`

---

## 🐳 Validação Docker GHCR

### Pull da Imagem

```bash
# Pull
docker pull ghcr.io/jonathancarvalho39/scrum-center-api:latest

# Verificar
docker inspect ghcr.io/jonathancarvalho39/scrum-center-api:latest
```

### Validações

- [ ] Imagem publica (sem necessidade de auth para pull)
- [ ] Multi-arch disponível (amd64 e arm64)
- [ ] Tamanho razoável (~15-20MB)
- [ ] Labels corretos (metadata)
- [ ] Healthcheck funciona (se aplicável)

### Testar Imagem

```bash
# Rodar
docker run -d --name test-api \
  -p 8080:8080 \
  -e DB_HOST=host.docker.internal \
  ghcr.io/jonathancarvalho39/scrum-center-api:latest

# Esperar
sleep 10

# Testar
curl http://localhost:8080/health

# Logs
docker logs test-api

# Limpar
docker stop test-api
docker rm test-api
```

---

## 🔐 Validação Security

### GitHub Security Tab

- [ ] Code scanning habilitado
- [ ] Dependabot alerts habilitado (se quiser)
- [ ] Secret scanning habilitado

### SARIF Uploads

- [ ] gosec results uploadados
- [ ] Trivy results uploadados
- [ ] Sem vulnerabilidades críticas

### Manual Check

```bash
# Local
golangci-lint run --enable=gosec
govulncheck ./...

# Docker
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock \
  aquasec/trivy image ghcr.io/jonathancarvalho39/scrum-center-api:latest
```

---

## 📊 Validação Workflows Específicos

### CI Workflow
- [ ] Lint job passou
- [ ] Test matrix (1.20, 1.21, 1.22) todos verdes
- [ ] Build job passou
- [ ] Security job passou
- [ ] Dependencies job passou
- [ ] Artifact de build uploadado
- [ ] Cache funcionando (build subsequente mais rápido)

### Docker Workflow
- [ ] Build multi-arch passou
- [ ] Push para GHCR bem-sucedido
- [ ] Tags corretas aplicadas
- [ ] Security scan passou
- [ ] Test image job passou
- [ ] Cache de layers funcionando

### Release Workflow
- [ ] Release criado automaticamente
- [ ] Changelog gerado
- [ ] Todos os binários built
- [ ] Checksums corretos
- [ ] Docker tags de versão criadas

### Cleanup Workflow
- [ ] Workflow existe
- [ ] Cron configurado
- [ ] workflow_dispatch habilitado (pode rodar manualmente)

---

## ✅ Checklist Final

### Estrutura
- [x] Todos os workflows criados
- [x] Dockerfile otimizado
- [x] docker-compose atualizado
- [x] Makefile estendido
- [x] Linter configurado
- [x] Documentação completa

### Funcionalidade
- [ ] CI passa em push
- [ ] Docker build funciona
- [ ] Imagem disponível no GHCR
- [ ] Release workflow funciona
- [ ] Security scans rodando

### Documentação
- [x] README atualizado
- [x] Setup summary criado
- [x] CI/CD architecture documentado
- [x] GHCR guide criado
- [x] Quick reference criado
- [x] Validation checklist criado

### Próximos Passos
- [ ] Fazer primeiro commit
- [ ] Verificar CI rodar
- [ ] Verificar imagem no GHCR
- [ ] Criar primeiro release
- [ ] Configurar branch protection
- [ ] (Opcional) Configurar Codecov

---

## 🎯 Critérios de Sucesso

### ✅ Setup Bem-Sucedido Se:

1. **CI passa** em todo push/PR
2. **Imagens Docker** publicadas automaticamente no GHCR
3. **Releases** criadas automaticamente com binários
4. **Security scans** rodando sem vulnerabilidades críticas
5. **Cache** funcionando (builds subsequentes ~2-3x mais rápidos)
6. **Documentação** acessível e completa

### 📈 Métricas Esperadas:

- **CI Time:** 3-5 minutos
- **Docker Build:** 5-8 minutos (primeira vez), 2-3 minutos (com cache)
- **Release Time:** 8-12 minutos
- **Image Size:** 15-25 MB
- **Coverage:** >80% (meta)

---

## 🐛 Se Algo Falhar

### CI não roda
1. Verificar: GitHub Actions habilitado?
2. Verificar: YAML syntax correto?
3. Ver: Actions → Failed workflow → Logs

### Docker push falha
1. Verificar: `GITHUB_TOKEN` tem permissão?
2. Verificar: Packages habilitado no repo?
3. Ver: Workflow permissions configuradas?

### Testes falham no CI mas passam local
1. Rodar: `make test-race` (detectar race conditions)
2. Verificar: Variáveis de ambiente?
3. Ver: Logs do workflow

### Imagem muito grande
1. Verificar: Stage final é distroless?
2. Verificar: Não tem arquivos desnecessários?
3. Rodar: `docker history <image>` para ver layers

---

**🎉 Se tudo na checklist está ✅, o setup está completo e funcional!**

**📌 Mantenha este arquivo atualizado conforme o projeto evolui.**

