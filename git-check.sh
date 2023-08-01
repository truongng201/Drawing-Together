#!/bin/bash

# Define the folder you want to monitor
FOLDER_TO_MONITOR="client"

# Go to the repository root directory
REPO_ROOT=$(git rev-parse --show-toplevel)
cd "$REPO_ROOT" || exit 1

# Get the current folder name
CURRENT_FOLDER_NAME=$(basename "$FOLDER_TO_MONITOR")

# Check for changes in the folder name
CHANGED=$(git diff --name-only -- "$FOLDER_TO_MONITOR")

if [[ -n "$CHANGED" ]]; then
  echo "Folder name has changed in the repository!"
  echo "Old folder name: $CURRENT_FOLDER_NAME"
  echo "New folder name: $(basename "$FOLDER_TO_MONITOR")"
else
  echo "No folder name changes found in the repository."
fi
