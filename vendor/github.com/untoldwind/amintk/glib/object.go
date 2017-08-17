package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include "object.go.h"
import "C"
import (
	"runtime"
	"unsafe"
)

type SignalHandle uint

// Object is a representation of GLib's GObject.
type Object struct {
	GObject *C.GObject
}

// WrapObject creates a new Object from a GObject pointer.
func WrapObject(p unsafe.Pointer) *Object {
	if p == nil {
		return nil
	}
	obj := &Object{GObject: (*C.GObject)(p)}

	if C.g_object_is_floating(C.gpointer(obj.GObject)) != 0 {
		C.g_object_ref_sink(C.gpointer(obj.GObject))
	} else {
		C.g_object_ref(C.gpointer(obj.GObject))
	}

	runtime.SetFinalizer(obj, finalizeObject)

	return obj
}

func finalizeObject(obj *Object) {
	C.g_object_unref(C.gpointer(obj.GObject))
}

func NewObject(gType Type) *Object {
	gObj := C._g_object_new(C.GType(gType))
	obj := WrapObject(unsafe.Pointer(gObj))
	C.g_object_unref(C.gpointer(gObj))
	return obj
}

func (v *Object) Native() unsafe.Pointer {
	return unsafe.Pointer(v.GObject)
}

// StopEmission is a wrapper around g_signal_stop_emission_by_name().
func (v *Object) StopEmission(s string) {
	if v == nil {
		return
	}
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.g_signal_stop_emission_by_name((C.gpointer)(v.GObject),
		(*C.gchar)(cstr))
}

// GetPropertyType returns the Type of a property of the underlying GObject.
// If the property is missing it will return TYPE_INVALID and an error.
func (v *Object) GetPropertyType(name string) Type {
	if v == nil {
		return TYPE_INVALID
	}
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	paramSpec := C.g_object_class_find_property(C._g_object_get_class(v.GObject), (*C.gchar)(cstr))
	if paramSpec == nil {
		return TYPE_INVALID
	}
	return Type(paramSpec.value_type)
}

// GetProperty is a wrapper around g_object_get_property().
func (v *Object) GetProperty(name string) interface{} {
	if v == nil {
		return nil
	}
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	t := v.GetPropertyType(name)

	p := ValueInit(t)
	if p == nil {
		return nil
	}
	C.g_object_get_property(v.GObject, (*C.gchar)(cstr), p.GValue)
	return p.GoValue()
}

// SetProperty is a wrapper around g_object_set_property().
func (v *Object) SetProperty(name string, value interface{}) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	if _, ok := value.(Object); ok {
		value = value.(Object).GObject
	}

	if p := GValue(value); p != nil {
		C.g_object_set_property(v.GObject, (*C.gchar)(cstr), p.GValue)
	}
}

func (v *Object) connectClosure(after bool, detailedSignal string, f interface{}) (SignalHandle, error) {
	cstr := C.CString(detailedSignal)
	defer C.free(unsafe.Pointer(cstr))

	closure, err := ClosureNew(f)
	if err != nil {
		return 0, err
	}
	c := C.g_signal_connect_closure(C.gpointer(v.GObject), (*C.gchar)(cstr), closure, gbool(after))
	return SignalHandle(c), nil
}

func (v *Object) Connect(detailedSignal string, f func()) (SignalHandle, error) {
	return v.connectClosure(false, detailedSignal, f)
}

func (v *Object) ConnectAfter(detailedSignal string, f func()) (SignalHandle, error) {
	return v.connectClosure(true, detailedSignal, f)
}
