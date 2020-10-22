#!/bin/bash
#
# A script to run the DB for the test

exec sudo docker run -ti --rm \
    -e MYSQL_RANDOM_ROOT_PASSWORD=1 \
    -e MYSQL_PASSWORD=bench \
    -e MYSQL_USER=test \
    -e MYSQL_DATABASE=testbulk \
    -v "${PWD}"/init.sql:/docker-entrypoint-initdb.d/1-init.sql \
    --net host \
    mysql:8
