#!/bin/bash
#
# A workaround script for https://github.com/actions/upload-artifact/pull/133
set -euxo pipefail

readonly WORKSPACE_ROOT=$(git rev-parse --show-toplevel)
cd ${WORKSPACE_ROOT}
rm -rf ./artifacts
mkdir -p ./artifacts/{debs,test-results}

(
    cd bazel-testlogs/
    find . -iname test.xml -exec cp --parents \{\} ${WORKSPACE_ROOT}/artifacts/test-results/ \;
)
(
    cd bazel-bin/
    find . -name \*debian.deb -exec cp --parents \{\} ${WORKSPACE_ROOT}/artifacts/debs/ \;
)
