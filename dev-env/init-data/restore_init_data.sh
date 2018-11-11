#!/usr/bin/env bash

set -xe

PSQL_CONTAINER=$(docker ps --format '{{ .Names }}' | grep db)
docker cp init.dump $PSQL_CONTAINER:/
docker exec $PSQL_CONTAINER sh -c 'pg_restore -Upostgres -dpostgres init.dump'
