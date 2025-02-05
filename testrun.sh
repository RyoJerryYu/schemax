#!/bin/bash

set -e
set -x
set -o pipefail

go install ./cmd/schemax
schemax -c ./cmd/schemax/config.yaml -v=3 -logtostderr -schema=schemax
# schemax -v=3 -logtostderr -dsn="mysql://schemaxuser:123456@127.0.0.1:13306/schemax?loc=Local&parseTime=true&charset=utf8mb4" > a.bin
