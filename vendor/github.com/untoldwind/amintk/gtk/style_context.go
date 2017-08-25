package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"github.com/untoldwind/amintk/glib"
)

type StyleProviderPriority int

const (
	StyleProviderPriorityFallback    StyleProviderPriority = C.GTK_STYLE_PROVIDER_PRIORITY_FALLBACK
	StyleProviderPriorityTheme                             = C.GTK_STYLE_PROVIDER_PRIORITY_THEME
	StyleProviderPrioritySettings                          = C.GTK_STYLE_PROVIDER_PRIORITY_SETTINGS
	StyleProviderPriorityApplication                       = C.GTK_STYLE_PROVIDER_PRIORITY_APPLICATION
	StyleProviderPriorityUser                              = C.GTK_STYLE_PROVIDER_PRIORITY_USER
)

// StateFlags is a representation of GTK's GtkStateFlags.
type StateFlags int

const (
	StateFlagsNormal       StateFlags = C.GTK_STATE_FLAG_NORMAL
	StateFlagsActive       StateFlags = C.GTK_STATE_FLAG_ACTIVE
	StateFLagsPreflight    StateFlags = C.GTK_STATE_FLAG_PRELIGHT
	StateFlagsSelected     StateFlags = C.GTK_STATE_FLAG_SELECTED
	StateFlagsInsensitive  StateFlags = C.GTK_STATE_FLAG_INSENSITIVE
	StateFlagsInconsistent StateFlags = C.GTK_STATE_FLAG_INCONSISTENT
	StateFlagsFocused      StateFlags = C.GTK_STATE_FLAG_FOCUSED
	StateFlagsBackdrop     StateFlags = C.GTK_STATE_FLAG_BACKDROP
)

// StyleContext is a representation of GTK's GtkStyleContext.
type StyleContext struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkStyleContext.
func (v *StyleContext) native() *C.GtkStyleContext {
	if v == nil {
		return nil
	}
	return (*C.GtkStyleContext)(v.Native())
}

func wrapStyleContext(obj *glib.Object) *StyleContext {
	return &StyleContext{Object: obj}
}

// AddProvider is a wrapper around gtk_style_context_add_provider().
func (v *StyleContext) AddProvider(provider IStyleProvider, prio uint) {
	C.gtk_style_context_add_provider(v.native(), provider.toStyleProvider(), C.guint(prio))
}

// SetState is a wrapper around gtk_style_context_set_state().
func (v *StyleContext) SetState(state StateFlags) {
	C.gtk_style_context_set_state(v.native(), C.GtkStateFlags(state))
}
