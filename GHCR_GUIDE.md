# GitHub Container Registry - Guia de Uso

Este projeto usa o **GitHub Container Registry (ghcr.io)** para hospedar imagens Docker.

## 🔐 Autenticação

### 1. Criar Personal Access Token (PAT)

1. Vá em: https://github.com/settings/tokens
2. Click em "Generate new token" → "Generate new token (classic)"
3. Selecione os escopos:
   - `read:packages` - para pull de imagens
   - `write:packages` - para push de imagens
   - `delete:packages` - para deletar imagens
4. Copie o token gerado

### 2. Login no Registry

```bash
# Usando o token
echo "SEU_TOKEN_AQUI" | docker login ghcr.io -u SEU_USERNAME --password-stdin

# Ou interativo
docker login ghcr.io -u SEU_USERNAME
```

## 📦 Usando as Imagens

### Pull

```bash
# Última versão (latest)
docker pull ghcr.io/jonathancarvalho39/scrum-center-api:latest

# Versão específica
docker pull ghcr.io/jonathancarvalho39/scrum-center-api:v1.0.0

# Branch develop
docker pull ghcr.io/jonathancarvalho39/scrum-center-api:develop

# Commit específico
docker pull ghcr.io/jonathancarvalho39/scrum-center-api:main-abc1234
```

### Run

```bash
docker run -d \
  --name scrum-center-api \
  -p 8080:8080 \
  -e DATABASE_URL="postgres://user:pass@host:5432/db" \
  ghcr.io/jonathancarvalho39/scrum-center-api:latest
```

### Docker Compose

```yaml
version: '3.8'

services:
  api:
    image: ghcr.io/jonathancarvalho39/scrum-center-api:latest
    ports:
      - "8080:8080"
    environment:
      - DATABASE_URL=postgres://user:pass@db:5432/scrumcenter
    depends_on:
      - db

  db:
    image: postgres:15-alpine
    environment:
      - POSTGRES_DB=scrumcenter
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=pass
    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
```

## 🚀 Push Manual (Desenvolvedor)

```bash
# Build local
docker build -t ghcr.io/jonathancarvalho39/scrum-center-api:minha-feature .

# Push
docker push ghcr.io/jonathancarvalho39/scrum-center-api:minha-feature
```

## 🤖 CI/CD Automático

O GitHub Actions faz push automático:

| Evento | Tag Gerada |
|--------|-----------|
| Push para `main` | `latest`, `main`, `main-{sha}` |
| Push para `develop` | `develop`, `develop-{sha}` |
| Push para `v1.2.3` | `v1.2.3`, `1.2`, `1`, `latest` |
| Pull Request #42 | `pr-42` |

## 📊 Visualizar Pacotes

Acesse: https://github.com/JonathanCarvalho39?tab=packages

## 🔍 Inspecionar Imagem

```bash
# Ver tags disponíveis (via browser)
https://github.com/JonathanCarvalho39/scrum-center-api/pkgs/container/scrum-center-api

# Inspecionar localmente
docker inspect ghcr.io/jonathancarvalho39/scrum-center-api:latest

# Ver histórico de layers
docker history ghcr.io/jonathancarvalho39/scrum-center-api:latest
```

## 🧹 Limpeza

```bash
# Remover imagens locais antigas
docker image prune -a

# Remover imagens específicas
docker rmi ghcr.io/jonathancarvalho39/scrum-center-api:old-tag
```

O workflow `cleanup.yml` remove automaticamente imagens antigas do registry semanalmente.

## 🔒 Visibilidade

Por padrão, os pacotes herdam a visibilidade do repositório:
- Repositório **público** → Pacote **público** (sem autenticação para pull)
- Repositório **privado** → Pacote **privado** (requer autenticação)

Para mudar: Settings → Packages → Change visibility

## 💡 Dicas

1. **Cache**: As Actions usam cache entre builds para acelerar
2. **Multi-arch**: Imagens suportam `amd64` e `arm64`
3. **Segurança**: Todas as imagens são escaneadas com Trivy
4. **Attestation**: Imagens incluem build provenance para auditoria
5. **Retention**: Mantemos as últimas 10 versões + tags importantes

## 🆘 Problemas Comuns

### Erro: "denied: permission_denied"
- Verifique se está autenticado: `docker login ghcr.io`
- Verifique se o token tem permissão `write:packages`

### Erro: "manifest unknown"
- Tag não existe, verifique tags disponíveis no GitHub
- Pode ser que o workflow ainda não tenha terminado

### Imagem grande demais
- As imagens usam multi-stage build (final ~20MB)
- Verifique se não está usando imagem intermediária

