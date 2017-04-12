package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/pkg/errors"
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
	store               *Store
	secretID            string
}

func newSecretDetail(store *Store, logger logging.Logger) (*secretDetail, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create box")
	}
	stack, err := gtk.StackNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create stack")
	}
	newButton, err := gtk.MenuButtonNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create new button")
	}
	changeButton, err := gtk.ButtonNewFromIconName("document-open-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create changeButton")
	}
	deleteButton, err := gtk.MenuButtonNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create deleteButton")
	}
	abortEditButton, err := gtk.ButtonNewFromIconName("window-close-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create abourtEditButton")
	}
	saveEditButton, err := gtk.ButtonNewFromIconName("document-save-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create saveEditButton")
	}

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

	buttonBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create button box")
	}
	w.Add(buttonBox)

	newImage, err := gtk.ImageNewFromIconName("list-add-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create newImage")
	}

	newMenu, err := w.newSecretMenu()
	if err != nil {
		return nil, err
	}

	w.newButton.SetImage(newImage)
	w.newButton.SetHAlign(gtk.ALIGN_START)
	w.newButton.SetMarginTop(2)
	w.newButton.SetMarginStart(2)
	w.newButton.SetMarginEnd(2)
	w.newButton.SetMarginBottom(2)
	w.newButton.SetPopup(newMenu)
	w.newButton.SetDirection(gtk.ARROW_UP)
	buttonBox.Add(w.newButton)

	w.changeButton.SetLabel("Change")
	w.changeButton.SetAlwaysShowImage(true)
	w.changeButton.SetHAlign(gtk.ALIGN_START)
	w.changeButton.SetMarginTop(2)
	w.changeButton.SetMarginStart(2)
	w.changeButton.SetMarginEnd(2)
	w.changeButton.SetMarginBottom(2)
	w.changeButton.SetNoShowAll(true)
	w.changeButton.Connect("clicked", w.store.actionEditCurrent)
	buttonBox.Add(w.changeButton)

	w.abortEditButton.SetLabel("Abort")
	w.abortEditButton.SetAlwaysShowImage(true)
	w.abortEditButton.SetHAlign(gtk.ALIGN_START)
	w.abortEditButton.SetMarginTop(2)
	w.abortEditButton.SetMarginStart(2)
	w.abortEditButton.SetMarginEnd(2)
	w.abortEditButton.SetMarginBottom(2)
	w.abortEditButton.SetNoShowAll(true)
	w.abortEditButton.Connect("clicked", w.store.actionEditAbort)
	buttonBox.Add(w.abortEditButton)

	deleteConfirm, err := gtk.ButtonNewWithLabel("Confirm")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create confirm button")
	}
	deleteConfirm.Connect("clicked", w.onDelete)
	deleteConfirm.Show()
	confirmPopover, err := gtk.PopoverNew(w.deleteButton)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create confirm popover")
	}
	confirmPopover.Add(deleteConfirm)
	confirmPopover.SetBorderWidth(5)

	deleteImage, err := gtk.ImageNewFromIconName("edit-delete-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create deleteImage")
	}
	w.deleteButton.SetImage(deleteImage)
	w.deleteButton.SetLabel("Delete")
	w.deleteButton.SetAlwaysShowImage(true)
	w.deleteButton.SetHAlign(gtk.ALIGN_END)
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
	w.saveEditButton.SetHAlign(gtk.ALIGN_END)
	w.saveEditButton.SetHExpand(true)
	w.saveEditButton.SetMarginTop(2)
	w.saveEditButton.SetMarginStart(2)
	w.saveEditButton.SetMarginEnd(2)
	w.saveEditButton.SetMarginBottom(2)
	w.saveEditButton.SetNoShowAll(true)
	w.saveEditButton.Connect("clicked", w.onEditSave)
	buttonBox.Add(w.saveEditButton)

	placeholder, err := newSecretDetailPlaceholder()
	if err != nil {
		return nil, err
	}
	w.stack.AddNamed(placeholder, "placeholder")

	w.secretDetailDisplay, err = newSecretDetailDisplay(logger)
	if err != nil {
		return nil, err
	}
	w.stack.AddNamed(w.secretDetailDisplay, "display")

	w.secretDetailEdit, err = newSecretDetailEdit(logger)
	if err != nil {
		return nil, err
	}
	w.stack.AddNamed(w.secretDetailEdit, "edit")

	w.store.addListener(w.onStateChanged)

	return w, nil
}

func (w *secretDetail) newSecretMenu() (*gtk.Menu, error) {
	menu, err := gtk.MenuNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create menu")
	}

	for _, secretTypeDefinition := range api.SecretTypes {
		item, err := gtk.MenuItemNewWithLabel(secretTypeDefinition.Display)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to create menu item")
		}
		item.Show()
		menu.Append(item)
	}

	return menu, nil
}

func (w *secretDetail) onDelete() {
	if w.secretID != "" {
		w.store.actionMarkDeleted(w.secretID)
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
	w.store.actionEditStore(w.secretID, *version)
}

func (w *secretDetail) onStateChanged(prev, next *State) {
	if next.currentSecret == nil {
		w.secretID = ""
		w.newButton.Show()
		w.changeButton.Hide()
		w.deleteButton.Hide()
		w.abortEditButton.Hide()
		w.saveEditButton.Hide()
		w.stack.SetVisibleChildName("placeholder")
		return
	} else if next.currentEdit {
		w.secretID = next.currentSecret.ID
		w.newButton.Hide()
		w.changeButton.Hide()
		w.deleteButton.Hide()
		w.abortEditButton.Show()
		w.saveEditButton.Show()
		w.stack.SetVisibleChildName("edit")
		w.secretDetailEdit.setEdit(&next.currentSecret.SecretCurrent)
		return
	}
	w.secretID = next.currentSecret.ID
	w.newButton.Show()
	w.changeButton.Show()
	w.deleteButton.Show()
	w.abortEditButton.Hide()
	w.saveEditButton.Hide()
	w.stack.SetVisibleChildName("display")
	w.secretDetailDisplay.display(next.currentSecret)
}
