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

// Retrieve set size.
// Returns number of symbols in the set.
func (ss *SymbolSet) GetSize() int {
	return int(C.zbar_symbol_set_get_size(ss.c_symbol_set))
}

//  Set iterator.
// Returns first decoded symbol result in a set. Nil if the set is empty
func (ss *SymbolSet) SetFirstSymbol() *Symbol {
	s := Symbol{}
	s.c_symbol = C.zbar_symbol_set_first_symbol(ss.c_symbol_set)
	if s.c_symbol != nil {
		return &s
	}
	return nil
}

type Symbol struct {
	c_symbol *C.zbar_symbol_t
}

// Symbol reference count manipulation.
// Increment the reference count when you store a new reference to the symbol.
// Decrement when the reference is no longer used.
// Do not refer to the symbol once the count is decremented and the containing image has been recycled or destroyed.
// Note:
// 		the containing image holds a reference to the symbol,
// 		so you only need to use this if you keep a symbol after the image has been destroyed or reused.
func (s *Symbol) Ref(refs int) {
	C.zbar_symbol_ref(s.c_symbol, C.int(refs))
}

// Retrieve type of decoded symbol.
func (s *Symbol) GetType() int {
	return int(C.zbar_symbol_get_type(s.c_symbol))
}

// Retrieve data decoded from symbol.
func (s *Symbol) GetData() string {
	cstr := C.zbar_symbol_get_data(s.c_symbol)
	// defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// Retrieve length of binary data.
func (s *Symbol) GetDataLength() uint {
	return uint(C.zbar_symbol_get_data_length(s.c_symbol))
}

// Retrieve string name for symbol encoding.
func (s *Symbol) GetName() string {
	cstr := C.zbar_get_symbol_name(C.zbar_symbol_get_type(s.c_symbol))
	return C.GoString(cstr)
}

// Retrieve string name for addon encoding
func (s *Symbol) GetAddonName() string {
	return C.GoString(C.zbar_get_addon_name(C.zbar_symbol_get_type(s.c_symbol)))
}

// Retrieve a symbol confidence metric.
// Returns an unscaled, relative quantity: larger values are better than smaller values,
// where "large" and "small" are application dependent.
// Note:
//		expect the exact definition of this quantity to change as the metric is refined.
// 		Currently, only the ordered relationship between two values is defined and will remain stable in the future
func (s *Symbol) GetQuality() int {
	return int(C.zbar_symbol_get_quality(s.c_symbol))
}

// Retrieve current cache count.
// When the cache is enabled for the image_scanner this provides inter-frame reliability and redundancy information for video streams.
// Returns:
//		< 0 if symbol is still uncertain.
//		0 if symbol is newly verified.
//		> 0 for duplicate symbols
func (s *Symbol) GetCount() int {
	return int(C.zbar_symbol_get_count(s.c_symbol))
}

// Retrieve the number of points in the location polygon.
// The location polygon defines the image area that the symbol was extracted from.
// Returns the number of points in the location polygon
// Note:
//    this is currently not a polygon, but the scan locations where the symbol was decoded
func (s *Symbol) GetLocSize() uint {
	return uint(C.zbar_symbol_get_loc_size(s.c_symbol))
}

// Retrieve location polygon x-coordinates.
// Points are specified by 0-based index.
// Returns:
//		the x-coordinate for a point in the location polygon.
//		-1 if index is out of range
func (s *Symbol) GetLocX(index uint) int {
	return int(C.zbar_symbol_get_loc_x(s.c_symbol, C.unsigned(index)))
}

// Retrieve location polygon y-coordinates.
// Points are specified by 0-based index.
// Returns:
//		the y-coordinate for a point in the location polygon.
//		-1 if index is out of range
func (s *Symbol) GetLocY(index uint) int {
	return int(C.zbar_symbol_get_loc_y(s.c_symbol, C.unsigned(index)))
}

// Iterate the set to which this symbol belongs (there can be only one).
// Returns:
//		the next symbol in the set, or
//		nil when no more results are available
func (s *Symbol) Next() *Symbol {
	newS := Symbol{}
	newS.c_symbol = C.zbar_symbol_next(s.c_symbol)
	if newS.c_symbol != nil {
		return &newS
	}
	return nil
}

// Retrieve components of a composite result.
// Returns:
//		the symbol set containing the components
//		nil if the symbol is already a physical symbol
func (s *Symbol) GetComponents() *SymbolSet {
	ss := SymbolSet{}
	ss.c_symbol_set = C.zbar_symbol_get_components(s.c_symbol)
	if ss.c_symbol_set != nil {
		return &ss
	}
	return nil
}

// Iterate components of a composite result.
// Returns:
//		the first physical component symbol of a composite result
//		nil if the symbol is already a physical symbol
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
