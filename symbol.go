//
//
// TODO:
// 1. Memory management ?!

package zbar

// #include <stdlib.h>
// #include <zbar.h>
import "C"

type SymbolSet struct {
	c_symbol_set *C.zbar_symbol_set_t
}

func (ss *SymbolSet) GetSize() int {
	return int(C.zbar_symbol_set_get_size(ss.c_symbol_set))
}

func (ss *SymbolSet) SetFirstSymbol() *Symbol {
	s := Symbol{}
	s.c_symbol = C.zbar_symbol_set_first_symbol(ss.c_symbol_set)
	return &s
}

type Symbol struct {
	c_symbol *C.zbar_symbol_t
}

func (s *Symbol) Ref(refs int) {
	C.zbar_symbol_ref(s.c_symbol, C.int(refs))
}

func (s *Symbol) GetType() int {
	return int(C.zbar_symbol_get_type(s.c_symbol))
}

func (s *Symbol) GetData() string {
	cstr := C.zbar_symbol_get_data(s.c_symbol)
	// defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

func (s *Symbol) GetDataLength() uint {
	return uint(C.zbar_symbol_get_data_length(s.c_symbol))
}

func (s *Symbol) GetName() string {
	cstr := C.zbar_get_symbol_name(C.zbar_symbol_get_type(s.c_symbol))
	// defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

func (s *Symbol) GetAddonName() string {
	return C.GoString(C.zbar_get_addon_name(C.zbar_symbol_get_type(s.c_symbol)))
}

func (s *Symbol) GetQuality() int {
	return int(C.zbar_symbol_get_quality(s.c_symbol))
}

func (s *Symbol) GetCount() int {
	return int(C.zbar_symbol_get_count(s.c_symbol))
}

func (s *Symbol) GetLocSize() uint {
	return uint(C.zbar_symbol_get_loc_size(s.c_symbol))
}

func (s *Symbol) GetLocX(index uint) int {
	return int(C.zbar_symbol_get_loc_x(s.c_symbol, C.unsigned(index)))
}

func (s *Symbol) GetLocY(index uint) int {
	return int(C.zbar_symbol_get_loc_y(s.c_symbol, C.unsigned(index)))
}

func (s *Symbol) Next() *Symbol {
	newS := Symbol{}
	newS.c_symbol = C.zbar_symbol_next(s.c_symbol)
	if newS.c_symbol != nil {
		return &newS
	}
	return nil
}

func (s *Symbol) GetComponents() *SymbolSet {
	ss := SymbolSet{}
	ss.c_symbol_set = C.zbar_symbol_get_components(s.c_symbol)
	if ss.c_symbol_set != nil {
		return &ss
	}
	return nil
}

func (s *Symbol) FirstComponent() *Symbol {
	newS := Symbol{}
	newS.c_symbol = C.zbar_symbol_first_component(s.c_symbol)
	if newS.c_symbol != nil {
		return &newS
	}
	return nil
}

// func (s *Symbol) XML(buffSize uint) string {

// }
