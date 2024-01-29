#!/bin/sh

targetOS=linux
targetARCH=amd64
cmd=auth-server
repository=rhzx3519/auth-server
workdir=build
isdocker=$1

GOOS=$targetOS GOARCH=$targetARCH go build -o bin/$cmd main/main.go

cp .env $workdir
cp ./bin/$cmd $workdir
cp Dockerfile $workdir
if [ $isdocker ]; then
  cp .env-docker $workdir/.env
fi
cd $workdir

docker build -t ${repository}:latest .