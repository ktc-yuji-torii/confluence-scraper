#!/bin/bash

# 検索するDockerネットワークの名前
NETWORK_NAME="common-external-network"

# Dockerネットワークが存在するか確認
if ! docker network ls | grep -q "$NETWORK_NAME"; then
  echo "Network $NETWORK_NAME does not exist. Creating it..."
  docker network create $NETWORK_NAME
else
  echo "Network $NETWORK_NAME already exists."
fi
