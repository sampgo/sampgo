# sampgo
sampgo is a SAMP gamemode SDK written in Go, based on Zeex's SAMPGDK.

CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -buildmode=c-shared -o build/libsampgo.so main.go

## Quickstart
```
package main

import (
	"fmt"

	"github.com/sampgo/sampgo/sampgo"
)

func init() {
	sampgo.Print("go init() called")

	sampgo.On("goModeInit", func() bool {
		sampgo.Print("Hello from Go!")

		v1 := sampgo.CreateVehicle(503, 0, 0, 5.0, 0, 1, 1, 30, true)
		sampgo.Print(fmt.Sprintf("Vehicle id %d created", v1))

		v2 := sampgo.CreateVehicle(504, 0, 2, 5, 0, 3, 3, 30, false)
		sampgo.Print(fmt.Sprintf("Vehicle id %d created", v2))

		return true
	})

	sampgo.On("goModeExit", func() bool {
		sampgo.Print("goModeExit!")
		return true
	})

	sampgo.On("onPlayerConnect", func(p sampgo.Player) bool {
		sampgo.Print(fmt.Sprintf("Player ID is %d", p.ID))
		return true
	})
}

func main() {}
```

## Installation
```
go get -u github.com/sampgo/sampgo
```

## Credits
AliLogic for quite a few things.
Dakyskye for helping out with some event handling logic, and a lot of motivation.