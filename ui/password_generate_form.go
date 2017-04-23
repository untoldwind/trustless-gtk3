package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless-gtk3/gtkextra"
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
	handleRefs                 gtkextra.HandleRefs
	takeHandler                func(string)
}

func newPasswordGenerateForm(store *state.Store, logger logging.Logger) (*passwordGenerateForm, error) {
	grid, err := gtk.GridNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create box")
	}

	entry, err := gtk.EntryNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create entry")
	}
	stack, err := gtk.StackNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create stack")
	}
	passwordGenerateCharParams, err := newPasswordGenerateCharParams(logger)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create passwordGenerateCharParams")
	}
	passwordGenerateWordParams, err := newPasswordGenerateWordParams(logger)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create passwordGenerateWordParams")
	}

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

	refreshButton, err := gtk.ButtonNewFromIconName("view-refresh-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create refreshButton")
	}
	w.Attach(refreshButton, 1, 0, 1, 1)
	w.handleRefs.SafeConnect(refreshButton.Object, "clicked", w.generate)

	stackSwitcher, err := gtk.StackSwitcherNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create stackSwitcher")
	}
	w.Attach(stackSwitcher, 0, 1, 2, 1)

	stackSwitcher.SetStack(w.stack)
	w.Attach(w.stack, 0, 2, 2, 1)

	w.stack.AddTitled(w.passwordGenerateCharParams, "chars", "Chars")
	w.stack.AddTitled(w.passwordGenerateWordParams, "words", "Words")

	takeButton, err := gtk.ButtonNewWithLabel("Take")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create takeButton")
	}
	w.handleRefs.SafeConnect(takeButton.Object, "clicked", w.onTake)
	w.Attach(takeButton, 0, 3, 2, 1)

	w.generate()

	return w, nil
}

func (w *passwordGenerateForm) generate() {
	var parameter api.GenerateParameter

	switch w.stack.GetVisibleChildName() {
	case "chars":
		charParameters, err := w.passwordGenerateCharParams.getParams()
		if err != nil {
			w.logger.ErrorErr(err)
			return
		}
		parameter.Chars = charParameters
	case "words":
		wordParameters, err := w.passwordGenerateWordParams.getParams()
		if err != nil {
			w.logger.ErrorErr(err)
			return
		}
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
		password, err := w.entry.GetText()
		if err != nil {
			w.logger.ErrorErr(err)
			return
		}
		w.takeHandler(password)
	}
}

func (w *passwordGenerateForm) connectTake(takeHandler func(string)) {
	w.takeHandler = takeHandler
}

func (w *passwordGenerateForm) Destroy() {
	w.handleRefs.Cleanup()
	w.Grid.Destroy()
}
