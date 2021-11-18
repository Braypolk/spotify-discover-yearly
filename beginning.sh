#!/bin/bash

if ! command -v go &> /dev/null
then
    echo "Golang not found, download here: https://golang.org/doc/install"
    echo "Please install go, then run this script agian to ensure dependencies can be called"
    exit 1
fi

if ! command -v node &> /dev/null
then
    echo "Node not found, download here: https://nodejs.org/en/download/"
    echo "Please install npm/node, then run this script agian to ensure dependencies can be called"
    exit 1
elif ! command -v npm &> /dev/null
then
    echo "NPM not found, your node install may be broken. Download node here: https://nodejs.org/en/download/"
    echo "Please fix the npm/node issue then run this script agian to ensure dependencies can be called"
    exit 1
fi

echo "Create Spotify Dev Account if you don't already have one: https://developer.spotify.com/"
echo "Create AWS Account here if you want: https://aws.amazon.com/free/"
read -p "Press enter to continue"


# to avoid annoying github sync garbage
echo "Checking env files"
FILE=go/env.go
if [ ! -f "$FILE" ]; then
    echo "Creating go env file"
    touch go/env.go

    echo 'package main

var ClientId string = "your_spotify_client_id"
var ClientSecret string = "your_spotify_client_secret"
var RefreshToken string = "your_spotify_refresh_token"' >> go/env.go

fi

FILE=js/env.js
if [ ! -f "$FILE" ]; then
    echo "Creating js env file"
    touch js/env.js
    echo "exports.client_id = 'your_spotify_client_id';
exports.client_secret = 'your_spotify_client_secret';
exports.refresh_token = 'your_spotify_client_secret';" >> js/env.js

fi

cd js
npm install
cd ../
