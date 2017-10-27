package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// MessageType is a representation of GTK's GtkMessageType.
type MessageType int

const (
	MessageTypeInfo     MessageType = C.GTK_MESSAGE_INFO
	MessageTypeWarning  MessageType = C.GTK_MESSAGE_WARNING
	MessageTypeQuestion MessageType = C.GTK_MESSAGE_QUESTION
	MessageTypeError    MessageType = C.GTK_MESSAGE_ERROR
	MessageTypeOther    MessageType = C.GTK_MESSAGE_OTHER
)

type InfoBar struct {
	Box
}

func (v *InfoBar) native() *C.GtkInfoBar {
	if v == nil {
		return nil
	}
	return (*C.GtkInfoBar)(v.Native())
}

func InfoBarNew() *InfoBar {
	c := C.gtk_info_bar_new()
	return wrapInfoBar(unsafe.Pointer(c))
}

func wrapInfoBar(p unsafe.Pointer) *InfoBar {
	if box := wrapBox(p); box != nil {
		return &InfoBar{Box: *box}
	}
	return nil
}

func (v *InfoBar) SetMessageType(messageType MessageType) {
	C.gtk_info_bar_set_message_type(v.native(), C.GtkMessageType(messageType))
}

func (v *InfoBar) GetMessageType() MessageType {
	messageType := C.gtk_info_bar_get_message_type(v.native())
	return MessageType(messageType)
}

func (v *InfoBar) GetShowCloseButton() bool {
	b := C.gtk_info_bar_get_show_close_button(v.native())
	return gobool(b)
}

func (v *InfoBar) SetShowCloseButton(setting bool) {
	C.gtk_info_bar_set_show_close_button(v.native(), gbool(setting))
}

func (v *InfoBar) GetContentArea() *Box {
	c := C.gtk_info_bar_get_content_area(v.native())
	if c == nil {
		return nil
	}
	return wrapBox(unsafe.Pointer(c))
}

func (v *InfoBar) GetActionArea() *Widget {
	c := C.gtk_info_bar_get_action_area(v.native())
	if c == nil {
		return nil
	}
	return wrapWidget(unsafe.Pointer(c))
}

func (v *InfoBar) OnClose(callback func()) *glib.SignalHandle {
	return v.Connect("close", glib.CallbackVoidVoid(callback))
}

func (v *InfoBar) OnResponse(callback func(responseId int)) *glib.SignalHandle {
	return v.Connect("close", glib.CallbackIntVoid(callback))
}
