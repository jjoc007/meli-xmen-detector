#!/bin/bash

pwd
CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o $(pwd)/build/xmen-dna-process/main xmen/functions/dna/process/main.go
CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o $(pwd)/build/xmen-stats-get/main xmen/functions/stats/get/main.go

build-lambda-zip -o $(pwd)/build/xmen-dna-process/main.zip $(pwd)/build/xmen-dna-process/main
build-lambda-zip -o $(pwd)/build/xmen-stats-get/main.zip $(pwd)/build/xmen-stats-get/main