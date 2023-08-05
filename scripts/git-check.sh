#!/bin/bash

FOLDER_TO_MONITOR=$1

# Go to the repository root directory
# REPO_ROOT=$(git rev-parse --show-toplevel)
# cd "$REPO_ROOT" || exit 1

# Get the current folder name
CURRENT_FOLDER_NAME=$(basename "$FOLDER_TO_MONITOR")

  # Check for changes in the folder name
CHANGED=$(git checkout HEAD^ && git diff --name-only HEAD~ HEAD | grep "$CURRENT_FOLDER_NAME")

if [[ -n "$CHANGED" ]]; then
  echo "$FOLDER_TO_MONITOR changed"
  exit 0
else
  exit 0
fi