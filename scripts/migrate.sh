#!/bin/bash

# Get the project root directory (parent of scripts directory)
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

# Load environment variables from .env file in project root
if [ -f "${PROJECT_ROOT}/.env" ]; then
    source "${PROJECT_ROOT}/.env"
fi

# Construct the database URL from environment variables
DB_URL="postgres://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT:-5432}/${DB_NAME}?sslmode=disable"

# Default to "up" if no direction is specified
DIRECTION=${1:-up}

case $DIRECTION in
    up)
        echo "Running migrations up..."
        migrate -path "${PROJECT_ROOT}/misc/migrations" -database "${DB_URL}" up
        ;;
    down)
        echo "Running migrations down..."
        migrate -path "${PROJECT_ROOT}/misc/migrations" -database "${DB_URL}" down
        ;;
    *)
        echo "Invalid direction. Use 'up' or 'down'"
        exit 1
        ;;
esac