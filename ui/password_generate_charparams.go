package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
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

func newPasswordGenerateCharParams(logger logging.Logger) (*passwordGenerateCharParams, error) {
	grid, err := gtk.GridNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create grid")
	}

	numChars, err := gtk.SpinButtonNewWithRange(4, 30, 1)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create numChars")
	}
	includeUpper, err := gtk.SwitchNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create includeUpper")
	}
	includeNumbers, err := gtk.SwitchNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create includeNumbers")
	}
	includeSymbols, err := gtk.SwitchNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create includeSymbols")
	}
	requireUpper, err := gtk.SwitchNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create requireUpper")
	}
	requireNumber, err := gtk.SwitchNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create requireNumber")
	}
	requireSymbol, err := gtk.SwitchNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create requireSymbol")
	}
	excludeSimilar, err := gtk.SwitchNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create excludeSimilar")
	}
	excludeAmbiguous, err := gtk.SwitchNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create excludeAmbigous")
	}

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

	numCharsLabel, err := gtk.LabelNew("Chars")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create numCharsLabel")
	}
	numCharsLabel.SetHAlign(gtk.ALIGN_START)
	w.Attach(numCharsLabel, 0, 0, 1, 1)
	w.numChars.SetHExpand(true)
	w.numChars.SetValue(14)
	w.Attach(w.numChars, 1, 0, 1, 1)

	includeUpperLabel, err := gtk.LabelNew("Include upper")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create includeUpperLabel")
	}
	includeUpperLabel.SetHAlign(gtk.ALIGN_START)
	w.Attach(includeUpperLabel, 0, 1, 1, 1)
	w.includeUpper.SetHAlign(gtk.ALIGN_CENTER)
	w.includeUpper.SetActive(true)
	w.Attach(w.includeUpper, 1, 1, 1, 1)

	includeNumbersLabel, err := gtk.LabelNew("Include numbers")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create includeNumbersLabel")
	}
	includeNumbersLabel.SetHAlign(gtk.ALIGN_START)
	w.Attach(includeNumbersLabel, 0, 2, 1, 1)
	w.includeNumbers.SetHAlign(gtk.ALIGN_CENTER)
	w.includeNumbers.SetActive(true)
	w.Attach(w.includeNumbers, 1, 2, 1, 1)

	includeSymbolsLabel, err := gtk.LabelNew("Include symbols")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create includeSymbolsLabel")
	}
	includeSymbolsLabel.SetHAlign(gtk.ALIGN_START)
	w.Attach(includeSymbolsLabel, 0, 3, 1, 1)
	w.includeSymbols.SetHAlign(gtk.ALIGN_CENTER)
	w.includeSymbols.SetActive(true)
	w.Attach(w.includeSymbols, 1, 3, 1, 1)

	requireUpperLabel, err := gtk.LabelNew("Require upper")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create requireUpperLabel")
	}
	requireUpperLabel.SetHAlign(gtk.ALIGN_START)
	w.Attach(requireUpperLabel, 0, 4, 1, 1)
	w.requireUpper.SetHAlign(gtk.ALIGN_CENTER)
	w.Attach(w.requireUpper, 1, 4, 1, 1)

	requireSymbolLabel, err := gtk.LabelNew("Require symbol")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create requireSymbolLabel")
	}
	requireSymbolLabel.SetHAlign(gtk.ALIGN_START)
	w.Attach(requireSymbolLabel, 0, 5, 1, 1)
	w.requireSymbol.SetHAlign(gtk.ALIGN_CENTER)
	w.Attach(w.requireSymbol, 1, 5, 1, 1)

	requireNumberLabel, err := gtk.LabelNew("Require number")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create requireNumberLabel")
	}
	requireNumberLabel.SetHAlign(gtk.ALIGN_START)
	w.Attach(requireNumberLabel, 0, 6, 1, 1)
	w.requireNumber.SetHAlign(gtk.ALIGN_CENTER)
	w.Attach(w.requireNumber, 1, 6, 1, 1)

	excludeSimilarLabel, err := gtk.LabelNew("Exclude similiar")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create excludeSimilarLabel")
	}
	excludeSimilarLabel.SetHAlign(gtk.ALIGN_START)
	w.Attach(excludeSimilarLabel, 0, 7, 1, 1)
	w.excludeSimilar.SetHAlign(gtk.ALIGN_CENTER)
	w.Attach(w.excludeSimilar, 1, 7, 1, 1)

	excludeAmbiguousLabel, err := gtk.LabelNew("Exclude ambiguous")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create excludeAmbiguousLabel")
	}
	excludeAmbiguousLabel.SetHAlign(gtk.ALIGN_START)
	w.Attach(excludeAmbiguousLabel, 0, 8, 1, 1)
	w.excludeAmbiguous.SetActive(true)
	w.excludeAmbiguous.SetHAlign(gtk.ALIGN_CENTER)
	w.Attach(w.excludeAmbiguous, 1, 8, 1, 1)

	return w, nil
}

func (w *passwordGenerateCharParams) getParams() (*api.CharsParameter, error) {
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
	}, nil
}
