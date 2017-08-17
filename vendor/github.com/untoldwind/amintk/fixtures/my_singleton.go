package fixtures

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include "my_singleton.go.h"
import "C"
import "github.com/untoldwind/amintk/glib"

func MySingleObjectGetType() glib.Type {
	return glib.Type(C.my_singleton_object_get_type())
}
