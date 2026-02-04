#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o dist/qrcode-linux-amd64 src/main.go
sha256sum dist/qrcode-linux-amd64 > dist/qrcode-linux-amd64.sha256sum

GOOS=windows GOARCH=amd64 go build -o dist/qrcode-windows-amd64.exe src/main.go
sha256sum dist/qrcode-windows-amd64.exe > dist/qrcode-windows-amd64.exe.sha256sum

GOOS=darwin GOARCH=arm64 go build -o dist/qrcode-darwin-arm64 src/main.go
sha256sum dist/qrcode-darwin-arm64 > dist/qrcode-darwin-arm64.sha256sum

echo "Build completed"