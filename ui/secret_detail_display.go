package ui

import (
	"html"

	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless/api"
)

type secretDetailDisplay struct {
	*gtk.ScrolledWindow
	grid              *gtk.Grid
	nameLabel         *gtk.Label
	versionSelect     *secretVersionSelect
	typeLabel         *gtk.Label
	propertiesDisplay *secretPropertiesDisplay
	logger            logging.Logger
	typeNameMap       map[api.SecretType]string
	displayedSecret   *api.Secret
	displayedVersion  *api.SecretVersion
}

func newSecretDetailDisplay(logger logging.Logger) (*secretDetailDisplay, error) {
	typeNameMap := map[api.SecretType]string{}
	for _, secretType := range api.SecretTypes {
		typeNameMap[secretType.Type] = secretType.Display
	}

	scrolledWindow, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create scrolled window")
	}
	grid, err := gtk.GridNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create grid")
	}
	nameLabel, err := gtk.LabelNew("")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create nameLabel")
	}
	typeLabel, err := gtk.LabelNew("")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create typeLabel")
	}
	versionSelect, err := newSecretVersionSelect(logger)
	if err != nil {
		return nil, err
	}
	propertiesDisplay, err := newSecretPropertiesDisplay(logger)
	if err != nil {
		return nil, err
	}

	w := &secretDetailDisplay{
		ScrolledWindow:    scrolledWindow,
		grid:              grid,
		nameLabel:         nameLabel,
		typeLabel:         typeLabel,
		propertiesDisplay: propertiesDisplay,
		versionSelect:     versionSelect,
		logger:            logger.WithField("package", "ui").WithField("component", "secretDetailDisplay"),
		typeNameMap:       typeNameMap,
	}
	w.Add(w.grid)

	w.grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	w.typeLabel.SetMarginStart(5)
	w.typeLabel.SetMarginEnd(5)
	w.grid.Attach(w.typeLabel, 0, 0, 1, 1)

	w.nameLabel.SetHExpand(true)
	w.grid.Attach(w.nameLabel, 1, 0, 1, 1)

	w.versionSelect.SetHAlign(gtk.ALIGN_END)
	w.versionSelect.onSelect(w.onSelectVersion)
	w.grid.Attach(w.versionSelect, 0, 1, 2, 1)

	w.propertiesDisplay.SetHExpand(true)
	w.propertiesDisplay.SetVExpand(true)
	w.grid.Attach(w.propertiesDisplay, 0, 2, 2, 1)

	return w, nil
}

func (w *secretDetailDisplay) display(secret *api.Secret) {
	typeNameDisplay, ok := w.typeNameMap[secret.Type]
	if !ok {
		typeNameDisplay = string(secret.Type)
	}
	w.typeLabel.SetText(typeNameDisplay)
	w.displayedSecret = secret
	w.displayedVersion = secret.Versions[0]
	w.versionSelect.setVersions(secret.Versions)
	w.updateView()
}

func (w *secretDetailDisplay) onSelectVersion(version *api.SecretVersion) {
	if version == w.displayedVersion {
		return
	}
	w.displayedVersion = version
	w.updateView()
}

func (w *secretDetailDisplay) updateView() {
	w.nameLabel.SetMarkup("<span font=\"20\">" + html.EscapeString(w.displayedVersion.Name) + "</span>")
	if w.displayedSecret.Versions[0] == w.displayedVersion {
		w.propertiesDisplay.display(w.displayedVersion, w.displayedSecret.PasswordStrengths)
	} else {
		w.propertiesDisplay.display(w.displayedVersion, nil)
	}
}
