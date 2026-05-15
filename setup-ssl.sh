#!/bin/bash

# Script para configurar SSL com Let's Encrypt no DuckDNS
# Executar na VM: sudo bash setup-ssl.sh seu-email@example.com

set -e

EMAIL=$1
DOMAIN="scrumcenter.duckdns.org"
DEPLOY_DIR="/opt/scrum-center-api"

if [ -z "$EMAIL" ]; then
    echo "❌ Erro: Email é obrigatório"
    echo "Uso: sudo bash setup-ssl.sh seu-email@example.com"
    exit 1
fi

echo "🔐 Configurando SSL para $DOMAIN"
echo "📧 Email: $EMAIL"
echo ""

# 1. Parar Nginx temporariamente (liberar porta 80)
echo "⏸️  Parando Nginx..."
cd $DEPLOY_DIR
docker compose stop nginx || true

# 2. Instalar certbot
echo "📦 Instalando certbot..."
apt update
apt install -y certbot

# 3. Gerar certificado
echo "🔑 Gerando certificado SSL..."
certbot certonly \
    --standalone \
    --non-interactive \
    --agree-tos \
    --email "$EMAIL" \
    -d "$DOMAIN"

# 4. Copiar certificados para o projeto
echo "📋 Copiando certificados..."
mkdir -p $DEPLOY_DIR/nginx/ssl
cp /etc/letsencrypt/live/$DOMAIN/fullchain.pem $DEPLOY_DIR/nginx/ssl/
cp /etc/letsencrypt/live/$DOMAIN/privkey.pem $DEPLOY_DIR/nginx/ssl/
chown -R ubuntu:ubuntu $DEPLOY_DIR/nginx/ssl/
chmod 644 $DEPLOY_DIR/nginx/ssl/fullchain.pem
chmod 600 $DEPLOY_DIR/nginx/ssl/privkey.pem

# 5. Ativar HTTPS no Nginx
echo "🔧 Habilitando HTTPS no Nginx..."
NGINX_CONF="$DEPLOY_DIR/nginx/conf.d/api.conf"

# Criar backup
cp $NGINX_CONF $NGINX_CONF.bak

# Comentar bloco HTTP do domínio (linhas 23-58, mantém default_server)
sed -i '/^# HTTP Server para o domínio (ANTES de configurar SSL)/,/^}$/ {
    /^# HTTP Server para o domínio/! {
        /^}$/! s/^/# /
    }
}' $NGINX_CONF

# Descomentar blocos HTTPS (depois da linha com ====)
sed -i '/^# ============================================================/,$ {
    /^# server {$/,/^# }$/ s/^# //
    /^#     / s/^# //
    /^# Redirect/ s/^# //
}' $NGINX_CONF

echo ""
echo "✅ SSL configurado com sucesso!"
echo ""
echo "📝 Próximos passos:"
echo "   1. Reiniciar containers: cd $DEPLOY_DIR && docker compose up -d"
echo "   2. Testar HTTP→HTTPS redirect: curl -I http://$DOMAIN"
echo "   3. Testar HTTPS: curl https://$DOMAIN/health"
echo "   4. Configurar renovação: sudo bash $DEPLOY_DIR/setup-ssl-renew.sh"
echo ""


