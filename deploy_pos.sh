#!/usr/bin/env bash

env GOOS=linux GOARCH=amd64 go build

mv resto-be pos

sudo scp -i ~/Documents/project/dev\ aws/aws-lightsail.pem pos ubuntu@54.251.137.12:/home/ubuntu