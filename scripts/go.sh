#!/bin/bash

CLEANUP_DIRS=(api docs)
readonly CLEANUP_DIRS

OPENAPI_GENERATOR_CLI_VER=latest
readonly OPENAPI_GENERATOR_CLI_VER

GIT_USER=${GIT_USER:-apache}
readonly GIT_USER

function validate_input {
    if [ "$#" -ne 2 ]; then
        echo "USAGE: $0 SPEC_PATH OUTPUT_DIR"
        exit 1
    fi

    if ! [ -x "$(command -v realpath)" ]; then
      echo 'Error: realpath is not installed.' >&2
      exit 1
    fi

    SPEC_PATH=$(realpath "$1")
    readonly SPEC_PATH

    if [ ! -d "$2" ]; then
        echo "$2 is not a valid directory or does not exist."
        exit 1
    fi

    OUTPUT_DIR=$(realpath "$2")
    readonly OUTPUT_DIR

    # cleanup the existing generated code, otherwise generator would skip them
    for dir in "${CLEANUP_DIRS[@]}"
    do
        local dirToClean="${OUTPUT_DIR}/${dir}"
        echo "Cleaning up ${dirToClean}"
        rm -rf "${dirToClean:?}"
    done

    # create openapi ignore file to keep generated code clean
    cat <<EOF > "${OUTPUT_DIR}/.openapi-generator-ignore"
.travis.yml
git_push.sh
.gitlab-ci.yml
requirements.txt
setup.cfg
setup.py
test-requirements.txt
tox.ini
EOF
}

function gen_client {
    lang=$1
    shift
    set -ex
    IFS=','
    docker run --rm \
        -u "$(id -u):$(id -g)" \
        -v "${SPEC_PATH}:/spec" \
        -v "${OUTPUT_DIR}:/output" \
        openapitools/openapi-generator-cli:v${OPENAPI_GENERATOR_CLI_VER} \
        generate \
        --input-spec "/spec" \
        --generator-name "${lang}" \
        --git-user-id "${GIT_USER}" \
        --additional-properties=isGoSubmodule=true, \
        --output "/output" "$@"
}

VERSION=2.3.2
readonly VERSION

go_config=(
    "packageVersion=${VERSION}"
    "enumClassPrefix=true"
)

validate_input "$@"

# additional-properties key value tuples need to be separated by comma, not space
IFS=,
gen_client go \
    --package-name airflow \
    --git-repo-id airflow-client-go/airflow \
    --additional-properties "${go_config[*]}"
