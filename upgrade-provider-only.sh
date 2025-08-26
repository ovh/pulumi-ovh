#!/bin/bash

#
# Usage: upgrade-provider-light.sh <upstream-provider-version>
#

# MAIN
# Usage
usage="upgrade-provider-light.sh <upstream-provider-version>"
if [[ "$1" == "" || "$1" == "-h" || "$1" == "--help" ]]
then
    echo "$usage"
    exit 1
fi

# Change <upstream-provider-version> in go.mod
export TARGET_PROVIDER_VERSION="$1" #v2.7.0 for example

echo "update target provider version to $TARGET_PROVIDER_VERSION in go.mod"
cd provider
sed -i "s|^\([[:space:]]*\)\(github.com/ovh/terraform-provider-ovh/v2\)[[:space:]].*|\1\2 ${TARGET_PROVIDER_VERSION}|" go.mod


# Generate dependencies

go get github.com/ovh/pulumi-ovh/provider/v2
cd ..

# Re-generate the SDKs

make build_sdks