#!/bin/bash
# Code generated by Project Forge, see https://projectforge.dev for details.

## Runs all the tests

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

go test ./app/...
