#!/bin/sh

# pubsubplus-opentelemetry-go-integration
#
# Copyright 2024 Solace Corporation. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Generates the SEMPv2 action, config and monitor Go clients from the specs in
# ./spec using swagger-codegen running in Docker.
#
# Inputs and outputs travel through a named Docker volume and `docker cp`
# rather than bind mounts of $PWD. A bind mount is resolved on the daemon
# host, so it silently breaks when DOCKER_HOST points at a remote daemon
# (as it does for the Mac legs of the Jenkins pipeline, which run against an
# Otterbox docker-server over WireGuard). Volumes and `docker cp` work the
# same way for local and remote daemons.

set -e

SWAGGER_VER="3.0.27"
IMAGE_TAG="solace-semp-swagger-codegen-cli:${SWAGGER_VER}"
VOLUME_NAME="semp-generation-vol-$$"
SETUP_CONTAINER="semp-generation-setup-$$"
OUTPUT_CONTAINER="semp-generation-output-$$"
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

cleanup() {
    docker rm -f "$SETUP_CONTAINER" "$OUTPUT_CONTAINER" >/dev/null 2>&1 || true
    docker volume rm "$VOLUME_NAME" >/dev/null 2>&1 || true
}
trap cleanup EXIT

echo "Building ${IMAGE_TAG}..."
docker build -f "$SCRIPT_DIR/Dockerfile" -t "$IMAGE_TAG" --build-arg "SWAGGER_VER=${SWAGGER_VER}" "$SCRIPT_DIR"

echo "Creating volume ${VOLUME_NAME} and copying specs into it..."
docker volume create "$VOLUME_NAME" >/dev/null
docker run --name "$SETUP_CONTAINER" -v "$VOLUME_NAME:/workspace" alpine:latest true
docker cp "$SCRIPT_DIR/spec" "$SETUP_CONTAINER:/workspace/"
docker rm "$SETUP_CONTAINER" >/dev/null

generate() {
    api="$1"
    echo "Generating ${api} client..."
    docker run --rm -v "$VOLUME_NAME:/workspace" -w /workspace \
        --entrypoint java "$IMAGE_TAG" \
        -jar /opt/swagger-codegen-cli/swagger-codegen-cli.jar \
        generate -l go -i "/workspace/spec/spec_${api}.json" -o "/workspace/${api}" \
        --type-mappings 'boolean=*bool' --additional-properties "packageName=${api}"
}

generate action
generate config
generate monitor

echo "Copying generated clients back to ${SCRIPT_DIR}..."
docker run --name "$OUTPUT_CONTAINER" -v "$VOLUME_NAME:/workspace" alpine:latest true
for api in action config monitor; do
    docker cp "$OUTPUT_CONTAINER:/workspace/${api}" "$SCRIPT_DIR/"
done
docker rm "$OUTPUT_CONTAINER" >/dev/null

echo "SEMPv2 client generation complete."
