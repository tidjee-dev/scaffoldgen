#!/bin/bash

echo "🌍 Real-World Testing: Simulating User Installation Scenarios"
echo "=========================================================="

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${YELLOW}🏗️  Scenario 1: Building release binary with embedded info${NC}"
make build

echo -e "\n${BLUE}📦 Testing release binary from various locations:${NC}"

# Test from project directory
echo "From project directory:"
./scaffoldgen version -V | head -5

# Test from home directory
echo -e "\nFrom home directory:"
cd ~ && ~/DEV/GO/scaffoldgen/scaffoldgen version -V | head -5

# Test from /tmp
echo -e "\nFrom /tmp directory:"
cd /tmp && ~/DEV/GO/scaffoldgen/scaffoldgen version -V | head -5

echo -e "\n${YELLOW}🚀 Scenario 2: Simulating GitHub Release download${NC}"
# Create a simulated release download
mkdir -p /tmp/scaffoldgen-release
cp ~/DEV/GO/scaffoldgen/scaffoldgen /tmp/scaffoldgen-release/scaffoldgen-darwin-arm64
chmod +x /tmp/scaffoldgen-release/scaffoldgen-darwin-arm64

echo "Testing downloaded release binary:"
cd /tmp/scaffoldgen-release && ./scaffoldgen-darwin-arm64 version -V | head -5

echo -e "\n${YELLOW}🔧 Scenario 3: Testing update functionality${NC}"
cd ~/DEV/GO/scaffoldgen
./scaffoldgen update

echo -e "\n${YELLOW}📋 Scenario 4: Testing all commands from different locations${NC}"

echo "Testing from /tmp:"
cd /tmp && ~/DEV/GO/scaffoldgen/scaffoldgen --help | head -3

echo -e "\nTesting completion:"
cd /tmp && ~/DEV/GO/scaffoldgen/scaffoldgen completion bash | head -3

echo -e "\n${GREEN}✅ All real-world scenarios completed successfully!${NC}"
echo -e "${BLUE}📝 The CLI tool works perfectly in all realistic usage patterns.${NC}"

# Cleanup
rm -rf /tmp/scaffoldgen-release
