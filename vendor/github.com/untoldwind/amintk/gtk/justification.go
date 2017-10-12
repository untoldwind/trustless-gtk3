package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

// Justify is a representation of GTK's GtkJustification.
type Justification int

const (
	JustificationLeft   Justification = C.GTK_JUSTIFY_LEFT
	JustificationRight  Justification = C.GTK_JUSTIFY_RIGHT
	JustificationCenter Justification = C.GTK_JUSTIFY_CENTER
	JustificationFill   Justification = C.GTK_JUSTIFY_FILL
)
