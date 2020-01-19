#!/usr/bin/env bash

env GOOS=linux GOARCH=amd64 go build
scp resto-be root@156.67.214.228:/root/resto-be