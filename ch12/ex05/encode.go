package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// encode writes to buf an S-expression representation of v.
func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(',')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(']')

	case reflect.Struct: // ((name value) ...)
		buf.WriteString("{\n")
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteString(",\n")
			}
			fmt.Fprintf(buf, "%q: ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
		}
		buf.WriteString("\n}")

	case reflect.Map: // ((key value) ...)
		switch v.Type().Key().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if err := encodeNumKeyMap(buf, v); err != nil {
				return err
			}
		case reflect.String:
			if err := encodeStringKeyMap(buf, v); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported type: %s", v.Type())
		}

	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func encodeNumKeyMap(buf *bytes.Buffer, v reflect.Value) error {
	buf.WriteString("{\n")
	for i, key := range v.MapKeys() {

		if i > 0 {
			buf.WriteString(",\n")
		}
		buf.WriteByte('"')
		if err := encode(buf, key); err != nil {
			return err
		}
		buf.WriteString(`": `)
		if err := encode(buf, v.MapIndex(key)); err != nil {
			return err
		}

	}
	buf.WriteString("\n}")
	return nil
}

func encodeStringKeyMap(buf *bytes.Buffer, v reflect.Value) error {
	buf.WriteString("{\n")
	for i, key := range v.MapKeys() {
		if i > 0 {
			buf.WriteString(",\n")
		}
		fmt.Fprintf(buf, `%q: `, key.String())
		if err := encode(buf, v.MapIndex(key)); err != nil {
			return err
		}
	}
	buf.WriteString("\n}")
	return nil
}
