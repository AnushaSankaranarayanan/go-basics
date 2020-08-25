#! /bin/sh
echo "\n Removing docker container for MYSQL\n"
docker stop some-mysql
docker rm some-mysql
