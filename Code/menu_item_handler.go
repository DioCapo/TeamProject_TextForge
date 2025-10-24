package main

import (
	"fyne.io/fyne/v2"
)

/*
Used for making the creation of menu items in the menus in the main menu bar easier
*/
type MenuItemHandler struct {
	label string
	action func()
}


/*
Creates a new MenuItemHandler
*/
func NewMenuItemHandler(label string, action func()) *MenuItemHandler {
	return &MenuItemHandler {
		label: label,
		action: action,
	}
}

/*
Converts the MenuItemHandler into a *fyne.MenuItem
*/
func (mIHandler *MenuItemHandler) GenerateMenuItem() *fyne.MenuItem{
	return fyne.NewMenuItem(mIHandler.label, mIHandler.action)
}
