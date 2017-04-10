package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless/api"
)

type secretPropertiesEdit struct {
	*gtk.Grid
	logger  logging.Logger
	widgets []destroyable
	rows    int
}

func newSecretPropertiesEdit(logger logging.Logger) (*secretPropertiesEdit, error) {
	grid, err := gtk.GridNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create grid")
	}

	w := &secretPropertiesEdit{
		Grid:   grid,
		logger: logger.WithField("package", "ui").WithField("component", "secretPropertiesEdit"),
	}

	return w, nil
}

func (w *secretPropertiesEdit) setEdit(secretVersion *api.SecretVersion) {
	w.destroyAllChildren()

	w.renderUrls(secretVersion.URLs)

	w.ShowAll()
}

func (w *secretPropertiesEdit) renderUrls(urls []string) {
	label, err := gtk.LabelNew("URLs")
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
	label.SetHAlign(gtk.ALIGN_START)
	label.SetVAlign(gtk.ALIGN_START)
	w.widgets = append(w.widgets, label)
	w.Attach(label, 0, w.rows, 1, 1)

	urlsEdit, err := newUrlsEdit(w.logger)
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
	urlsEdit.SetHExpand(true)
	urlsEdit.setUrls(urls)
	w.widgets = append(w.widgets, urlsEdit)
	w.Attach(urlsEdit, 1, w.rows, 1, 1)

	w.rows++
}

func (w *secretPropertiesEdit) destroyAllChildren() {
	for _, widget := range w.widgets {
		w.Remove(widget)
		widget.Destroy()
	}
	w.widgets = nil
	w.rows = 0
}
