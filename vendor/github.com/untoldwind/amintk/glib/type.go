package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <glib.h>
// #include <glib-object.h>
import "C"

// Type is a representation of GLib's GType.
type Type uint

const (
	TypeInvalid   Type = C.G_TYPE_INVALID
	TypeNone      Type = C.G_TYPE_NONE
	TypeInterface Type = C.G_TYPE_INTERFACE
	TypeChar      Type = C.G_TYPE_CHAR
	TypeUChar     Type = C.G_TYPE_UCHAR
	TypeBoolean   Type = C.G_TYPE_BOOLEAN
	TypeInt       Type = C.G_TYPE_INT
	TypeUInt      Type = C.G_TYPE_UINT
	TypeLong      Type = C.G_TYPE_LONG
	TypeULong     Type = C.G_TYPE_ULONG
	TypeInt64     Type = C.G_TYPE_INT64
	TypeUInt64    Type = C.G_TYPE_UINT64
	TypeEnum      Type = C.G_TYPE_ENUM
	TypeFlags     Type = C.G_TYPE_FLAGS
	TypeFloat     Type = C.G_TYPE_FLOAT
	TypeDouble    Type = C.G_TYPE_DOUBLE
	TypeString    Type = C.G_TYPE_STRING
	TypePointer   Type = C.G_TYPE_POINTER
	TypeBoxed     Type = C.G_TYPE_BOXED
	TypeParam     Type = C.G_TYPE_PARAM
	TypeObject    Type = C.G_TYPE_OBJECT
	TypeVariant   Type = C.G_TYPE_VARIANT
)

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}
