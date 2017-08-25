package ui

import (
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless-gtk3/state"
	"github.com/untoldwind/trustless/api"
)

type secretDetail struct {
	*gtk.Box
	stack               *gtk.Stack
	secretDetailDisplay *secretDetailDisplay
	secretDetailEdit    *secretDetailEdit
	newButton           *gtk.MenuButton
	changeButton        *gtk.Button
	deleteButton        *gtk.MenuButton
	abortEditButton     *gtk.Button
	saveEditButton      *gtk.Button
	logger              logging.Logger
	store               *state.Store
	secretID            string
	editing             bool
}

func newSecretDetail(store *state.Store, logger logging.Logger) *secretDetail {
	box := gtk.BoxNew(gtk.OrientationVertical, 0)
	stack := gtk.StackNew()
	newButton := gtk.MenuButtonNew()
	changeButton := gtk.ButtonNewFromIconName("document-open-symbolic", gtk.IconSizeButton)
	deleteButton := gtk.MenuButtonNew()
	abortEditButton := gtk.ButtonNewFromIconName("window-close-symbolic", gtk.IconSizeButton)
	saveEditButton := gtk.ButtonNewFromIconName("document-save-symbolic", gtk.IconSizeButton)

	w := &secretDetail{
		Box:             box,
		stack:           stack,
		newButton:       newButton,
		changeButton:    changeButton,
		deleteButton:    deleteButton,
		abortEditButton: abortEditButton,
		saveEditButton:  saveEditButton,
		logger:          logger.WithField("package", "ui").WithField("component", "secretDetail"),
		store:           store,
	}

	w.stack.SetHExpand(true)
	w.stack.SetVExpand(true)
	w.Add(w.stack)

	buttonBox := gtk.BoxNew(gtk.OrientationHorizontal, 5)
	w.Add(buttonBox)

	newImage := gtk.ImageNewFromIconName("list-add-symbolic", gtk.IconSizeButton)

	newMenu := w.newSecretMenu()

	w.newButton.SetImage(newImage)
	w.newButton.SetHAlign(gtk.AlignStart)
	w.newButton.SetMarginTop(2)
	w.newButton.SetMarginStart(2)
	w.newButton.SetMarginEnd(2)
	w.newButton.SetMarginBottom(2)
	w.newButton.SetPopup(newMenu)
	w.newButton.SetDirection(gtk.ArrowTypeUp)
	buttonBox.Add(w.newButton)

	w.changeButton.SetLabel("Change")
	w.changeButton.SetAlwaysShowImage(true)
	w.changeButton.SetHAlign(gtk.AlignStart)
	w.changeButton.SetMarginTop(2)
	w.changeButton.SetMarginStart(2)
	w.changeButton.SetMarginEnd(2)
	w.changeButton.SetMarginBottom(2)
	w.changeButton.SetNoShowAll(true)
	w.changeButton.OnClicked(w.store.ActionEditCurrent)
	buttonBox.Add(w.changeButton)

	w.abortEditButton.SetLabel("Abort")
	w.abortEditButton.SetAlwaysShowImage(true)
	w.abortEditButton.SetHAlign(gtk.AlignStart)
	w.abortEditButton.SetMarginTop(2)
	w.abortEditButton.SetMarginStart(2)
	w.abortEditButton.SetMarginEnd(2)
	w.abortEditButton.SetMarginBottom(2)
	w.abortEditButton.SetNoShowAll(true)
	w.abortEditButton.OnClicked(w.store.ActionEditAbort)
	buttonBox.Add(w.abortEditButton)

	deleteConfirm := gtk.ButtonNewWithLabel("Confirm")
	deleteConfirm.OnClicked(w.onDelete)
	deleteConfirm.Show()
	confirmPopover := gtk.PopoverNew(w.deleteButton)
	confirmPopover.Add(deleteConfirm)
	confirmPopover.SetBorderWidth(5)

	deleteImage := gtk.ImageNewFromIconName("edit-delete-symbolic", gtk.IconSizeButton)
	w.deleteButton.SetImage(deleteImage)
	w.deleteButton.SetLabel("Delete")
	w.deleteButton.SetAlwaysShowImage(true)
	w.deleteButton.SetHAlign(gtk.AlignEnd)
	w.deleteButton.SetHExpand(true)
	w.deleteButton.SetMarginTop(2)
	w.deleteButton.SetMarginStart(2)
	w.deleteButton.SetMarginEnd(2)
	w.deleteButton.SetMarginBottom(2)
	w.deleteButton.SetPopover(confirmPopover)
	w.deleteButton.SetNoShowAll(true)
	buttonBox.Add(w.deleteButton)

	w.saveEditButton.SetLabel("Save")
	w.saveEditButton.SetAlwaysShowImage(true)
	w.saveEditButton.SetHAlign(gtk.AlignEnd)
	w.saveEditButton.SetHExpand(true)
	w.saveEditButton.SetMarginTop(2)
	w.saveEditButton.SetMarginStart(2)
	w.saveEditButton.SetMarginEnd(2)
	w.saveEditButton.SetMarginBottom(2)
	w.saveEditButton.SetNoShowAll(true)
	w.saveEditButton.OnClicked(w.onEditSave)
	buttonBox.Add(w.saveEditButton)

	placeholder := newSecretDetailPlaceholder()
	w.stack.AddNamed(placeholder, "placeholder")

	w.secretDetailDisplay = newSecretDetailDisplay(logger)
	w.stack.AddNamed(w.secretDetailDisplay, "display")

	w.secretDetailEdit = newSecretDetailEdit(store, logger)
	w.stack.AddNamed(w.secretDetailEdit, "edit")

	w.store.AddListener(w.onStateChanged)

	return w
}

func (w *secretDetail) newSecretMenu() *gtk.Menu {
	menu := gtk.MenuNew()

	for _, secretTypeDefinition := range api.SecretTypes {
		item := gtk.MenuItemNewWithLabel(secretTypeDefinition.Display)
		secretType := secretTypeDefinition.Type
		item.OnActivate(func() {
			w.store.ActionEditNew(secretType)
		})
		item.Show()
		menu.Append(item)
	}

	return menu
}

func (w *secretDetail) onDelete() {
	if w.secretID != "" {
		w.store.ActionMarkDeleted(w.secretID)
	}
}

func (w *secretDetail) onEditSave() {
	if w.secretID == "" {
		return
	}
	version, err := w.secretDetailEdit.getEdit()
	if err != nil {
		w.logger.ErrorErr(err)
		return
	}
	w.store.ActionEditStore(w.secretID, *version)
}

func (w *secretDetail) onStateChanged(prev, next *state.State) {
	if next.CurrentSecret == nil && w.secretID != "" {
		w.secretID = ""
		w.newButton.Show()
		w.changeButton.Hide()
		w.deleteButton.Hide()
		w.abortEditButton.Hide()
		w.saveEditButton.Hide()
		w.stack.SetVisibleChildName("placeholder")
	} else if next.CurrentSecret != nil && next.CurrentEdit && (next.CurrentSecret.ID != w.secretID || !w.editing) {
		w.secretID = next.CurrentSecret.ID
		w.editing = true
		w.newButton.Hide()
		w.changeButton.Hide()
		w.deleteButton.Hide()
		w.abortEditButton.Show()
		w.saveEditButton.Show()
		w.stack.SetVisibleChildName("edit")
		w.secretDetailEdit.setEdit(&next.CurrentSecret.SecretCurrent)
	} else if next.CurrentSecret != nil && !next.CurrentEdit && (next.CurrentSecret.ID != w.secretID || w.editing) {
		w.secretID = next.CurrentSecret.ID
		w.editing = false
		w.newButton.Show()
		w.changeButton.Show()
		w.deleteButton.Show()
		w.abortEditButton.Hide()
		w.saveEditButton.Hide()
		w.stack.SetVisibleChildName("display")
		w.secretDetailDisplay.display(next.CurrentSecret)
	}
}
