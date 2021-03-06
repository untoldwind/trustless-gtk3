package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// ArrowType is a representation of GTK's GtkArrowType.
type ArrowType int

const (
	ArrowTypeUp    ArrowType = C.GTK_ARROW_UP
	ArrowTypeDown  ArrowType = C.GTK_ARROW_DOWN
	ArrowTypeLeft  ArrowType = C.GTK_ARROW_LEFT
	ArrowTypeRight ArrowType = C.GTK_ARROW_RIGHT
	ArrowTypeNone  ArrowType = C.GTK_ARROW_NONE
)

// MenuButton is a representation of GTK's GtkMenuButton.
type MenuButton struct {
	ToggleButton
}

// native() returns a pointer to the underlying GtkButton.
func (v *MenuButton) native() *C.GtkMenuButton {
	if v == nil {
		return nil
	}
	return (*C.GtkMenuButton)(v.Native())
}

// MenuButtonNew is a wrapper around gtk_menu_button_new().
func MenuButtonNew() *MenuButton {
	c := C.gtk_menu_button_new()
	return wrapMenuButton(unsafe.Pointer(c))
}

func wrapMenuButton(p unsafe.Pointer) *MenuButton {
	if toggleButton := wrapToggleButton(p); toggleButton != nil {
		return &MenuButton{ToggleButton: *toggleButton}
	}
	return nil
}

// SetPopup is a wrapper around gtk_menu_button_set_popup().
func (v *MenuButton) SetPopup(menu IMenu) {
	C.gtk_menu_button_set_popup(v.native(), menu.toWidget())
}

// GetPopup is a wrapper around gtk_menu_button_get_popup().
func (v *MenuButton) GetPopup() *Menu {
	c := C.gtk_menu_button_get_popup(v.native())
	if c == nil {
		return nil
	}
	return wrapMenu(unsafe.Pointer(c))
}

// SetDirection is a wrapper around gtk_menu_button_set_direction().
func (v *MenuButton) SetDirection(direction ArrowType) {
	C.gtk_menu_button_set_direction(v.native(), C.GtkArrowType(direction))
}

// GetDirection is a wrapper around gtk_menu_button_get_direction().
func (v *MenuButton) GetDirection() ArrowType {
	c := C.gtk_menu_button_get_direction(v.native())
	return ArrowType(c)
}

// SetPopover is a wrapper around gtk_menu_button_set_popover().
func (v *MenuButton) SetPopover(popover *Popover) {
	C.gtk_menu_button_set_popover(v.native(), popover.toWidget())
}

// GetPopover is a wrapper around gtk_menu_button_get_popover().
func (v *MenuButton) GetPopover() *Popover {
	c := C.gtk_menu_button_get_popover(v.native())
	return wrapPopover(unsafe.Pointer(c))
}
