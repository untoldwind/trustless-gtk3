package gtk

// #cgo pkg-config: gtk+-3.0
// #include "main.go.h"
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/gdk"
)

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func gobool(b C.gboolean) bool {
	return b != C.FALSE
}

// Init is a wrapper around gtk_init() and must be called before any
// other GTK calls and is used to initialize everything necessary.
//
// In addition to setting up GTK for usage, a pointer to a slice of
// strings may be passed in to parse standard GTK command line arguments.
// args will be modified to remove any flags that were handled.
// Alternatively, nil may be passed in to not perform any command line
// parsing.
func Init(args *[]string) {
	if args != nil {
		argc := C.int(len(*args))
		argv := C.make_strings(argc)
		defer C.destroy_strings(argv)

		for i, arg := range *args {
			cstr := C.CString(arg)
			C.set_string(argv, C.int(i), (*C.gchar)(cstr))
		}

		C.gtk_init((*C.int)(unsafe.Pointer(&argc)),
			(***C.char)(unsafe.Pointer(&argv)))

		unhandled := make([]string, argc)
		for i := 0; i < int(argc); i++ {
			cstr := C.get_string(argv, C.int(i))
			unhandled[i] = C.GoString((*C.char)(cstr))
			C.free(unsafe.Pointer(cstr))
		}
		*args = unhandled
	} else {
		C.gtk_init(nil, nil)
	}
}

// Main is a wrapper around gtk_main() and runs the GTK main loop,
// blocking until MainQuit() is called.
func Main() {
	C.gtk_main()
}

// MainIteration is a wrapper around gtk_main_iteration.
func MainIteration() bool {
	return gobool(C.gtk_main_iteration())
}

// MainIterationDo is a wrapper around gtk_main_iteration_do.
func MainIterationDo(blocking bool) bool {
	return gobool(C.gtk_main_iteration_do(gbool(blocking)))
}

// EventsPending is a wrapper around gtk_events_pending.
func EventsPending() bool {
	return gobool(C.gtk_events_pending())
}

// MainQuit is a wrapper around gtk_main_quit() is used to terminate
// the GTK main loop (started by Main()).
func MainQuit() {
	C.gtk_main_quit()
}

func CurrentEvent() *gdk.Event {
	return gdk.WrapEvent(unsafe.Pointer(C.gtk_get_current_event()))
}
