#!/bin/bash

docker build -t harvest-db-sync . && docker run -it --rm harvest-db-sync

#docker run -it --rm harvest-db-sync /bin/ash