# Build miniTranscoder server bin for pack docker image

## 1. Prepare ENV

ubuntu:

```bash
# install tools
apt-get update
apt-get install -y gcc
apt-get install -y wget

# get golang package
wget https://studygolang.com/dl/golang/go1.15.3.linux-amd64.tar.gz
tar -xf go1.15.3.linux-amd64.tar.gz
mv go /usr/local/

export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
c

```

Centos:

```bash
# install tools
yum install -y gcc
yum install -y wget

# get golang package
wget https://studygolang.com/dl/golang/go1.15.3.linux-amd64.tar.gz
tar -xf go1.15.3.linux-amd64.tar.gz
mv go /usr/local/

export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:/usr/local/go/bin

```

alpine:

```bash
# change mirror
echo "http://mirrors.ustc.edu.cn/alpine/v3.13/main/" > /etc/apk/repositories

# install tools
apk add gcc

# get golang package
wget https://studygolang.com/dl/golang/go1.15.3.linux-amd64.tar.gz
tar -xf go1.15.3.linux-amd64.tar.gz
mv go /usr/local/

export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:/usr/local/go/bin
```

## 2. Build

```bash

go mod vendor
GOOS=linux GOARCH=amd64 go build -mod=vendor -tags netgo -ldflags "-s -w" -o mtserver main.go

```



```sh
docker buildx build --push --platform linux/amd64,linux/arm64 -t shenhailuanma/minitranscoder:1.0.11 .
```