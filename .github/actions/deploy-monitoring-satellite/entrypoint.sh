#!/usr/bin/env bash

set -euo pipefail

export HOME=/home/khulnasoft
# shellcheck disable=SC2155
export BLAZEDOCK_WORKSPACE_ROOT="$(pwd)"
export PATH="$PATH:$HOME/bin"

mkdir $HOME/bin

gcloud auth login --cred-file="$GOOGLE_APPLICATION_CREDENTIALS" --activate --quiet
blazedock run dev/preview/previewctl:install

echo "previewctl install-context"
previewctl install-context --log-level debug --timeout 10m

echo "blazedock run dev/preview:deploy-monitoring-satellite"
blazedock run dev/preview:deploy-monitoring-satellite

{
    echo '<p>Monitoring satellite has been installed in your preview environment.</p>'
    echo '<ul>'
    echo '<li><b>📚 Documentation</b> - See our <a href="https://www.notion.so/khulnasoft/f2938b2bcb0c4c8c99afe1d2b872380e" target="_blank">internal documentation</a> on how to use it.</li>'
    echo '</ul>'
} >> "${GITHUB_STEP_SUMMARY}"
