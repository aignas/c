#!/bin/bash
set -euo pipefail

readonly WORKSPACE_ROOT=$(git rev-parse --show-toplevel)
cd "${WORKSPACE_ROOT}"

bazelisk() {
    local -r url="https://github.com/bazelbuild/bazelisk/releases/download"
    local -r version="$1"
    local -r name="bazelisk-$2"

    rm -f ./tools/bin/bazel
    curl -L "$url/$version/$name" \
        --output ./tools/bin/bazel
    chmod +x ./tools/bin/bazel
}

main() {
    mkdir -p ./tools/bin
    bazelisk v1.5.0 linux-amd64
}

main
