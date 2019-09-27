#!/bin/bash

set -eux

BASE_DIR="/var/tmp"

cd ${BASE_DIR}

git clone https://github.com/bats-core/bats-core.git

cd bats-core
sudo ./install.sh /usr/local
cd -
rm -rf ${BASE_DIR}/bats