package gdk

// #cgo pkg-config: gdk-3.0
// #include "pixbuf.go.h"
import "C"
import (
	"errors"
	"reflect"
	"runtime"
	"strconv"
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// Colorspace is a representation of GDK's GdkColorspace.
type Colorspace int

const (
	ColorspaceRGB Colorspace = C.GDK_COLORSPACE_RGB
)

// InterpType is a representation of GDK's GdkInterpType.
type InterpType int

const (
	InterpTypeNearest  InterpType = C.GDK_INTERP_NEAREST
	InterpTypeTiles    InterpType = C.GDK_INTERP_TILES
	InterpTypeBilinear InterpType = C.GDK_INTERP_BILINEAR
	InterpTypeHyper    InterpType = C.GDK_INTERP_HYPER
)

// PixbufRotation is a representation of GDK's GdkPixbufRotation.
type PixbufRotation int

const (
	PixbufRotationNone              PixbufRotation = C.GDK_PIXBUF_ROTATE_NONE
	PixbufRotationCounterclockwise  PixbufRotation = C.GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE
	PixbufRotationCounterUpsidedown PixbufRotation = C.GDK_PIXBUF_ROTATE_UPSIDEDOWN
	PixbufRotationClockwise         PixbufRotation = C.GDK_PIXBUF_ROTATE_CLOCKWISE
)

// Pixbuf is a representation of GDK's GdkPixbuf.
type Pixbuf struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkBox.
func (v *Pixbuf) native() *C.GdkPixbuf {
	if v == nil {
		return nil
	}
	return (*C.GdkPixbuf)(v.Native())
}

func WrapPixbuf(p unsafe.Pointer) *Pixbuf {
	if obj := glib.WrapObject(p); obj != nil {
		return &Pixbuf{Object: obj}
	}
	return nil
}

// PixbufNew is a wrapper around gdk_pixbuf_new().
func PixbufNew(colorspace Colorspace, hasAlpha bool, bitsPerSample, width, height int) *Pixbuf {
	c := C.gdk_pixbuf_new(C.GdkColorspace(colorspace), gbool(hasAlpha),
		C.int(bitsPerSample), C.int(width), C.int(height))
	return WrapPixbuf(unsafe.Pointer(c))
}

// PixbufNewFromFile is a wrapper around gdk_pixbuf_new_from_file().
func PixbufNewFromFile(filename string) (*Pixbuf, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError
	res := C.gdk_pixbuf_new_from_file((*C.char)(cstr), &err)
	if res == nil {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}
	return WrapPixbuf(unsafe.Pointer(res)), nil
}

// PixbufNewFromFileAtSize is a wrapper around gdk_pixbuf_new_from_file_at_size().
func PixbufNewFromFileAtSize(filename string, width, height int) (*Pixbuf, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	res := C.gdk_pixbuf_new_from_file_at_size(cstr, C.int(width), C.int(height), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}
	return WrapPixbuf(unsafe.Pointer(res)), nil
}

// PixbufNewFromFileAtScale is a wrapper around gdk_pixbuf_new_from_file_at_scale().
func PixbufNewFromFileAtScale(filename string, width, height int, preserveAspectRatio bool) (*Pixbuf, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	res := C.gdk_pixbuf_new_from_file_at_scale(cstr, C.int(width), C.int(height),
		gbool(preserveAspectRatio), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}
	return WrapPixbuf(unsafe.Pointer(res)), nil
}

// GetColorspace is a wrapper around gdk_pixbuf_get_colorspace().
func (v *Pixbuf) GetColorspace() Colorspace {
	c := C.gdk_pixbuf_get_colorspace(v.native())
	return Colorspace(c)
}

// GetNChannels is a wrapper around gdk_pixbuf_get_n_channels().
func (v *Pixbuf) GetNChannels() int {
	c := C.gdk_pixbuf_get_n_channels(v.native())
	return int(c)
}

// GetHasAlpha is a wrapper around gdk_pixbuf_get_has_alpha().
func (v *Pixbuf) GetHasAlpha() bool {
	c := C.gdk_pixbuf_get_has_alpha(v.native())
	return gobool(c)
}

// GetBitsPerSample is a wrapper around gdk_pixbuf_get_bits_per_sample().
func (v *Pixbuf) GetBitsPerSample() int {
	c := C.gdk_pixbuf_get_bits_per_sample(v.native())
	return int(c)
}

// GetPixels is a wrapper around gdk_pixbuf_get_pixels_with_length().
// A Go slice is used to represent the underlying Pixbuf data array, one
// byte per channel.
func (v *Pixbuf) GetPixels() (channels []byte) {
	var length C.guint
	c := C.gdk_pixbuf_get_pixels_with_length(v.native(), &length)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&channels))
	sliceHeader.Data = uintptr(unsafe.Pointer(c))
	sliceHeader.Len = int(length)
	sliceHeader.Cap = int(length)
	// To make sure the slice doesn't outlive the Pixbuf, add a reference
	v.Ref()
	runtime.SetFinalizer(&channels, func(_ *[]byte) {
		v.Unref()
	})
	return
}

// GetWidth is a wrapper around gdk_pixbuf_get_width().
func (v *Pixbuf) GetWidth() int {
	c := C.gdk_pixbuf_get_width(v.native())
	return int(c)
}

// GetHeight is a wrapper around gdk_pixbuf_get_height().
func (v *Pixbuf) GetHeight() int {
	c := C.gdk_pixbuf_get_height(v.native())
	return int(c)
}

// GetRowstride is a wrapper around gdk_pixbuf_get_rowstride().
func (v *Pixbuf) GetRowstride() int {
	c := C.gdk_pixbuf_get_rowstride(v.native())
	return int(c)
}

// GetByteLength is a wrapper around gdk_pixbuf_get_byte_length().
func (v *Pixbuf) GetByteLength() int {
	c := C.gdk_pixbuf_get_byte_length(v.native())
	return int(c)
}

// GetOption is a wrapper around gdk_pixbuf_get_option().  ok is true if
// the key has an associated value.
func (v *Pixbuf) GetOption(key string) (value string, ok bool) {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_pixbuf_get_option(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return "", false
	}
	return C.GoString((*C.char)(c)), true
}

func (v *Pixbuf) Copy() *Pixbuf {
	c := C.gdk_pixbuf_copy(v.native())
	return WrapPixbuf(unsafe.Pointer(c))
}

// ScaleSimple is a wrapper around gdk_pixbuf_scale_simple().
func (v *Pixbuf) ScaleSimple(destWidth, destHeight int, interpType InterpType) *Pixbuf {
	c := C.gdk_pixbuf_scale_simple(v.native(), C.int(destWidth),
		C.int(destHeight), C.GdkInterpType(interpType))
	return WrapPixbuf(unsafe.Pointer(c))
}

// RotateSimple is a wrapper around gdk_pixbuf_rotate_simple().
func (v *Pixbuf) RotateSimple(angle PixbufRotation) *Pixbuf {
	c := C.gdk_pixbuf_rotate_simple(v.native(), C.GdkPixbufRotation(angle))
	return WrapPixbuf(unsafe.Pointer(c))
}

// ApplyEmbeddedOrientation is a wrapper around gdk_pixbuf_apply_embedded_orientation().
func (v *Pixbuf) ApplyEmbeddedOrientation() *Pixbuf {
	c := C.gdk_pixbuf_apply_embedded_orientation(v.native())
	return WrapPixbuf(unsafe.Pointer(c))
}

// Flip is a wrapper around gdk_pixbuf_flip().
func (v *Pixbuf) Flip(horizontal bool) *Pixbuf {
	c := C.gdk_pixbuf_flip(v.native(), gbool(horizontal))
	return WrapPixbuf(unsafe.Pointer(c))
}

// SaveJPEG is a wrapper around gdk_pixbuf_save().
// Quality is a number between 0...100
func (v *Pixbuf) SaveJPEG(path string, quality int) error {
	cpath := C.CString(path)
	cquality := C.CString(strconv.Itoa(quality))
	defer C.free(unsafe.Pointer(cpath))
	defer C.free(unsafe.Pointer(cquality))

	var err *C.GError
	c := C._gdk_pixbuf_save_jpeg(v.native(), cpath, &err, cquality)
	if !gobool(c) {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// SavePNG is a wrapper around gdk_pixbuf_save().
// Compression is a number between 0...9
func (v *Pixbuf) SavePNG(path string, compression int) error {
	cpath := C.CString(path)
	ccompression := C.CString(strconv.Itoa(compression))
	defer C.free(unsafe.Pointer(cpath))
	defer C.free(unsafe.Pointer(ccompression))

	var err *C.GError
	c := C._gdk_pixbuf_save_png(v.native(), cpath, &err, ccompression)
	if !gobool(c) {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}
