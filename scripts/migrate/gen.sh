#!/bin/bash


PROJECT_ROOT=$(pwd)
SCRIPT_DIR=$(dirname $0)

cd $SCRIPT_DIR

mysql_DSN="mysql://schemaxuser:123456@127.0.0.1:13306/schemax?loc=Local&parseTime=true&charset=utf8mb4"
postgres_DSN="postgres://schemaxuser:123456@127.0.0.1:15432/schemax?sslmode=disable"
sqlite3_DSN="sqlite://./schemax.db"

DATABASES=("mysql" "postgres" "sqlite3")

for TYPE in ${DATABASES[@]}; do
    DSN_NAME="${TYPE}_DSN"
    DB=${!DSN_NAME}
    echo "Generating $TYPE schema for $DB"
    DIR="$PROJECT_ROOT/xo/$TYPE"
    mkdir -p $DIR
    rm -f $DIR/*.xo.*
    xo schema $DB -o $DIR
done