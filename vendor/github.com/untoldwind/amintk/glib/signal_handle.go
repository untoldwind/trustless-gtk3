package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

type SignalHandle struct {
	handle C.gulong
	object *Object
}

// Block is a wrapper around g_signal_handler_block().
func (v *SignalHandle) Block() {
	if v == nil {
		return
	}
	C.g_signal_handler_block(C.gpointer(v.object.GObject), v.handle)
}

// Unblock is a wrapper around g_signal_handler_unblock().
func (v *SignalHandle) Unblock() {
	if v == nil {
		return
	}
	C.g_signal_handler_unblock(C.gpointer(v.object.GObject), v.handle)
}

// Disconnect is a wrapper around g_signal_handler_disconnect().
func (v *SignalHandle) Disconnect() {
	if v == nil {
		return
	}
	C.g_signal_handler_disconnect(C.gpointer(v.object.GObject), v.handle)
}

type SignalHandles []*SignalHandle

func (v *SignalHandles) Add(handle *SignalHandle) {
	*v = append(*v, handle)
}

func (v SignalHandles) BlockAll() {
	for _, handle := range v {
		handle.Block()
	}
}

func (v SignalHandles) UnblockAll() {
	for _, handle := range v {
		handle.Unblock()
	}
}

func (v SignalHandles) DisconnectAll() {
	for _, handle := range v {
		handle.Disconnect()
	}
}
