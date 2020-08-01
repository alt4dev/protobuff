#!/usr/bin/bash

function golang_build(){
    echo "Building go lang"
    if [ -e proto ]; then
        rm -r proto
    fi
    mkdir proto
    protoc --proto_path=./ --go_out=proto/ --go_opt=paths=source_relative definitions.proto
}

function build_help(){
    echo "Invalid language provided. Below are the supported commands"
    echo "    ./build.sh go"
}

LANGUAGE="$1"
case $LANGUAGE in
    "go") golang_build;;
    *) build_help;;
esac
