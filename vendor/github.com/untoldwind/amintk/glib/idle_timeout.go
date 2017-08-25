package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"
import (
	"errors"
)

type SourceHandle uint

// IdleAdd adds an idle source to the default main event loop
// context.  After running once, the source func will be removed
// from the main event loop, unless f returns a single bool true.
//
// This function will cause a panic when f eventually runs if the
// types of args do not match those of f.
func IdleAdd(f func()) (SourceHandle, error) {
	// Create an idle source func to be added to the main loop context.
	idleSrc := C.g_idle_source_new()
	return sourceAttach(idleSrc, f)
}

// TimeoutAdd adds an timeout source to the default main event loop
// context.  After running once, the source func will be removed
// from the main event loop, unless f returns a single bool true.
//
// This function will cause a panic when f eventually runs if the
// types of args do not match those of f.
// timeout is in milliseconds
func TimeoutAdd(timeout uint, f func()) (SourceHandle, error) {
	// Create a timeout source func to be added to the main loop context.
	timeoutSrc := C.g_timeout_source_new(C.guint(timeout))
	return sourceAttach(timeoutSrc, f)
}

// sourceAttach attaches a source to the default main loop context.
func sourceAttach(src *C.struct__GSource, f func()) (SourceHandle, error) {
	if src == nil {
		return 0, errors.New("Source is nil")
	}

	// Create a new GClosure from f that invalidates itself when
	// f returns false.  The error is ignored here, as this will
	// always be a function.
	var closure *C.GClosure
	closure = ClosureNew(CallbackVoidVoid(func() {
		f()

		C.g_closure_invalidate(closure)
		C.g_source_destroy(src)
	}))

	// Set closure to run as a callback when the idle source runs.
	C.g_source_set_closure(src, closure)

	// Attach the idle source func to the default main event loop
	// context.
	cid := C.g_source_attach(src, nil)
	return SourceHandle(cid), nil
}
