#!/bin/bash

# Exit on error
set -e

# Define output directory
OUTPUT_DIR="release"

# Clean up previous builds
rm -rf $OUTPUT_DIR
mkdir -p $OUTPUT_DIR

# Define platforms
PLATFORMS=(
  "darwin arm64"
  "linux amd64"
  "windows amd64"
)

# Define binary name
BINARY_NAME="tflint-ruleset-hardcode"

# Build binaries for each platform
for PLATFORM in "${PLATFORMS[@]}"; do
  OS=$(echo $PLATFORM | awk '{print $1}')
  ARCH=$(echo $PLATFORM | awk '{print $2}')

  OUTPUT_BINARY="$OUTPUT_DIR/${BINARY_NAME}_${OS}_${ARCH}"

  if [ "$OS" == "windows" ]; then
    OUTPUT_BINARY+=".exe"
  fi

  echo "Building for $OS $ARCH..."
  GOOS=$OS GOARCH=$ARCH go build -o $OUTPUT_BINARY

  # Zip the binary
  echo "Zipping $OUTPUT_BINARY..."
  zip -j "$OUTPUT_BINARY.zip" "$OUTPUT_BINARY"
  rm "$OUTPUT_BINARY"
done

# Generate checksums.txt
CHECKSUMS_FILE="checksums.txt"
echo "Generating checksums.txt..."
cd $OUTPUT_DIR
shasum -a 256 *.zip > $CHECKSUMS_FILE

# Output success message
echo "Release artifacts created in $OUTPUT_DIR:" 
cd ..
ls -1 $OUTPUT_DIR

