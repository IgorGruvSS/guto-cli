#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${YELLOW}🧹 Guto CLI - Interactive Uninstaller (Safe)${NC}"
echo "----------------------------------------------------"

confirm_remove() {
    local path=$1
    local desc=$2
    local use_sudo=$3

    if [ -e "$path" ] || [ -d "$path" ]; then
        echo -e "${CYAN}❓ Do you want to remove $desc? ($path)${NC}"
        read -p "[y/N]: " choice
        case "$choice" in 
          y|Y ) 
            echo -e "${RED}🗑️  Removing $path...${NC}"
            if [ "$use_sudo" = "true" ]; then
                sudo rm -rf "$path"
            else
                rm -rf "$path"
            fi
            ;;
          * ) echo -e "${GREEN}✅ Kept: $path${NC}" ;;
        esac
    fi
}

# 1. Binaries in PATH
# Search for all occurrences of guto to ensure total cleanup
INSTANCES=$(which -a guto 2>/dev/null)
if [ ! -z "$INSTANCES" ]; then
    echo -e "${YELLOW}🔍 Guto instances found in PATH:${NC}"
    for inst in $INSTANCES; do
        confirm_remove "$inst" "the binary in path" "true"
    done
fi

# 2. Whisper Environment (Venv in /opt/guto)
# This removes the isolated Python environment and downloaded models
confirm_remove "/opt/guto" "the Whisper environment (Venv and models in /opt/guto)" "true"

# 3. Guto Settings (~/.config/guto)
confirm_remove "$HOME/.config/guto" "user settings (~/.config/guto)" "false"

# 4. Local Output Folder
confirm_remove "./Output" "the generated files folder in the project (Output/)" "false"

# 5. Local binaries in the project folder
confirm_remove "./guto" "the local compiled binary (./guto)" "false"
confirm_remove "./bin" "the local binary folder (./bin)" "false"

# 6. Go Cache
echo -e "${CYAN}❓ Do you want to clear the Go module cache?${NC}"
read -p "[y/N]: " choice
case "$choice" in 
  y|Y ) 
    echo -e "${RED}🗑️  Clearing Go cache...${NC}"
    go clean -cache -modcache
    ;;
  * ) echo -e "${GREEN}✅ Go cache kept.${NC}" ;;
esac

echo "----------------------------------------------------"
echo -e "${GREEN}🎉 Cleanup completed! Now you can run ./install.sh again.${NC}"
