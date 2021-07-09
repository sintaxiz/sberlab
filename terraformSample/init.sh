#!/bin/bash

# apt update
# apt upgrade -y
sudo apt-get update
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
apt install -y docker.io
pull sintaxiz/tf-frontend
pull sintaxiz/tf-backend
docker run -p 80:80 -e API_BASE_URL=http://45.9.24.240:8080/products/ -d sintaxiz/tf-frontend
docker run -p 8080:80 -d sintaxiz/tf-backend