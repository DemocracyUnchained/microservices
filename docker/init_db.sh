#!/usr/bin/env bash

# Ensure the MySQL container is started.
docker-compose up -d mysql

# Download the SQL data from the repo.
# Use a docker image containing curl.
docker run --rm tutum/curl curl https://raw.githubusercontent.com/DemocracyUnchained/dataset/master/democracy.sql > democracy.sql

# Create the databases and users
docker exec -i docker_mysql_1 mysql -uroot -h127.0.0.1 -P3306 < du_db.sql
docker exec -i docker_mysql_1 mysql -uroot -h127.0.0.1 -P3306 du_db < democracy.sql

# Remove the downloaded SQL file.
rm democracy.sql

# Stop the MySQL container.
docker-compose stop mysql
