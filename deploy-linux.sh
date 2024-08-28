#!/bin/bash

TARGET_DIR="/opt/jcloud"
CONFIG_DIR="/opt/jcloud/config"
ARCHIVE_PATH="bin/jcloud.tar.gz"
CONFIG_PATH="config/client.yaml"

sudo mkdir -p $TARGET_DIR
sudo mkdir -p $CONFIG_DIR

sudo mv $ARCHIVE_PATH $TARGET_DIR/
sudo mv $CONFIG_PATH $CONFIG_DIR

cd $TARGET_DIR
sudo tar -xzvf jcloud.tar.gz

sudo rm jcloud.tar.gz

if ! grep -q "$TARGET_DIR" <<< "$PATH"; then
    echo "export PATH=\$PATH:$TARGET_DIR" | sudo tee -a /etc/profile
    source /etc/profile
fi

echo "export CLIENT_CONFIG_PATH="/opt/jcloud/config/client.yaml"" >> ~/.bashrc && source ~/.bashrc


echo "Install successfully completed, jcloud installed in $TARGET_DIR"
