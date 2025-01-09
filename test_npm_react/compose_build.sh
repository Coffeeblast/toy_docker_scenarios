#!/bin/bash
# log stdout and stderr to console as well as file
docker compose build --progress=plain $@ > >(tee ./logs/build.log) 2>&1
