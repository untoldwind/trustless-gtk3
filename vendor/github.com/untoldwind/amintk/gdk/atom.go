package gdk

// #cgo pkg-config: gdk-3.0
// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"
import "unsafe"

// Atom is a representation of GDK's GdkAtom.
type Atom uintptr

const (
	AtomSelectionPrimary      Atom = 1
	AtomSelectionSecondary    Atom = 2
	AtomSelectionClipboard    Atom = 69
	AtomTargetBitmap          Atom = 5
	AtomTagetColormap         Atom = 7
	AtomTargetDrawable        Atom = 17
	AtomTargetPixmap          Atom = 20
	AtomTargetString          Atom = 31
	AtomSelectionTypeAtom     Atom = 4
	AtomSelectionTypeBitmap   Atom = 5
	AtomSelectionTypeColormap Atom = 7
	AtomSelectionTypeDrawable Atom = 17
	AtomSelectionTypeInteger  Atom = 19
	AtomSelectionTypePixmap   Atom = 20
	AtomSelectionTypeWindow   Atom = 33
	AtomSelectionTypeString   Atom = 31
)

// native returns the underlying GdkAtom.
func (v Atom) native() C.GdkAtom {
	return C.GdkAtom(unsafe.Pointer(uintptr(v)))
}

func (v Atom) Name() string {
	c := C.gdk_atom_name(v.native())
	defer C.g_free(C.gpointer(c))
	return C.GoString((*C.char)(c))
}

// GdkAtomIntern is a wrapper around gdk_atom_intern
func GdkAtomIntern(atomName string, onlyIfExists bool) Atom {
	cstr := C.CString(atomName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_atom_intern((*C.gchar)(cstr), gbool(onlyIfExists))
	return Atom(uintptr(unsafe.Pointer(c)))
}

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}
