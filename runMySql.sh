#! /bin/sh
echo "\n Starting docker container for MYSQL\n"
docker run -p 3306:3306 --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:latest 
