#!/usr/bin/env bash

set -euo pipefail

# Fixes "MountWrapper Type cannot implement 'MountWrapper' as it has a non-exported method and is defined in a different package"
# See https://github.com/kubernetes/mount-utils/commit/a20fcfb15a701977d086330b47b7efad51eb608e for context.
sed -i '/type MockMountWrapper struct {/a \\tmount.Interface' pkg/utils/mount/mock_mountutils_unix.go
