package glib

import "unsafe"

// InitiallyUnowned is a representation of GLib's GInitiallyUnowned.
type InitiallyUnowned struct {
	// This must be a pointer so copies of the ref-sinked object
	// do not outlive the original object, causing an unref
	// finalizer to prematurely run.
	*Object
}

func WrapInitiallyUnowned(p unsafe.Pointer) *InitiallyUnowned {
	if obj := WrapObject(p); obj != nil {
		return &InitiallyUnowned{Object: obj}
	}
	return nil
}
