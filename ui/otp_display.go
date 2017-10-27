package ui

import (
	"time"

	"github.com/untoldwind/amintk/gdk"

	qrcode "github.com/skip2/go-qrcode"
	"github.com/untoldwind/amintk/glib"

	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless/secrets/otp"
)

type otpDisplay struct {
	*gtk.Box
	userCodeLabel   *gtk.Label
	expirationLevel *gtk.LevelBar
	qrCodeButton    *gtk.MenuButton
	copyButton      *gtk.Button
	otp             otp.OTP
	userCode        string
	showUrl         bool
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
		qrCodeButton:    gtk.MenuButtonNew(),
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

	qrCodeIcon := gtk.ImageNewFromIconName("phone-symbolic", gtk.IconSizeButton)
	w.qrCodeButton.SetImage(qrCodeIcon)
	w.qrCodeButton.SetTooltipText("QR code")
	w.qrCodeButton.SetHAlign(gtk.AlignFill)
	w.qrCodeButton.SetVAlign(gtk.AlignStart)
	w.Add(w.qrCodeButton)

	qrCodePopover := gtk.PopoverNew(w.qrCodeButton)
	qrCodePopover.SetBorderWidth(5)
	w.qrCodeButton.SetPopover(qrCodePopover)

	qrCode, err := generateQRCodeImage(otpUrl)
	if err != nil {
		logger.ErrorErr(err)
	} else {
		qrCodePopover.Add(qrCode)
		qrCode.ShowAll()
	}

	blurredStack := gtk.StackNew()
	blurredStack.SetHAlign(gtk.AlignFill)
	blurredStack.SetVAlign(gtk.AlignStart)
	w.Add(blurredStack)

	revealButton := gtk.ButtonNewFromIconName("changes-allow-symbolic", gtk.IconSizeButton)

	revealButton.SetTooltipText("Reveal")
	revealButton.OnClicked(func() {
		blurredStack.SetVisibleChildName("hide")
		w.showUrl = true
		w.updateUserCode()
	})
	blurredStack.AddNamed(revealButton, "reveal")

	hideButton := gtk.ButtonNewFromIconName("changes-prevent-symbolic", gtk.IconSizeButton)

	hideButton.SetTooltipText("Hide")
	hideButton.OnClicked(func() {
		blurredStack.SetVisibleChildName("reveal")
		w.showUrl = false
		w.updateUserCode()
	})
	blurredStack.AddNamed(hideButton, "hide")

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
	if w.showUrl {
		w.userCodeLabel.SetText(w.otp.GetURL().String())
		w.expirationLevel.Hide()
		return
	}
	w.userCode, w.validFor = w.otp.GetUserCode()
	w.userCodeLabel.SetText(w.userCode)
	w.expirationLevel.SetMaxValue(float64(w.otp.MaxDuration()))
	w.expirationLevel.SetValue(float64(w.validFor))
	w.expirationLevel.Show()

	w.timerHandle.Remove()
	w.timerHandle, _ = glib.TimeoutAdd(1000, w.updateUserCode)
}

func generateQRCodeImage(url string) (*gtk.Image, error) {
	qr, err := qrcode.Encode(url, qrcode.Medium, 384)
	if err != nil {
		return nil, err
	}
	loader := gdk.PixbufLoaderNew()
	if _, err := loader.Write(qr); err != nil {
		return nil, err
	}
	if err := loader.Close(); err != nil {
		return nil, err
	}
	return gtk.ImageNewFromPixbuf(loader.GetPixbuf()), nil
}
