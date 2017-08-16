package ui

import (
	"fmt"

	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless/api"
)

type passwordStrengthBar struct {
	*gtk.LevelBar
}

func newPasswordStrengthBar(passwordStrength *api.PasswordStrength) *passwordStrengthBar {
	levelBar := gtk.LevelBarNew()

	w := &passwordStrengthBar{
		LevelBar: levelBar,
	}

	w.SetMinValue(0)
	w.SetMaxValue(80)
	if passwordStrength != nil {
		w.setPasswordStrength(passwordStrength)
	}

	return w
}

func (w *passwordStrengthBar) setPasswordStrength(passwordStrength *api.PasswordStrength) {
	w.SetValue(passwordStrength.Entropy)
	w.SetTooltipText(fmt.Sprintf("Entropy: %.1f Cracktime: %s", passwordStrength.Entropy, passwordStrength.CrackTimeDisplay))
}
