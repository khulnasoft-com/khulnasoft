#!/bin/bash
# Copyright (c) 2022 Khulnasoft GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License.AGPL.txt in the project root for license information.


COMPONENT_PATH="$(dirname "$0")/.."
echo "Component Path: ${COMPONENT_PATH}"

if [ "${BLAZEDOCK_BUILD-}" == "true" ]; then
    CONFIG_PATH="./_deps/components-khulnasoft-protocol--khulnasoft-schema/khulnasoft-schema.json"
else
    CONFIG_PATH="$COMPONENT_PATH/../data/khulnasoft-schema.json"
fi
echo "Config Path: ${CONFIG_PATH}"

KHULNASOFT_CONFIG_TYPE_PATH="$COMPONENT_PATH/khulnasoft-config-types.go"
echo "Config Types Path: ${KHULNASOFT_CONFIG_TYPE_PATH}"
if [ "${BLAZEDOCK_BUILD-}" == "true" ]; then
    git init -q
    git add "$KHULNASOFT_CONFIG_TYPE_PATH"
fi

go install github.com/a-h/generate/...@latest

schema-generate -p protocol "$CONFIG_PATH" > "$KHULNASOFT_CONFIG_TYPE_PATH"

# remove custom marshal logic to allow additional properties
sed -i '/func /,$d' "$KHULNASOFT_CONFIG_TYPE_PATH" #functions
sed -i '5,10d' "$KHULNASOFT_CONFIG_TYPE_PATH" #imports
# support yaml and json
sed -i -E 's/(json:)(".*")/yaml:\2 \1\2/g' "$KHULNASOFT_CONFIG_TYPE_PATH"
gofmt -w "$KHULNASOFT_CONFIG_TYPE_PATH"

if [ "${BLAZEDOCK_BUILD-}" == "true" ]; then
    ./_deps/dev-addlicense--app/addlicense "$KHULNASOFT_CONFIG_TYPE_PATH"
else
    blazedock run components:update-license-header
fi

git diff --exit-code "$KHULNASOFT_CONFIG_TYPE_PATH"
