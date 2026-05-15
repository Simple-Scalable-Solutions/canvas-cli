#!/bin/sh
set -e

REPO="simple-scalable-solutions/canvas-cli"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

# Detect OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$OS" in
  linux|darwin) ;;
  *) echo "Unsupported OS: $OS" >&2; exit 1 ;;
esac

# Detect architecture
ARCH=$(uname -m)
case "$ARCH" in
  x86_64)         ARCH="amd64" ;;
  arm64|aarch64)  ARCH="arm64" ;;
  *)              echo "Unsupported architecture: $ARCH" >&2; exit 1 ;;
esac

# Resolve version
if [ -z "$VERSION" ]; then
  VERSION=$(curl -sSf "https://api.github.com/repos/${REPO}/releases/latest" \
    | grep '"tag_name"' | sed 's/.*"tag_name": *"\([^"]*\)".*/\1/')
fi
if [ -z "$VERSION" ]; then
  echo "Could not determine latest version. Set VERSION=vX.Y.Z to override." >&2
  exit 1
fi

TARBALL="canvas-cli_${VERSION#v}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/${VERSION}/${TARBALL}"

echo "Installing canvas-cli ${VERSION} (${OS}/${ARCH}) → ${INSTALL_DIR}"

TMP=$(mktemp -d)
trap 'rm -rf "$TMP"' EXIT

curl -sSfL "$URL" | tar -xz -C "$TMP"
install -m755 "$TMP/canvas-cli" "${INSTALL_DIR}/canvas-cli"
install -m755 "$TMP/canvas-mcp" "${INSTALL_DIR}/canvas-mcp"

echo "Done. Run: canvas-cli --version"
