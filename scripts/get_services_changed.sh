#!/bin/bash

FOLDER_TO_MONITOR=$1

# Get the current folder name
CURRENT_FOLDER_NAME=$(basename "$FOLDER_TO_MONITOR")
SERVICES_CHANGED=()
  # Check for changes in the folder name
CHANGED=$(git diff --name-only HEAD~ HEAD | grep "$CURRENT_FOLDER_NAME")
if [[ -n "$CHANGED" ]]; then
  # read line by line of the changed files and check which services are affected 
  while read -r line; do
    # Get the service name
    SERVICE_NAME=$(echo "$line" | cut -d'/' -f1)
    # Check if the service name is not empty
    if [[ -n "$SERVICE_NAME" ]]; then
      # Check if the service name is not already in the array
      if [[ ! " ${SERVICES_CHANGED[@]} " =~ " ${SERVICE_NAME} " ]]; then
        # Add the service name to the array
        SERVICES_CHANGED+=("$SERVICE_NAME")
      fi
    fi
  done <<< "$CHANGED"
  echo ${SERVICES_CHANGED}
  exit 0
else
  exit 0
fi