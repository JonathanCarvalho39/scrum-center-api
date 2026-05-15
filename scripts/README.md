# DuckDNS Auto-Update

## GitHub Actions (Recomendado)

Workflow automático atualiza o IP a cada 5 minutos.

### Configurar:
1. Adicione secret `TOKEN_DNS` no GitHub
2. Workflow `.github/workflows/update-dns.yml` roda automaticamente

## Alternativamente: Crontab na VM

Se preferir atualizar direto da VM:

### 1. Configurar script
```bash
# Na VM
cd /opt/scrum-center-api
nano scripts/update-duckdns.sh

# Substituir SEU_TOKEN_AQUI pelo token real
# Salvar e sair (Ctrl+X, Y, Enter)

# Tornar executável
chmod +x scripts/update-duckdns.sh
```

### 2. Adicionar ao crontab
```bash
crontab -e

# Adicionar linha (atualiza a cada 5 minutos):
*/5 * * * * /opt/scrum-center-api/scripts/update-duckdns.sh
```

### 3. Testar
```bash
# Executar manualmente
./scripts/update-duckdns.sh

# Ver log
tail -f /var/log/duckdns.log
```

## Verificar IP atual
```bash
# Ver IP do domínio
nslookup scrumcenter.duckdns.org

# Ver IP da VM
curl ifconfig.me
```

## Atualização Manual
```bash
curl "https://www.duckdns.org/update?domains=scrumcenter&token=SEU_TOKEN&ip="
```

