#!/bin/bash -e

if [ $# -ne 2 ]; then
  echo "Wrong number of arguments"
  echo "Usage: $0 [prj_name] [module_name]"
  exit 1
fi

mkdir -p $1/$2
cd $1/$2 && go mod init example.com/$1 && touch $2.go && code . && cd ../..

