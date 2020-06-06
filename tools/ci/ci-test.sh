#!/bin/bash
set -eu

die() {
    echo "$@"
    exit 1
}

ensure() {
    "$1"
    files="$(git status --porcelain)"
    if [ -n "$files" ]; then
        echo "Detected changes to:"
        echo "$files"
        die "Please run '$1' to cleanup the build files"
    fi
}

readonly WORKSPACE_ROOT=$(git rev-parse --show-toplevel)
cd "${WORKSPACE_ROOT}"

ensure tools/buildifier
ensure tools/gazelle
