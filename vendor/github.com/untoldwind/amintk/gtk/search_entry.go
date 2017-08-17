package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// SearchEntry is a reprensentation of GTK's GtkSearchEntry.
type SearchEntry struct {
	Entry
}

// native returns a pointer to the underlying GtkSearchEntry.
func (v *SearchEntry) native() *C.GtkSearchEntry {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkSearchEntry)(p)
}

// SearchEntryNew is a wrapper around gtk_search_entry_new().
func SearchEntryNew() *SearchEntry {
	c := C.gtk_search_entry_new()
	return wrapSearchEntry(glib.WrapObject(unsafe.Pointer(c)))
}

func wrapSearchEntry(obj *glib.Object) *SearchEntry {
	return &SearchEntry{Entry{Widget{glib.InitiallyUnowned{Object: obj}}}}
}
