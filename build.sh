#!/bin/sh

targetARCH=amd64
cmd=auth-server
repository=rhzx3519/auth-server
workdir=build

isdocker=${1:-true}
arch=${2:-arm64}
os=${3:-linux}

GOOS=$os GOARCH=$targetARCH go build -o bin/$cmd main/main.go

cp .env $workdir
cp ./bin/$cmd $workdir
cp Dockerfile $workdir
if [ $isdocker ]; then
  cp .env-docker $workdir/.env
fi
cd $workdir

docker build --platform=$os/$arch  -t ${repository}:latest .
