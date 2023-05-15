#!/bin/bash -e

if [ $# -ne 2 ]; then
  echo "Wrong number of arguments"
  echo "Usage: $0 [prj_name] [module_name]"
  exit 1
fi

mkdir -p $1/$2-v0
cd $1 && echo "# $2 " >> README.md && cd $2-v0 && go mod init example.com/$2 && touch $2.go && code . && cd ../..
