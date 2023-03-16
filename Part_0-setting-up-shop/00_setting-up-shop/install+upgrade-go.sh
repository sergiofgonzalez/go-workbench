#!/bin/bash -e

GO_VERSION="1.20.2"


mkdir -p ~/Downloads/go
rm -f ~/Downloads/go/go${GO_VERSION}.linux-amd64.tar.gz
wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz --directory-prefix=$(readlink -f ~/Downloads/go)
sudo rm -rf /usr/local/go

sudo tar -C /usr/local -xzf ~/Downloads/go/go${GO_VERSION}.linux-amd64.tar.gz

echo "Installed GO v${GO_VERSION}"
go version
