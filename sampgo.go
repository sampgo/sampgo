package sampgo

/*
#cgo CFLAGS: -I./vendor/sdk -I./vendor/gdk -I./interop

#ifndef GOLANG_APP
#define GOLANG_APP

#include "interop.h"

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

//export onTick
func onTick() {
	evt, ok := events["tick"]
	if !ok {
		return
	}

	fn, ok := evt.Handler.(func())
	if !ok {
		return
	}
	fn()
	return
}

// On registers an event with a handler.
func On(eventName string, handler interface{}) error {
	_, ok := events[eventName]
	if ok {
		return fmt.Errorf("this handler already exists")
	}

	events[eventName] = event{Handler: handler}
	_ = Print(fmt.Sprintf("Registered %s event", eventName))

	return nil
}

// Print allows you to print to the SAMP console.
func Print(msg string) error {
	cstr := C.CString(msg)
	defer C.free(unsafe.Pointer(cstr))
	C.goLogprintf(cstr)

	return nil
}
