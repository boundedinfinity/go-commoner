#!/bin/bash

function ensure_fish() {
    sudo apt-get update
    sudo apt install -y software-properties-common
    sudo apt-add-repository -y ppa:fish-shell/release-3
    sudo apt-get update
    sudo apt-get -y install fish
    sudo apt-get autoremove -y
    sudo apt-get clean -y
    sudo rm -rf /var/lib/apt/lists/*
}

function ensure_enumer() {
    go install github.com/boundedinfinity/enumer/cmd/enumer@v1.0.12
}

ensure_fish
ensure_enumer
