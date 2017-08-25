package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include "closure.go.h"
import "C"
import (
	"reflect"
	"sync"
	"unsafe"
)

var closures = struct {
	sync.RWMutex
	m map[*C.GClosure]Callback
}{
	m: make(map[*C.GClosure]Callback),
}

func RegisteredClosures() int {
	closures.RLock()
	defer closures.RUnlock()
	return len(closures.m)
}

// ClosureNew creates a new GClosure and adds its callback function
// to the internally-maintained map. It's exported for visibility to other
// gotk3 packages and shouldn't be used in application code.
func ClosureNew(callback Callback) *C.GClosure {
	c := C._g_closure_new()

	// Associate the GClosure with rf.  rf will be looked up in this
	// map by the closure when the closure runs.
	closures.Lock()
	closures.m[c] = callback
	closures.Unlock()

	C._g_closure_add_finalize_notifier(c)

	return c
}

// removeClosure removes a closure from the internal closures map.  This is
// needed to prevent a leak where Go code can access the closure context
// (along with rf and userdata) even after an object has been destroyed and
// the GClosure is invalidated and will never run.
//
//export callbackRemoveClosure
func callbackRemoveClosure(_ C.gpointer, closure *C.GClosure) {
	closures.Lock()
	delete(closures.m, closure)
	closures.Unlock()
}

// goMarshal is called by the GLib runtime when a closure needs to be invoked.
// The closure will be invoked with as many arguments as it can take, from 0 to
// the full amount provided by the call. If the closure asks for more parameters
// than there are to give, a warning is printed to stderr and the closure is
// not run.
//
//export callbackGoMarshal
func callbackGoMarshal(closure *C.GClosure, retValue *C.GValue,
	nParams C.guint, params *C.GValue,
	invocationHint C.gpointer, marshalData *C.GValue) {

	// Get the context associated with this callback closure.
	closures.RLock()
	callback := closures.m[closure]
	closures.RUnlock()

	gValues := gValueSlice(params, int(nParams))
	args := make([]Value, len(gValues))
	for i, gValue := range gValues {
		args[i] = Value{&gValue}
	}

	ret := callback.Call(args)

	if ret != nil && retValue != nil {
		*retValue = *ret.GValue
	}
}

// gValueSlice converts a C array of GValues to a Go slice.
func gValueSlice(values *C.GValue, nValues int) (slice []C.GValue) {
	header := (*reflect.SliceHeader)((unsafe.Pointer(&slice)))
	header.Cap = nValues
	header.Len = nValues
	header.Data = uintptr(unsafe.Pointer(values))
	return
}
