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

raze() {
  local -r version="$1"

  [[ -f ./tools/bin/cargo-raze ]] && return

  cargo install cargo-raze --root=./tools --version="$1"
}

main() {
  mkdir -p ./tools/bin
  bazelisk 1.7.1 linux-amd64
  raze 0.5.0
}

main
