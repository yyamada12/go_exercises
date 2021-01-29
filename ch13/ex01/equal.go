// Package equal provides a deep equivalence relation for arbitrary values.
package equal

import (
	"math"
	"reflect"
	"unsafe"
)

var numTypes = []reflect.Kind{reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64}

func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() && !(contains(numTypes, x.Kind()) && contains(numTypes, y.Kind())) {
		return false
	}

	// cycle check
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true // identical references
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()

	case reflect.String:
		return x.String() == y.String()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return compareNumber(float64(x.Int()), y)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return compareNumber(float64(x.Uint()), y)

	case reflect.Float32, reflect.Float64:
		return compareNumber(x.Float(), y)

	case reflect.Complex64, reflect.Complex128:
		xc := x.Complex()
		yc := y.Complex()
		return floatEqual(real(xc), real(yc)) && floatEqual(imag(xc), imag(yc))

	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()

	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)

	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if !equal(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true

	case reflect.Map:
		if x.Len() != y.Len() {
			return false
		}
		for _, k := range x.MapKeys() {
			if !equal(x.MapIndex(k), y.MapIndex(k), seen) {
				return false
			}
		}
		return true
	}
	panic("unreachable")
}

// Equal reports whether x and y are deeply equal.
//
// Map keys are always compared with ==, not deeply.
// (This matters for keys containing pointers or interfaces.)
//!+comparison
func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}

func compareNumber(x float64, y reflect.Value) bool {
	switch y.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return floatEqual(x, float64(y.Int()))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return floatEqual(x, float64(y.Uint()))

	case reflect.Float32, reflect.Float64:
		return floatEqual(x, y.Float())
	}
	panic("unreachable")
}

func floatEqual(x, y float64) bool {
	return math.Abs(x-y) < 0.000000001
}

func contains(s []reflect.Kind, e reflect.Kind) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
