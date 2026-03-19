#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

WHISPER_DIR="/opt/guto/whisper-env"

# 4. LLM Instructions (Gemini)
echo -e "${YELLOW}⚠️  AI Configuration (Guto Press):${NC}"
echo "Guto uses Google Gemini for summaries."
echo "1. Obtain an API key at: https://aistudio.google.com/app/apikey"
echo "2. Configure it in Guto:"
echo -e "${CYAN}   guto config set press.api_key YOUR_TOKEN_HERE${NC}"
echo ""
echo "3. List and choose a model (optional, default: gemini-2.0-flash):"
echo -e "${CYAN}   guto config models${NC}"
echo -e "${CYAN}   guto config set press.model MODEL_NAME${NC}"
echo ""
echo "4. Configure the path for the Whisper installed now:"
echo -e "${CYAN}   guto config set scribe.python_bin $WHISPER_DIR/bin/python3${NC}"
echo ""

# 5. Audio and Output Configuration
echo -e "${YELLOW}🎙️  Audio and Output Configuration:${NC}"
echo "1. List available audio devices (System/Mic):"
echo -e "${CYAN}   guto config audio-devices${NC}"
echo "2. Set your preferred sources:"
echo -e "${CYAN}   guto config set audio.output_monitor <NAME>${NC}"
echo -e "${CYAN}   guto config set audio.input_source <NAME>${NC}"
echo ""
echo "3. (Optional) Set a fixed Output directory (default: ./Output):"
echo -e "${CYAN}   guto config set output.base_dir /path/to/your/project/Output${NC}"

# 6. Finalization
echo ""
echo -e "${GREEN}🎉 Installation completed!${NC}"
echo "Verba volant, scripta manent."
