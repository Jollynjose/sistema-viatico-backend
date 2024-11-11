PG_FILE=${1}

if [[ -z "$PG_FILE" ]]; then
    echo "Please specify a PostgreSQL file"
    exit 1
fi

CONTAINER_ID=$(docker ps --format='{{.ID}}' --filter name=^/sistema-viatico-backend$ --filter name=^/postgres$)

docker cp ${PG_FILE} ${CONTAINER_ID}:/dump.sql

docker exec -it ${CONTAINER_ID} pg_restore -U postgres -d universidad /dump.sql