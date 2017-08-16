package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless/api"
)

type passwordGenerateWordParams struct {
	*gtk.Grid
	numWords *gtk.SpinButton
	delim    *gtk.Entry
}

func newPasswordGenerateWordParams(logger logging.Logger) *passwordGenerateWordParams {
	grid := gtk.GridNew()
	numWords := gtk.SpinButtonNewWithRange(4, 30, 1)
	delim := gtk.EntryNew()

	w := &passwordGenerateWordParams{
		Grid:     grid,
		numWords: numWords,
		delim:    delim,
	}

	numWordsLabel := gtk.LabelNew("Words")
	numWordsLabel.SetHAlign(gtk.AlignStart)
	w.Attach(numWordsLabel, 0, 0, 1, 1)
	w.numWords.SetHExpand(true)
	w.numWords.SetValue(4)
	w.Attach(w.numWords, 1, 0, 1, 1)

	delimLabel := gtk.LabelNew("Delim")
	delimLabel.SetHAlign(gtk.AlignStart)
	w.Attach(delimLabel, 0, 1, 1, 1)
	w.delim.SetHExpand(true)
	w.delim.SetText(".")
	w.Attach(w.delim, 1, 1, 1, 1)

	return w
}

func (w *passwordGenerateWordParams) getParams() *api.WordsParameter {
	delim := w.delim.GetText()
	return &api.WordsParameter{
		NumWords: int(w.numWords.GetValue()),
		Delim:    delim,
	}
}
