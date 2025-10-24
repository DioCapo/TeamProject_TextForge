package main

import (
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
)

/*
Represents an open file in a DocumentTab
*/
type DocumentTabItem struct {
	text string
	content *widget.Entry
}

/*
Creates a new DocumentTabItem
*/
func NewDocumentTabItem(text string) *DocumentTabItem {
	entry := widget.NewMultiLineEntry();
	return &DocumentTabItem{
		text: text,
		content: entry,
	}
}

/*
Converts the DocumentTabItem into a *container.TabItem
*/
func (dTItem *DocumentTabItem) GenerateDocumentTabItem() *container.TabItem {
	return container.NewTabItem(dTItem.text, dTItem.content)
}
