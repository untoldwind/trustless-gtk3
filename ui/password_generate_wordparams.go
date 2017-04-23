package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless/api"
)

type passwordGenerateWordParams struct {
	*gtk.Grid
	numWords *gtk.SpinButton
	delim    *gtk.Entry
}

func newPasswordGenerateWordParams(logger logging.Logger) (*passwordGenerateWordParams, error) {
	grid, err := gtk.GridNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create grid")
	}
	numWords, err := gtk.SpinButtonNewWithRange(4, 30, 1)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create numWords")
	}
	delim, err := gtk.EntryNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create delim")
	}

	w := &passwordGenerateWordParams{
		Grid:     grid,
		numWords: numWords,
		delim:    delim,
	}

	numWordsLabel, err := gtk.LabelNew("Words")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create numWordsLabel")
	}
	numWordsLabel.SetHAlign(gtk.ALIGN_START)
	w.Attach(numWordsLabel, 0, 0, 1, 1)
	w.numWords.SetHExpand(true)
	w.numWords.SetValue(4)
	w.Attach(w.numWords, 1, 0, 1, 1)

	delimLabel, err := gtk.LabelNew("Delim")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create delimLabel")
	}
	delimLabel.SetHAlign(gtk.ALIGN_START)
	w.Attach(delimLabel, 0, 1, 1, 1)
	w.delim.SetHExpand(true)
	w.delim.SetText(".")
	w.Attach(w.delim, 1, 1, 1, 1)

	return w, nil
}

func (w *passwordGenerateWordParams) getParams() (*api.WordsParameter, error) {
	delim, err := w.delim.GetText()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get delim")
	}
	return &api.WordsParameter{
		NumWords: int(w.numWords.GetValue()),
		Delim:    delim,
	}, nil
}
