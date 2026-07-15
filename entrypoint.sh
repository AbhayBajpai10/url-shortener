#!/bin/sh

set -e

echo "Waiting for PostgreSQL..."

sleep 10

echo "Running migrations..."

goose -dir ./migrations postgres \
"postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" \
up

echo "Starting server..."

./main