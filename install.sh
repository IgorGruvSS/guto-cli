#!/bin/bash
set -e

# Cores para output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}📦 Guto CLI - Script de Instalação (Debian/Ubuntu)${NC}"

# 1. Verifica sudo
if [ "$EUID" -ne 0 ]; then 
  echo -e "${RED}Por favor, rode como root (sudo ./install.sh)${NC}"
  exit 1
fi

# 2. Instala dependências do sistema
echo "🔄 Atualizando repositórios e instalando dependências de sistema..."
apt-get update -qq
apt-get install -y ffmpeg pulseaudio-utils python3 python3-pip python3-venv git golang-go

# 3. Verifica Go
if ! command -v go &> /dev/null; then
    echo -e "${RED}Go não encontrado mesmo após tentativa de instalação. Instale Go 1.25+ manualmente.${NC}"
    exit 1
fi

# 4. Build do Guto
echo "🔨 Compilando Guto CLI..."
# Assume que estamos na raiz do projeto
if [ ! -f "go.mod" ]; then
    echo -e "${RED}Erro: go.mod não encontrado. Rode o script na raiz do projeto.${NC}"
    exit 1
fi

# Ajuste de permissão para build se rodado com sudo mas em diretório de usuário
# Melhor rodar o build como usuário original se possível, mas aqui vamos compilar e mover
go build -o guto main.go
mv guto /usr/local/bin/
echo -e "${GREEN}✅ Binário instalado em /usr/local/bin/guto${NC}"

# 5. Configuração do Whisper (Opcional mas recomendado)
# Cria um venv em /opt/guto-whisper para ser global, ou instrui o usuário
echo "⚠️  Configuração do Whisper:"
echo "O Guto espera um ambiente Python com whisper-ctranslate2."
echo "Recomendamos criar um virtualenv dedicado:"
echo "  python3 -m venv ~/.local/share/whisper-env"
echo "  ~/.local/share/whisper-env/bin/pip install whisper-ctranslate2"
echo ""
echo "Depois, configure no Guto:"
echo "  guto config set scribe.python_bin ~/.local/share/whisper-env/bin/python3"

# 6. Finalização
echo ""
echo -e "${GREEN}🎉 Instalação concluída!${NC}"
echo "Teste com: guto --help"
echo "Configure seu áudio se necessário: guto config set audio.input_source <source_name>"
