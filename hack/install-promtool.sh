#!/bin/sh
# Install the pinned PromQL test tool without adding Go module dependencies.
set -eu

destination=${1:?usage: install-promtool.sh DESTINATION_DIRECTORY}
version=3.5.0
platform="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "$(uname -m)" in
    x86_64) arch=amd64 ;;
    aarch64|arm64) arch=arm64 ;;
    *) echo "unsupported architecture" >&2; exit 1 ;;
esac
case "$platform-$arch" in
    linux-amd64) checksum=e811827af26d822afb09a4f28314f61b618b12cff5369835a67f674d8b46f39a ;;
    linux-arm64) checksum=173389cc42bf09c4e6e54cb53fa07a5a835d7c261e14775d2183181d6e385d1c ;;
    darwin-amd64) checksum=5856d0c306e7d852d9a79788b20acb630501b3fab120b69f5317d6b0732dea57 ;;
    darwin-arm64) checksum=c17bc2c992f0b5515eebe2b4ccd626214f8e309a0fe6ec40ed13c42ec13d05fb ;;
    *) echo "unsupported platform: $platform" >&2; exit 1 ;;
esac

temporary=$(mktemp -d)
trap 'rm -rf "$temporary"' EXIT HUP INT TERM
archive="prometheus-$version.$platform-$arch"
curl -fsSL "https://github.com/prometheus/prometheus/releases/download/v$version/$archive.tar.gz" -o "$temporary/archive.tar.gz"
if command -v sha256sum >/dev/null 2>&1; then
    echo "$checksum  $temporary/archive.tar.gz" | sha256sum -c -
else
    echo "$checksum  $temporary/archive.tar.gz" | shasum -a 256 -c -
fi
tar -xzf "$temporary/archive.tar.gz" -C "$temporary" "$archive/promtool"
mkdir -p "$destination"
cp "$temporary/$archive/promtool" "$destination/promtool"
