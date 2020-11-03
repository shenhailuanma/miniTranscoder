#!/bin/bash

# build frontend
pushd frontend
npm run build
popd

# build backend
pushd backend
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o mtserver main.go
popd


# prepare dir
mkdir -p ./miniTranscoder/bin
mkdir -p ./miniTranscoder/www
cp -rf ./backend/mtserver ./miniTranscoder/bin
chmod +x ./miniTranscoder/bin/mtserver

cp -rf ./frontend/dist ./miniTranscoder
mv ./miniTranscoder/dist ./miniTranscoder/www

# docker build
docker build -t minitranscoder:v1.0.0 .

# run 
# docker run -p 9000:9000 -d minitranscoder:v1.0.0