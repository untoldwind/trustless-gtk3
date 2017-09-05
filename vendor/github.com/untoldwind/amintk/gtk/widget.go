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

type SizeRequestMode int

const (
	SizeRequestModeHeightForWidth SizeRequestMode = C.GTK_SIZE_REQUEST_HEIGHT_FOR_WIDTH
	SizeRequestModeWidthForHeight SizeRequestMode = C.GTK_SIZE_REQUEST_WIDTH_FOR_HEIGHT
	SizeRequestModeConstantSize   SizeRequestMode = C.GTK_SIZE_REQUEST_CONSTANT_SIZE
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
	if v == nil {
		return nil
	}
	return (*C.GtkWidget)(v.Native())
}

func (v *Widget) toWidget() *C.GtkWidget {
	if v == nil {
		return nil
	}
	return v.native()
}

func wrapWidget(p unsafe.Pointer) *Widget {
	if obj := glib.WrapInitiallyUnowned(p); obj != nil {
		return &Widget{InitiallyUnowned: *obj}
	}
	return nil
}

// GetCanFocus is a wrapper around gtk_widget_get_can_focus().
func (v *Widget) GetCanFocus() bool {
	c := C.gtk_widget_get_can_focus(v.native())
	return gobool(c)
}

// SetCanFocus is a wrapper around gtk_widget_set_can_focus().
func (v *Widget) SetCanFocus(canFocus bool) {
	C.gtk_widget_set_can_focus(v.native(), gbool(canFocus))
}

// GetCanDefault is a wrapper around gtk_widget_get_can_default().
func (v *Widget) GetCanDefault() bool {
	c := C.gtk_widget_get_can_default(v.native())
	return gobool(c)
}

// SetCanDefault is a wrapper around gtk_widget_set_can_default().
func (v *Widget) SetCanDefault(canDefault bool) {
	C.gtk_widget_set_can_default(v.native(), gbool(canDefault))
}

// GetMapped is a wrapper around gtk_widget_get_mapped().
func (v *Widget) GetMapped() bool {
	c := C.gtk_widget_get_mapped(v.native())
	return gobool(c)
}

// SetMapped is a wrapper around gtk_widget_set_mapped().
func (v *Widget) SetMapped(mapped bool) {
	C.gtk_widget_set_can_focus(v.native(), gbool(mapped))
}

// GetRealized is a wrapper around gtk_widget_get_realized().
func (v *Widget) GetRealized() bool {
	c := C.gtk_widget_get_realized(v.native())
	return gobool(c)
}

// SetRealized is a wrapper around gtk_widget_set_realized().
func (v *Widget) SetRealized(realized bool) {
	C.gtk_widget_set_realized(v.native(), gbool(realized))
}

// GetHasWindow is a wrapper around gtk_widget_get_has_window().
func (v *Widget) GetHasWindow() bool {
	c := C.gtk_widget_get_has_window(v.native())
	return gobool(c)
}

// SetHasWindow is a wrapper around gtk_widget_set_has_window().
func (v *Widget) SetHasWindow(hasWindow bool) {
	C.gtk_widget_set_has_window(v.native(), gbool(hasWindow))
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
	return wrapWidget(unsafe.Pointer(c))
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

// SetStateFlags is a wrapper around gtk_widget_set_state_flags().
func (v *Widget) SetStateFlags(stateFlags StateFlags, clear bool) {
	C.gtk_widget_set_state_flags(v.native(), C.GtkStateFlags(stateFlags), gbool(clear))
}

// Event is a wrapper around gtk_widget_event().
func (v *Widget) Event(event *gdk.Event) bool {
	c := C.gtk_widget_event(v.native(), (*C.GdkEvent)(event.Native()))
	return gobool(c)
}

// Activate is a wrapper around gtk_widget_activate().
func (v *Widget) Activate() bool {
	return gobool(C.gtk_widget_activate(v.native()))
}

// IsFocus is a wrapper around gtk_widget_is_focus().
func (v *Widget) IsFocus() bool {
	return gobool(C.gtk_widget_is_focus(v.native()))
}

// GrabFocus is a wrapper around gtk_widget_grab_focus().
func (v *Widget) GrabFocus() {
	C.gtk_widget_grab_focus(v.native())
}

// GrabDefault is a wrapper around gtk_widget_grab_default().
func (v *Widget) GrabDefault() {
	C.gtk_widget_grab_default(v.native())
}

// SetName is a wrapper around gtk_widget_set_name().
func (v *Widget) SetName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_widget_set_name(v.native(), (*C.gchar)(cstr))
}

// GetName is a wrapper around gtk_widget_get_name().  A non-nil
// error is returned in the case that gtk_widget_get_name returns NULL to
// differentiate between NULL and an empty string.
func (v *Widget) GetName() string {
	c := C.gtk_widget_get_name(v.native())
	return C.GoString((*C.char)(c))
}

// GetSensitive is a wrapper around gtk_widget_get_sensitive().
func (v *Widget) GetSensitive() bool {
	c := C.gtk_widget_get_sensitive(v.native())
	return gobool(c)
}

// IsSensitive is a wrapper around gtk_widget_is_sensitive().
func (v *Widget) IsSensitive() bool {
	c := C.gtk_widget_is_sensitive(v.native())
	return gobool(c)
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

// SetVisible is a wrapper around gtk_widget_set_visible().
func (v *Widget) SetVisible(visible bool) {
	C.gtk_widget_set_visible(v.native(), gbool(visible))
}

// SetParent is a wrapper around gtk_widget_set_parent().
func (v *Widget) SetParent(parent IWidget) {
	C.gtk_widget_set_parent(v.native(), parent.toWidget())
}

// GetParent is a wrapper around gtk_widget_get_parent().
func (v *Widget) GetParent() *Widget {
	c := C.gtk_widget_get_parent(v.native())
	return wrapWidget(unsafe.Pointer(c))
}

// SetSizeRequest is a wrapper around gtk_widget_set_size_request().
func (v *Widget) SetSizeRequest(width, height int) {
	C.gtk_widget_set_size_request(v.native(), C.gint(width), C.gint(height))
}

// GetSizeRequest is a wrapper around gtk_widget_get_size_request().
func (v *Widget) GetSizeRequest() (width, height int) {
	var w, h C.gint
	C.gtk_widget_get_size_request(v.native(), &w, &h)
	return int(w), int(h)
}

// SetParentWindow is a wrapper around gtk_widget_set_parent_window().
func (v *Widget) SetParentWindow(parentWindow *gdk.Window) {
	C.gtk_widget_set_parent_window(v.native(),
		(*C.GdkWindow)(unsafe.Pointer(parentWindow.Native())))
}

// GetParentWindow is a wrapper around gtk_widget_get_parent_window().
func (v *Widget) GetParentWindow() *gdk.Window {
	c := C.gtk_widget_get_parent_window(v.native())
	return gdk.WrapWindow(unsafe.Pointer(c))
}

// SetEvents is a wrapper around gtk_widget_set_events().
func (v *Widget) SetEvents(events int) {
	C.gtk_widget_set_events(v.native(), C.gint(events))
}

// GetEvents is a wrapper around gtk_widget_get_events().
func (v *Widget) GetEvents() int {
	return int(C.gtk_widget_get_events(v.native()))
}

// AddEvents is a wrapper around gtk_widget_add_events().
func (v *Widget) AddEvents(events int) {
	C.gtk_widget_add_events(v.native(), C.gint(events))
}

// HasDefault is a wrapper around gtk_widget_has_default().
func (v *Widget) HasDefault() bool {
	c := C.gtk_widget_has_default(v.native())
	return gobool(c)
}

// HasFocus is a wrapper around gtk_widget_has_focus().
func (v *Widget) HasFocus() bool {
	c := C.gtk_widget_has_focus(v.native())
	return gobool(c)
}

// HasVisibleFocus is a wrapper around gtk_widget_has_visible_focus().
func (v *Widget) HasVisibleFocus() bool {
	c := C.gtk_widget_has_visible_focus(v.native())
	return gobool(c)
}

// HasGrab is a wrapper around gtk_widget_has_grab().
func (v *Widget) HasGrab() bool {
	c := C.gtk_widget_has_grab(v.native())
	return gobool(c)
}

// IsDrawable is a wrapper around gtk_widget_is_drawable().
func (v *Widget) IsDrawable() bool {
	c := C.gtk_widget_is_drawable(v.native())
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

// QueueDraw is a wrapper around gtk_widget_queue_draw().
func (v *Widget) QueueDraw() {
	C.gtk_widget_queue_draw(v.native())
}

// GetAllocation is a wrapper around gtk_widget_get_allocation().
func (v *Widget) GetAllocation() *Allocation {
	var a Allocation
	C.gtk_widget_get_allocation(v.native(), a.native())
	return &a
}

// SetAllocation is a wrapper around gtk_widget_set_allocation().
func (v *Widget) SetAllocation(allocation *Allocation) {
	C.gtk_widget_set_allocation(v.native(), allocation.native())
}

// SizeAllocate is a wrapper around gtk_widget_size_allocate().
func (v *Widget) SizeAllocate(allocation *Allocation) {
	C.gtk_widget_size_allocate(v.native(), allocation.native())
}

// GetPreferredHeight is a wrapper around gtk_widget_get_preferred_height().
func (v *Widget) GetPreferredHeight() (int, int) {
	var minimum, natural C.gint
	C.gtk_widget_get_preferred_height(v.native(), &minimum, &natural)
	return int(minimum), int(natural)
}

// GetPreferredWidth is a wrapper around gtk_widget_get_preferred_width().
func (v *Widget) GetPreferredWidth() (int, int) {
	var minimum, natural C.gint
	C.gtk_widget_get_preferred_width(v.native(), &minimum, &natural)
	return int(minimum), int(natural)
}

// GetWindow is a wrapper around gtk_widget_get_window().
func (v *Widget) GetWindow() *gdk.Window {
	c := C.gtk_widget_get_window(v.native())
	return gdk.WrapWindow(unsafe.Pointer(c))
}

func (v *Widget) OnDestroy(callback func()) {
	if v != nil {
		v.Connect("destroy", glib.CallbackVoidVoid(callback))
	}
}

func (v *Widget) OnAfterShow(callback func()) {
	if v != nil {
		v.ConnectAfter("show", glib.CallbackVoidVoid(callback))
	}
}

func (v *Widget) OnAfterHide(callback func()) {
	if v != nil {
		v.ConnectAfter("hide", glib.CallbackVoidVoid(callback))
	}
}

func (v *Widget) OnAfterRealize(callback func()) {
	if v != nil {
		v.ConnectAfter("realize", glib.CallbackVoidVoid(callback))
	}
}

func (v *Widget) OnPopupMenu(callback func()) {
	if v != nil {
		v.ConnectAfter("popup-menu", glib.CallbackVoidVoid(callback))
	}
}

func (v *Widget) OnKeyPressEvent(callback func(*gdk.Event) bool) {
	if v != nil {
		v.Connect("key-press-event", gdk.CallbackEventBoolean(callback))
	}
}

func (v *Widget) OnKeyReleaseEvent(callback func(*gdk.Event) bool) {
	if v != nil {
		v.Connect("key-release-event", gdk.CallbackEventBoolean(callback))
	}
}

func (v *Widget) OnButtonPressEvent(callback func(*gdk.Event) bool) {
	if v != nil {
		v.Connect("button-press-event", gdk.CallbackEventBoolean(callback))
	}
}

func (v *Widget) OnButtonReleaseEvent(callback func(*gdk.Event) bool) {
	if v != nil {
		v.Connect("button-release-event", gdk.CallbackEventBoolean(callback))
	}
}

func (v *Widget) OnEnterNotifyEvent(callback func(*gdk.Event) bool) {
	if v != nil {
		v.Connect("enter-notify-event", gdk.CallbackEventBoolean(callback))
	}
}

func (v *Widget) OnLeaveNotifyEvent(callback func(*gdk.Event) bool) {
	if v != nil {
		v.Connect("leave-notify-event", gdk.CallbackEventBoolean(callback))
	}
}

func (v *Widget) OnSizeAllocate(callback func(rect *gdk.Rectangle)) {
	if v != nil {
		v.Connect("size-allocate", gdk.CallbackRectangleVoid(callback))
	}
}
