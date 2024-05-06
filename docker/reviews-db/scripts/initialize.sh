#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "$0")

mongod --fork --logpath /var/log/mongodb.log --dbpath /data/db/

FILES=$SCRIPT_DIR/*.js
for f in $FILES; do mongo localhost:27017/reviews $f; done
