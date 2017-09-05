package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <glib.h>
// #include <glib-object.h>
import "C"
import (
	"fmt"
	"os"
)

type Callback interface {
	Call(args []Value) *Value
}

type CallbackVoidVoid func()

func (c CallbackVoidVoid) Call(args []Value) *Value {
	c()
	return nil
}

type CallbackIntVoid func(int)

func (c CallbackIntVoid) Call(args []Value) *Value {
	var arg0 int
	var arg0Ok bool
	for _, value := range args {
		if _, fundamental := value.Type(); fundamental == TypeInt {
			arg0, arg0Ok = value.GetInt()
			break
		}
	}
	if !arg0Ok {
		fmt.Fprintln(os.Stderr, "WARNING: CallbackIntVoid: No int found in args")
	}
	c(arg0)
	return nil
}
