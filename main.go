package main

import "fmt"

/*
#cgo LDFLAGS: -L. -l:sampgdk/build/bin/Release/libsampgdk.a
#define GOLANG_APP
#include "sampgdk/src/main.h"
*/

func goModeInit() bool {
	fmt.Println("Hello From Go!")
	return true
}

func main() {}
