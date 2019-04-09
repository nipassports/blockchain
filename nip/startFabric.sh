#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error
set -e

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1
starttime=$(date +%s)
CC_SRC_LANGUAGE=${1:-"go"}
CC_SRC_LANGUAGE=`echo "$CC_SRC_LANGUAGE" | tr [:upper:] [:lower:]`
CC_RUNTIME_LANGUAGE=golang
CC_SRC_PATH=github.com/passport/go

# clean the keystore
rm -rf ./hfc-key-store

cd javascript/wallet
rm -rf *

cd ../..

cd ../first-network
./byfn.sh down
./byfn.sh up

cd scripts
./connexion_update.sh
