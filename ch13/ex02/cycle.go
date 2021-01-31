package cycle

import (
	"reflect"
	"unsafe"
)

func hasCycle(v reflect.Value, seen map[target]bool) bool {
	// cycle check
	if v.CanAddr() {
		vptr := unsafe.Pointer(v.UnsafeAddr())
		t := target{vptr, v.Type()}
		if seen[t] {
			return true // already seen
		}
		seen[t] = true
	}

	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return hasCycle(v.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if hasCycle(v.Index(i), copySeen(seen)) {
				return true
			}
		}
		return false

	case reflect.Struct:
		for i, n := 0, v.NumField(); i < n; i++ {
			if hasCycle(v.Field(i), copySeen(seen)) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range v.MapKeys() {
			if hasCycle(v.MapIndex(k), copySeen(seen)) {
				return true
			}
		}
		return false
	}
	return false
}

// Equal reports whether x and y are deeply equal.
//
// Map keys are always compared with ==, not deeply.
// (This matters for keys containing pointers or interfaces.)
func HasCycle(v interface{}) bool {
	seen := make(map[target]bool)
	return hasCycle(reflect.ValueOf(v), seen)
}

type target struct {
	p unsafe.Pointer
	t reflect.Type
}

func copySeen(m map[target]bool) map[target]bool {
	cp := make(map[target]bool)
	for k, v := range m {
		cp[k] = v
	}
	return cp
}
