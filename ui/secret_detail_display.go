package ui

import (
	"html"

	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
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

func newSecretDetailDisplay(logger logging.Logger) *secretDetailDisplay {
	typeNameMap := map[api.SecretType]string{}
	for _, secretType := range api.SecretTypes {
		typeNameMap[secretType.Type] = secretType.Display
	}

	scrolledWindow := gtk.ScrolledWindowNew(nil, nil)
	grid := gtk.GridNew()
	nameLabel := gtk.LabelNew("")
	typeLabel := gtk.LabelNew("")
	versionSelect := newSecretVersionSelect(logger)
	propertiesDisplay := newSecretPropertiesDisplay(logger)

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

	w.grid.SetOrientation(gtk.OrientationVertical)

	w.typeLabel.SetMarginStart(5)
	w.typeLabel.SetMarginEnd(5)
	w.grid.Attach(w.typeLabel, 0, 0, 1, 1)

	w.nameLabel.SetHExpand(true)
	w.grid.Attach(w.nameLabel, 1, 0, 1, 1)

	w.versionSelect.SetHAlign(gtk.AlignEnd)
	w.versionSelect.onSelect(w.onSelectVersion)
	w.grid.Attach(w.versionSelect, 0, 1, 2, 1)

	w.propertiesDisplay.SetHExpand(true)
	w.propertiesDisplay.SetVExpand(true)
	w.grid.Attach(w.propertiesDisplay, 0, 2, 2, 1)

	return w
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
