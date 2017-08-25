package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
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

func newSecretVersionSelect(logger logging.Logger) *secretVersionSelect {
	stack := gtk.StackNew()
	timestampLabel := gtk.LabelNew("")
	backButton := gtk.ButtonNewFromIconName("media-seek-backward-symbolic", gtk.IconSizeButton)
	forwardButton := gtk.ButtonNewFromIconName("media-seek-forward-symbolic", gtk.IconSizeButton)
	versionsSelect := gtk.ComboBoxTextNew()
	currentButton := gtk.ButtonNewFromIconName("media-skip-forward-symbolic", gtk.IconSizeButton)

	w := &secretVersionSelect{
		Stack:          stack,
		logger:         logger.WithField("package", "ui").WithField("component", "secretVersionSelect"),
		timestampLabel: timestampLabel,
		backButton:     backButton,
		versionsSelect: versionsSelect,
		forwardButton:  forwardButton,
		currentButton:  currentButton,
	}

	box := gtk.BoxNew(gtk.OrientationHorizontal, 2)
	w.AddNamed(timestampLabel, "label")
	w.AddNamed(box, "box")

	w.backButton.OnClicked(w.onBack)
	box.Add(w.backButton)
	w.versionsSelect.SetHExpand(true)
	w.versionsSelect.OnChanged(w.onVersionSelect)
	box.Add(w.versionsSelect)
	w.forwardButton.OnClicked(w.onForward)
	box.Add(w.forwardButton)
	w.currentButton.OnClicked(w.onCurrent)
	box.Add(w.currentButton)

	return w
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
