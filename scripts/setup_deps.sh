#!/bin/bash
set -e

# Colors for output
GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${CYAN}🔄 Detecting package manager...${NC}"

if command -v apt-get &> /dev/null; then
    PKG_MANAGER="apt-get"
    INSTALL_CMD="apt-get install -y"
    UPDATE_CMD="apt-get update -qq"
    DEPS="ffmpeg pulseaudio-utils python3 python3-pip python3-venv git golang-go make"
elif command -v dnf &> /dev/null; then
    PKG_MANAGER="dnf"
    INSTALL_CMD="dnf install -y"
    UPDATE_CMD="dnf check-update"
    DEPS="ffmpeg pulseaudio-utils python3 python3-pip python3-virtualenv git golang make"
elif command -v pacman &> /dev/null; then
    PKG_MANAGER="pacman"
    INSTALL_CMD="pacman -S --noconfirm"
    UPDATE_CMD="pacman -Sy"
    DEPS="ffmpeg libpulse python python-pip git go make"
elif command -v zypper &> /dev/null; then
    PKG_MANAGER="zypper"
    INSTALL_CMD="zypper install -y"
    UPDATE_CMD="zypper refresh"
    DEPS="ffmpeg pulseaudio-utils python3 python3-pip python3-venv git go1.25 make"
else
    echo -e "${RED}Unsupported package manager. Please install dependencies manually: ffmpeg, pulseaudio-utils, python3, golang, make.${NC}"
    exit 1
fi

echo -e "${CYAN}🔄 Detected: $PKG_MANAGER. Installing dependencies...${NC}"
$UPDATE_CMD || true
$INSTALL_CMD $DEPS
echo -e "${GREEN}✅ System dependencies installed.${NC}"
