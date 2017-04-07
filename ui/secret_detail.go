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
	changeButton        *gtk.Button
	deleteButton        *gtk.MenuButton
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
	changeButton, err := gtk.ButtonNewFromIconName("document-open-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create changeButton")
	}
	deleteButton, err := gtk.MenuButtonNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create deleteButton")
	}

	w := &secretDetail{
		Box:          box,
		stack:        stack,
		changeButton: changeButton,
		deleteButton: deleteButton,
		logger:       logger.WithField("package", "ui").WithField("component", "secretDetail"),
		store:        store,
	}

	w.stack.SetHExpand(true)
	w.stack.SetVExpand(true)
	w.Add(w.stack)

	buttonBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create button box")
	}
	w.Add(buttonBox)

	newButton, err := gtk.MenuButtonNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create new button")
	}
	newImage, err := gtk.ImageNewFromIconName("list-add-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create newImage")
	}

	newMenu, err := w.newSecretMenu()
	if err != nil {
		return nil, err
	}

	newButton.SetImage(newImage)
	newButton.SetHAlign(gtk.ALIGN_START)
	newButton.SetMarginTop(2)
	newButton.SetMarginStart(2)
	newButton.SetMarginEnd(2)
	newButton.SetMarginBottom(2)
	newButton.SetPopup(newMenu)
	newButton.SetDirection(gtk.ARROW_UP)
	buttonBox.Add(newButton)

	w.changeButton.SetLabel("Change")
	w.changeButton.SetAlwaysShowImage(true)
	w.changeButton.SetHAlign(gtk.ALIGN_START)
	w.changeButton.SetMarginTop(2)
	w.changeButton.SetMarginStart(2)
	w.changeButton.SetMarginEnd(2)
	w.changeButton.SetMarginBottom(2)
	w.changeButton.SetNoShowAll(true)
	buttonBox.Add(w.changeButton)

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

func (w *secretDetail) onStateChanged(prev, next *State) {
	if next.currentSecret == nil {
		w.secretID = ""
		w.changeButton.Hide()
		w.deleteButton.Hide()
		w.stack.SetVisibleChildName("placeholder")
		return
	}
	w.secretID = next.currentSecret.ID
	w.changeButton.Show()
	w.deleteButton.Show()
	w.stack.SetVisibleChildName("display")
	w.secretDetailDisplay.display(next.currentSecret)
}
