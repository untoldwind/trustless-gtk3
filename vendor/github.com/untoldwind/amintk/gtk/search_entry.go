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
	if v == nil {
		return nil
	}
	return (*C.GtkSearchEntry)(v.Native())
}

// SearchEntryNew is a wrapper around gtk_search_entry_new().
func SearchEntryNew() *SearchEntry {
	c := C.gtk_search_entry_new()
	return wrapSearchEntry(unsafe.Pointer(c))
}

func wrapSearchEntry(p unsafe.Pointer) *SearchEntry {
	if entry := wrapEntry(p); entry != nil {
		return &SearchEntry{Entry: *entry}
	}
	return nil
}

func (v *Widget) OnSearchChanged(callback func()) *glib.SignalHandle {
	return v.ConnectAfter("search-changed", glib.CallbackVoidVoid(callback))
}
