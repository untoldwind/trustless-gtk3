package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless/api"
)

type destroyable interface {
	gtk.IWidget

	Destroy()
}

type secretPropertiesDisplay struct {
	*gtk.Grid
	widgets []destroyable
	rows    int
	logger  logging.Logger
}

func newSecretPropertiesDisplay(logger logging.Logger) (*secretPropertiesDisplay, error) {
	grid, err := gtk.GridNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create grid")
	}

	w := &secretPropertiesDisplay{
		Grid:   grid,
		logger: logger.WithField("package", "ui").WithField("component", "secretPropertiesDisplay"),
	}

	w.SetOrientation(gtk.ORIENTATION_HORIZONTAL)
	w.SetBorderWidth(5)
	w.SetColumnSpacing(5)
	w.SetRowSpacing(5)

	return w, nil
}

func (w *secretPropertiesDisplay) display(secretVersion *api.SecretVersion) {
	w.destroyAllChildren()

	w.renderUrls(secretVersion.URLs)

	knownNames := w.renderProperties(api.SecretProperties, secretVersion.Properties)

	var unknownPropertyDefs api.SecretPropertyList
	for name := range secretVersion.Properties {
		if _, ok := knownNames[name]; ok {
			continue
		}
		unknownPropertyDefs = append(unknownPropertyDefs, api.SecretProperty{
			Name:    name,
			Display: name,
		})
	}
	unknownPropertyDefs.Sort()

	w.renderProperties(unknownPropertyDefs, secretVersion.Properties)

	w.ShowAll()
}

func (w *secretPropertiesDisplay) renderUrls(urls []string) {
	if len(urls) == 0 {
		return
	}
	label, err := gtk.LabelNew("URLs")
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
	label.SetHAlign(gtk.ALIGN_START)
	label.SetVAlign(gtk.ALIGN_START)
	w.widgets = append(w.widgets, label)
	w.Attach(label, 0, w.rows, 1, 1)

	for _, url := range urls {
		urlLabel, err := newUrlLabel(w.logger, url)
		if err != nil {
			w.logger.ErrorErr(err)
			continue
		}
		urlLabel.SetHExpand(true)
		urlLabel.SetHAlign(gtk.ALIGN_START)

		w.widgets = append(w.widgets, urlLabel)
		w.Attach(urlLabel, 1, w.rows, 1, 1)
		w.rows++
	}
}

func (w *secretPropertiesDisplay) renderProperties(propertyDefs api.SecretPropertyList, properties map[string]string) map[string]bool {
	renderedNames := map[string]bool{}
	for _, propertyDef := range propertyDefs {
		value, ok := properties[propertyDef.Name]
		if !ok {
			continue
		}
		renderedNames[propertyDef.Name] = true
		label, err := gtk.LabelNew(propertyDef.Display)
		if err != nil {
			w.logger.ErrorErr(err)
			continue
		}
		label.SetHAlign(gtk.ALIGN_START)
		label.SetVAlign(gtk.ALIGN_START)
		w.widgets = append(w.widgets, label)
		w.Attach(label, 0, w.rows, 1, 1)

		valueDisplay, err := newSecretValueDisplay(value, propertyDef.Blurred, w.logger)
		if err != nil {
			w.logger.ErrorErr(err)
			continue
		}
		valueDisplay.SetHExpand(true)
		w.widgets = append(w.widgets, valueDisplay)
		w.Attach(valueDisplay, 1, w.rows, 1, 1)

		w.rows++
	}

	return renderedNames
}

func (w *secretPropertiesDisplay) destroyAllChildren() {
	for _, widget := range w.widgets {
		w.Remove(widget)
		widget.Destroy()
	}
	w.widgets = nil
	w.rows = 0
}
