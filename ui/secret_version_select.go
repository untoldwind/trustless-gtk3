package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless/api"
)

type secretVersionSelect struct {
	*gtk.Stack
	logger         logging.Logger
	timestampLabel *gtk.Label
	backButton     *gtk.Button
	versionsSelect *gtk.ComboBoxText
	forwardButton  *gtk.Button
	currentButton  *gtk.Button
	selectHandler  func(*api.SecretVersion)
	versions       api.SecretVersions
}

func newSecretVersionSelect(logger logging.Logger) (*secretVersionSelect, error) {
	stack, err := gtk.StackNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create stack")
	}

	timestampLabel, err := gtk.LabelNew("")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create timeStampLabel")
	}

	backButton, err := gtk.ButtonNewFromIconName("media-seek-backward-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create backButton")
	}

	forwardButton, err := gtk.ButtonNewFromIconName("media-seek-forward-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create forwardButton")
	}

	versionsSelect, err := gtk.ComboBoxTextNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create versionsSelect")
	}

	currentButton, err := gtk.ButtonNewFromIconName("media-skip-forward-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create currentButton")
	}

	w := &secretVersionSelect{
		Stack:          stack,
		logger:         logger.WithField("package", "ui").WithField("component", "secretVersionSelect"),
		timestampLabel: timestampLabel,
		backButton:     backButton,
		versionsSelect: versionsSelect,
		forwardButton:  forwardButton,
		currentButton:  currentButton,
	}

	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 2)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create box")
	}
	w.AddNamed(timestampLabel, "label")
	w.AddNamed(box, "box")

	w.backButton.Connect("clicked", w.onBack)
	box.Add(w.backButton)
	w.versionsSelect.SetHExpand(true)
	w.versionsSelect.Connect("changed", w.onVersionSelect)
	box.Add(w.versionsSelect)
	w.forwardButton.Connect("clicked", w.onForward)
	box.Add(w.forwardButton)
	w.currentButton.Connect("clicked", w.onCurrent)
	box.Add(w.currentButton)

	return w, nil
}

func (w *secretVersionSelect) onSelect(selectHandler func(*api.SecretVersion)) {
	w.selectHandler = selectHandler
}

func (w *secretVersionSelect) setVersions(versions api.SecretVersions) {
	w.versions = versions

	if len(w.versions) == 0 {
		w.SetVisibleChildName("label")
		w.timestampLabel.SetText("")
	} else if len(w.versions) == 1 {
		w.SetVisibleChildName("label")
		w.timestampLabel.SetText(w.versions[0].Timestamp.String())
	} else {
		w.SetVisibleChildName("box")
		w.versionsSelect.RemoveAll()
		for _, version := range w.versions {
			w.versionsSelect.AppendText(version.Timestamp.String())
		}
		w.versionsSelect.SetActive(0)
		w.updateButtons(0)
	}
}

func (w *secretVersionSelect) onBack() {
	idx := w.versionsSelect.GetActive()
	if idx+1 < len(w.versions) {
		w.versionsSelect.SetActive(idx + 1)
	}
}

func (w *secretVersionSelect) onVersionSelect() {
	idx := w.versionsSelect.GetActive()
	if idx < 0 || idx >= len(w.versions) {
		return
	}
	w.updateButtons(idx)
	if w.selectHandler != nil {
		w.selectHandler(w.versions[idx])
	}
}

func (w *secretVersionSelect) onForward() {
	idx := w.versionsSelect.GetActive()
	if idx > 0 {
		w.versionsSelect.SetActive(idx - 1)
	}

}

func (w *secretVersionSelect) onCurrent() {
	w.versionsSelect.SetActive(0)
}

func (w *secretVersionSelect) updateButtons(idx int) {
	w.versionsSelect.SetSensitive(len(w.versions) > 1)
	w.forwardButton.SetSensitive(idx > 0)
	w.currentButton.SetSensitive(idx > 0)
	w.backButton.SetSensitive(idx+1 < len(w.versions))
}
