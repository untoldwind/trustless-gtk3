package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

// IconSize is a representation of GTK's GtkIconSize.
type IconSize int

const (
	IconSizeInvalid      IconSize = C.GTK_ICON_SIZE_INVALID
	IconSizeMenu         IconSize = C.GTK_ICON_SIZE_MENU
	IconSizeSmallToolbar IconSize = C.GTK_ICON_SIZE_SMALL_TOOLBAR
	IconSizeLargeToolbar IconSize = C.GTK_ICON_SIZE_LARGE_TOOLBAR
	IconSizeButton       IconSize = C.GTK_ICON_SIZE_BUTTON
	IconSizeDND          IconSize = C.GTK_ICON_SIZE_DND
	IconSizeDialog       IconSize = C.GTK_ICON_SIZE_DIALOG
)
