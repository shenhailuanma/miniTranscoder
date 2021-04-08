#!/bin/bash

# build frontend
pushd frontend
# npm run build
popd

# build backend
pushd backend
# GOOS=linux GOARCH=amd64 go build -mod=vendor -tags netgo -ldflags "-s -w" -o mtserver main.go
popd


# prepare dir
rm -rf miniTranscoder
mkdir -p ./miniTranscoder/bin
mkdir -p ./miniTranscoder/www
cp -rf ./backend/mtserver ./miniTranscoder/bin
chmod +x ./miniTranscoder/bin/mtserver

cp -rf ./frontend/dist ./miniTranscoder
mv ./miniTranscoder/dist ./miniTranscoder/www

# docker build
docker build -t shenhailuanma/minitranscoder:1.0.7 .

# run service
# docker run -p 9000:9000 -d shenhailuanma/minitranscoder:1.0.7

# push 
# docker push shenhailuanma/minitranscoder:1.0.7