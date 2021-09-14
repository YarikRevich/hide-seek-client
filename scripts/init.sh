#!/bin/bash

# echo "root='$PWD/..'" > /usr/local/etc/HideSeek_system.toml

copy_assets(){
    if [[ "$OSTYPE" == "darwin"* || "$OSTYPE" == "linux"* ]]; then
        mkdir -p /usr/local/share/games/HideSeek
        cp -r $(PWD)/../assets/* /usr/local/share/games/HideSeek
    fi
}

copy_assets

