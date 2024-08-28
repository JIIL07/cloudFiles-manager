#!/bin/bash

TARGET_DIR="/opt/jcloud"
ARCHIVE_PATH="bin/jcloud.tar.gz"
CONFIG_PATH="config/client.yaml"

sudo mkdir -p $TARGET_DIR

sudo mv $ARCHIVE_PATH $TARGET_DIR/
sudo mv $CONFIG_PATH $TARGET_DIR/

cd $TARGET_DIR
sudo tar -xzvf jcloud.tar.gz

sudo rm jcloud.tar.gz

if ! grep -q "$TARGET_DIR" <<< "$PATH"; then
    echo "export PATH=\$PATH:$TARGET_DIR" | sudo tee -a /etc/profile
    source /etc/profile
fi

echo "Install successfully completed, jcloud installed in $TARGET_DIR"
