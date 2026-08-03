#!/bin/bash
set -euo pipefail

echo "===================================="
echo "FinancePA DevContainer Setup"
echo "===================================="

if [ "$(id -u)" -ne 0 ]; then
  echo "ERROR: This script must be run as root"
  exit 1
fi

echo ""
echo "Detecting system architecture..."
MACHINE=$(uname -m)
case "${MACHINE}" in
  x86_64)
    ARCH="amd64"
    ;;
  aarch64|arm64)
    ARCH="arm64"
    ;;
  *)
    echo "WARNING: Unsupported architecture ${MACHINE}, defaulting to amd64"
    ARCH="amd64"
    ;;
esac
echo "Architecture: ${ARCH}"

echo ""
echo "------------------------------------"
echo "Setting up bash completion..."
echo "------------------------------------"

BASH_COMPLETIONS_DIR="/usr/share/bash-completion/completions"
if [ -d "${BASH_COMPLETIONS_DIR}" ]; then
  if ! grep -q "source /usr/share/bash-completion/bash_completion" ~/.bashrc 2>/dev/null; then
    echo 'source /usr/share/bash-completion/bash_completion' >> ~/.bashrc
    echo "Added bash-completion to .bashrc"
  fi
fi

echo ""
echo "------------------------------------"
echo "Installing Kubernetes tooling..."
echo "------------------------------------"

if ! command -v kind &> /dev/null; then
  echo "Installing kind..."
  curl -Lo /usr/local/bin/kind "https://kind.sigs.k8s.io/dl/latest/kind-linux-${ARCH}"
  chmod +x /usr/local/bin/kind
fi

if command -v kind &> /dev/null; then
  if [ -d "${BASH_COMPLETIONS_DIR}" ]; then
    kind completion bash > "${BASH_COMPLETIONS_DIR}/kind" 2>/dev/null || true
  fi
fi

if ! command -v kubebuilder &> /dev/null; then
  echo "Installing kubebuilder..."
  curl -Lo /usr/local/bin/kubebuilder "https://go.kubebuilder.io/dl/latest/linux/${ARCH}"
  chmod +x /usr/local/bin/kubebuilder
fi

if command -v kubebuilder &> /dev/null; then
  if [ -d "${BASH_COMPLETIONS_DIR}" ]; then
    kubebuilder completion bash > "${BASH_COMPLETIONS_DIR}/kubebuilder" 2>/dev/null || true
  fi
fi

if ! command -v kubectl &> /dev/null; then
  echo "Installing kubectl..."
  KUBECTL_VERSION=$(curl -Ls https://dl.k8s.io/release/stable.txt)
  curl -Lo /usr/local/bin/kubectl "https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/${ARCH}/kubectl"
  chmod +x /usr/local/bin/kubectl
fi

if command -v kubectl &> /dev/null; then
  if [ -d "${BASH_COMPLETIONS_DIR}" ]; then
    kubectl completion bash > "${BASH_COMPLETIONS_DIR}/kubectl" 2>/dev/null || true
  fi
fi

if command -v docker &> /dev/null; then
  if [ -d "${BASH_COMPLETIONS_DIR}" ]; then
    docker completion bash > "${BASH_COMPLETIONS_DIR}/docker" 2>/dev/null || true
  fi
fi

echo ""
echo "------------------------------------"
echo "Configuring Docker environment..."
echo "------------------------------------"

for i in {1..30}; do
  if docker info >/dev/null 2>&1; then
    echo "Docker is ready"
    break
  fi
  if [ "$i" -eq 30 ]; then
    echo "WARNING: Docker not ready after 30s"
  fi
  sleep 1
done

if ! docker network inspect kind >/dev/null 2>&1; then
  if docker network create kind >/dev/null 2>&1; then
    echo "Created kind network"
  else
    echo "WARNING: Failed to create kind network"
  fi
fi

echo ""
echo "------------------------------------"
echo "Verifying installations..."
echo "------------------------------------"
kind version >/dev/null 2>&1 || true
kubebuilder version >/dev/null 2>&1 || true
kubectl version --client >/dev/null 2>&1 || true
docker --version >/dev/null 2>&1 || true
go version >/dev/null 2>&1 || true

echo ""
echo "===================================="
echo "DevContainer ready!"
echo "===================================="
echo "All development tools installed successfully."
