#!/bin/bash
docker build -t builder -f Dockerfile.build .
GOARCH=arm GOARM=7 GOOS=linux CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc ARCH=armhf \
docker run -ti --rm -e ARCH -e GOOS -e GOARCH -e GOARM -e CC -v "$PWD:/build" -w /build builder make $@