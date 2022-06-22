#!/usr/bin/env bash

mkdir -p out

pushd golang > /dev/null
GOOS=windows GOARCH=386 go build -o ../out/winreg-tasks-386.exe -trimpath ./cmd
GOOS=windows GOARCH=amd64 go build -o ../out/winreg-tasks-amd64.exe -trimpath ./cmd
popd > /dev/null
