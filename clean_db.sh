#!/bin/sh

docker rm -f sakila-mysql
docker volume rm -f docker_sakilastore
cd .docker && docker compose up -d

