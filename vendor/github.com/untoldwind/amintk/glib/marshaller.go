package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <glib.h>
// #include <glib-object.h>
import "C"
import "unsafe"

// GValueMarshaler is a marshal function to convert a GValue into an
// appropiate Go type.  The uintptr parameter is a *C.GValue.
type GValueMarshaler func(unsafe.Pointer) interface{}

type marshalMap map[Type]GValueMarshaler

// gValueMarshalers is a map of Glib types to functions to marshal a
// GValue to a native Go type.
var gValueMarshalers = marshalMap{
	TYPE_INVALID:   marshalInvalid,
	TYPE_NONE:      marshalNone,
	TYPE_INTERFACE: marshalInterface,
	TYPE_CHAR:      marshalChar,
	TYPE_UCHAR:     marshalUchar,
	TYPE_BOOLEAN:   marshalBoolean,
	TYPE_INT:       marshalInt,
	TYPE_LONG:      marshalLong,
	TYPE_ENUM:      marshalEnum,
	TYPE_INT64:     marshalInt64,
	TYPE_UINT:      marshalUint,
	TYPE_ULONG:     marshalUlong,
	TYPE_FLAGS:     marshalFlags,
	TYPE_UINT64:    marshalUint64,
	TYPE_FLOAT:     marshalFloat,
	TYPE_DOUBLE:    marshalDouble,
	TYPE_STRING:    marshalString,
	TYPE_POINTER:   marshalPointer,
	TYPE_BOXED:     marshalBoxed,
	TYPE_OBJECT:    marshalObject,
	TYPE_VARIANT:   marshalVariant,
}

// TypeMarshaler represents an actual type and it's associated marshaler.
type TypeMarshaler struct {
	T Type
	F GValueMarshaler
}

func (m marshalMap) register(tm []TypeMarshaler) {
	for i := range tm {
		m[tm[i].T] = tm[i].F
	}
}

func (m marshalMap) lookup(v *Value) GValueMarshaler {
	actual, fundamental := v.Type()

	if f, ok := m[actual]; ok {
		return f
	}
	if f, ok := m[fundamental]; ok {
		return f
	}
	return nil
}

func marshalInvalid(unsafe.Pointer) interface{} {
	return nil
}

func marshalNone(unsafe.Pointer) interface{} {
	return nil
}

func marshalInterface(unsafe.Pointer) interface{} {
	return nil
}

func marshalChar(p unsafe.Pointer) interface{} {
	c := C.g_value_get_schar((*C.GValue)(p))
	return int8(c)
}

func marshalUchar(p unsafe.Pointer) interface{} {
	c := C.g_value_get_uchar((*C.GValue)(p))
	return uint8(c)
}

func marshalBoolean(p unsafe.Pointer) interface{} {
	c := C.g_value_get_boolean((*C.GValue)(p))
	return c != 0
}

func marshalInt(p unsafe.Pointer) interface{} {
	c := C.g_value_get_int((*C.GValue)(p))
	return int(c)
}

func marshalLong(p unsafe.Pointer) interface{} {
	c := C.g_value_get_long((*C.GValue)(p))
	return int(c)
}

func marshalEnum(p unsafe.Pointer) interface{} {
	c := C.g_value_get_enum((*C.GValue)(p))
	return int(c)
}

func marshalInt64(p unsafe.Pointer) interface{} {
	c := C.g_value_get_int64((*C.GValue)(p))
	return int64(c)
}

func marshalUint(p unsafe.Pointer) interface{} {
	c := C.g_value_get_uint((*C.GValue)(p))
	return uint(c)
}

func marshalUlong(p unsafe.Pointer) interface{} {
	c := C.g_value_get_ulong((*C.GValue)(p))
	return uint(c)
}

func marshalFlags(p unsafe.Pointer) interface{} {
	c := C.g_value_get_flags((*C.GValue)(p))
	return uint(c)
}

func marshalUint64(p unsafe.Pointer) interface{} {
	c := C.g_value_get_uint64((*C.GValue)(p))
	return uint64(c)
}

func marshalFloat(p unsafe.Pointer) interface{} {
	c := C.g_value_get_float((*C.GValue)(p))
	return float32(c)
}

func marshalDouble(p unsafe.Pointer) interface{} {
	c := C.g_value_get_double((*C.GValue)(p))
	return float64(c)
}

func marshalString(p unsafe.Pointer) interface{} {
	c := C.g_value_get_string((*C.GValue)(p))
	return C.GoString((*C.char)(c))
}

func marshalBoxed(p unsafe.Pointer) interface{} {
	c := C.g_value_get_boxed((*C.GValue)(p))
	return uintptr(unsafe.Pointer(c))
}

func marshalPointer(p unsafe.Pointer) interface{} {
	c := C.g_value_get_pointer((*C.GValue)(p))
	return unsafe.Pointer(c)
}

func marshalObject(p unsafe.Pointer) interface{} {
	c := C.g_value_get_object((*C.GValue)(p))
	return WrapObject(unsafe.Pointer(c))
}

func marshalVariant(p unsafe.Pointer) interface{} {
	return nil
}
