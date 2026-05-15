# Stage 1: Build
FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git make

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build

# Stage 2: Runtime
FROM golang:1.22-alpine

WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/bin/api /usr/local/bin/api

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

USER appuser

EXPOSE 8080

CMD ["api"]
