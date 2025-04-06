# Copyright (c) 2020 Khulnasoft GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License.AGPL.txt in the project root for license information.
FROM node:18 as builder

ARG CODE_EXTENSION_COMMIT

RUN apt update -y \
    && apt install jq --no-install-recommends -y

RUN mkdir /khulnasoft-code-web \
    && cd /khulnasoft-code-web \
    && git init \
    && git remote add origin https://github.com/khulnasoft-com/khulnasoft-code \
    && git fetch origin $CODE_EXTENSION_COMMIT --depth=1 \
    && git reset --hard FETCH_HEAD
WORKDIR /khulnasoft-code-web
RUN yarn --frozen-lockfile --network-timeout 180000

# update package.json
RUN cd khulnasoft-web && \
    setSegmentKey="setpath([\"segmentKey\"]; \"untrusted-dummy-key\")" && \
    jqCommands="${setSegmentKey}" && \
    cat package.json | jq "${jqCommands}" > package.json.tmp && \
    mv package.json.tmp package.json
RUN yarn build:khulnasoft-web && yarn --cwd khulnasoft-web/ inject-commit-hash


FROM scratch

COPY --from=builder --chown=33333:33333 /khulnasoft-code-web/khulnasoft-web/out /ide/extensions/khulnasoft-web/out/
COPY --from=builder --chown=33333:33333 /khulnasoft-code-web/khulnasoft-web/public /ide/extensions/khulnasoft-web/public/
COPY --from=builder --chown=33333:33333 /khulnasoft-code-web/khulnasoft-web/resources /ide/extensions/khulnasoft-web/resources/
COPY --from=builder --chown=33333:33333 /khulnasoft-code-web/khulnasoft-web/package.json /khulnasoft-code-web/khulnasoft-web/package.nls.json /khulnasoft-code-web/khulnasoft-web/README.md /khulnasoft-code-web/khulnasoft-web/LICENSE.txt /ide/extensions/khulnasoft-web/
