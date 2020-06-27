#!/bin/bash
set -eu

die() {
    echo "$@"
    exit 1
}

ensure() {
    "$@"
    files="$(git status --porcelain)"
    if [ -n "$files" ]; then
        echo "Detected changes to:"
        echo "$files"
        die "Please run '$*' to cleanup the build files"
    fi
}

readonly WORKSPACE_ROOT=$(git rev-parse --show-toplevel)
cd "${WORKSPACE_ROOT}"

ensure tools/buildifier
ensure tools/gazelle
ensure tools/gazelle update-repos -from_file=src/go.mod -prune
bazel test //:verify-all
bazel test //src/...
