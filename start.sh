#!/bin/sh

set -e

echo "Run DB migrations"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "Start the application"
exec "$@"