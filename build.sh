#!/usr/bin/bash

function golang_build(){
    echo "Building go lang"
    if [ -e proto ]; then
        rm -r proto
    fi
    mkdir proto
    protoc --proto_path=./ \
        --go_out=plugins=grpc:proto/ \
        --go_opt=paths=source_relative definitions.proto
    echo "Done."
}

function python_build(){
    OUTPUT_DIR="$1"
    echo "Building python to $OUTPUT_DIR"
    if [ -e $OUTPUT_DIR ]; then
        echo "Overwriting $OUTPUT_DIR"
        rm -r $OUTPUT_DIR
    fi
    mkdir $OUTPUT_DIR
    protoc --proto_path=./ --python_out=$OUTPUT_DIR definitions.proto
    echo "Done."
}

function build_help(){
    echo "Invalid language provided. Below are the supported commands"
    echo "    ./build.sh go"
    echo "    ./build.sh python [OUTPUT_DIR]"
}

LANGUAGE="$1"
OUTPUT_DIR="$2"
case $LANGUAGE in
    "go") golang_build;;
    "python") python_build $OUTPUT_DIR;;
    *) build_help;;
esac
