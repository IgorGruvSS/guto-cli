#!/bin/bash
set -e

# Colors for output
GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Get the absolute path of the script directory and project root
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
HOOKS_DIR="$PROJECT_ROOT/.git/hooks"

echo -e "${CYAN}🎣 Setting up Git Hooks...${NC}"

if [ ! -d "$HOOKS_DIR" ]; then
    echo -e "${RED}Error: .git/hooks directory not found. Is this a git repository?${NC}"
    exit 1
fi

PRE_PUSH_HOOK="$HOOKS_DIR/pre-push"

cat > "$PRE_PUSH_HOOK" <<EOL
#!/bin/bash
# Pre-push hook to run CI checks
echo -e "${CYAN}🚀 Running pre-push CI checks...${NC}"

# Run make ci from the project root
PROJECT_ROOT=\$(git rev-parse --show-toplevel)
cd "\$PROJECT_ROOT"

make ci

if [ \$? -ne 0 ]; then
    echo -e "${RED}❌ CI checks failed. Push aborted.${NC}"
    exit 1
fi

echo -e "${GREEN}✅ CI checks passed. Pushing...${NC}"
exit 0
EOL

chmod +x "$PRE_PUSH_HOOK"
echo -e "${GREEN}✅ Pre-push hook installed at $PRE_PUSH_HOOK${NC}"
echo "Run 'make ci' manually to test your changes before pushing."
