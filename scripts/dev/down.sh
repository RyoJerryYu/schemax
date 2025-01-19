#!/bin/bash

set -e
set -o errexit
set -o pipefail

cd $(dirname "$0")

echo "Shutting down the development environment..."

docker-compose down