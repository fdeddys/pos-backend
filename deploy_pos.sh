#!/usr/bin/env bash

env GOOS=linux GOARCH=amd64  go build -o pos-be

# mv resto-be pos-be

# sudo scp -i ~/Documents/project/dev\ aws/aws-lightsail.pem pos-be ubuntu@54.251.137.12:/home/ubuntu


scp -i ~/Documents/pos/aws/aws-lightsail.pem pos-be ubuntu@52.221.255.231:/home/ubuntu