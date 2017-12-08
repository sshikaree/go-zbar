package zbar

// #include <stdlib.h>
// #include <zbar.h>
import "C"
import "unsafe"

type Image struct {
	c_image *C.zbar_image_t
}

func NewImage() *Image {
	img := Image{}
	img.c_image = C.zbar_image_create()
	return &img
}

func (i *Image) Destroy() {
	C.zbar_image_destroy(i.c_image)
	i = nil
}

func (i *Image) Ref(refs int) {
	C.zbar_image_ref(i.c_image, C.int(refs))
}

// Image format conversion.
func (i *Image) Convert(format uint64) *Image {
	newImg := Image{}
	newImg.c_image = C.zbar_image_convert(i.c_image, C.ulong(format))
	return &newImg
}

// Image format conversion with crop/pad.
func (i *Image) ConvertResize(format uint64, width, height uint) *Image {
	newImg := Image{}
	newImg.c_image = C.zbar_image_convert_resize(i.c_image, C.ulong(format), C.unsigned(width), C.unsigned(height))
	return &newImg
}

// Retrieve the image format.
func (i *Image) GetFormat() uint64 {
	return uint64(C.zbar_image_get_format(i.c_image))
}

// Retrieve a "sequence" (page/frame) number associated with this image.
func (i *Image) GetSequence() uint {
	return uint(C.zbar_image_get_sequence(i.c_image))
}

// Retrieve the width of the image.
func (i *Image) GetWidth() uint {
	return uint(C.zbar_image_get_width(i.c_image))
}

// Retrieve the height of the image.
func (i *Image) GetHeight() uint {
	return uint(C.zbar_image_get_height(i.c_image))
}

// Return the image sample data.
func (i *Image) GetData() interface{} {
	return C.zbar_image_get_data(i.c_image)
}

// Return the size of image data.
func (i *Image) GetDataLength() uint64 {
	return uint64(C.zbar_image_get_data_length(i.c_image))
}

// Retrieve the decoded results.
func (i *Image) GetSymbols() *SymbolSet {
	ss := SymbolSet{}
	ss.c_symbol_set = C.zbar_image_get_symbols(i.c_image)
	if ss.c_symbol_set != nil {
		return &ss
	}
	return nil
}

// Associate the specified symbol set with the image, replacing any existing results.
func (i *Image) SetSymbols(symbols *SymbolSet) {
	C.zbar_image_set_symbols(i.c_image, symbols.c_symbol_set)
}

// Image_scanner decode result iterator.
func (i *Image) FirstSymbol() *Symbol {
	s := Symbol{}
	s.c_symbol = C.zbar_image_first_symbol(i.c_image)
	if s.c_symbol != nil {
		return &s
	}
	return nil
}

// Specify the fourcc image format code for image sample data.
func (i *Image) SetFormat(format uint64) {
	C.zbar_image_set_format(i.c_image, C.ulong(format))
}

// Associate a "sequence" (page/frame) number with this image.
func (i *Image) SetSequence(sequenceNum uint) {
	C.zbar_image_set_sequence(i.c_image, C.unsigned(sequenceNum))
}

// Specify the pixel size of the image.
func (i *Image) SetSize(width, height uint) {
	C.zbar_image_set_size(i.c_image, C.unsigned(width), C.unsigned(height))
}

// Specify image sample data.
// func (i *Image) SetData() {
//
// }

// Built-in cleanup handler.
func (i *Image) FreeData() {
	C.zbar_image_free_data(i.c_image)
}

// Dump raw image data to a file for debug.
func (i *Image) Write(filebase string) int {
	cfilebase := C.CString(filebase)
	defer C.free(unsafe.Pointer(cfilebase))
	return int(C.zbar_image_write(i.c_image, cfilebase))
}

// Read back an image in the format written by zbar_image_write()
// func (i *Image) Read(filename string) *Image {
// 	cfilename := C.CString(filename)
// 	defer C.free(unsafe.Pointer(cfilename))
// 	i := Image{}
// 	i.c_image = C.zbar_image_read(cfilename)
// 	return &i
// }
