#!/bin/bash

set -e
set -x
set -o pipefail

go install ./cmd/schemax-gen-debug
schemax-gen-debug -v=3 -logtostderr < a.bin