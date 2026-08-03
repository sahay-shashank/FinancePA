#!/usr/bin/env bash

set -eax

sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin
# Enable bash-completion for task
/usr/local/bin/task --completion bash | sudo tee /etc/bash_completion.d/task > /dev/null

uname_os() {
os=$(uname -s | tr \'[:upper:]\' \'[:lower:]\')
case "$os" in
cygwin_nt*) os="windows" ;;
mingw*) os="windows" ;;
msys_nt*) os="windows" ;;
esac
echo "$os"
}

uname_arch() {
arch=$(uname -m)
case $arch in
x86_64) arch="amd64" ;;
x86) arch="386" ;;
i686) arch="386" ;;
i386) arch="386" ;;
aarch64) arch="arm64" ;;
armv5*) arch="arm" ;;
armv6*) arch="arm" ;;
armv7*) arch="arm" ;;
esac
echo ${arch}
}

curl -L -o kubebuilder "https://go.kubebuilder.io/dl/latest/$(uname_os)/$(uname_arch)"
chmod +x kubebuilder && sudo mv kubebuilder /usr/local/bin/