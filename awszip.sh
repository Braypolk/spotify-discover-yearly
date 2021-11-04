#!/bin/bash

cd go
env GOOS=linux GOARCH=arm64 go build .
mv spotify-discover-yearly ../aws/spotify-discover-yearly
cd ../aws
chmod -R 777 spotify-discover-yearly
chmod -R 777 lambda_function.py
zip awsSpotify.zip spotify-discover-yearly lambda_function.py