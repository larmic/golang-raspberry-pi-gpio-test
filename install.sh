#!/usr/bin/env zsh

declare -r BINARY_NAME="gpio-test"
declare -r RASPI_SCP_FOLDER="pi@raspberrypi.local:/home/pi/test"

# CGO_ENABLED=0   -> Disable interoperate with C libraries -> speed up build time! Enable it, if dependencies use C libraries!
# GOOS=linux      -> compile to linux because scratch docker file is linux
# GOARCH=amd64    -> because, hmm, everthing works fine with 64 bit :)
# -a              -> force rebuilding of packages that are already up-to-date.
# -o gpio-test    -> force to build an executable app file (instead of default https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies)

echo "Building arm binary '$BINARY_NAME-arm'..."
env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -a -o $BINARY_NAME-arm

echo "Building amd64 binary '$BINARY_NAME-amd64'..."
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o $BINARY_NAME-amd64

echo "Copy arm binary '$BINARY_NAME-arm' to $RASPI_SCP_FOLDER"
scp $BINARY_NAME-arm $RASPI_SCP_FOLDER