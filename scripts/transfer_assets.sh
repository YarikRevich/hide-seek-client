#!/bin/bash

if [[ "$OSTYPE" == "darwin"* || "$OSTYPE" == "linux"* ]]; then
    mkdir -p /usr/local/share/games/HideSeek
    cp -r assets/* /usr/local/share/games/HideSeek
fi

