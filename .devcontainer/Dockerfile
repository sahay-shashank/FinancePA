ARG VARIANT=4-24-trixie
FROM mcr.microsoft.com/devcontainers/typescript-node:${VARIANT}

## Copy all the scripts
#
COPY scripts/*.sh /tmp/scripts/

## Install taskfile
#
RUN echo "$(go env GOOS)"
RUN echo "$(go env GOARCH)"
RUN bash /tmp/scripts/install_dependencies.sh
