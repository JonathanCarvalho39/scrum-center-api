# Nginx SSL Configuration - scrumcenter.duckdns.org

## Configuração Atual
- HTTP na porta 80 (sem SSL)
- Domínio: scrumcenter.duckdns.org
- Reverse proxy para API na porta 8080

## Para Habilitar HTTPS

### 1. Gerar Certificado SSL com Let's Encrypt

**Certbot com DuckDNS:**
```bash
# Na VM, parar o nginx temporariamente
cd /opt/scrum-center-api
docker compose stop nginx

# Instalar certbot
sudo apt update
sudo apt install certbot -y

# Gerar certificado (standalone mode)
sudo certbot certonly --standalone -d scrumcenter.duckdns.org --agree-tos --email seu-email@example.com

# Certificados serão salvos em:
# /etc/letsencrypt/live/scrumcenter.duckdns.org/fullchain.pem
# /etc/letsencrypt/live/scrumcenter.duckdns.org/privkey.pem
```

### 2. Copiar Certificados para o Projeto
```bash
# Na VM
sudo cp /etc/letsencrypt/live/scrumcenter.duckdns.org/fullchain.pem \
   /opt/scrum-center-api/nginx/ssl/

sudo cp /etc/letsencrypt/live/scrumcenter.duckdns.org/privkey.pem \
   /opt/scrum-center-api/nginx/ssl/

# Ajustar permissões
sudo chown ubuntu:ubuntu /opt/scrum-center-api/nginx/ssl/*.pem
sudo chmod 644 /opt/scrum-center-api/nginx/ssl/fullchain.pem
sudo chmod 600 /opt/scrum-center-api/nginx/ssl/privkey.pem
```

### 3. Habilitar HTTPS no Nginx
Edite `/opt/scrum-center-api/nginx/conf.d/api.conf`:
- Descomente os blocos `server { listen 443 ssl http2; ... }`
- Descomente o bloco de redirect HTTP → HTTPS no final

### 4. Reiniciar Nginx
```bash
cd /opt/scrum-center-api
docker compose up -d nginx
```

### 5. Testar
```bash
# Verificar HTTP redireciona para HTTPS
curl -I http://scrumcenter.duckdns.org

# Testar HTTPS
curl https://scrumcenter.duckdns.org/health
```

## Renovação Automática

Let's Encrypt expira em 90 dias. Configure renovação automática:

```bash
# Criar script de renovação
sudo nano /opt/scrum-center-api/renew-ssl.sh
```

Conteúdo do script:
```bash
#!/bin/bash
cd /opt/scrum-center-api
docker compose stop nginx
certbot renew --quiet
cp /etc/letsencrypt/live/scrumcenter.duckdns.org/fullchain.pem nginx/ssl/
cp /etc/letsencrypt/live/scrumcenter.duckdns.org/privkey.pem nginx/ssl/
chown ubuntu:ubuntu nginx/ssl/*.pem
chmod 644 nginx/ssl/fullchain.pem
chmod 600 nginx/ssl/privkey.pem
docker compose up -d nginx
```

Tornar executável e adicionar ao crontab:
```bash
sudo chmod +x /opt/scrum-center-api/renew-ssl.sh

# Crontab (roda todo domingo às 3h da manhã)
sudo crontab -e
# Adicionar linha:
0 3 * * 0 /opt/scrum-center-api/renew-ssl.sh >> /var/log/ssl-renew.log 2>&1
```

## Verificar Certificado
```bash
# Ver validade
sudo certbot certificates

# Testar renovação (dry-run)
sudo certbot renew --dry-run
```


