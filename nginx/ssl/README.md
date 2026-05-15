# Nginx SSL Configuration

## Configuração Atual
- HTTP na porta 80 (sem SSL)
- Reverse proxy para API na porta 8080

## Para Habilitar HTTPS

### 1. Gerar Certificado SSL

**Opção A: Let's Encrypt (Recomendado - Grátis)**
```bash
# Na VM
sudo apt install certbot
sudo certbot certonly --standalone -d seu-dominio.com
```

**Opção B: Certificado Self-Signed (Desenvolvimento)**
```bash
# Na VM
sudo openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout /opt/scrum-center-api/nginx/ssl/privkey.pem \
  -out /opt/scrum-center-api/nginx/ssl/fullchain.pem
```

### 2. Copiar Certificados
```bash
# Se usar Let's Encrypt
sudo cp /etc/letsencrypt/live/seu-dominio.com/fullchain.pem nginx/ssl/
sudo cp /etc/letsencrypt/live/seu-dominio.com/privkey.pem nginx/ssl/
```

### 3. Descomentar Bloco HTTPS
Edite `nginx/conf.d/api.conf`:
- Descomente o bloco `server { listen 443 ssl http2; ... }`
- Descomente a linha `return 301 https://...` no bloco HTTP
- Troque `seu-dominio.com` pelo seu domínio real

### 4. Reiniciar Nginx
```bash
docker compose restart nginx
```

## Renovação Automática (Let's Encrypt)
```bash
# Crontab na VM
0 0 * * 0 certbot renew --quiet && docker compose -f /opt/scrum-center-api/docker-compose.yml restart nginx
```

