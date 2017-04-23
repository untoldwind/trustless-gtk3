package ui

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/pkg/errors"
	"github.com/untoldwind/trustless/api"
)

type passwordStrengthBar struct {
	*gtk.LevelBar
}

func newPasswordStrengthBar(passwordStrength *api.PasswordStrength) (*passwordStrengthBar, error) {
	levelBar, err := gtk.LevelBarNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create levelBar")
	}

	w := &passwordStrengthBar{
		LevelBar: levelBar,
	}

	w.SetMinValue(0)
	w.SetMaxValue(80)
	if passwordStrength != nil {
		w.setPasswordStrength(passwordStrength)
	}

	return w, nil
}

func (w *passwordStrengthBar) setPasswordStrength(passwordStrength *api.PasswordStrength) {
	w.SetValue(passwordStrength.Entropy)
	w.SetTooltipText(fmt.Sprintf("Entropy: %.1f Cracktime: %s", passwordStrength.Entropy, passwordStrength.CrackTimeDisplay))
}
