package main

import (
	"github.com/untoldwind/amintk/gtk"
)

func main() {
	gtk.Init(nil)

	win := gtk.WindowNew(gtk.WindowToplevel)

	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	label := gtk.LabelNew("Hello, gotk3!")

	win.Add(label)

	win.SetDefaultSize(800, 600)

	win.ShowAll()

	gtk.Main()
}
