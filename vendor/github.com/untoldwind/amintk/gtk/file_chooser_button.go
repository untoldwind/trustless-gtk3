package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// FileChooserButton is a representation of GTK's GtkFileChooserButton.
type FileChooserButton struct {
	Box

	FileChooser
}

// native returns a pointer to the underlying GtkFileChooserButton.
func (v *FileChooserButton) native() *C.GtkFileChooserButton {
	if v == nil {
		return nil
	}
	return (*C.GtkFileChooserButton)(v.Native())
}

// FileChooserButtonNew is a wrapper around gtk_file_chooser_button_new().
func FileChooserButtonNew(title string, action FileChooserAction) *FileChooserButton {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_file_chooser_button_new((*C.gchar)(cstr),
		(C.GtkFileChooserAction)(action))
	return wrapFileChooserButton(unsafe.Pointer(c))
}

func wrapFileChooserButton(p unsafe.Pointer) *FileChooserButton {
	if box := wrapBox(p); box != nil {
		fc := wrapFileChooser(box.Object)
		return &FileChooserButton{Box: *box, FileChooser: *fc}
	}
	return nil
}

func (v *FileChooserButton) OnFileSet(callback func()) *glib.SignalHandle {
	return v.Connect("file-set", glib.CallbackVoidVoid(callback))
}
