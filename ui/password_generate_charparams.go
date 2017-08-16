package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless/api"
)

type passwordGenerateCharParams struct {
	*gtk.Grid
	numChars         *gtk.SpinButton
	includeUpper     *gtk.Switch
	includeNumbers   *gtk.Switch
	includeSymbols   *gtk.Switch
	requireUpper     *gtk.Switch
	requireNumber    *gtk.Switch
	requireSymbol    *gtk.Switch
	excludeSimilar   *gtk.Switch
	excludeAmbiguous *gtk.Switch
}

func newPasswordGenerateCharParams(logger logging.Logger) *passwordGenerateCharParams {
	grid := gtk.GridNew()

	numChars := gtk.SpinButtonNewWithRange(4, 30, 1)
	includeUpper := gtk.SwitchNew()
	includeNumbers := gtk.SwitchNew()
	includeSymbols := gtk.SwitchNew()
	requireUpper := gtk.SwitchNew()
	requireNumber := gtk.SwitchNew()
	requireSymbol := gtk.SwitchNew()
	excludeSimilar := gtk.SwitchNew()
	excludeAmbiguous := gtk.SwitchNew()

	w := &passwordGenerateCharParams{
		Grid:             grid,
		numChars:         numChars,
		includeUpper:     includeUpper,
		includeNumbers:   includeNumbers,
		includeSymbols:   includeSymbols,
		requireUpper:     requireUpper,
		requireNumber:    requireNumber,
		requireSymbol:    requireSymbol,
		excludeSimilar:   excludeSimilar,
		excludeAmbiguous: excludeAmbiguous,
	}

	numCharsLabel := gtk.LabelNew("Chars")
	numCharsLabel.SetHAlign(gtk.AlignStart)
	w.Attach(numCharsLabel, 0, 0, 1, 1)
	w.numChars.SetHExpand(true)
	w.numChars.SetValue(14)
	w.Attach(w.numChars, 1, 0, 1, 1)

	includeUpperLabel := gtk.LabelNew("Include upper")
	includeUpperLabel.SetHAlign(gtk.AlignStart)
	w.Attach(includeUpperLabel, 0, 1, 1, 1)
	w.includeUpper.SetHAlign(gtk.AlignCenter)
	w.includeUpper.SetActive(true)
	w.Attach(w.includeUpper, 1, 1, 1, 1)

	includeNumbersLabel := gtk.LabelNew("Include numbers")
	includeNumbersLabel.SetHAlign(gtk.AlignStart)
	w.Attach(includeNumbersLabel, 0, 2, 1, 1)
	w.includeNumbers.SetHAlign(gtk.AlignCenter)
	w.includeNumbers.SetActive(true)
	w.Attach(w.includeNumbers, 1, 2, 1, 1)

	includeSymbolsLabel := gtk.LabelNew("Include symbols")
	includeSymbolsLabel.SetHAlign(gtk.AlignStart)
	w.Attach(includeSymbolsLabel, 0, 3, 1, 1)
	w.includeSymbols.SetHAlign(gtk.AlignCenter)
	w.includeSymbols.SetActive(true)
	w.Attach(w.includeSymbols, 1, 3, 1, 1)

	requireUpperLabel := gtk.LabelNew("Require upper")
	requireUpperLabel.SetHAlign(gtk.AlignStart)
	w.Attach(requireUpperLabel, 0, 4, 1, 1)
	w.requireUpper.SetHAlign(gtk.AlignCenter)
	w.Attach(w.requireUpper, 1, 4, 1, 1)

	requireSymbolLabel := gtk.LabelNew("Require symbol")
	requireSymbolLabel.SetHAlign(gtk.AlignStart)
	w.Attach(requireSymbolLabel, 0, 5, 1, 1)
	w.requireSymbol.SetHAlign(gtk.AlignCenter)
	w.Attach(w.requireSymbol, 1, 5, 1, 1)

	requireNumberLabel := gtk.LabelNew("Require number")
	requireNumberLabel.SetHAlign(gtk.AlignStart)
	w.Attach(requireNumberLabel, 0, 6, 1, 1)
	w.requireNumber.SetHAlign(gtk.AlignCenter)
	w.Attach(w.requireNumber, 1, 6, 1, 1)

	excludeSimilarLabel := gtk.LabelNew("Exclude similiar")
	excludeSimilarLabel.SetHAlign(gtk.AlignStart)
	w.Attach(excludeSimilarLabel, 0, 7, 1, 1)
	w.excludeSimilar.SetHAlign(gtk.AlignCenter)
	w.Attach(w.excludeSimilar, 1, 7, 1, 1)

	excludeAmbiguousLabel := gtk.LabelNew("Exclude ambiguous")
	excludeAmbiguousLabel.SetHAlign(gtk.AlignStart)
	w.Attach(excludeAmbiguousLabel, 0, 8, 1, 1)
	w.excludeAmbiguous.SetActive(true)
	w.excludeAmbiguous.SetHAlign(gtk.AlignCenter)
	w.Attach(w.excludeAmbiguous, 1, 8, 1, 1)

	return w
}

func (w *passwordGenerateCharParams) getParams() *api.CharsParameter {
	return &api.CharsParameter{
		NumChars:         int(w.numChars.GetValue()),
		IncludeUpper:     w.includeUpper.GetActive(),
		IncludeNumbers:   w.includeNumbers.GetActive(),
		IncludeSymbols:   w.includeSymbols.GetActive(),
		RequireUpper:     w.requireUpper.GetActive(),
		RequireSymbol:    w.requireSymbol.GetActive(),
		RequireNumber:    w.requireNumber.GetActive(),
		ExcludeSimilar:   w.excludeSimilar.GetActive(),
		ExcludeAmbiguous: w.excludeAmbiguous.GetActive(),
	}
}
