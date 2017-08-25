package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// PolicyType is a representation of GTK's GtkPolicyType.
type PolicyType int

const (
	PolicyTypeAlways    PolicyType = C.GTK_POLICY_ALWAYS
	PolicyTypeAutomatic PolicyType = C.GTK_POLICY_AUTOMATIC
	PolicyTypeNever     PolicyType = C.GTK_POLICY_NEVER
)

// ScrolledWindow is a representation of GTK's GtkScrolledWindow.
type ScrolledWindow struct {
	Bin
}

// native returns a pointer to the underlying GtkScrolledWindow.
func (v *ScrolledWindow) native() *C.GtkScrolledWindow {
	if v == nil {
		return nil
	}
	return (*C.GtkScrolledWindow)(v.Native())
}

// ScrolledWindowNew is a wrapper around gtk_scrolled_window_new().
func ScrolledWindowNew(hadjustment, vadjustment *Adjustment) *ScrolledWindow {
	c := C.gtk_scrolled_window_new(hadjustment.native(),
		vadjustment.native())
	return wrapScrolledWindow(unsafe.Pointer(c))
}

func wrapScrolledWindow(p unsafe.Pointer) *ScrolledWindow {
	if bin := wrapBin(p); bin != nil {
		return &ScrolledWindow{Bin: *bin}
	}
	return nil
}

// SetPolicy is a wrapper around gtk_scrolled_window_set_policy().
func (v *ScrolledWindow) SetPolicy(hScrollbarPolicy, vScrollbarPolicy PolicyType) {
	C.gtk_scrolled_window_set_policy(v.native(),
		C.GtkPolicyType(hScrollbarPolicy),
		C.GtkPolicyType(vScrollbarPolicy))
}
