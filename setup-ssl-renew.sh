#!/bin/bash

# Script para configurar renovação automática de SSL
# Executar na VM: sudo bash setup-ssl-renew.sh

set -e

DOMAIN="scrumcenter.duckdns.org"
DEPLOY_DIR="/opt/scrum-center-api"
RENEW_SCRIPT="$DEPLOY_DIR/renew-ssl.sh"

echo "🔄 Configurando renovação automática de SSL"

# Criar script de renovação
cat > $RENEW_SCRIPT << 'EOF'
#!/bin/bash

DOMAIN="scrumcenter.duckdns.org"
DEPLOY_DIR="/opt/scrum-center-api"

# Parar Nginx (liberar porta 80 para certbot)
cd $DEPLOY_DIR
docker compose stop nginx

# Renovar certificado
certbot renew --quiet

# Copiar certificados atualizados
cp /etc/letsencrypt/live/$DOMAIN/fullchain.pem $DEPLOY_DIR/nginx/ssl/
cp /etc/letsencrypt/live/$DOMAIN/privkey.pem $DEPLOY_DIR/nginx/ssl/
chown -R ubuntu:ubuntu $DEPLOY_DIR/nginx/ssl/
chmod 644 $DEPLOY_DIR/nginx/ssl/fullchain.pem
chmod 600 $DEPLOY_DIR/nginx/ssl/privkey.pem

# Reiniciar Nginx
docker compose up -d nginx

echo "✅ SSL renovado em $(date)" >> /var/log/ssl-renew.log
EOF

# Tornar executável
chmod +x $RENEW_SCRIPT

# Adicionar ao crontab (todo domingo às 3h da manhã)
CRON_CMD="0 3 * * 0 $RENEW_SCRIPT >> /var/log/ssl-renew.log 2>&1"

# Verificar se já existe
if ! crontab -l 2>/dev/null | grep -q "$RENEW_SCRIPT"; then
    (crontab -l 2>/dev/null; echo "$CRON_CMD") | crontab -
    echo "✅ Crontab configurado"
else
    echo "ℹ️  Crontab já configurado"
fi

echo ""
echo "✅ Renovação automática configurada!"
echo ""
echo "📝 Detalhes:"
echo "   • Script: $RENEW_SCRIPT"
echo "   • Frequência: Todos os domingos às 3h"
echo "   • Log: /var/log/ssl-renew.log"
echo ""
echo "🧪 Testar renovação (dry-run):"
echo "   sudo certbot renew --dry-run"
echo ""
echo "📋 Ver quando expira:"
echo "   sudo certbot certificates"
echo ""
echo "📄 Ver logs de renovação:"
echo "   sudo tail -f /var/log/ssl-renew.log"
echo ""

