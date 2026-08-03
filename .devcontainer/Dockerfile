ARG VARIANT=4-24-trixie
FROM mcr.microsoft.com/devcontainers/typescript-node:${VARIANT}

## Copy all the scripts
#
COPY scripts/*.sh /tmp/scripts/

## Install dependencies
#
RUN bash /tmp/scripts/install_dependencies.sh
RUN bash /tmp/scripts/post-install.sh