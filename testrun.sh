#!/bin/bash

go install ./cmd/schemax
schemax -c ./cmd/schemax/config.yaml -v=3 -logtostderr
