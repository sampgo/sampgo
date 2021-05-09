package interop

/*
#cgo windows CFLAGS: -I./lib -I./lib/amx -Wno-attributes -Wno-implicit-function-declaration
#cgo windows CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DWIN32
#cgo windows LDFLAGS: -Wl,--subsystem,windows,--kill-at

#cgo linux CFLAGS: -I./lib -I./lib/amx -Wno-attributes -Wno-implicit-function-declaration
#cgo linux CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DLINUX -D_GNU_SOURCE
#cgo linux LDFLAGS: -ldl

#ifndef GOLANG_APP
#define GOLANG_APP

#include "../lib/unitybuild.c"
#include "../lib/main.h"

#endif
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type Native struct {
	Ptr  C.AMX_NATIVE
	Name string
}

func NewNative(name string) (Native, error) {
	var native Native

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	ptr := C.sampgdk_FindNative(C.constToNonConst(cname))

	if ptr == nil {
		return native, fmt.Errorf("native could not be found")
	}

	native.Ptr = ptr
	native.Name = name

	return native, nil
}

func (n *Native) Invoke() (int, error) {
	if n.Ptr == nil {
		return -1, fmt.Errorf("native is invalid")
	}

	cparams := C.CString("")
	defer C.free(unsafe.Pointer(cparams))

	ret := C.sampgdk_InvokeNative(n.Ptr, cparams)

	return ret, nil
}
