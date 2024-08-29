#!/bin/bash

# Переменные
TARGET_DIR="/opt"
BIN_DIR="/opt/jcloud/bin"
CONFIG_DIR="/opt/jcloud/config"
ARCHIVE_PATH="./jcloud.tar.gz"
CONFIG_PATH="config/client.yaml"
PROFILE_FILE="/etc/profile"

sudo mkdir -p $TARGET_DIR
sudo mkdir -p $CONFIG_DIR

sudo mv $ARCHIVE_PATH /opt/
sudo mv $CONFIG_PATH $CONFIG_DIR

cd /opt
sudo tar -xzvf jcloud.tar.gz -C $TARGET_DIR

sudo rm jcloud.tar.gz

sudo sh -c "echo 'export CLIENT_CONFIG_PATH=\"$CONFIG_DIR/client.yaml\"' >> $PROFILE_FILE"
sudo sh -c "echo 'export PATH=\$PATH:$BIN_DIR' >> $PROFILE_FILE"

source $PROFILE_FILE

echo "Environment variables updated in $PROFILE_FILE"
echo "Install successfully completed, jcloud installed in $TARGET_DIR/jcloud"
