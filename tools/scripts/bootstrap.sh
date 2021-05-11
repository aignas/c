#!/bin/bash
#
# Depends on:
# - curl
# - cargo
set -euo pipefail

readonly WORKSPACE_ROOT=$(git rev-parse --show-toplevel)
cd "${WORKSPACE_ROOT}"

bazelisk() {
    local -r url="https://github.com/bazelbuild/bazelisk/releases/download"
    local -r version="$1"
    local -r name="bazelisk-$2"

    [[ -f ./tools/bin/bazel ]] && return

    curl -L "$url/v$version/$name" \
        --output ./tools/bin/bazel
    chmod +x ./tools/bin/bazel
}

main() {
    mkdir -p ./tools/bin
    if [ "$(uname -s)" == "Darwin" ]; then
        bazelisk 1.8.1 darwin-amd64
    else
        bazelisk 1.8.1 linux-amd64
    fi
}

main
