#!/bin/bash
set -e -o pipefail

: ${WORKDIR:="/root/go/src/github.com/vmware/dispatch"}
: ${CI_IMAGE:="vmware/dispatch-golang-ci:1.10"}
: ${QUIET:="-q"}

echo Using image ${CI_IMAGE}

PACKAGE=${1}
APP=${2}
SWAGGER=${3}


mkdir -p ./pkg/${PACKAGE}/gen
if [[ -z ${CI} ]]; then
    docker run --rm -v `pwd`:${WORKDIR} ${CI_IMAGE} bash -c "cd ${WORKDIR} && swagger generate server ${QUIET} -A ${APP} -t ./pkg/${PACKAGE}/gen -f ./swagger/${SWAGGER} --exclude-main"
    docker run --rm -v `pwd`:${WORKDIR} ${CI_IMAGE} bash -c "cd ${WORKDIR} && swagger generate client ${QUIET} -A ${APP} -t ./pkg/${PACKAGE}/gen -f ./swagger/${SWAGGER}"
else
    echo "CI is set to ${CI}"
    swagger generate server ${QUIET} -A ${APP} -t ./pkg/${PACKAGE}/gen -f ./swagger/${SWAGGER} --exclude-main
    swagger generate client ${QUIET} -A ${APP} -t ./pkg/${PACKAGE}/gen -f ./swagger/${SWAGGER}
fi
