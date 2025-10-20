package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
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
}
