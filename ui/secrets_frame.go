package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless-gtk3/state"
)

type secretsFrame struct {
	*gtk.Box
	logger logging.Logger
}

func newSecretsFrame(store *state.Store, logger logging.Logger) (*secretsFrame, error) {
	box := gtk.BoxNew(gtk.OrientationHorizontal, 0)
	w := &secretsFrame{
		Box:    box,
		logger: logger.WithField("package", "ui").WithField("component", "secretsFrame"),
	}
	sidebar := newSidebar(store, logger)
	w.Add(sidebar)

	right := gtk.BoxNew(gtk.OrientationVertical, 0)
	w.Add(right)

	headerBar := newHeaderBar(store, logger)
	right.Add(headerBar)
	right.SetFocusChild(headerBar)
	w.SetFocusChild(right)

	paned := gtk.PanedNew(gtk.OrientationHorizontal)
	paned.SetVExpand(true)
	right.Add(paned)

	entryList := newEntryList(store, logger)
	paned.Add1(entryList)

	secretDetail := newSecretDetail(store, logger)

	paned.Add2(secretDetail)
	paned.OnAfterRealize(func() {
		paned.SetPosition(300)
	})

	return w, nil
}
