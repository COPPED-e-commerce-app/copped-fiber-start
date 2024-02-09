#!/bin/sh

go mod init

go mod tidy

go version

echo "Done initializing modules"
