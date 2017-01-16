#!/usr/bin/env bash

# Ensure the MySQL container is started.
echo "Starting the MySQL server container..."
docker-compose up -d mysql

# Download the SQL data from the repo.
# Use a docker image containing curl.
echo "Downloading the Democracy Unchained SQL data..."
docker run --rm tutum/curl curl https://raw.githubusercontent.com/DemocracyUnchained/dataset/master/democracy.sql > democracy.sql

# Attempt the SQL data load until it succeeds, with a limit on attempts.
echo -n "Waiting for MySQL to fully start..."
MAX_LOAD_ATTEMPTS=50
for ((i=1;i<=MAX_LOAD_ATTEMPTS;i++)); do
    result=$(docker exec -i democracy_mysql_1 mysql -uroot du_db < democracy.sql 2>&1)
    if [ $? -ne 0 ]; then
        echo -n "."
        sleep 1
    else
        echo
        echo "Successfully loaded SQL data into MySQL server."
        break
    fi
done

# Remove the downloaded SQL file.
rm democracy.sql

# Stop the MySQL container.
echo "Shutting down the MySQL server container..."
docker-compose stop mysql
