package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless-gtk3/state"
	"github.com/untoldwind/trustless/api"
)

type propertyValueGetter func() (string, error)

type secretPropertiesEdit struct {
	*gtk.Grid
	logger          logging.Logger
	store           *state.Store
	widgets         []destroyable
	urlsEdit        *urlsEdit
	propertyGetters map[string]propertyValueGetter
	rows            int
}

func newSecretPropertiesEdit(store *state.Store, logger logging.Logger) (*secretPropertiesEdit, error) {
	grid, err := gtk.GridNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create grid")
	}
	w := &secretPropertiesEdit{
		Grid:   grid,
		logger: logger.WithField("package", "ui").WithField("component", "secretPropertiesEdit"),
		store:  store,
	}
	w.SetRowSpacing(2)
	w.SetColumnSpacing(2)

	return w, nil
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
		value, err := getter()
		if err != nil {
			return nil, nil, errors.Wrapf(err, "Failed to get value for: %s", name)
		}
		properties[name] = value
	}

	return urls, properties, nil
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

	w.urlsEdit, err = newUrlsEdit(w.logger)
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
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
		label, err := gtk.LabelNew(propertyDef.Display)
		if err != nil {
			w.logger.ErrorErr(err)
			continue
		}
		label.SetHAlign(gtk.ALIGN_START)
		label.SetVAlign(gtk.ALIGN_START)
		w.widgets = append(w.widgets, label)
		w.Attach(label, 0, w.rows, 1, 1)

		if propertyDef.MultiLine {
			valueEdit, err := gtk.TextViewNew()
			if err != nil {
				w.logger.ErrorErr(err)
				continue
			}
			valueEdit.SetHExpand(true)
			valueEdit.SetVExpand(true)
			w.widgets = append(w.widgets, valueEdit)
			w.Attach(valueEdit, 1, w.rows, 1, 1)

			value, ok := properties[propertyDef.Name]
			if ok {
				buffer, err := valueEdit.GetBuffer()
				if err != nil {
					w.logger.ErrorErr(err)
					continue
				}
				buffer.SetText(value)
			}
			w.propertyGetters[propertyDef.Name] = func() (string, error) {
				buffer, err := valueEdit.GetBuffer()
				if err != nil {
					return "", err
				}
				start := buffer.GetStartIter()
				end := buffer.GetEndIter()
				return buffer.GetText(start, end, true)
			}
		} else if propertyDef.Blurred {
			passwordEdit, err := newSecretPasswordEdit(w.store, w.logger)
			if err != nil {
				w.logger.ErrorErr(err)
				continue
			}
			passwordEdit.SetHExpand(true)
			w.widgets = append(w.widgets, passwordEdit)
			w.Attach(passwordEdit, 1, w.rows, 1, 1)

			value, ok := properties[propertyDef.Name]
			if ok {
				passwordEdit.setText(value)
			}
			w.propertyGetters[propertyDef.Name] = passwordEdit.getText
		} else {
			valueEdit, err := gtk.EntryNew()
			if err != nil {
				w.logger.ErrorErr(err)
				continue
			}
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
