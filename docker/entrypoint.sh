#!/bin/sh

echo "SQL_ADDR:"
echo $SQL_ADDR
exec /usr/local/bin/wait-for-it $SQL_ADDR -- \
    "$@"