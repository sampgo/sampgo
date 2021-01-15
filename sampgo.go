package sampgo

/*
#cgo linux CFLAGS: -I./sampgdk -I./sampgdk/amx -DLINUX -D_GNU_SOURCE -Wno-implicit-function-declaration
#cgo linux LDFLAGS: -ldl
#cgo windows CFLAGS: -I./sampgdk -I./sampgdk/amx -DWINDOWS

#cgo CFLAGS: -Wno-attributes

#cgo linux LDFLAGS: -ldl

#ifndef GOLANG_APP
#define GOLANG_APP

#include "unitybuild.c"

#endif
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type event struct {
	Handler interface{}
}

var events = make(map[string]event)
var mainEvent func() = nil

// On registers an event with a handler.
func On(eventName string, handler interface{}) error {
	_, ok := events[eventName]
	if ok {
		return fmt.Errorf("this handler already exists")
	}

	events[eventName] = event{Handler: handler}
	Print(fmt.Sprintf("Registered %s event", eventName))

	return nil
}

// Print allows you to print to the SAMP console.
func Print(msg string) error {
	cstr := C.CString(msg)
	defer C.free(unsafe.Pointer(cstr))
	C.goLogprintf(cstr)

	return nil
}
