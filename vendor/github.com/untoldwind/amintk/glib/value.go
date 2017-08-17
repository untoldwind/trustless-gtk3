package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include "value.go.h"
import "C"
import (
	"reflect"
	"runtime"
	"unsafe"
)

// Value is a representation of GLib's GValue.
//
// Don't allocate Values on the stack or heap manually as they may not
// be properly unset when going out of scope. Instead, use ValueAlloc(),
// which will set the runtime finalizer to unset the Value after it has
// left scope.
type Value struct {
	GValue *C.GValue
}

// ValueAlloc allocates a Value and sets a runtime finalizer to call
// g_value_unset() on the underlying GValue after leaving scope.
// ValueAlloc() returns a non-nil error if the allocation failed.
func ValueAlloc() *Value {
	c := C._g_value_alloc()
	if c == nil {
		return nil
	}

	v := &Value{c}

	//An allocated GValue is not guaranteed to hold a value that can be unset
	//We need to double check before unsetting, to prevent:
	//`g_value_unset: assertion 'G_IS_VALUE (value)' failed`
	runtime.SetFinalizer(v, func(f *Value) {
		if t, _ := f.Type(); t == TYPE_INVALID || t == TYPE_NONE {
			C.g_free(C.gpointer(f.GValue))
			return
		}

		f.unset()
	})

	return v
}

// ValueInit is a wrapper around g_value_init() and allocates and
// initializes a new Value with the Type t.  A runtime finalizer is set
// to call g_value_unset() on the underlying GValue after leaving scope.
// ValueInit() returns a non-nil error if the allocation failed.
func ValueInit(t Type) *Value {
	c := C._g_value_init(C.GType(t))
	if c == nil {
		return nil
	}

	v := &Value{c}

	runtime.SetFinalizer(v, (*Value).unset)
	return v
}

// GValue converts a Go type to a comparable GValue.  GValue()
// returns a non-nil error if the conversion was unsuccessful.
func GValue(v interface{}) *Value {
	if v == nil {
		val := ValueInit(TYPE_POINTER)
		val.SetPointer(unsafe.Pointer(nil))
		return val
	}

	switch e := v.(type) {
	case bool:
		val := ValueInit(TYPE_BOOLEAN)
		val.SetBool(e)
		return val

	case int8:
		val := ValueInit(TYPE_CHAR)
		val.SetSChar(e)
		return val

	case int64:
		val := ValueInit(TYPE_INT64)
		val.SetInt64(e)
		return val

	case int:
		val := ValueInit(TYPE_INT)
		val.SetInt(e)
		return val

	case uint8:
		val := ValueInit(TYPE_UCHAR)
		val.SetUChar(e)
		return val

	case uint64:
		val := ValueInit(TYPE_UINT64)
		val.SetUInt64(e)
		return val

	case uint:
		val := ValueInit(TYPE_UINT)
		val.SetUInt(e)
		return val

	case float32:
		val := ValueInit(TYPE_FLOAT)
		val.SetFloat(e)
		return val

	case float64:
		val := ValueInit(TYPE_DOUBLE)
		val.SetDouble(e)
		return val

	case string:
		val := ValueInit(TYPE_STRING)
		val.SetString(e)
		return val

	case *Object:
		val := ValueInit(TYPE_OBJECT)
		val.SetInstance(unsafe.Pointer(e.GObject))
		return val

	default:
		/* Try this since above doesn't catch constants under other types */
		rval := reflect.ValueOf(v)
		switch rval.Kind() {
		case reflect.Int8:
			val := ValueInit(TYPE_CHAR)
			val.SetSChar(int8(rval.Int()))
			return val

		case reflect.Int16:
			val := ValueInit(TYPE_INT64)
			val.SetInt64(rval.Int())
			return val

		case reflect.Int32:
			val := ValueInit(TYPE_INT64)
			val.SetInt64(rval.Int())
			return val

		case reflect.Int64:
			val := ValueInit(TYPE_INT64)
			val.SetInt64(rval.Int())
			return val

		case reflect.Int:
			val := ValueInit(TYPE_INT)
			val.SetInt(int(rval.Int()))
			return val

		case reflect.Uintptr, reflect.Ptr:
			val := ValueInit(TYPE_POINTER)
			val.SetPointer(unsafe.Pointer(rval.Pointer()))
			return val
		}
	}

	return nil
}

// Type is a wrapper around the G_VALUE_HOLDS_GTYPE() macro and
// the g_value_get_gtype() function.  GetType() returns TYPE_INVALID if v
// does not hold a Type, or otherwise returns the Type of v.
func (v *Value) Type() (actual Type, fundamental Type) {
	if v == nil || C._g_is_value(v.GValue) == 0 {
		return TYPE_INVALID, TYPE_INVALID
	}
	cActual := C._g_value_type(v.GValue)
	cFundamental := C._g_value_fundamental(cActual)
	return Type(cActual), Type(cFundamental)
}

// GoValue converts a Value to comparable Go type.  GoValue()
// returns a non-nil error if the conversion was unsuccessful.  The
// returned interface{} must be type asserted as the actual Go
// representation of the Value.
//
// This function is a wrapper around the many g_value_get_*()
// functions, depending on the type of the Value.
func (v *Value) GoValue() interface{} {
	if v == nil {
		return nil
	}
	f := gValueMarshalers.lookup(v)
	if f == nil {
		return nil
	}

	//No need to add finalizer because it is already done by ValueAlloc and ValueInit
	return f(unsafe.Pointer(v.GValue))
}

// SetBool is a wrapper around g_value_set_boolean().
func (v *Value) SetBool(val bool) {
	if val {
		C.g_value_set_boolean(v.GValue, C.gboolean(1))
	} else {
		C.g_value_set_boolean(v.GValue, C.gboolean(0))
	}
}

// SetSChar is a wrapper around g_value_set_schar().
func (v *Value) SetSChar(val int8) {
	C.g_value_set_schar(v.GValue, C.gint8(val))
}

// SetInt64 is a wrapper around g_value_set_int64().
func (v *Value) SetInt64(val int64) {
	C.g_value_set_int64(v.GValue, C.gint64(val))
}

// SetInt is a wrapper around g_value_set_int().
func (v *Value) SetInt(val int) {
	C.g_value_set_int(v.GValue, C.gint(val))
}

// SetUChar is a wrapper around g_value_set_uchar().
func (v *Value) SetUChar(val uint8) {
	C.g_value_set_uchar(v.GValue, C.guchar(val))
}

// SetUInt64 is a wrapper around g_value_set_uint64().
func (v *Value) SetUInt64(val uint64) {
	C.g_value_set_uint64(v.GValue, C.guint64(val))
}

// SetUInt is a wrapper around g_value_set_uint().
func (v *Value) SetUInt(val uint) {
	C.g_value_set_uint(v.GValue, C.guint(val))
}

// SetFloat is a wrapper around g_value_set_float().
func (v *Value) SetFloat(val float32) {
	C.g_value_set_float(v.GValue, C.gfloat(val))
}

// SetDouble is a wrapper around g_value_set_double().
func (v *Value) SetDouble(val float64) {
	C.g_value_set_double(v.GValue, C.gdouble(val))
}

// SetString is a wrapper around g_value_set_string().
func (v *Value) SetString(val string) {
	cstr := C.CString(val)
	defer C.free(unsafe.Pointer(cstr))
	C.g_value_set_string(v.GValue, (*C.gchar)(cstr))
}

// SetInstance is a wrapper around g_value_set_instance().
func (v *Value) SetInstance(instance unsafe.Pointer) {
	C.g_value_set_instance(v.GValue, C.gpointer(instance))
}

// SetPointer is a wrapper around g_value_set_pointer().
func (v *Value) SetPointer(p unsafe.Pointer) {
	C.g_value_set_pointer(v.GValue, C.gpointer(p))
}

// GetPointer is a wrapper around g_value_get_pointer().
func (v *Value) GetPointer() unsafe.Pointer {
	return unsafe.Pointer(C.g_value_get_pointer(v.GValue))
}

// GetString is a wrapper around g_value_get_string().
func (v *Value) GetString() string {
	c := C.g_value_get_string(v.GValue)
	if c == nil {
		return ""
	}
	return C.GoString((*C.char)(c))
}

func (v *Value) unset() {
	if v == nil {
		return
	}
	C.g_value_unset(v.GValue)
}
