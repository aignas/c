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

# shellcheck source=/dev/null
source .envrc

ensure buildifier
ensure gazelle
ensure gazelle update-repos -from_file=src/go.mod -prune
bazel test //:verify-all
bazel test //src/...
