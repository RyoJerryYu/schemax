#!/bin/bash

set -e
set -o errexit
set -o pipefail

cd $(dirname "$0")

echo "Starting up the development environment..."

docker-compose up -d
