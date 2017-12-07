package zbar

// #cgo LDFLAGS: -lzbar
// #include <stdlib.h>
// #include <stdio.h>
// #include <zbar.h>
// #include "simplifier.h"
import "C"

import (
	"errors"
	"unsafe"
)

// Recognize single code from webcam
func ScanSingleSymbol(device string) (result, symbol_type string, err error) {
	dev := C.CString(device)
	defer C.free(unsafe.Pointer(dev))
	var (
		res   *C.char = (*C.char)(C.malloc(1024)) // * (C.size_t(unsafe.Sizeof(C.char))))
		sym_t *C.char = (*C.char)(C.malloc(128))  // * C.size_t(C.sizeof(C.char))))
	)
	defer C.free(unsafe.Pointer(res))
	defer C.free(unsafe.Pointer(sym_t))

	if C.scan_single_symbol(dev, res, sym_t) < 0 { // error occured
		return "", "", errors.New("ZBar processing error")
	}

	result = C.GoString(res)
	symbol_type = C.GoString(sym_t)

	return result, symbol_type, nil
}
