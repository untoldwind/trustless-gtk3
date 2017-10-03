package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless-gtk3/state"
	"github.com/untoldwind/trustless/api"
)

type propertyValueGetter func() string

type secretPropertiesEdit struct {
	*gtk.Grid
	logger          logging.Logger
	store           *state.Store
	widgets         []destroyable
	urlsEdit        *urlsEdit
	propertyGetters map[string]propertyValueGetter
	rows            int
}

func newSecretPropertiesEdit(store *state.Store, logger logging.Logger) *secretPropertiesEdit {
	grid := gtk.GridNew()
	w := &secretPropertiesEdit{
		Grid:   grid,
		logger: logger.WithField("package", "ui").WithField("component", "secretPropertiesEdit"),
		store:  store,
	}
	w.SetRowSpacing(2)
	w.SetColumnSpacing(2)

	return w
}

func (w *secretPropertiesEdit) setEdit(secret *api.SecretCurrent) {
	secretVersion := secret.Current

	w.destroyAllChildren()

	w.renderUrls(secretVersion.URLs)

	knownNames := map[string]bool{}
	for _, secretType := range api.SecretTypes {
		if secretType.Type != secret.Type {
			continue
		}
		knownNames = w.renderProperties(secretType.Properties, secretVersion.Properties)
		break
	}
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

func (w *secretPropertiesEdit) getEdit() ([]string, map[string]string, error) {
	var urls []string

	if w.urlsEdit != nil {
		urls = w.urlsEdit.getUrls()
	}
	properties := map[string]string{}
	for name, getter := range w.propertyGetters {
		properties[name] = getter()
	}

	return urls, properties, nil
}

func (w *secretPropertiesEdit) renderUrls(urls []string) {
	label := gtk.LabelNew("URLs")
	label.SetHAlign(gtk.AlignStart)
	label.SetVAlign(gtk.AlignStart)
	w.widgets = append(w.widgets, label)
	w.Attach(label, 0, w.rows, 1, 1)

	w.urlsEdit = newUrlsEdit(w.logger)
	w.urlsEdit.SetHExpand(true)
	w.urlsEdit.setUrls(urls)
	w.widgets = append(w.widgets, w.urlsEdit)
	w.Attach(w.urlsEdit, 1, w.rows, 1, 1)

	w.rows++
}

func (w *secretPropertiesEdit) renderProperties(propertyDefs api.SecretPropertyList, properties map[string]string) map[string]bool {
	renderedNames := map[string]bool{}
	for _, propertyDef := range propertyDefs {
		renderedNames[propertyDef.Name] = true
		label := gtk.LabelNew(propertyDef.Display)
		label.SetHAlign(gtk.AlignStart)
		label.SetVAlign(gtk.AlignStart)
		w.widgets = append(w.widgets, label)
		w.Attach(label, 0, w.rows, 1, 1)

		if propertyDef.MultiLine {
			valueEdit := gtk.TextViewNew()
			valueEdit.SetHExpand(true)
			valueEdit.SetVExpand(true)
			w.widgets = append(w.widgets, valueEdit)
			w.Attach(valueEdit, 1, w.rows, 1, 1)

			value, ok := properties[propertyDef.Name]
			if ok {
				buffer := valueEdit.GetBuffer()
				buffer.SetText(value)
			}
			w.propertyGetters[propertyDef.Name] = func() string {
				buffer := valueEdit.GetBuffer()
				start := buffer.GetStartIter()
				end := buffer.GetEndIter()
				return buffer.GetText(start, end, true)
			}
		} else if propertyDef.OTPParams {
			otpEdit := newOTPEdit(w.store, w.logger)
			w.widgets = append(w.widgets, otpEdit)
			w.Attach(otpEdit, 1, w.rows, 1, 1)

			value, ok := properties[propertyDef.Name]
			if ok {
				otpEdit.setValue(value)
			}
			w.propertyGetters[propertyDef.Name] = otpEdit.getValue
		} else if propertyDef.Blurred {
			passwordEdit := newSecretPasswordEdit(w.store, w.logger)
			passwordEdit.SetHExpand(true)
			w.widgets = append(w.widgets, passwordEdit)
			w.Attach(passwordEdit, 1, w.rows, 1, 1)

			value, ok := properties[propertyDef.Name]
			if ok {
				passwordEdit.setText(value)
			}
			w.propertyGetters[propertyDef.Name] = passwordEdit.getText
		} else {
			valueEdit := gtk.EntryNew()
			valueEdit.SetHExpand(true)
			w.widgets = append(w.widgets, valueEdit)
			w.Attach(valueEdit, 1, w.rows, 1, 1)

			value, ok := properties[propertyDef.Name]
			if ok {
				valueEdit.SetText(value)
			}
			w.propertyGetters[propertyDef.Name] = valueEdit.GetText
		}

		w.rows++
	}

	return renderedNames
}

func (w *secretPropertiesEdit) destroyAllChildren() {
	w.propertyGetters = map[string]propertyValueGetter{}
	w.urlsEdit = nil
	for _, widget := range w.widgets {
		w.Remove(widget)
		widget.Destroy()
	}
	w.widgets = nil
	w.rows = 0
}
