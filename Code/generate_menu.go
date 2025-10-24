package main

import (
	"log"
	"fyne.io/fyne/v2"
)

/*
Generates the main menu

This exists for keeping the code for generating the main menu out of main
*/
func GenerateMenu() *fyne.MainMenu {
	// Generate file menu
	// Create menu Items
	openFileMenuItem := NewMenuItemHandler("Open File", func() {
		log.Println("TODO: open file here");
	})
	
	openFolderMenuItem := NewMenuItemHandler("Open Folder", func() {
		log.Println("TODO: open folder here");
	})

	// Merge into menu
	fileMenuHdr := NewMenuHandler("File", 2)
	fileMenuHdr.AddMenuItem(openFileMenuItem.GenerateMenuItem());
	fileMenuHdr.AddMenuItem(openFolderMenuItem.GenerateMenuItem());
	
	// Convert to correct format
	fileMenu, err := fileMenuHdr.GenerateMenu();
	if DEBUG {
		if err != nil {
			log.Println("Too many menu items in file menu!");
		}
	}

	// Generate edit menu
	// Create menu items
	undoItem := NewMenuItemHandler("Undo", func() {
		log.Println("TODO: undo code here");
	})
	
	redoItem := NewMenuItemHandler("Redo", func() {
		log.Println("TODO: redo code here");
	})
	
	cutItem := NewMenuItemHandler("Cut", func() {
		log.Println("TODO: cut code here");
	})
	
	copyItem := NewMenuItemHandler("Copy", func() {
		log.Println("TODO: copy code here");
	})
	
	pasteItem := NewMenuItemHandler("Paste", func() {
		log.Println("TODO: paste code here");
	})

	// Merge into menu
	editMenuHdr := NewMenuHandler("Edit", 5);
	editMenuHdr.AddMenuItem(undoItem.GenerateMenuItem());
	editMenuHdr.AddMenuItem(redoItem.GenerateMenuItem());
	editMenuHdr.AddMenuItem(cutItem.GenerateMenuItem());
	editMenuHdr.AddMenuItem(copyItem.GenerateMenuItem());
	editMenuHdr.AddMenuItem(pasteItem.GenerateMenuItem());
	
	// Convert to correct format
	editMenu, err := editMenuHdr.GenerateMenu();
	if DEBUG {
		if err != nil {
			log.Println("Too many menu items in edit menu!");
		}
	}


	return fyne.NewMainMenu(fileMenu, editMenu);
}
