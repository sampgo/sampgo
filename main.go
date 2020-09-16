package main

/*
#cgo LDFLAGS: -L. -l:sampgdk/build/bin/Release/libsampgdk.a
#include "sampgdk/sampgdk.h"

typedef const char* str;

PLUGIN_EXPORT bool PLUGIN_CALL OnPublicCall(AMX *amx, const char *name, cell *params, cell *retval) {
	sampgdk_logprintf("OnPublicCall from C side, public called was %s\n", name);
	return true;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

/*func OnPublicCall(amx *C.AMX, name C.str) bool {
	fmt.Printf("Function %s called\n", C.GoString(name))
	return true
}*/

//export goModeInit
func goModeInit() {
	fmt.Println("goModeInit called")
}

//export Supports
func Supports() uint {
	fmt.Println("Called main.go#Supports")
	return C.SUPPORTS_VERSION | C.SUPPORTS_AMX_NATIVES
}

//export Load
func Load(ppData *unsafe.Pointer) bool {
	fmt.Println("Called main.go#Load")
	/*pAMXFunctions := unsafe.Pointer(uintptr(*ppData) + C.PLUGIN_DATA_AMX_EXPORTS)
	logprintf := (*func(string, ...string))(unsafe.Pointer(uintptr(pAMXFunctions) + C.PLUGIN_DATA_LOGPRINTF))
	log.Printf("logprintf: %v", logprintf)
	(*logprintf)("Hello")*/
	return true
}

func main() {}
