#!/bin/bash

git config pull.rebase false
git config --global commit.template .gitmessage

# Check if an argument is provided
if [ -z "$1" ]; then
  echo "Error: No directory provided."
  exit 1
fi

containerWorkspaceFolder=$1

# Use quotes around the variable to handle spaces in the path
git config --global --add safe.directory "${containerWorkspaceFolder}"
