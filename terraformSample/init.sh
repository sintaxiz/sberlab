#!/bin/bash

# apt update
# apt upgrade -y
sudo apt-get update > log.txt
sudo apt-get install docker-ce docker-ce-cli containerd.io >> log.txt
apt install -y docker.io >> log.txt
pull sintaxiz/tf-frontend >> log.txt
pull sintaxiz/tf-backend >> log.txt
docker run -p 9999:80 -e API_BASE_URL=http://178.170.195.175/9990/validate/ -d sintaxiz/tf-frontend >> log.txt
# docker run -p 9999:80 -d sintaxiz/tf-backend >> log.txt