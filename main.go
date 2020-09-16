package main

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>

typedef struct {
	const char *dli_fname;
	void       *dli_fbase;
	const char *dli_sname;
	void       *dli_saddr;
} Dl_info;

#include "sampgdk/sampgdk.c"

typedef const char* str;
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func OnPublicCall(amx *C.AMX, name C.str) bool {
	fmt.Printf("Function %s called\n", C.GoString(name))
	return true
}

func OnGameModeInit() bool {
	fmt.Println("Called main.go#OnGameModeInit")
	return true
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
