#!/bin/bash

# Script de setup inicial da VM para deploy
# Executar UMA VEZ na VM antes do primeiro deploy:
# curl -fsSL https://raw.githubusercontent.com/SEU_USUARIO/scrum-center-api/main/setup-vm.sh | bash

set -e

echo "🚀 Configurando VM para deploy..."

# 1. Atualizar sistema
echo "📦 Atualizando sistema..."
sudo apt-get update
sudo apt-get upgrade -y

# 2. Instalar Docker
echo "🐳 Instalando Docker..."
if ! command -v docker &> /dev/null; then
    sudo apt-get install -y ca-certificates curl
    sudo install -m 0755 -d /etc/apt/keyrings
    sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
    sudo chmod a+r /etc/apt/keyrings/docker.asc

    echo \
      "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
      $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
      sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

    sudo apt-get update
    sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

    echo "✅ Docker instalado"
else
    echo "✅ Docker já está instalado"
fi

# 3. Adicionar usuário ao grupo docker
echo "👤 Configurando permissões do usuário..."
sudo usermod -aG docker $USER

# 4. Criar diretório de deploy
echo "📁 Criando diretório de deploy..."
sudo mkdir -p /opt/scrum-center-api
sudo chown $USER:$USER /opt/scrum-center-api

# 5. Configurar firewall
echo "🔥 Configurando firewall..."
sudo ufw allow 22/tcp    # SSH
sudo ufw allow 80/tcp    # HTTP
sudo ufw allow 443/tcp   # HTTPS
sudo ufw --force enable

echo ""
echo "✅ Setup completo!"
echo ""
echo "⚠️  IMPORTANTE: Faça logout e login novamente para aplicar permissões do Docker"
echo "   Comando: exit"
echo "   Depois: ssh novamente na VM"
echo ""
echo "Após re-login, teste:"
echo "   docker ps"
echo "   docker compose version"
echo ""

