#!/bin/bash

set -e
set -o pipefail

NODE_IMAGE='forta-protocol/forta-node:latest'

docker build -t "$NODE_IMAGE" -f Dockerfile.node .

go build -o forta .
