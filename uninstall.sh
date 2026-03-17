#!/bin/bash

# Cores para output
GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${YELLOW}🧹 Guto CLI - Desinstalador Interativo (Seguro)${NC}"
echo "----------------------------------------------------"

confirm_remove() {
    local path=$1
    local desc=$2
    local use_sudo=$3

    if [ -e "$path" ] || [ -d "$path" ]; then
        echo -e "${CYAN}❓ Deseja remover $desc? ($path)${NC}"
        read -p "[y/N]: " choice
        case "$choice" in 
          y|Y ) 
            echo -e "${RED}🗑️  Removendo $path...${NC}"
            if [ "$use_sudo" = "true" ]; then
                sudo rm -rf "$path"
            else
                rm -rf "$path"
            fi
            ;;
          * ) echo -e "${GREEN}✅ Mantido: $path${NC}" ;;
        esac
    fi
}

# 1. Binários no PATH
# Busca todas as ocorrências do guto para garantir limpeza total
INSTANCES=$(which -a guto 2>/dev/null)
if [ ! -z "$INSTANCES" ]; then
    echo -e "${YELLOW}🔍 Instâncias do guto encontradas no PATH:${NC}"
    for inst in $INSTANCES; do
        confirm_remove "$inst" "o binário no path" "true"
    done
fi

# 2. Ambiente Whisper (Venv em /opt/guto)
# Isso remove o ambiente Python isolado e os modelos baixados
confirm_remove "/opt/guto" "o ambiente Whisper (Venv e modelos em /opt/guto)" "true"

# 3. Configurações do Guto (~/.config/guto)
confirm_remove "$HOME/.config/guto" "as configurações do usuário (~/.config/guto)" "false"

# 4. Pasta de Output Local
confirm_remove "./Output" "a pasta de arquivos gerados no projeto (Output/)" "false"

# 5. Binários locais na pasta do projeto
confirm_remove "./guto" "o binário compilado local (./guto)" "false"
confirm_remove "./bin" "a pasta de binários local (./bin)" "false"

# 6. Cache do Go
echo -e "${CYAN}❓ Deseja limpar o cache de módulos do Go?${NC}"
read -p "[y/N]: " choice
case "$choice" in 
  y|Y ) 
    echo -e "${RED}🗑️  Limpando cache do Go...${NC}"
    go clean -cache -modcache
    ;;
  * ) echo -e "${GREEN}✅ Cache do Go mantido.${NC}" ;;
esac

echo "----------------------------------------------------"
echo -e "${GREEN}🎉 Limpeza concluída! Agora você pode rodar o ./install.sh novamente.${NC}"
