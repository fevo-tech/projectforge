#!/bin/bash
# Code generated by projectforge.dev using code from the [notarize] module, which is under license [CC0]

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

if [ "$PUBLISH_TEST" != "true" ]
then
  time gon ./tools/notarize/gon.amd64.hcl
fi
