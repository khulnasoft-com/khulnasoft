# Copyright (c) 2023 Khulnasoft GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License.AGPL.txt in the project root for license information.

FROM cgr.dev/chainguard/wolfi-base:latest@sha256:e3daef4eecf6f1cc4f505d6384871208fe8b476252ae92e109f7f8701f275544

WORKDIR /app
COPY components-local-app--app-with-manifest/bin/* ./

ARG __GIT_COMMIT
ARG VERSION

ENV KHULNASOFT_BUILD_GIT_COMMIT=${__GIT_COMMIT}
ENV KHULNASOFT_BUILD_VERSION=${VERSION}
CMD ["/bin/sh", "-c", "cp /app/* /out"]
