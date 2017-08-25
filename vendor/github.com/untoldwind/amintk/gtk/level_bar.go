package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

type LevelBar struct {
	Widget
}

// native returns a pointer to the underlying GtkLevelBar.
func (v *LevelBar) native() *C.GtkLevelBar {
	if v == nil {
		return nil
	}
	return (*C.GtkLevelBar)(v.Native())
}

// LevelBarNew is a wrapper around gtk_level_bar_new().
func LevelBarNew() *LevelBar {
	c := C.gtk_level_bar_new()
	return wrapLevelBar(unsafe.Pointer(c))
}

func wrapLevelBar(p unsafe.Pointer) *LevelBar {
	if widget := wrapWidget(p); widget != nil {
		return &LevelBar{Widget: *widget}
	}
	return nil
}

// SetValue is a wrapper around gtk_level_bar_set_value().
func (v *LevelBar) SetValue(value float64) {
	C.gtk_level_bar_set_value(v.native(), C.gdouble(value))
}

// GetValue is a wrapper around gtk_level_bar_get_value().
func (v *LevelBar) GetValue() float64 {
	c := C.gtk_level_bar_get_value(v.native())
	return float64(c)
}

// SetMinValue is a wrapper around gtk_level_bar_set_min_value().
func (v *LevelBar) SetMinValue(value float64) {
	C.gtk_level_bar_set_min_value(v.native(), C.gdouble(value))
}

// GetMinValue is a wrapper around gtk_level_bar_get_min_value().
func (v *LevelBar) GetMinValue() float64 {
	c := C.gtk_level_bar_get_min_value(v.native())
	return float64(c)
}

// SetMaxValue is a wrapper around gtk_level_bar_set_max_value().
func (v *LevelBar) SetMaxValue(value float64) {
	C.gtk_level_bar_set_max_value(v.native(), C.gdouble(value))
}

// GetMaxValue is a wrapper around gtk_level_bar_get_max_value().
func (v *LevelBar) GetMaxValue() float64 {
	c := C.gtk_level_bar_get_max_value(v.native())
	return float64(c)
}
