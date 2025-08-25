#!/bin/bash

TARGET_SDK_BRANCH=auth-settings

set -e

echo "Creating temporary directory..."
TMP_DIR=$(mktemp -d)

echo "Cloning cortex-cloud-go repository..."
git clone -b $TARGET_SDK_BRANCH https://github.com/mdboynton/cortex-cloud-go.git "$TMP_DIR"

#echo "Clearing vendor directory..."
#rm -rf vendor

echo "Creating vendor directory structure..."
mkdir -p vendor/github.com/mdboynton/cortex-cloud-go/internal

echo "Copying sub-packages..."
cp -r "$TMP_DIR"/api vendor/github.com/mdboynton/cortex-cloud-go/
cp -r "$TMP_DIR"/appsec vendor/github.com/mdboynton/cortex-cloud-go/
cp -r "$TMP_DIR"/cloudonboarding vendor/github.com/mdboynton/cortex-cloud-go/
cp -r "$TMP_DIR"/errors vendor/github.com/mdboynton/cortex-cloud-go/
cp -r "$TMP_DIR"/enums vendor/github.com/mdboynton/cortex-cloud-go/
cp -r "$TMP_DIR"/log vendor/github.com/mdboynton/cortex-cloud-go/
cp -r "$TMP_DIR"/platform vendor/github.com/mdboynton/cortex-cloud-go/
cp -r "$TMP_DIR"/internal/app vendor/github.com/mdboynton/cortex-cloud-go/internal/
cp -r "$TMP_DIR"/internal/util vendor/github.com/mdboynton/cortex-cloud-go/internal/

#echo "Running go mod vendor..."
#go mod vendor

echo "Cleaning up..."
rm -rf "$TMP_DIR"

echo "Done."
