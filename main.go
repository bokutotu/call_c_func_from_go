package main

/*
#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/lib/bin
#cgo LDFLAGS: -L${SRCDIR}/lib/bin/ 
#cgo LDFLAGS: -lhoge 
#include "./lib/include/hoge.h"
*/
import "C"

func main() {
    C.hoge()
}
