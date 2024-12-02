#!/bin/bash

# build frontend
# pushd frontend
# npm install --froce
# npm run build
# popd

# build backend
# pushd backend
# GOOS=linux GOARCH=amd64 go build -mod=vendor -tags netgo -ldflags "-s -w" -o mtserver-amd64 main.go
# GOOS=linux GOARCH=arm64 go build -mod=vendor -tags netgo -ldflags "-s -w" -o mtserver-arm64 main.go
# popd


# prepare build dir
rm -rf .build
mkdir -p .build/amd64/minitranscoder/bin
mkdir -p .build/amd64/minitranscoder/ui
mkdir -p .build/arm64/minitranscoder/bin
mkdir -p .build/arm64/minitranscoder/ui


# copy files
cp -rf ./backend/mtserver-amd64 .build/amd64/minitranscoder/bin/mtserver
chmod +x .build/amd64/minitranscoder/bin/mtserver

cp -rf ./backend/mtserver-arm64 .build/arm64/minitranscoder/bin/mtserver
chmod +x .build/arm64/minitranscoder/bin/mtserver

cp -rf ./frontend/dist/* .build/amd64/minitranscoder/ui/
cp -rf ./frontend/dist/* .build/arm64/minitranscoder/ui/

# docker build
# docker buildx build --platform linux/amd64,linux/arm64 -t shenhailuanma/minitranscoder:1.0.11 . --push

