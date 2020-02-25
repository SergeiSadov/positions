#!/bin/bash -eux

docker-compose up -d domains-web domains-elasticsearch domains-kibana domains-prometheus

docker-compose build --no-cache domains-go_green domains-go_blue

docker-compose stop domains-go_green
docker-compose up -d --force-recreate domains-go_green

sleep 5

docker-compose stop domains-go_blue
docker-compose up -d --force-recreate domains-go_blue