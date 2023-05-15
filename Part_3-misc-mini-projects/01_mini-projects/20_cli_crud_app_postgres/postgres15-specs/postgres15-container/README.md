# Docker Container for PostgreSQL Database Server
> containerized PostgreSQL 15.2 with volume mapping to enable interacting with data from outside the container and foster stateless containers

This is the working directory for a container running **PostgreSQL 15.2** database server official image. It is intended for general purpose development and does not include any predefined database/schema.

The host directory `./pgdata` is intended to be mapped to the `/var/lib/postgresql/data` so that the container is completely stateless.

## Hints and Tips

If you want to create the container with an empty database run the following command from the containers working directory:
```bash
$ pwd
[...]/postgres15-container
$ rm -rf ./pgdata/*
```

## Build Instructions
First, make sure that you're on the container working dir (that is, you should be in the same directory where the `Dockerfile` is defined).

Then type:
```bash
docker build -t="sergiofgonzalez/postgres:v15.2-0.1.0" .
```

## Run Instructions

First, make sure that you're on the container working dir (that is, you should be in the same directory where the `Dockerfile` is defined). Otherwise, you will need to adapt the paths of the mapped volumes.
```bash
docker run -d -p 5432:5432 \
  -v $PWD/pgdata:/var/lib/postgresql/data \
  --name postgres-container \
  sergiofgonzalez/postgres:v15.2-0.1.0
```

To open an interactive session within the container (for inspection):
```bash
docker exec -i -t postgres-container /bin/bash
```