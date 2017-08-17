package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/gdk"
	"github.com/untoldwind/amintk/glib"
)

// Align is a representation of GTK's GtkAlign.
type Align int

const (
	AlignFill   Align = C.GTK_ALIGN_FILL
	AlignStart  Align = C.GTK_ALIGN_START
	AlignEnd    Align = C.GTK_ALIGN_END
	AlignCenter Align = C.GTK_ALIGN_CENTER
)

// Widget is a representation of GTK's GtkWidget.
type Widget struct {
	glib.InitiallyUnowned
}

type IWidget interface {
	toWidget() *C.GtkWidget
}

// native returns a pointer to the underlying GtkWidget.
func (v *Widget) native() *C.GtkWidget {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkWidget)(p)
}

func (v *Widget) toWidget() *C.GtkWidget {
	if v == nil {
		return nil
	}
	return v.native()
}

func wrapWidget(obj *glib.Object) *Widget {
	return &Widget{glib.InitiallyUnowned{Object: obj}}
}

// Show is a wrapper around gtk_widget_show().
func (v *Widget) Show() {
	C.gtk_widget_show(v.native())
}

// Hide is a wrapper around gtk_widget_hide().
func (v *Widget) Hide() {
	C.gtk_widget_hide(v.native())
}

// ShowNow is a wrapper around gtk_widget_show_now().
func (v *Widget) ShowNow() {
	C.gtk_widget_show_now(v.native())
}

// ShowAll is a wrapper around gtk_widget_show_all().
func (v *Widget) ShowAll() {
	C.gtk_widget_show_all(v.native())
}

// Destroy is a wrapper around gtk_widget_destroy().
func (v *Widget) Destroy() {
	C.gtk_widget_destroy(v.native())
}

// GetToplevel is a wrapper around gtk_widget_get_toplevel().
func (v *Widget) GetToplevel() *Widget {
	c := C.gtk_widget_get_toplevel(v.native())
	return wrapWidget(glib.WrapObject(unsafe.Pointer(c)))
}

// IsToplevel is a wrapper around gtk_widget_is_toplevel().
func (v *Widget) IsToplevel() bool {
	c := C.gtk_widget_is_toplevel(v.native())
	return gobool(c)
}

// SetNoShowAll is a wrapper around gtk_widget_set_no_show_all().
func (v *Widget) SetNoShowAll(noShowAll bool) {
	C.gtk_widget_set_no_show_all(v.native(), gbool(noShowAll))
}

// GetNoShowAll is a wrapper around gtk_widget_get_no_show_all().
func (v *Widget) GetNoShowAll() bool {
	c := C.gtk_widget_get_no_show_all(v.native())
	return gobool(c)
}

// GetTooltipText is a wrapper around gtk_widget_get_tooltip_text().
// A non-nil error is returned in the case that
// gtk_widget_get_tooltip_text returns NULL to differentiate between NULL
// and an empty string.
func (v *Widget) GetTooltipText() string {
	c := C.gtk_widget_get_tooltip_text(v.native())
	return C.GoString((*C.char)(c))
}

// SetTooltipText is a wrapper around gtk_widget_set_tooltip_text().
func (v *Widget) SetTooltipText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_widget_set_tooltip_text(v.native(), (*C.gchar)(cstr))
}

// GetHAlign is a wrapper around gtk_widget_get_halign().
func (v *Widget) GetHAlign() Align {
	c := C.gtk_widget_get_halign(v.native())
	return Align(c)
}

// SetHAlign is a wrapper around gtk_widget_set_halign().
func (v *Widget) SetHAlign(align Align) {
	C.gtk_widget_set_halign(v.native(), C.GtkAlign(align))
}

// GetVAlign is a wrapper around gtk_widget_get_valign().
func (v *Widget) GetVAlign() Align {
	c := C.gtk_widget_get_valign(v.native())
	return Align(c)
}

// SetVAlign is a wrapper around gtk_widget_set_valign().
func (v *Widget) SetVAlign(align Align) {
	C.gtk_widget_set_valign(v.native(), C.GtkAlign(align))
}

// GetMarginTop is a wrapper around gtk_widget_get_margin_top().
func (v *Widget) GetMarginTop() int {
	c := C.gtk_widget_get_margin_top(v.native())
	return int(c)
}

// SetMarginTop is a wrapper around gtk_widget_set_margin_top().
func (v *Widget) SetMarginTop(margin int) {
	C.gtk_widget_set_margin_top(v.native(), C.gint(margin))
}

// GetMarginBottom is a wrapper around gtk_widget_get_margin_bottom().
func (v *Widget) GetMarginBottom() int {
	c := C.gtk_widget_get_margin_bottom(v.native())
	return int(c)
}

// SetMarginBottom is a wrapper around gtk_widget_set_margin_bottom().
func (v *Widget) SetMarginBottom(margin int) {
	C.gtk_widget_set_margin_bottom(v.native(), C.gint(margin))
}

// GetHExpand is a wrapper around gtk_widget_get_hexpand().
func (v *Widget) GetHExpand() bool {
	c := C.gtk_widget_get_hexpand(v.native())
	return gobool(c)
}

// SetHExpand is a wrapper around gtk_widget_set_hexpand().
func (v *Widget) SetHExpand(expand bool) {
	C.gtk_widget_set_hexpand(v.native(), gbool(expand))
}

// GetVExpand is a wrapper around gtk_widget_get_vexpand().
func (v *Widget) GetVExpand() bool {
	c := C.gtk_widget_get_vexpand(v.native())
	return gobool(c)
}

// SetVExpand is a wrapper around gtk_widget_set_vexpand().
func (v *Widget) SetVExpand(expand bool) {
	C.gtk_widget_set_vexpand(v.native(), gbool(expand))
}

func (v *Widget) SetMarginStart(margin int) {
	C.gtk_widget_set_margin_start(v.native(), C.gint(margin))
}

func (v *Widget) GetMarginStart() int {
	c := C.gtk_widget_get_margin_start(v.native())
	return int(c)
}

func (v *Widget) SetMarginEnd(margin int) {
	C.gtk_widget_set_margin_end(v.native(), C.gint(margin))
}

func (v *Widget) GetMarginEnd() int {
	c := C.gtk_widget_get_margin_end(v.native())
	return int(c)
}

// SetSensitive is a wrapper around gtk_widget_set_sensitive().
func (v *Widget) SetSensitive(sensitive bool) {
	C.gtk_widget_set_sensitive(v.native(), gbool(sensitive))
}

// GetVisible is a wrapper around gtk_widget_get_visible().
func (v *Widget) GetVisible() bool {
	c := C.gtk_widget_get_visible(v.native())
	return gobool(c)
}

func fromNativeStyleContext(c *C.GtkStyleContext) *StyleContext {
	if c == nil {
		return nil
	}

	obj := glib.WrapObject(unsafe.Pointer(c))
	return wrapStyleContext(obj)
}

// GetStyleContext is a wrapper around gtk_widget_get_style_context().
func (v *Widget) GetStyleContext() *StyleContext {
	return fromNativeStyleContext(C.gtk_widget_get_style_context(v.native()))
}

// GetWindow is a wrapper around gtk_widget_get_window().
func (v *Widget) GetWindow() *gdk.Window {
	c := C.gtk_widget_get_window(v.native())
	w := &gdk.Window{glib.WrapObject(unsafe.Pointer(c))}
	return w
}
