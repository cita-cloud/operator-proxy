#!/bin/bash
#
# Copyright Rivtower Technologies LLC.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

set -o errexit
export OWNER=cita-cloud
export REPO=operator-proxy
export BIN_LOCATION="/usr/local/bin"

cli_version=$(curl -s https://api.github.com/repos/$OWNER/$REPO/releases/latest | grep 'tag_name' | cut -d '"' -f 4 | tr -d 'v')

function sudocmd(){
    if [ $(whoami) == "root" ]; then
        echo ""
    else
        echo "sudo"
    fi
}
SUDO=$(sudocmd)

getPackage() {
    uname=$(uname)
    userid=$(id -u)

    suffix=""
    case $uname in
    "Darwin")
        arch=$(uname -m)
        case $arch in
        "x86_64")
        suffix="darwin-amd64"
        ;;
        esac
        case $arch in
        "arm64")
        suffix="darwin-arm64"
        ;;
        esac
    ;;

    "MINGW"*)
    suffix=".exe"
    BINLOCATION="$HOME/bin"
    mkdir -p $BINLOCATION

    ;;
    "Linux")
        arch=$(uname -m)
        echo $arch
        case $arch in
        "x86_64")
        suffix="linux-amd64"
        ;;
        esac
        case $arch in
        "aarch64")
        suffix="linux-arm64"
        ;;
        esac
    ;;
    esac

    download_dir="/tmp"
    targetFile="$download_dir/cco-cli-$cli_version-$suffix"

    if [ -e "$targetFile" ]; then
        rm "$targetFile"
    fi

    url=https://github.com/$OWNER/$REPO/releases/download/v$cli_version/cco-cli-$cli_version-$suffix
    echo "Downloading package $url as $targetFile"

    curl -sSL $url --output "$targetFile"

    if [ "$?" = "0" ]; then
      echo "Download complete."
      sh -c "$SUDO mv $targetFile $BIN_LOCATION/cco-cli"
      sh -c "$SUDO chmod +x $BIN_LOCATION/cco-cli"

      echo "Install successful!"
      echo
      echo "=============================================================="
      echo "  You should first set the OPERATOR_PROXY_ENDPOINT"
      echo "  environment variable to use this client, like:"
      echo "  export OPERATOR_PROXY_ENDPOINT=IP:PORT."
      echo "  IP is your kubernetes work node's ip, and you can"
      echo "  find the PORT from below command in 'PORT(S)' column:"
      echo "  kubectl get svc cita-cloud-operator-proxy -n{YOUR_NAMESPACE}"
      echo "=============================================================="
      echo
    fi
}

getPackage