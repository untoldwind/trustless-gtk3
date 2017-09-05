package gdk

// #cgo pkg-config: gdk-3.0
// #include "pixbuf.go.h"
import "C"
import (
	"fmt"
	"os"
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

var RectangleType = glib.Type(C.gdk_rectangle_get_type())

// Rectangle is a representation of GDK's GdkRectangle type.
type Rectangle struct {
	GdkRectangle C.GdkRectangle
}

func WrapRectangle(p unsafe.Pointer) *Rectangle {
	if p != nil {
		return &Rectangle{GdkRectangle: *(*C.GdkRectangle)(p)}
	}
	return nil
}

func (r *Rectangle) native() *C.GdkRectangle {
	return &r.GdkRectangle
}

// GetX returns x field of the underlying GdkRectangle.
func (r *Rectangle) GetX() int {
	return int(r.native().x)
}

// GetY returns y field of the underlying GdkRectangle.
func (r *Rectangle) GetY() int {
	return int(r.native().y)
}

// GetWidth returns width field of the underlying GdkRectangle.
func (r *Rectangle) GetWidth() int {
	return int(r.native().width)
}

// GetHeight returns height field of the underlying GdkRectangle.
func (r *Rectangle) GetHeight() int {
	return int(r.native().height)
}

type CallbackRectangleVoid func(rect *Rectangle)

func (c CallbackRectangleVoid) Call(args []glib.Value) *glib.Value {
	var arg0 *Rectangle
	var arg0Ok bool
	for _, value := range args {
		if actual, fundamental := value.Type(); actual == RectangleType && fundamental == glib.TypeBoxed {
			if p, ok := value.GetBoxed(); ok {
				arg0 = WrapRectangle(p)
				arg0Ok = true
			}
		}
	}
	if !arg0Ok {
		fmt.Fprintln(os.Stderr, "WARNING: CallbackRectangleVoid: No Rectangle found in args")
	}
	c(arg0)
	return nil
}
