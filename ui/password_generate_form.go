package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless-gtk3/state"
	"github.com/untoldwind/trustless/api"
)

type passwordGenerateForm struct {
	*gtk.Grid
	logger                     logging.Logger
	store                      *state.Store
	entry                      *gtk.Entry
	stack                      *gtk.Stack
	passwordGenerateCharParams *passwordGenerateCharParams
	passwordGenerateWordParams *passwordGenerateWordParams
	takeHandler                func(string)
}

func newPasswordGenerateForm(store *state.Store, logger logging.Logger) *passwordGenerateForm {
	grid := gtk.GridNew()

	entry := gtk.EntryNew()
	stack := gtk.StackNew()
	passwordGenerateCharParams := newPasswordGenerateCharParams(logger)
	passwordGenerateWordParams := newPasswordGenerateWordParams(logger)

	w := &passwordGenerateForm{
		Grid:   grid,
		logger: logger.WithField("package", "ui").WithField("component", "passwordGenerateForm"),
		store:  store,
		entry:  entry,
		stack:  stack,
		passwordGenerateCharParams: passwordGenerateCharParams,
		passwordGenerateWordParams: passwordGenerateWordParams,
	}
	w.entry.SetHExpand(true)
	w.Attach(w.entry, 0, 0, 1, 1)

	refreshButton := gtk.ButtonNewFromIconName("view-refresh-symbolic", gtk.IconSizeButton)
	w.Attach(refreshButton, 1, 0, 1, 1)
	refreshButton.OnClicked(w.generate)

	stackSwitcher := gtk.StackSwitcherNew()
	w.Attach(stackSwitcher, 0, 1, 2, 1)

	stackSwitcher.SetStack(w.stack)
	w.Attach(w.stack, 0, 2, 2, 1)

	w.stack.AddTitled(w.passwordGenerateCharParams, "chars", "Chars")
	w.stack.AddTitled(w.passwordGenerateWordParams, "words", "Words")

	takeButton := gtk.ButtonNewWithLabel("Take")
	takeButton.OnClicked(w.onTake)
	w.Attach(takeButton, 0, 3, 2, 1)

	w.generate()

	return w
}

func (w *passwordGenerateForm) generate() {
	var parameter api.GenerateParameter

	switch w.stack.GetVisibleChildName() {
	case "chars":
		charParameters := w.passwordGenerateCharParams.getParams()
		parameter.Chars = charParameters
	case "words":
		wordParameters := w.passwordGenerateWordParams.getParams()
		parameter.Words = wordParameters
	}
	generated, err := w.store.GeneratePassword(parameter)
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
	w.entry.SetText(generated)
}

func (w *passwordGenerateForm) onTake() {
	if w.takeHandler != nil {
		password := w.entry.GetText()
		w.takeHandler(password)
	}
}

func (w *passwordGenerateForm) connectTake(takeHandler func(string)) {
	w.takeHandler = takeHandler
}
