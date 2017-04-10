package gtkextra

import "github.com/gotk3/gotk3/glib"

type HandleRef struct {
	obj    *glib.Object
	handle glib.SignalHandle
}

type HandleRefs []HandleRef

func (r *HandleRefs) SafeConnect(obj *glib.Object, signal string, handler func()) error {
	handle, err := obj.Connect(signal, handler)
	if err != nil {
		return err
	}
	*r = append(*r, HandleRef{
		obj:    obj,
		handle: handle,
	})
	return nil
}

func (r *HandleRefs) Cleanup() {
	for _, ref := range *r {
		ref.obj.HandlerDisconnect(ref.handle)
	}

	*r = nil
}
