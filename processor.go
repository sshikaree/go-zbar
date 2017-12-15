package zbar

// #cgo LDFLAGS: -lzbar
// #include <stdlib.h>
// #include <zbar.h>
// #include "image_data_handler.h"
import "C"
import (
	"unsafe"
)

// var callBacksMap = map[uintptr]func(img *Image){}

// Constructor.
// If threaded is set and threading is available the processor will spawn threads
// where appropriate to avoid blocking and improve responsiveness
func NewProcessor(threaded int) *Processor {
	p := Processor{}
	p.c_processor = C.zbar_processor_create(C.int(threaded))

	return &p
}

type Processor struct {
	c_processor *C.zbar_processor_t
	dataHandler func(image *Image)
	userdata    unsafe.Pointer
}

// Set config for indicated symbology (0 for all) to specified value.
// Returns 0 for success, non-0 for failure (config does not apply to specified symbology, or value out of range)
func (p *Processor) SetConfig(symbology int, config int, value int) int {
	return int(C.zbar_processor_set_config(p.c_processor, C.zbar_symbol_type_t(symbology), C.zbar_config_t(config), C.int(value)))
}

// (re)Initialization.
// Opens a video input device and/or prepares to display output
func (p *Processor) Init(device string, enableDisplay int) int {
	c_device := C.CString(device)
	defer C.free(unsafe.Pointer(c_device))
	return int(C.zbar_processor_init(p.c_processor, c_device, C.int(enableDisplay)))
}

// Destructor.
// Cleans up all resources associated with the processor
func (p *Processor) Destroy() {
	C.zbar_processor_destroy(p.c_processor)
	p = nil
}

// Request a preferred size for the video image from the device
func (p *Processor) RequestSize(width, height uint) int {
	return int(C.zbar_processor_request_size(p.c_processor, C.unsigned(width), C.unsigned(height)))
}

// Request a preferred video driver interface version for debug/testing.
func (p *Processor) RequestInterface(version int) int {
	return int(C.zbar_processor_request_interface(p.c_processor, C.int(version)))
}

// Request a preferred video I/O mode for debug/testing.
func (p *Processor) RequestIOMode(iomode int) int {
	return int(C.zbar_processor_request_iomode(p.c_processor, C.int(iomode)))
}

// Force specific input and output formats for debug/testing.
func (p *Processor) ForceFormat(inputFormat, outputFormat uint64) int {
	return int(C.zbar_processor_force_format(p.c_processor, C.ulong(inputFormat), C.ulong(outputFormat)))
}

//export image_handler_callback
func image_handler_callback(image *C.zbar_image_t, userdata unsafe.Pointer) {
	img := Image{}
	img.c_image = image
	p := (*Processor)(userdata)
	p.dataHandler(&img)
}

// Setup result handler callback.
// The specified function will be called by the processor whenever new results are available from the video stream or a static image.
func (p *Processor) SetDataHandler(fn func(img *Image)) {
	p.dataHandler = fn
	C.zbar_processor_set_data_handler(p.c_processor, (*C.zbar_image_data_handler_t)(C.image_data_handler), unsafe.Pointer(p))
}

// Associate user specified data value with the processor.
func (p *Processor) SetUserData(userdata unsafe.Pointer) {
	C.zbar_processor_set_userdata(p.c_processor, userdata)
}

// Return user specified data value associated with the processor.
func (p *Processor) GetUserData() unsafe.Pointer {
	return unsafe.Pointer(C.zbar_processor_get_userdata(p.c_processor))
}

// Show or hide the display window owned by the library.
func (p *Processor) SetVisible(visible int) int {
	return int(C.zbar_processor_set_visible(p.c_processor, C.int(visible)))
}

// Control the processor in free running video mode.
func (p *Processor) SetActive(active int) int {
	return int(C.zbar_processor_set_active(p.c_processor, C.int(active)))
}

// Process from the video stream until a result is available, or the timeout (in milliseconds) expires.
func (p *Processor) ProcessOne(timeout int) int {
	return int(C.zbar_process_one(p.c_processor, C.int(timeout)))
}

// Retrieve decode results for last scanned image/frame.
func (p *Processor) GetResults() *SymbolSet {
	ss := SymbolSet{}
	ss.c_symbol_set = C.zbar_processor_get_results(p.c_processor)
	if ss.c_symbol_set != nil {
		return &ss

	}
	return nil
}

// Wait for input to the display window from the user (via mouse or keyboard).
func (p *Processor) UserWait(timeout int) int {
	return int(C.zbar_processor_user_wait(p.c_processor, C.int(timeout)))
}

// Retrieve the detail string for the last processor error
func (p *Processor) ErrorString(verbosity int) string {
	return C.GoString(C.zbar_processor_error_string(p.c_processor, C.int(verbosity)))
}

// Retrieve the type code for the last processor error.
func (p *Processor) GetErrorCode() int {
	return int(C.zbar_processor_get_error_code(p.c_processor))
}
