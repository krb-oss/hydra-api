#!/bin/bash
set -o errexit
psql \
    -v ON_ERROR_STOP=1 \
    --user "${POSTGRES_USER}" \
    --dbname "${POSTGRESDB}" <<-EOSQL
BEGIN;
COMMIT;
EOSQL
