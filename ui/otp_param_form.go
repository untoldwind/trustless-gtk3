package ui

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"strings"
	"time"

	"github.com/untoldwind/amintk/glib"

	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless-gtk3/state"
	"github.com/untoldwind/trustless/secrets/otp"
)

type otpParamForm struct {
	*gtk.Grid
	secretEntry    *gtk.Entry
	digitsSpin     *gtk.SpinButton
	periodSpin     *gtk.SpinButton
	algorithmCombo *gtk.ComboBoxText
	labelEntry     *gtk.Entry
	issuerEntry    *gtk.Entry
	totp           *otp.TOTP
	changeCallback func(string)
	logger         logging.Logger
	handles        glib.SignalHandles
}

func newOTPParamForm(store *state.Store, logger logging.Logger) *otpParamForm {
	w := &otpParamForm{
		Grid:           gtk.GridNew(),
		secretEntry:    gtk.EntryNew(),
		digitsSpin:     gtk.SpinButtonNewWithRange(6, 9, 1),
		periodSpin:     gtk.SpinButtonNewWithRange(30, 60, 10),
		algorithmCombo: gtk.ComboBoxTextNew(),
		labelEntry:     gtk.EntryNew(),
		issuerEntry:    gtk.EntryNew(),
		totp:           otp.NewTOTP([]byte{}),
		logger:         logger.WithField("package", "ui").WithField("component", "otpParamForm"),
	}

	w.Attach(gtk.LabelNew("Secret"), 0, 0, 1, 1)
	w.secretEntry.SetHExpand(true)
	w.handles.Add(w.secretEntry.OnChanged(w.onEntryChanged))
	w.Attach(w.secretEntry, 1, 0, 1, 1)

	w.Attach(gtk.LabelNew("Digits"), 0, 1, 1, 1)
	w.digitsSpin.SetHExpand(true)
	w.handles.Add(w.digitsSpin.OnChanged(w.onEntryChanged))
	w.Attach(w.digitsSpin, 1, 1, 1, 1)

	w.Attach(gtk.LabelNew("Period"), 0, 2, 1, 1)
	w.periodSpin.SetHExpand(true)
	w.handles.Add(w.periodSpin.OnChanged(w.onEntryChanged))
	w.Attach(w.periodSpin, 1, 2, 1, 1)

	w.Attach(gtk.LabelNew("Algorithm"), 0, 3, 1, 1)
	w.algorithmCombo.SetHExpand(true)
	w.handles.Add(w.algorithmCombo.OnChanged(w.onEntryChanged))
	w.Attach(w.algorithmCombo, 1, 3, 1, 1)
	w.algorithmCombo.AppendText("SHA-1")
	w.algorithmCombo.AppendText("SHA-256")
	w.algorithmCombo.AppendText("SHA-512")

	w.Attach(gtk.LabelNew("Label"), 0, 4, 1, 1)
	w.labelEntry.SetHExpand(true)
	w.handles.Add(w.labelEntry.OnChanged(w.onEntryChanged))
	w.Attach(w.labelEntry, 1, 4, 1, 1)

	w.Attach(gtk.LabelNew("Issuer"), 0, 5, 1, 1)
	w.issuerEntry.SetHExpand(true)
	w.handles.Add(w.issuerEntry.OnChanged(w.onEntryChanged))
	w.Attach(w.issuerEntry, 1, 5, 1, 1)

	w.updateForm()

	return w
}

func (w *otpParamForm) setChangeCallback(changeCallback func(string)) {
	w.changeCallback = changeCallback
}

func (w *otpParamForm) onEntryChanged() {
	if w.totp == nil {
		return
	}
	w.totp.SetEncodedSecret(w.secretEntry.GetText())
	w.totp.Digits = uint8(w.digitsSpin.GetValue())
	w.totp.TimeStep = time.Duration(w.periodSpin.GetValue()) * time.Second
	switch w.algorithmCombo.GetActive() {
	case 0:
		w.totp.Hash = sha1.New
	case 1:
		w.totp.Hash = sha256.New
	case 2:
		w.totp.Hash = sha512.New
	}
	label := w.labelEntry.GetText()
	issuer := w.issuerEntry.GetText()
	if issuer != "" {
		w.totp.Label = issuer + ":" + label
		w.totp.Issuer = issuer
	} else {
		w.totp.Label = label
		w.totp.Issuer = ""
	}
	if w.changeCallback != nil {
		w.changeCallback(w.totp.GetURL().String())
	}
}

func (w *otpParamForm) setParams(otpUrl string) {
	if otpUrl == "" {
		w.totp = otp.NewTOTP([]byte{})
	} else {
		parsed, err := otp.ParseURL(otpUrl)
		if err != nil {
			w.logger.ErrorErr(err)
			return
		}
		totp, ok := parsed.(*otp.TOTP)
		if !ok {
			w.logger.Warn("Can handle anything but TOTP yet")
			return
		}
		w.totp = totp
	}
	w.updateForm()
}

func (w *otpParamForm) updateForm() {
	w.handles.BlockAll()
	w.secretEntry.SetText(w.totp.GetEncodedSecret())
	w.secretEntry.SetSensitive(true)
	w.digitsSpin.SetValue(float64(w.totp.Digits))
	w.digitsSpin.SetSensitive(true)
	w.periodSpin.SetValue(float64(w.totp.TimeStep / time.Second))
	w.periodSpin.SetSensitive(true)
	switch w.totp.Hash().Size() {
	case 20:
		w.algorithmCombo.SetActive(0)
	case 32:
		w.algorithmCombo.SetActive(1)
	case 64:
		w.algorithmCombo.SetActive(2)
	}
	splitLabel := strings.Split(w.totp.Label, ":")
	if len(splitLabel) == 2 {
		w.labelEntry.SetText(splitLabel[1])
	} else {
		w.labelEntry.SetText(w.totp.Label)
	}
	w.labelEntry.SetSensitive(true)
	w.issuerEntry.SetText(w.totp.Issuer)
	w.issuerEntry.SetSensitive(true)
	w.handles.UnblockAll()
}

func (w *otpParamForm) disable() {
	w.secretEntry.SetSensitive(false)
	w.digitsSpin.SetSensitive(false)
	w.periodSpin.SetSensitive(false)
	w.algorithmCombo.SetSensitive(false)
	w.labelEntry.SetSensitive(false)
	w.issuerEntry.SetSensitive(false)
}
