package ui

import (
	"time"

	"github.com/untoldwind/amintk/glib"

	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless/secrets/otp"
)

type otpDisplay struct {
	*gtk.Box
	userCodeLabel   *gtk.Label
	expirationLevel *gtk.LevelBar
	copyButton      *gtk.Button
	otp             otp.OTP
	userCode        string
	validFor        time.Duration
	timerHandle     glib.SourceHandle
	logger          logging.Logger
}

func newOTPDisplay(otpUrl string, parent logging.Logger) *otpDisplay {
	logger := parent.WithField("package", "ui").WithField("component", "otpDisplay")
	otp, err := otp.ParseURL(otpUrl)
	if err != nil {
		logger.ErrorErr(err)
	}

	w := &otpDisplay{
		Box:             gtk.BoxNew(gtk.OrientationHorizontal, 5),
		userCodeLabel:   gtk.LabelNew(""),
		expirationLevel: gtk.LevelBarNew(),
		copyButton:      gtk.ButtonNewFromIconName("edit-copy-symbolic", gtk.IconSizeButton),
		otp:             otp,
		logger:          logger,
	}

	labelBox := gtk.BoxNew(gtk.OrientationVertical, 0)
	labelBox.SetHExpand(true)
	w.Add(labelBox)

	labelBox.Add(w.userCodeLabel)
	labelBox.Add(w.expirationLevel)

	w.expirationLevel.SetMinValue(0)

	w.copyButton.SetTooltipText("Copy")
	w.copyButton.SetHAlign(gtk.AlignFill)
	w.copyButton.SetVAlign(gtk.AlignStart)
	w.copyButton.OnClicked(func() {
		safeCopy(w.logger, w.userCode)
	})
	w.Add(w.copyButton)

	w.updateUserCode()

	return w
}

func (w *otpDisplay) Destroy() {
	w.timerHandle.Remove()
	w.Box.Destroy()
}

func (w *otpDisplay) updateUserCode() {
	if w.otp == nil {
		w.copyButton.SetSensitive(false)
		w.userCodeLabel.SetText("*** Invalid OTP params ***")
		return
	}
	w.userCode, w.validFor = w.otp.GetUserCode()
	w.userCodeLabel.SetText(w.userCode)
	w.expirationLevel.SetMaxValue(float64(w.otp.MaxDuration()))
	w.expirationLevel.SetValue(float64(w.validFor))

	w.timerHandle.Remove()
	w.timerHandle, _ = glib.TimeoutAdd(1000, w.updateUserCode)
}
