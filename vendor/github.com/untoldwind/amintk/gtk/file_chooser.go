package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// FileChooserAction is a representation of GTK's GtkFileChooserAction.
type FileChooserAction int

const (
	FileChooserActionOpen         FileChooserAction = C.GTK_FILE_CHOOSER_ACTION_OPEN
	FileChooserActionSave         FileChooserAction = C.GTK_FILE_CHOOSER_ACTION_SAVE
	FileChooserActionSelectFolder FileChooserAction = C.GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER
	FileChooserActionCreateFolder FileChooserAction = C.GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER
)

// FileChooser is a representation of GTK's GtkFileChooser GInterface.
type FileChooser struct {
	*glib.Object
}

// native returns a pointer to the underlying GObject as a GtkFileChooser.
func (v *FileChooser) native() *C.GtkFileChooser {
	if v == nil {
		return nil
	}
	return (*C.GtkFileChooser)(v.Native())
}

func wrapFileChooser(obj *glib.Object) *FileChooser {
	return &FileChooser{Object: obj}
}

// GetFilename is a wrapper around gtk_file_chooser_get_filename().
func (v *FileChooser) GetFilename() string {
	c := C.gtk_file_chooser_get_filename(v.native())
	s := C.GoString((*C.char)(c))
	defer C.g_free((C.gpointer)(c))
	return s
}

// SetCurrentName is a wrapper around gtk_file_chooser_set_current_name().
func (v *FileChooser) SetCurrentName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_file_chooser_set_current_name(v.native(), (*C.gchar)(cstr))
	return
}

// SetCurrentFolder is a wrapper around gtk_file_chooser_set_current_folder().
func (v *FileChooser) SetCurrentFolder(folder string) bool {
	cstr := C.CString(folder)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_file_chooser_set_current_folder(v.native(), (*C.gchar)(cstr))
	return gobool(c)
}

// GetCurrentFolder is a wrapper around gtk_file_chooser_get_current_folder().
func (v *FileChooser) GetCurrentFolder() string {
	c := C.gtk_file_chooser_get_current_folder(v.native())
	defer C.free(unsafe.Pointer(c))
	return C.GoString((*C.char)(c))
}
