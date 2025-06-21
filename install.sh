#!/bin/bash

set -e

echo "🔧 Installing EG Automation Engine..."

TMP_DIR="/tmp/eg-lm-automation"
TARGET_DIR="/opt/lm-automation"

echo "📥 Cloning from GitHub..."
git clone --depth 1 https://github.com/endgatesystemtechnologies/eg-lm-automation.git $TMP_DIR

echo "📂 Copying files to $TARGET_DIR..."
mkdir -p $TARGET_DIR
cp -r $TMP_DIR/internal $TMP_DIR/cmd $TMP_DIR/go.* $TARGET_DIR/

echo "✅ EG Automation Engine installed to $TARGET_DIR"
