#!/bin/sh
set -e

DOCKER_SOCKET=/var/run/docker.sock
DOCKER_GROUP=docker
BUILD_USER=goreleaser



if [ -S ${DOCKER_SOCKET} ]; then
	sudo chown ${BUILD_USER}:${DOCKER_GROUP} ${DOCKER_SOCKET}
fi
exec "$@"
