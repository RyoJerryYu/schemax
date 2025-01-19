#!/bin/bash
set -e

PROJECT_ROOT=$(pwd)
SCRIPT_DIR=$(dirname $0)

cd $SCRIPT_DIR

# env for selecting the configuration to apply
MIGRATE_ENV=${MIGRATE_ENV:-""}
# mode can be dry-run, auto-approve, apply
MIGRATE_MODE=${MIGRATE_MODE:-"dry-run"}

if [[ -z "$MIGRATE_ENV" ]]; then
    echo "MIGRATE_ENV is empty, please set MIGRATE_ENV"
    exit 1
fi

echo "Applying schema to $MIGRATE_ENV, Mode: $MIGRATE_MODE"

ATLAS_CHECK_FLAG=""
if [[ "$MIGRATE_MODE" == "dry-run" ]]; then
    ATLAS_CHECK_FLAG="--dry-run"
fi
if [[ "$MIGRATE_MODE" == "auto-approve" ]]; then
    ATLAS_CHECK_FLAG="--auto-approve"
fi

atlas schema apply \
    --env $MIGRATE_ENV \
    $ATLAS_CHECK_FLAG
