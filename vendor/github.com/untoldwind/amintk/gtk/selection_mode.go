package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

// SelectionMode is a representation of GTK's GtkSelectionMode.
type SelectionMode int

const (
	SelectionModeNone     SelectionMode = C.GTK_SELECTION_NONE
	SelectionModeSingle   SelectionMode = C.GTK_SELECTION_SINGLE
	SelectionModeBrowse   SelectionMode = C.GTK_SELECTION_BROWSE
	SelectionModeMultiple SelectionMode = C.GTK_SELECTION_MULTIPLE
)
