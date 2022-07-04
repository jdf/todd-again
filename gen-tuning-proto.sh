#!/bin/sh
dir=$(cd -P -- "$(dirname -- "$0")" && pwd -P)
pushd "$dir"
protoc --proto_path=./game/proto \
     --go_out=./game/proto --go_opt=paths=source_relative \
     game/proto/tuning.proto
popd