# Copyright (c) 2021 Khulnasoft GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License.AGPL.txt in the project root for license information.

FROM ghcr.io/khulnasoft-com/buildkit:v0.20.1-khulnasoft.2

USER root
RUN apk --no-cache add sudo bash \
    && addgroup -g 33333 khulnasoft \
    && adduser -D -h /home/khulnasoft -s /bin/sh -u 33333 -G khulnasoft khulnasoft \
    && echo "khulnasoft ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/khulnasoft \
    && chmod 0440 /etc/sudoers.d/khulnasoft

COPY components-image-builder-bob--runc-facade/bob /app/runc-facade
RUN chmod 4755 /app/runc-facade \
    && mv /usr/bin/buildkit-runc /usr/bin/bob-runc \
    && mv /app/runc-facade /usr/bin/buildkit-runc

COPY components-image-builder-bob--app/bob /app/
RUN chmod 4755 /app/bob

RUN mkdir /ide
COPY ide-startup.sh /ide/startup.sh
COPY supervisor-ide-config.json /ide/

ARG __GIT_COMMIT
ARG VERSION

ENV KHULNASOFT_BUILD_GIT_COMMIT=${__GIT_COMMIT}
ENV KHULNASOFT_BUILD_VERSION=${VERSION}
# sudo buildctl-daemonless.sh
ENTRYPOINT [ "/app/bob" ]
CMD [ "build" ]
