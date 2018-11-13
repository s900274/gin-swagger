#!/bin/bash

mkdir -p ./db
mkdir -p ./bin
mkdir -p ./log

cd src/httpservice/

./swag i -g httpserver.go

cd ../../

make

./bin/gin-test -config ./conf/gintest.$1.toml