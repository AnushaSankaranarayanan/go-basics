#! /bin/sh
echo "\n Executing into docker container for MYSQL\n"
#docker exec -it some-mysql bash
docker exec -it some-mysql mysql -uroot -pmy-secret-pw