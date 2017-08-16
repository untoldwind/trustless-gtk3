package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
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

func newSecretPropertiesDisplay(logger logging.Logger) *secretPropertiesDisplay {
	grid := gtk.GridNew()

	w := &secretPropertiesDisplay{
		Grid:   grid,
		logger: logger.WithField("package", "ui").WithField("component", "secretPropertiesDisplay"),
	}

	w.SetOrientation(gtk.OrientationHorizontal)
	w.SetBorderWidth(5)
	w.SetColumnSpacing(5)
	w.SetRowSpacing(5)

	return w
}

func (w *secretPropertiesDisplay) display(secretVersion *api.SecretVersion, passwordStrengths map[string]*api.PasswordStrength) {
	w.destroyAllChildren()

	w.renderUrls(secretVersion.URLs)

	knownNames := w.renderProperties(api.SecretProperties, secretVersion.Properties, passwordStrengths)

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

	w.renderProperties(unknownPropertyDefs, secretVersion.Properties, passwordStrengths)

	w.ShowAll()
}

func (w *secretPropertiesDisplay) renderUrls(urls []string) {
	if len(urls) == 0 {
		return
	}
	label := gtk.LabelNew("URLs")
	label.SetHAlign(gtk.AlignStart)
	label.SetVAlign(gtk.AlignStart)
	w.widgets = append(w.widgets, label)
	w.Attach(label, 0, w.rows, 1, 1)

	for _, url := range urls {
		urlLabel, err := newUrlLabel(w.logger, url)
		if err != nil {
			w.logger.ErrorErr(err)
			continue
		}
		urlLabel.SetHExpand(true)
		urlLabel.SetHAlign(gtk.AlignStart)

		w.widgets = append(w.widgets, urlLabel)
		w.Attach(urlLabel, 1, w.rows, 1, 1)
		w.rows++
	}
}

func (w *secretPropertiesDisplay) renderProperties(propertyDefs api.SecretPropertyList, properties map[string]string, passwordStrengths map[string]*api.PasswordStrength) map[string]bool {
	renderedNames := map[string]bool{}
	for _, propertyDef := range propertyDefs {
		value, ok := properties[propertyDef.Name]
		if !ok {
			continue
		}
		renderedNames[propertyDef.Name] = true
		label := gtk.LabelNew(propertyDef.Display)
		label.SetHAlign(gtk.AlignStart)
		label.SetVAlign(gtk.AlignStart)
		w.widgets = append(w.widgets, label)
		w.Attach(label, 0, w.rows, 1, 1)

		valueDisplay := newSecretValueDisplay(value, propertyDef.Blurred, passwordStrengths[propertyDef.Name], w.logger)
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
