#!/bin/bash
set -e

# Cores para output
GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}📦 Guto CLI - Script de Instalação Universal${NC}"

# 1. Verifica sudo
if [ "$EUID" -ne 0 ]; then 
  echo -e "${RED}Por favor, rode como root (sudo ./install.sh)${NC}"
  exit 1
fi

# Detecta o gerenciador de pacotes
if command -v apt-get &> /dev/null; then
    PKG_MANAGER="apt-get"
    INSTALL_CMD="apt-get install -y"
    UPDATE_CMD="apt-get update -qq"
    DEPS="ffmpeg pulseaudio-utils python3 python3-pip python3-venv git golang-go"
elif command -v dnf &> /dev/null; then
    PKG_MANAGER="dnf"
    INSTALL_CMD="dnf install -y"
    UPDATE_CMD="dnf check-update"
    DEPS="ffmpeg pulseaudio-utils python3 python3-pip python3-virtualenv git golang"
elif command -v pacman &> /dev/null; then
    PKG_MANAGER="pacman"
    INSTALL_CMD="pacman -S --noconfirm"
    UPDATE_CMD="pacman -Sy"
    DEPS="ffmpeg libpulse python python-pip git go"
elif command -v zypper &> /dev/null; then
    PKG_MANAGER="zypper"
    INSTALL_CMD="zypper install -y"
    UPDATE_CMD="zypper refresh"
    DEPS="ffmpeg pulseaudio-utils python3 python3-pip python3-venv git go1.25"
else
    echo -e "${RED}Gerenciador de pacotes não suportado. Instale as dependências manualmente: ffmpeg, pulseaudio-utils, python3, golang.${NC}"
    exit 1
fi

echo -e "${CYAN}🔄 Detectado: $PKG_MANAGER. Instalando dependências...${NC}"
$UPDATE_CMD || true
$INSTALL_CMD $DEPS

# 2. Build do Guto
echo -e "${CYAN}🔨 Compilando Guto CLI...${NC}"
if [ ! -f "go.mod" ]; then
    echo -e "${RED}Erro: go.mod não encontrado. Rode o script na raiz do projeto.${NC}"
    exit 1
fi

go build -o guto main.go
mv guto /usr/local/bin/
echo -e "${GREEN}✅ Binário instalado em /usr/local/bin/guto${NC}"

# 3. Configuração Automática do Whisper
WHISPER_DIR="/opt/guto/whisper-env"
echo -e "${CYAN}🤖 Configurando Guto Scribe (Whisper)...${NC}"
mkdir -p /opt/guto
if [ ! -d "$WHISPER_DIR" ]; then
    python3 -m venv "$WHISPER_DIR"
    "$WHISPER_DIR/bin/pip" install --upgrade pip
    "$WHISPER_DIR/bin/pip" install whisper-ctranslate2
    echo -e "${GREEN}✅ Ambiente Whisper criado em $WHISPER_DIR${NC}"
else
    echo -e "${YELLOW}ℹ️  Ambiente Whisper já existe em $WHISPER_DIR. Pulando instalação.${NC}"
fi

# 4. Instruções de LLM (Gemini)
echo -e "${YELLOW}⚠️  Configuração de IA (Guto Press):${NC}"
echo "O Guto utiliza o Google Gemini para sumários."
echo "1. Obtenha uma chave de API em: https://aistudio.google.com/app/apikey"
echo "2. Configure no Guto:"
echo -e "${CYAN}   guto config set press.api_key SEU_TOKEN_AQUI${NC}"
echo ""
echo "3. Liste e escolha um modelo (opcional, padrão: gemini-2.5-flash):"
echo -e "${CYAN}   guto config models${NC}"
echo -e "${CYAN}   guto config set press.model NOME_DO_MODELO${NC}"
echo ""
echo "4. Configure o caminho do Whisper instalado agora:"
echo -e "${CYAN}   guto config set scribe.python_bin $WHISPER_DIR/bin/python3${NC}"

# 5. Finalização
echo ""
echo -e "${GREEN}🎉 Instalação concluída!${NC}"
echo "Verba volant, scripta manent."
