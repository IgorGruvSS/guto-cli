#!/bin/bash
set -e

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# 1. Check for sudo
if [ "$EUID" -ne 0 ]; then 
  echo -e "${RED}Please run as root (sudo ./install.sh)${NC}"
  exit 1
fi

# Get the absolute path of the script directory and project root
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

cd "$PROJECT_ROOT"

if [ ! -f "Makefile" ]; then
    echo -e "${RED}Error: Makefile not found in $PROJECT_ROOT.${NC}"
    exit 1
fi

echo -e "${GREEN}📦 Guto CLI - Starting Full Setup via Make...${NC}"
make setup
