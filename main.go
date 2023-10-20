package main

// #cgo CFLAGS: -I/usr/local/include/rav1e
// #cgo LDFLAGS: trampoline.o -L/usr/local/lib -lrav1e
//#include "rav1e.h"
//#include "trampoline.h"
import "C"

import (
	"fmt"
	"unsafe"
)

// could be a generics candidate
func SetConfigValue(raConfig *C.RaConfig, key string, value interface{}) {
	ckey := C.CString("width")
	// will not work for all types - need to restrict to string, bool, int
	cvalue := C.CString(fmt.Sprintf("%v", value))
	C.rav1e_config_parse(raConfig, ckey, cvalue)
	C.free(unsafe.Pointer(ckey))
	C.free(unsafe.Pointer(cvalue))
}

const (
	width  = 64
	height = 96
	speed  = 9
)

func test_setup() {
	racfg := C.new_rav1e()
	C.t_rav1e_config_default(racfg)

	ret := C.t_rav1e_simple_setup(racfg)
	fmt.Printf("t_rav1e_simple_setup returned %v\n", ret)

	ret = C.t_rav1e_simple_chromaticity(racfg)
	fmt.Printf("t_rav1e_test returned %v\n", ret)

	ret = C.t_rav1e_context_and_frame(racfg)
	fmt.Printf("t_rav1e_context_and_frame returned %v\n", ret)

	ret = C.t_rav1e_send(racfg)
	fmt.Printf("t_rav1e_send returned %v\n", ret)
}

func main() {

	fmt.Printf("starting!\n")

	test_setup()

}
