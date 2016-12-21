#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

LIB_ROOT=$(dirname "${BASH_SOURCE}")/../../..
source "$LIB_ROOT/hack/libbuild/common/lib.sh"
source "$LIB_ROOT/hack/libbuild/common/public_image.sh"

GOPATH=$(go env GOPATH)
IMG=icinga
ICINGA_VER=2.4.8
K8S_VER=1.5
ICINGAWEB_VER=2.1.2
if [ -f "$GOPATH/src/github.com/appscode/searchlight/dist/.tag" ]; then
	export $(cat $GOPATH/src/github.com/appscode/searchlight/dist/.tag | xargs)
fi

clean() {
    pushd $GOPATH/src/github.com/appscode/searchlight/hack/docker/icinga
	rm -rf icingaweb2 plugins
	popd
}

build() {
    pushd $GOPATH/src/github.com/appscode/searchlight/hack/docker/icinga
    detect_tag $GOPATH/src/github.com/appscode/searchlight/dist/.tag

	rm -rf icingaweb2
	clone https://github.com/Icinga/icingaweb2.git
	cd icingaweb2
	git checkout tags/v$ICINGAWEB_VER
	cd ..

	rm -rf plugins; mkdir -p plugins
	gsutil cp gs://appscode-dev/binaries/hello_icinga/$TAG/hello_icinga-linux-amd64 plugins/hello_icinga
	gsutil cp gs://appscode-dev/binaries/hyperalert/$TAG/hyperalert-linux-amd64 plugins/hyperalert
	chmod 755 plugins/*

	local cmd="docker build -t appscode/$IMG:$TAG-k8s ."
	echo $cmd; $cmd
	popd
}

docker_push() {
	docker_up $IMG:$TAG-k8s
}

docker_release() {
	local cmd="docker push appscode/$IMG:$TAG-k8s"
	echo $cmd; $cmd
}

binary_repo $@
