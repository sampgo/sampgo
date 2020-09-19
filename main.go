package main

import (
	"fmt"

	"github.com/sampgo/sampgo/sampgo"
)

func init() {
	sampgo.Print("go init() called")

	sampgo.On("goModeInit", func() bool {
		sampgo.Print("Hello from Go!")
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
