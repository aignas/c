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
        die "Please run '$*' to cleanup files"
    fi
}

# shellcheck source=/dev/null
source .envrc

ensure buildifier
ensure gazelle
ensure mod-tidy
ensure raze
ensure rustfmt
bazel test //:verify-all
bazel build //src/...
bazel test //src/...
