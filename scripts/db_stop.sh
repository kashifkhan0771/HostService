#!/bin/bash
if [ "$(docker ps -q -f name=host-service-mysql-db)" ]; then
  docker rm -f host-service-mysql-db
fi

if [ "$(docker ps -q -f name=host-service-mongo-db)" ]; then
  docker rm -f host-service-mongo-db
fi
