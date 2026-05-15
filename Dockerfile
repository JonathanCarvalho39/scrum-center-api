# Stage 1: Build
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Instalar dependências necessárias para build
RUN apk add --no-cache git make ca-certificates

# Copiar go.mod e go.sum primeiro para aproveitar cache de layers
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copiar código fonte
COPY . .

# Build com flags de otimização e CGO desabilitado para binary estático
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a \
    -o /app/bin/api \
    ./cmd/api

# Stage 2: Runtime mínimo com distroless
FROM gcr.io/distroless/static-debian12:nonroot

# Copiar certificados CA para HTTPS funcionar
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copiar binário
COPY --from=builder /app/bin/api /api

# Usar usuário não-root (distroless já vem com nonroot user UID 65532)
USER nonroot:nonroot

EXPOSE 8080

# Healthcheck não funciona em distroless (sem shell), será feito no Kubernetes/Docker Compose

ENTRYPOINT ["/api"]
