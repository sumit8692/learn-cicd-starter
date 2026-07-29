#!/bin/bash
if [ -z "$DATABASE_URL" ] && [ -f .env ]; then
    source .env
fi
cd sql/schema
goose turso $DATABASE_URL up