package sampgo

/*
#cgo CFLAGS: -Wno-attributes
#cgo linux LDFLAGS: -L. -l:sampgdk/build/bin/Debug/libsampgdk.a
#cgo linux LDFLAGS: -ldl

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

func (p Player) SendMessage(colour int, msg string) error {
	if len(msg) < 1 {
		return fmt.Errorf("Msg not long enough")
	}

	SendClientMessage(p.ID, colour, msg)
	return nil
}
