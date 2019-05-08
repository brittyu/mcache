package main

// #cgo LDFLAGS: -L ./ -lfoo -Wl,-rpath=./
// #include <stdio.h>
// #include <stdlib.h>
// #include "foo.h"
import "C"
import "fmt"

func main() {
	fmt.Println(C.count)
	C.foo()
}
