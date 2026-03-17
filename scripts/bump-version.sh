#!/bin/bash

set -e

# Get current version
if [ ! -f "version.json" ]; then
    echo "Error: version.json not found"
    exit 1
fi

CURRENT_VERSION=$(jq -r '.version' version.json)
if [ "$CURRENT_VERSION" = "null" ]; then
    echo "Error: Could not read version from version.json"
    exit 1
fi

# Parse version parts
IFS='.' read -ra PARTS <<< "$CURRENT_VERSION"
MAJOR=${PARTS[0]}
MINOR=${PARTS[1]}
PATCH=${PARTS[2]}

# Determine bump type
BUMP_TYPE=${1:-patch}

case $BUMP_TYPE in
    major)
        MAJOR=$((MAJOR + 1))
        MINOR=0
        PATCH=0
        ;;
    minor)
        MINOR=$((MINOR + 1))
        PATCH=0
        ;;
    patch)
        PATCH=$((PATCH + 1))
        ;;
    *)
        echo "Usage: $0 [major|minor|patch]"
        exit 1
        ;;
esac

NEW_VERSION="$MAJOR.$MINOR.$PATCH"

echo "Bumping version: $CURRENT_VERSION -> $NEW_VERSION"

# Update version
go run scripts/update-version.go "$NEW_VERSION"

# Git commit and tag
echo "Creating git commit and tag..."
git add version.json internal/app/version.go
git commit -m "chore: bump version to $NEW_VERSION"
git tag -a "v$NEW_VERSION" -m "Release version $NEW_VERSION"

echo "✅ Version bumped to $NEW_VERSION"
echo "🏷️  Tag created: v$NEW_VERSION"
echo "📝 To push: git push && git push --tags"
