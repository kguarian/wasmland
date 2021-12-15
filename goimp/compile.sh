#!/usr/bin/bash

`GOOS=js GOARCH=wasm go build -o=wasm_lib.wasm wasmBuild`
errcode=$?
sudo cp wasm_lib.wasm ../pkg/wasm_lib.wasm
exit $errcode

git add *
git commit -m "update"
git push