#!/bin/bash

export $(cat .env) > /dev/null 2>&1; docker stack deploy -c server.yml speech