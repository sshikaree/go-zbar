package zbar

// #cgo LDFLAGS: -lzbar
// #include <stdlib.h>
// #include <zbar.h>
import "C"
import (
	"unsafe"
)

func NewProcessor(threaded int) *Processor {
	p := Processor{}
	p.c_processor = C.zbar_processor_create(C.int(threaded))

	return &p
}

type Processor struct {
	c_processor *C.zbar_processor_t
}

func (p *Processor) SetConfig(symbology int, config int, value int) int {
	return int(C.zbar_processor_set_config(p.c_processor, C.zbar_symbol_type_t(symbology), C.zbar_config_t(config), C.int(value)))
}

func (p *Processor) Init(device string, enableDisplay int) int {
	c_device := C.CString(device)
	defer C.free(unsafe.Pointer(c_device))
	return int(C.zbar_processor_init(p.c_processor, c_device, C.int(enableDisplay)))
}

func (p *Processor) Destroy() {
	C.zbar_processor_destroy(p.c_processor)
	p = nil
}

func (p *Processor) RequestSize(width, height uint) int {
	return int(C.zbar_processor_request_size(p.c_processor, C.unsigned(width), C.unsigned(height)))
}

func (p *Processor) RequestInterface(version int) int {
	return int(C.zbar_processor_request_interface(p.c_processor, C.int(version)))
}

func (p *Processor) SetVisible(visible int) int {
	return int(C.zbar_processor_set_visible(p.c_processor, C.int(visible)))
}

func (p *Processor) SetActive(active int) int {
	return int(C.zbar_processor_set_active(p.c_processor, C.int(active)))
}

func (p *Processor) ProcessOne(timeout int) int {
	return int(C.zbar_process_one(p.c_processor, C.int(timeout)))
}

func (p *Processor) GetResults() *SymbolSet {
	ss := SymbolSet{}
	ss.c_symbol_set = C.zbar_processor_get_results(p.c_processor)
	if ss.c_symbol_set != nil {
		return &ss

	}
	return nil
}
