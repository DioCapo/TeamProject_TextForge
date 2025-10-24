package main 

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

const DEBUG = true

func main() {
	// This code is just a start and probably will change
	app := app.New()
	window := app.NewWindow("TextForge")
	window.Resize(fyne.NewSize(800, 600))

	window.SetMainMenu(GenerateMenu());

	middle := GenerateMainEditorComponent()
	
	// Use the center layout to center widgets
	mainLayout := container.NewBorder(nil, nil, nil, nil, middle)
	window.SetContent(mainLayout)

	window.ShowAndRun()


	/*
	a := app.New()
	w := a.NewWindow("TextForge")
	w.Resize(fyne.NewSize(800, 600))
	menuItemNew := fyne.NewMenuItem("New", func() {
		log.Println("New menu item selected!")
	})

	menuItemQuit := fyne.NewMenuItemSeparator()
	menuItemQuitAction := fyne.NewMenuItem("Quit", func() {
		a.Quit()
	})

	fileMenu := fyne.NewMenu("File", menuItemNew, menuItemQuit, menuItemQuitAction)

	mainMenu := fyne.NewMainMenu(fileMenu)

	w.SetMainMenu(mainMenu)

	options := []string{"Option 1", "Option 2", "Option 3"}

	mySelect := widget.NewSelect(options, func(value string) {
		log.Println("Selected option is:", value)
	})

	w.SetContent(container.NewVBox(
		widget.NewLabel("Please select an option:"),
		mySelect,
	))

	w.ShowAndRun()
	*/
}
