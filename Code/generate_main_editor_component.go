package main

import (
	"fyne.io/fyne/v2/container"
)

/*
Generates the DocTabs used in the main editor window

This exists for keeping the code for generating it out of main
*/
func GenerateMainEditorComponent() *container.DocTabs {
	untitledItem := NewDocumentTabItem("(untitled)")
	documentTabItem := untitledItem.GenerateDocumentTabItem()

	return container.NewDocTabs(documentTabItem)
}
