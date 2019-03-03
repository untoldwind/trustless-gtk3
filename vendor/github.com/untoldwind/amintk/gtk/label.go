package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/pango"
)

// Label is a representation of GTK's GtkLabel.
type Label struct {
	Widget
}

// native returns a pointer to the underlying GtkLabel.
func (v *Label) native() *C.GtkLabel {
	if v == nil {
		return nil
	}
	return (*C.GtkLabel)(v.Native())
}

// LabelNew is a wrapper around gtk_label_new().
func LabelNew(str string) *Label {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_label_new((*C.gchar)(cstr))
	return wrapLabel(unsafe.Pointer(c))
}

func wrapLabel(p unsafe.Pointer) *Label {
	if widget := wrapWidget(p); widget != nil {
		return &Label{Widget: *widget}
	}
	return nil
}

// SetText is a wrapper around gtk_label_set_text().
func (v *Label) SetText(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_text(v.native(), (*C.gchar)(cstr))
}

// GetText is a wrapper around gtk_label_get_text().
func (v *Label) GetText() string {
	c := C.gtk_label_get_text(v.native())
	return C.GoString((*C.char)(c))
}

// SetMarkup is a wrapper around gtk_label_set_markup().
func (v *Label) SetMarkup(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_markup(v.native(), (*C.gchar)(cstr))
}

// SetMarkupWithMnemonic is a wrapper around
// gtk_label_set_markup_with_mnemonic().
func (v *Label) SetMarkupWithMnemonic(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_markup_with_mnemonic(v.native(), (*C.gchar)(cstr))
}

// SetPattern is a wrapper around gtk_label_set_pattern().
func (v *Label) SetPattern(patern string) {
	cstr := C.CString(patern)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_pattern(v.native(), (*C.gchar)(cstr))
}

// SetSelectable is a wrapper around gtk_label_set_selectable().
func (v *Label) SetSelectable(setting bool) {
	C.gtk_label_set_selectable(v.native(), gbool(setting))
}

// GetWidthChars is a wrapper around gtk_label_get_width_chars().
func (v *Label) GetWidthChars() int {
	c := C.gtk_label_get_width_chars(v.native())
	return int(c)
}

// SetWidthChars is a wrapper around gtk_label_set_width_chars().
func (v *Label) SetWidthChars(nChars int) {
	C.gtk_label_set_width_chars(v.native(), C.gint(nChars))
}

// GetMaxWidthChars is a wrapper around gtk_label_get_max_width_chars().
func (v *Label) GetMaxWidthChars() int {
	c := C.gtk_label_get_max_width_chars(v.native())
	return int(c)
}

// SetMaxWidthChars is a wrapper around gtk_label_set_max_width_chars().
func (v *Label) SetMaxWidthChars(nChars int) {
	C.gtk_label_set_max_width_chars(v.native(), C.gint(nChars))
}

// GetLineWrap is a wrapper around gtk_label_get_line_wrap().
func (v *Label) GetLineWrap() bool {
	c := C.gtk_label_get_line_wrap(v.native())
	return gobool(c)
}

// SetLineWrap is a wrapper around gtk_label_set_line_wrap().
func (v *Label) SetLineWrap(wrap bool) {
	C.gtk_label_set_line_wrap(v.native(), gbool(wrap))
}

// GetJustify is a wrapper around gtk_label_get_justify().
func (v *Label) GetJustify() Justification {
	c := C.gtk_label_get_justify(v.native())
	return Justification(c)
}

// SetJustify is a wrapper around gtk_label_set_justify().
func (v *Label) SetJustify(jtype Justification) {
	C.gtk_label_set_justify(v.native(), C.GtkJustification(jtype))
}

// GetEllipsize is a wrapper around gtk_label_get_ellipsize().
func (v *Label) GetEllipsize() pango.EllipsizeMode {
	c := C.gtk_label_get_ellipsize(v.native())
	return pango.EllipsizeMode(c)
}

// SetEllipsize is a wrapper around gtk_label_set_ellipsize().
func (v *Label) SetEllipsize(mode pango.EllipsizeMode) {
	C.gtk_label_set_ellipsize(v.native(), C.PangoEllipsizeMode(mode))
}

// GetCurrentUri is a wrapper around gtk_label_get_current_uri().
func (v *Label) GetCurrentUri() string {
	c := C.gtk_label_get_current_uri(v.native())
	return C.GoString((*C.char)(c))
}

// GetTrackVisitedLinks is a wrapper around gtk_label_get_track_visited_links().
func (v *Label) GetTrackVisitedLinks() bool {
	c := C.gtk_label_get_track_visited_links(v.native())
	return gobool(c)
}

// SetTrackVisitedLinks is a wrapper around gtk_label_set_track_visited_links().
func (v *Label) SetTrackVisitedLinks(trackLinks bool) {
	C.gtk_label_set_track_visited_links(v.native(), gbool(trackLinks))
}

// GetAngle is a wrapper around gtk_label_get_angle().
func (v *Label) GetAngle() float64 {
	c := C.gtk_label_get_angle(v.native())
	return float64(c)
}

// SetAngle is a wrapper around gtk_label_set_angle().
func (v *Label) SetAngle(angle float64) {
	C.gtk_label_set_angle(v.native(), C.gdouble(angle))
}

// GetSelectionBounds is a wrapper around gtk_label_get_selection_bounds().
func (v *Label) GetSelectionBounds() (start, end int, nonEmpty bool) {
	var cstart, cend C.gint
	c := C.gtk_label_get_selection_bounds(v.native(), &cstart, &cend)
	return int(cstart), int(cend), gobool(c)
}

// GetSingleLineMode is a wrapper around gtk_label_get_single_line_mode().
func (v *Label) GetSingleLineMode() bool {
	c := C.gtk_label_get_single_line_mode(v.native())
	return gobool(c)
}

// SetSingleLineMode is a wrapper around gtk_label_set_single_line_mode().
func (v *Label) SetSingleLineMode(mode bool) {
	C.gtk_label_set_single_line_mode(v.native(), gbool(mode))
}

// GetUseMarkup is a wrapper around gtk_label_get_use_markup().
func (v *Label) GetUseMarkup() bool {
	c := C.gtk_label_get_use_markup(v.native())
	return gobool(c)
}

// SetUseMarkup is a wrapper around gtk_label_set_use_markup().
func (v *Label) SetUseMarkup(use bool) {
	C.gtk_label_set_use_markup(v.native(), gbool(use))
}

// GetUseUnderline is a wrapper around gtk_label_get_use_underline().
func (v *Label) GetUseUnderline() bool {
	c := C.gtk_label_get_use_underline(v.native())
	return gobool(c)
}

// SetUseUnderline is a wrapper around gtk_label_set_use_underline().
func (v *Label) SetUseUnderline(use bool) {
	C.gtk_label_set_use_underline(v.native(), gbool(use))
}

// SetLabel is a wrapper around gtk_label_set_label().
func (v *Label) SetLabel(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_label(v.native(), (*C.gchar)(cstr))
}

// GetLabel is a wrapper around gtk_label_get_label().
func (v *Label) GetLabel() string {
	c := C.gtk_label_get_label(v.native())
	return C.GoString((*C.char)(c))
}
