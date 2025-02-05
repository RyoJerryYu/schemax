#!/bin/bash

set -e
set -x
set -o pipefail

go install ./cmd/schemax
schemax -c ./cmd/schemax/config.yaml -v=3 -logtostderr -schema=schemax > a.bin
