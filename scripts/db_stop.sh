#!/bin/bash

if [ "$(docker ps -q -f name=mongo-init-container)" ]; then
  docker rm -f mongo-init-container
fi
