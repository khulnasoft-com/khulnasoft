#!/usr/bin/env bash
# shellcheck disable=1091

set -euo pipefail

SCRIPT_PATH=$(realpath "$(dirname "$0")")

# shellcheck source=../lib/common.sh
source "$(realpath "${SCRIPT_PATH}/../lib/common.sh")"

import "ensure-gcloud-auth.sh"

blazedock run dev/preview:configure-workspace
ensure_gcloud_auth

if [[ "${VERSION:-}" == "" ]]; then
    VERSION="$(previewctl get name)-dev-$(date +%F_T%H-%M-%S)"
    log_info "VERSION is not set - using $VERSION"
    echo "$VERSION" > /tmp/local-dev-version
fi

blazedock build \
    -Dversion="${VERSION}" \
    --dont-test \
    dev/preview:deploy-dependencies
