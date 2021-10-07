#!/bin/bash

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

pwd

TGT=$1
[ "$TGT" ] || TGT="v0.0.0"

echo "packaging modules..."
mkdir -p build/dist/module

function z {
  echo "updating [$1] module"
  cd $1
  zip -r "../../build/dist/projectforge_module_$1.zip" .
  cd ..
}

cd module
for d in * ; do
  if [ -d "$d" ]; then
    z "$d"
  fi
done
