#!/bin/bash
set -e

# Colors for output
GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}📦 Guto CLI - Universal Installation Script${NC}"

# 1. Check for sudo
if [ "$EUID" -ne 0 ]; then 
  echo -e "${RED}Please run as root (sudo ./install.sh)${NC}"
  exit 1
fi

# Detect package manager
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
    echo -e "${RED}Unsupported package manager. Please install dependencies manually: ffmpeg, pulseaudio-utils, python3, golang.${NC}"
    exit 1
fi

echo -e "${CYAN}🔄 Detected: $PKG_MANAGER. Installing dependencies...${NC}"
$UPDATE_CMD || true
$INSTALL_CMD $DEPS

# 2. Build Guto
echo -e "${CYAN}🔨 Compiling Guto CLI...${NC}"
if [ ! -f "go.mod" ]; then
    echo -e "${RED}Error: go.mod not found. Run the script from the project root.${NC}"
    exit 1
fi

go build -o guto main.go
mv guto /usr/local/bin/
echo -e "${GREEN}✅ Binary installed to /usr/local/bin/guto${NC}"

# 3. Automatic Whisper Configuration
WHISPER_DIR="/opt/guto/whisper-env"
echo -e "${CYAN}🤖 Configuring Guto Scribe (Whisper)...${NC}"
mkdir -p /opt/guto
if [ ! -d "$WHISPER_DIR" ]; then
    python3 -m venv "$WHISPER_DIR"
    "$WHISPER_DIR/bin/pip" install --upgrade pip
    "$WHISPER_DIR/bin/pip" install whisper-ctranslate2
    echo -e "${GREEN}✅ Whisper environment created at $WHISPER_DIR${NC}"
else
    echo -e "${YELLOW}ℹ️  Whisper environment already exists at $WHISPER_DIR. Skipping installation.${NC}"
fi

# 4. LLM Instructions (Gemini)
echo -e "${YELLOW}⚠️  AI Configuration (Guto Press):${NC}"
echo "Guto uses Google Gemini for summaries."
echo "1. Obtain an API key at: https://aistudio.google.com/app/apikey"
echo "2. Configure it in Guto:"
echo -e "${CYAN}   guto config set press.api_key YOUR_TOKEN_HERE${NC}"
echo ""
echo "3. List and choose a model (optional, default: gemini-2.5-flash):"
echo -e "${CYAN}   guto config models${NC}"
echo -e "${CYAN}   guto config set press.model MODEL_NAME${NC}"
echo ""
echo "4. Configure the path for the Whisper installed now:"
echo -e "${CYAN}   guto config set scribe.python_bin $WHISPER_DIR/bin/python3${NC}"

# 5. Finalization
echo ""
echo -e "${GREEN}🎉 Installation completed!${NC}"
echo "Verba volant, scripta manent."
