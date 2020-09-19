package sampgo

/*
#cgo CFLAGS: -Wno-attributes

#cgo linux.386 linux LDFLAGS: -L. -l:sampgdk/build/bin/Debug/libsampgdk.a
#cgo linux.386 linux LDFLAGS: -ldl

#cgo windows.386 LDFLAGS: -L. -l:sampgdk/build/bin/Debug/sampgdk.lib

#ifndef GOLANG_APP
#define GOLANG_APP

#include "sampgdk/main.h"

#endif
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Player implements OO players.
type Player struct {
	ID int
}

// GetName returns the players name.
func (p Player) GetName() string {
	var name *C.char
	defer C.free(unsafe.Pointer(name))

	C.GetPlayerName(C.int(p.ID), name, C.int(MaxPlayerName))
	return C.GoString(name)
}

// SetName sets the players name.
func (p Player) SetName(name string) error {
	if len(name) > 24 {
		return fmt.Errorf("name length above 24 chars")
	}

	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))

	C.SetPlayerName(C.int(p.ID), n)
	return nil
}

// SendMessage allows you to send a player a message.
func (p Player) SendMessage(colour int, msg string) error {
	if len(msg) < 1 {
		return fmt.Errorf("Msg not long enough")
	}

	SendClientMessage(p.ID, colour, msg)
	return nil
}

// GetPos gets the player's current position.
func (p Player) GetPos() (float32, float32, float32, float32) {
	x, y, z := GetPlayerPos(p.ID)
	a := GetPlayerFacingAngle(p.ID)
	return x, y, z, a
}

// SetPos sets the player's current position.
func (p Player) SetPos(x, y, z, a float32) error {
	SetPlayerPos(p.ID, x, y, z)
	SetPlayerFacingAngle(p.ID, a)
	return nil
}

// Spawn spawns the player.
func (p Player) Spawn() error {
	err := SpawnPlayer(p.ID)
	if !err {
		return fmt.Errorf("player was unable to be spawned")
	}
	return nil
}
