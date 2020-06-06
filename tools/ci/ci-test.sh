#!/bin/bash
set -eu

readonly WORKSPACE_ROOT=$(git rev-parse --show-toplevel)

gazelle
if [ ! -z "$(git status --porcelain "${WORKSPACE_ROOT}"/**/BUILD.bazel)" ]; then
    echo "Please run 'gazelle' to cleanup the build files"
    exit 1
fi
