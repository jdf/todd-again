#!/bin/sh
dir=$(cd -P -- "$(dirname -- "$0")" && pwd -P)
pushd "$dir"
go install ./cmd/protoc-gen-imgui && \
  protoc --proto_path=./game/proto \
     --imgui_out=./game/proto --imgui_opt=paths=source_relative \
     game/proto/level.proto
popd