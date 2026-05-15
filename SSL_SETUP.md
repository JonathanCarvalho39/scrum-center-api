# 🔐 Configuração SSL/HTTPS

## Pré-requisitos ✅
- [x] Domínio: scrumcenter.duckdns.org
- [x] Porta 443 aberta na VM
- [x] Nginx configurado

## 🚀 Passo a Passo

### 1. Deploy Inicial (sem SSL)
```bash
# Faça o primeiro deploy
git push origin main

# Verifique se funciona via HTTP
curl http://scrumcenter.duckdns.org/health
```

### 2. Conectar na VM
```bash
ssh -i sua-chave.pem ubuntu@SEU_IP
```

### 3. Configurar SSL (1 comando!)
```bash
cd /opt/scrum-center-api
sudo bash setup-ssl.sh seu-email@example.com
```

**O que o script faz:**
- ✅ Para Nginx temporariamente
- ✅ Instala certbot
- ✅ Gera certificado Let's Encrypt
- ✅ Copia certificados para o projeto
- ✅ **Habilita HTTPS automaticamente no Nginx**
- ✅ Reinicia containers

### 4. Reiniciar Containers
```bash
cd /opt/scrum-center-api
docker compose up -d
```

### 5. Testar HTTPS
```bash
# Via curl
curl https://scrumcenter.duckdns.org/health

# Pelo browser
https://scrumcenter.duckdns.org/health
```

### 6. Configurar Renovação Automática
```bash
cd /opt/scrum-center-api
sudo bash setup-ssl-renew.sh
```

**Renovação configurada:**
- 🔄 Todo domingo às 3h da manhã
- 📄 Logs em `/var/log/ssl-renew.log`

---

## 🧪 Comandos Úteis

### Ver certificados
```bash
sudo certbot certificates
```

### Testar renovação (sem renovar de verdade)
```bash
sudo certbot renew --dry-run
```

### Forçar renovação manual
```bash
cd /opt/scrum-center-api
sudo bash renew-ssl.sh
```

### Ver logs de renovação
```bash
sudo tail -f /var/log/ssl-renew.log
```

### Verificar validade do certificado
```bash
echo | openssl s_client -servername scrumcenter.duckdns.org -connect scrumcenter.duckdns.org:443 2>/dev/null | openssl x509 -noout -dates
```

---

## 🔧 Troubleshooting

### Erro: "Port 80 already in use"
```bash
# Parar Nginx
cd /opt/scrum-center-api
docker compose stop nginx

# Tentar novamente
sudo bash setup-ssl.sh seu-email@example.com
```

### Certificado não renova automaticamente
```bash
# Ver crontab
crontab -l

# Reconfigurar
sudo bash setup-ssl-renew.sh
```

### HTTPS não funciona após configurar
```bash
# Verificar se certificados existem
ls -la /opt/scrum-center-api/nginx/ssl/

# Ver logs do Nginx
docker logs scrum-center-nginx

# Reiniciar Nginx
docker compose restart nginx
```

### Erro: "Challenge failed"
```bash
# Verificar se porta 80 está acessível externamente
curl http://scrumcenter.duckdns.org

# Verificar firewall
sudo ufw status
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
```

---

## 📋 Checklist Completo

- [ ] Deploy inicial funcionando (HTTP)
- [ ] Conectado na VM via SSH
- [ ] Executado `setup-ssl.sh`
- [ ] Containers reiniciados
- [ ] HTTPS funcionando
- [ ] Renovação automática configurada
- [ ] Testado renovação (dry-run)

---

## 🌐 URLs Finais

**HTTP (será redirecionado para HTTPS):**
- http://scrumcenter.duckdns.org

**HTTPS:**
- https://scrumcenter.duckdns.org/health
- https://scrumcenter.duckdns.org/api/v1/...

---

## ⚠️ Importante

- Certificado Let's Encrypt expira em **90 dias**
- Renovação automática roda **toda semana**
- Certbot só renova se faltar **menos de 30 dias** para expirar
- Você receberá **email de aviso** se renovação falhar

---

## 🔄 Depois de Configurar SSL

O Nginx automaticamente:
- ✅ Redireciona HTTP → HTTPS
- ✅ Aceita apenas HTTPS (exceto `/health` por IP)
- ✅ Usa TLS 1.2 e 1.3
- ✅ Ciphers seguros

---

**📝 Nota:** Scripts `setup-ssl.sh` e `setup-ssl-renew.sh` são enviados automaticamente em cada deploy!

