#!/bin/sh

# This script is used to build the documentation for the project.

# The first step is to ensure that the swag cli utility is installed. This is done by running the go install command, which will download and install the latest version of the swag cli utility.
# sudo GOBIN=/usr/local/bin/ go install github.com/swaggo/swag/cmd/swag@latest

cd app

# Generate the documentation
swag init --pd --parseDepth 3
