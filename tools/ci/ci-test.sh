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
ensure gazelle update-repos -from_file=src/go.mod -prune -to_macro=third_party/repositories.bzl%go_repositories
ensure raze
ensure rustfmt
bazel test //:verify-all
bazel build //src/...
bazel test //src/...
