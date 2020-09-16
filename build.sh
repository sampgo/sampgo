#!/bin/bash
CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -buildmode=c-shared -o build/plugin.so main.go