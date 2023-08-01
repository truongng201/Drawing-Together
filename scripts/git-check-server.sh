#!/bin/bash

# Define the folder you want to monitor
FOLDER_TO_MONITOR="../server"

# Go to the repository root directory
REPO_ROOT=$(git rev-parse --show-toplevel)
cd "$REPO_ROOT" || exit 1

# Get the current folder name
CURRENT_FOLDER_NAME=$(basename "$FOLDER_TO_MONITOR")

# Check for changes in the folder name
CHANGED=$(git diff --name-only HEAD^ HEAD | grep "$CURRENT_FOLDER_NAME")

if [[ -n "$CHANGED" ]]; then
  echo "Folder $CURRENT_FOLDER_NAME has changed in the repository!"
  exit 1
else
  echo "Folder $CURRENT_FOLDER_NAME hasn't changed in the repository!"
  exit 0
fi
