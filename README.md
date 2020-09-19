# sampgo
sampgo is a SAMP gamemode SDK written in Go, based on Zeex's SAMPGDK.

CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -buildmode=c-shared -o build/libsampgo.so main.go

## Quickstart
Clone the `https://github.com/sampgo/base.git` repo.

Or check the wiki for more information.

## Installation
```
go get -u github.com/sampgo/sampgo
```

## Credits
AliLogic for quite a few things.

Dakyskye for helping out with some event handling logic, and a lot of motivation.