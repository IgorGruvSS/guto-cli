#!/bin/bash
set -e

# Colors for output
GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

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
