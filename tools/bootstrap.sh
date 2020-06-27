#!/bin/bash
# Download bazelisk into tools
set -euo pipefail

readonly WORKSPACE_ROOT=$(git rev-parse --show-toplevel)
readonly VERSION=v1.5.0
readonly BAZEL=bazelisk-linux-amd64
cd "${WORKSPACE_ROOT}"

rm -f ./tools/bazel
curl -L "https://github.com/bazelbuild/bazelisk/releases/download/$VERSION/$BAZEL" \
    --output ./tools/bazel
chmod +x ./tools/bazel
