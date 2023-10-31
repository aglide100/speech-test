#!/bin/bash

if ! command -v docker &> /dev/null; then
  echo "Docker is not installed..."
  exit 1
fi

docker stack deploy -c <(docker-compose -f server.yml config) speech