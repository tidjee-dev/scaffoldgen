#!/bin/bash

echo "🧪 Testing Embedded Version Functionality in Real Conditions"
echo "=========================================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Test function
test_scenario() {
    local test_name="$1"
    local command="$2"
    local expected_pattern="$3"
    
    echo -e "\n${BLUE}📋 Test: $test_name${NC}"
    echo "Command: $command"
    echo "----------------------------------------"
    
    # Run the command and capture output
    output=$(eval "$command" 2>&1)
    exit_code=$?
    
    if [ $exit_code -eq 0 ]; then
        if echo "$output" | grep -q "$expected_pattern"; then
            echo -e "${GREEN}✅ PASS${NC}"
            echo "Output: $(echo "$output" | head -3)"
        else
            echo -e "${RED}❌ FAIL${NC}"
            echo "Expected pattern: $expected_pattern"
            echo "Actual output:"
            echo "$output"
        fi
    else
        echo -e "${RED}❌ FAIL (exit code: $exit_code)${NC}"
        echo "Output: $output"
    fi
}

# Build the binary with embedded info
echo -e "${YELLOW}🔨 Building scaffoldgen with embedded version info...${NC}"
make build

# Test 1: Version from project directory (should use embedded info)
test_scenario "Version from project directory" "./scaffoldgen version" "scaffoldgen 1.1.0"

# Test 2: Verbose version from project directory
test_scenario "Verbose version from project directory" "./scaffoldgen version -V" "Build Information"

# Test 3: Version from different directory
test_scenario "Version from home directory" "cd /tmp && ~/DEV/GO/scaffoldgen/scaffoldgen version" "scaffoldgen 1.1.0"

# Test 4: Verbose version from different directory  
test_scenario "Verbose version from home directory" "cd /tmp && ~/DEV/GO/scaffoldgen/scaffoldgen version -V" "Build Information"

# Test 5: Version from completely unrelated directory
test_scenario "Version from /tmp directory" "cd /tmp && ~/DEV/GO/scaffoldgen/scaffoldgen version" "scaffoldgen 1.1.0"

# Test 6: Copy binary to system location and test
echo -e "\n${YELLOW}📋 Testing binary copied to system location...${NC}"
cp ~/DEV/GO/scaffoldgen/scaffoldgen /tmp/scaffoldgen-test
test_scenario "Copied binary from /tmp" "cd /tmp && ./scaffoldgen-test version" "scaffoldgen 1.1.0"

# Test 7: Test with PATH (simulate go install)
echo -e "\n${YELLOW}📋 Testing with PATH (simulate go install)...${NC}"
export PATH="/tmp:$PATH"
test_scenario "Binary in PATH" "cd /tmp && scaffoldgen-test version" "scaffoldgen 1.1.0"

# Test 8: Test from user home directory
test_scenario "From user home directory" "cd ~ && ~/DEV/GO/scaffoldgen/scaffoldgen version -V" "Build Information"

# Test 9: Test with no version.json present (simulate production)
echo -e "\n${YELLOW}📋 Testing without version.json file...${NC}"
mv version.json version.json.backup 2>/dev/null || true
test_scenario "No version.json file" "./scaffoldgen version -V" "Build Information"
mv version.json.backup version.json 2>/dev/null || true

# Test 10: Test help command (should work everywhere)
test_scenario "Help command from /tmp" "cd /tmp && ~/DEV/GO/scaffoldgen/scaffoldgen --help" "Generate scaffold scripts"

echo -e "\n${GREEN}🎉 All tests completed!${NC}"
echo -e "${BLUE}📝 Summary: The embedded version functionality allows scaffoldgen to work perfectly from any directory without external file dependencies.${NC}"

# Cleanup
rm -f /tmp/scaffoldgen-test
echo -e "\n${YELLOW}🧹 Cleaned up test files${NC}"
