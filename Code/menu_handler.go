package main 

import (
	"errors"
	"fyne.io/fyne/v2"
)

/*
A struct used to make creating a top level menu in fyne easier

It has two components a label and a slice of type fyne.Menuitem containing all
of the desired items in the menu
*/
type MenuHandler struct {
	label string
	menuItems []*fyne.MenuItem
	numMenuItems int
	usedMenuItems int
}

/*
Creates a new MenuHandler with the provided label and number of menuItems
*/
func NewMenuHandler(label string, numMenuItems int) MenuHandler {
	return MenuHandler{
		label: label,
		menuItems: make([]*fyne.MenuItem, numMenuItems),
		numMenuItems: numMenuItems - 1,
		usedMenuItems: 0,
	}
}

/*
Adds a menu item to the MenuHandler, throwing an error if trying to insert
a menu item even if the container is full
*/
func (mHandler *MenuHandler) AddMenuItem(mItem *fyne.MenuItem) error {
	if mHandler.usedMenuItems <= mHandler.numMenuItems {
		mHandler.menuItems[mHandler.usedMenuItems] = mItem;
		mHandler.usedMenuItems += 1;
		return nil;
	} else {
		return errors.New("Trying to insert too many menu items!");
	}

}

func (mHandler *MenuHandler) GenerateMenu() (*fyne.Menu, error) {
	if (mHandler.usedMenuItems - 1) == mHandler.numMenuItems {
		return fyne.NewMenu(mHandler.label, mHandler.menuItems...), nil
	} else {
		return nil, errors.New("Not enough menu items!");
	}
}
