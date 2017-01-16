# Why Docker?

Docker is a containerization technology that wraps software components in a standardized environment. For more information, visit:

https://www.docker.com/what-docker

For the Democracy Unchained project, Docker offers these advantages:

* No native install of Golang or MySQL needed.
* No snowflake servers.
  * Same standard development environment for all developers.
  * Same versions of all dependencies.
* Easy to start hacking on the project:
  * Install Docker
  * Clone the repo
  * Issue a few `make` commands.
  * Start hacking.
* Exact same Docker containers can be used in development and production.

# Docker components

This `docker` directory contains a sub-directory per Docker container:
* `mysql` (the MySQL instance)
* `DemocracyUnchained` (the Go application)

The contents of each sub-directory:

* a `Dockerfile`, which holds all the instructions on how to build the container.
* any other files to be copied into the container's filesystem, such as configuration data or scripts

Docker containers are meant to be upgraded via deletion and re-creation. To persist data across containers, you need to mount disk locations as volumes. There's another sub-directory called `.dev`, which holds the MySQL data directory that's mounted as a volume to the MySQL container.

# Docker Compose

Docker Compose is a tool for defining and running multi-container Docker applications. More info here:

https://docs.docker.com/compose/overview/

This directory contains a `docker-compose.yml` file which defines the two containers - `mysql` and `web` - which compose the microservice. Compose allows the entire service to be started with a single command and also handles the creation of a network for the components to use in communicating with one another.

# Docker Resources

How to get started with Docker:

* https://docs.docker.com/engine/getstarted/

A good Docker beginning tutorial:

* https://docs.docker.com/engine/tutorials/dockerizing/

The main Docker docs:

* https://docs.docker.com/

Docker Compose:

* https://docs.docker.com/compose/overview/

Networking in Compose:

* https://docs.docker.com/compose/networking/

# Example: Adding redis To the Democracy Unchained Service

Suppose you wanted to add redis as another component of the Democracy Unchained service. First, you'd create a `redis` subdirectory and add a simple Dockerfile, such as follows:

```
FROM        ubuntu:16.04
RUN         apt-get update && apt-get install -y redis-server
EXPOSE      6379
ENTRYPOINT  ["/usr/bin/redis-server"]
```

Then, under the "services:" section of the `docker-compose.yml` file, you'd add the following minimal section:

```yaml
redis:
  build: redis
  ports: 9000:6379
```

This section instructs Docker Compose to start up a container based on the Dockerfile in the `redis` subdirectory - and to expose the container's port 6379 as port 9000 on the local host. Once Compose starts up the container via `docker-compose up`, the components in the other containers started by Compose can now access the redis server by using the host:port of "redis:6379". Also, you can reach the redis server locally by using the host:port of "127.0.0.1:9000".

For more details:

* https://docs.docker.com/engine/examples/running_redis_service/
