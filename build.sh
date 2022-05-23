#!/usr/bin/env bash

mkdir -p out

pushd golang > /dev/null
GOOS=windows GOARCH=amd64 go build -o ../out/winreg-tasks.exe -trimpath ./cmd
popd > /dev/null
