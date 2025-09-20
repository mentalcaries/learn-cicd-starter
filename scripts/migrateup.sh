#!/bin/bash

if [ -f .env ]; then
    source .env
fi

cd sql/schema
goose up sqlite3 $DATABASE_URL
