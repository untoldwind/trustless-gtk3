package gtk

import "unsafe"

// Bin is a representation of GTK's GtkBin.
type Bin struct {
	Container
}

func wrapBin(p unsafe.Pointer) *Bin {
	if container := wrapContainer(p); container != nil {
		return &Bin{Container: *container}
	}
	return nil
}
