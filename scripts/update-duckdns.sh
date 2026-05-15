#!/bin/bash

# Script para atualizar DuckDNS na VM
# Executar via crontab a cada 5 minutos

DOMAIN="scrumcenter"
TOKEN="SEU_TOKEN_AQUI"

# Atualiza IP
RESPONSE=$(curl -s "https://www.duckdns.org/update?domains=$DOMAIN&token=$TOKEN&ip=")

if [ "$RESPONSE" = "OK" ]; then
    echo "$(date): DuckDNS updated successfully" >> /var/log/duckdns.log
else
    echo "$(date): DuckDNS update failed: $RESPONSE" >> /var/log/duckdns.log
fi

