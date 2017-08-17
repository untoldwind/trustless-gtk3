package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
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
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkScrolledWindow)(p)
}

// ScrolledWindowNew() is a wrapper around gtk_scrolled_window_new().
func ScrolledWindowNew(hadjustment, vadjustment *Adjustment) *ScrolledWindow {
	c := C.gtk_scrolled_window_new(hadjustment.native(),
		vadjustment.native())
	return wrapScrolledWindow(glib.WrapObject(unsafe.Pointer(c)))
}

func wrapScrolledWindow(obj *glib.Object) *ScrolledWindow {
	return &ScrolledWindow{Bin{Container{Widget{glib.InitiallyUnowned{Object: obj}}}}}
}

// SetPolicy() is a wrapper around gtk_scrolled_window_set_policy().
func (v *ScrolledWindow) SetPolicy(hScrollbarPolicy, vScrollbarPolicy PolicyType) {
	C.gtk_scrolled_window_set_policy(v.native(),
		C.GtkPolicyType(hScrollbarPolicy),
		C.GtkPolicyType(vScrollbarPolicy))
}
