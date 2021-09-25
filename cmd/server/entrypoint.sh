#!/bin/sh

exec /usr/local/bin/wait-for-it $DB_ADDR -- \
    "$@"