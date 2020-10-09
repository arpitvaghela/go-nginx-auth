#!/usr/bin/env bash

if [ -x "$(command -v docker-compose)" ] && [ -x "$(command -v docker)" ]; then
    docker-compose down
    docker-compose build --parallel
    printf "\nDeleting old builds\n"
    docker image prune -f
else
    echo "Please verify you have docker-compose and docker installed in PATH"
    exit 1
fi
