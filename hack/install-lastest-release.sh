#!/usr/bin/env bash
#
# Install the latest relase, or informed version, of the 'path-helper' binary.
# This may # require "sudo" to install on the target directory, set by environment
# variable "INSTALL_DIR". By default installs the latest release.
#
# Usage:
#   $ hack/install-lastest-release.sh [version]
#

set -Eeu -o pipefail

declare -r APP="path-helper"

# The version to install, or 'latest' to install the latest release.
declare -r VERSION="${1:-latest}"
declare -r TMP_DIR="${TMP_DIR:-/tmp}"
# Installation directory, where the binary will be installed.
declare -r INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

# The repository owner and name, and the base domain path.
declare -r REPO_OWNER="${REPO_OWNER:-otaviof}"
declare -r REPO_NAME="${REPO_NAME:-${APP}}"
declare -r BASE_DOMAIN_PATH="github.com/repos/${REPO_OWNER}/${REPO_NAME}"

# Get the operating system and CPU type.
declare -r LOWER_OSTYPE=$(uname -s | tr '[:upper:]' '[:lower:]')
declare -r CPUTYPE=$(uname -m)

# Search for the specific relase artifact URL, using the operating system and CPU
# type. Returns the full URL to the artifact, or empty if not found.
get_release_artifact_url() {
    local _url="https://api.${BASE_DOMAIN_PATH}/releases"
    local release_part="${LOWER_OSTYPE}_${CPUTYPE}.tar.gz"
    if [[ "${VERSION}" == "latest" ]]; then
        echo $(
            curl -s ${_url}/latest |
                jq -r '.assets[].browser_download_url' |
                grep -i ${release_part} |
                head -n 1
        )
    else
        echo $(
            curl -s ${_url} |
                jq -r ".[] | select(.tag_name == \"${VERSION}\") | .assets[].browser_download_url" |
                grep -i ${release_part} |
                head -n 1
        )
    fi
}

# Installs the application from the artifact URL, if found.
install_release() {
    echo "Installing '${APP}' (${VERSION}) for ${LOWER_OSTYPE} (${CPUTYPE})..."
    local _artifact_url=$(get_release_artifact_url)
    if [[ -z "${_artifact_url}" ]]; then
        echo "Release not found for version '${VERSION}' " \
            "on ${LOWER_OSTYPE} (${CPUTYPE})!"
        exit 1
    fi

    local _tarball="${REPO_NAME}.tar.gz"
    echo "Installing '${APP}' ${VERSION} version from '${_artifact_url}'..."
    curl -sL ${_artifact_url} -o "${_tarball}" >/dev/null &&
        tar -zxvpf "${_tarball}" ${APP} >/dev/null &&
        install -m 0755 ${APP} ${INSTALL_DIR}/${APP} >/dev/null &&
        echo "The '${APP}' ${VERSION} version is installed at '${INSTALL_DIR}'"
}

#
# Main
#

pushd ${TMP_DIR} >/dev/null &&
    install_release
popd
