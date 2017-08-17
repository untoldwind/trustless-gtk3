package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include "closure.go.h"
import "C"
import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"sync"
	"unsafe"
)

var closures = struct {
	sync.RWMutex
	m map[*C.GClosure]reflect.Value
}{
	m: make(map[*C.GClosure]reflect.Value),
}

func RegisteredClosures() int {
	closures.RLock()
	defer closures.RUnlock()
	return len(closures.m)
}

// ClosureNew creates a new GClosure and adds its callback function
// to the internally-maintained map. It's exported for visibility to other
// gotk3 packages and shouldn't be used in application code.
func ClosureNew(f interface{}) (*C.GClosure, error) {
	// Create a reflect.Value from f.  This is called when the
	// returned GClosure runs.
	rf := reflect.ValueOf(f)

	// Closures can only be created from funcs.
	if rf.Type().Kind() != reflect.Func {
		return nil, errors.New("value is not a func")
	}

	c := C._g_closure_new()

	// Associate the GClosure with rf.  rf will be looked up in this
	// map by the closure when the closure runs.
	closures.Lock()
	closures.m[c] = rf
	closures.Unlock()

	C._g_closure_add_finalize_notifier(c)

	return c, nil
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
	rf := closures.m[closure]
	closures.RUnlock()

	// Get number of parameters passed in.  If user data was saved with the
	// closure context, increment the total number of parameters.
	nGLibParams := int(nParams)

	// Get number of parameters from the callback closure.  If this exceeds
	// the total number of marshaled parameters, a warning will be printed
	// to stderr, and the callback will not be run.
	nCbParams := rf.Type().NumIn()
	if nCbParams > nGLibParams {
		fmt.Fprintf(os.Stderr,
			"too many closure args: have %d, max allowed %d\n",
			nCbParams, nGLibParams)
		return
	}

	// Create a slice of reflect.Values as arguments to call the function.
	gValues := gValueSlice(params, nCbParams)
	args := make([]reflect.Value, 0, nCbParams)

	// Fill beginning of args, up to the minimum of the total number of callback
	// parameters and parameters from the glib runtime.
	for i := 0; i < nCbParams && i < nGLibParams; i++ {
		v := &Value{&gValues[i]}
		val := v.GoValue()
		rv := reflect.ValueOf(val)
		args = append(args, rv.Convert(rf.Type().In(i)))
	}

	// Call closure with args. If the callback returns one or more
	// values, save the GValue equivalent of the first.
	rv := rf.Call(args)
	if retValue != nil && len(rv) > 0 {
		if g := GValue(rv[0].Interface()); g != nil && g.GValue != nil {
			*retValue = *g.GValue
		}
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
