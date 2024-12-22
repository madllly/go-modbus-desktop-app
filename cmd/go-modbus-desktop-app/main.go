package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/madllly/go-modbus-desktop-app/internal/app/datastore"
)

func main() {

	dataStore := datastore.New()

	a := app.New()

	w := a.NewWindow("ModbusReader")
	w.Resize(fyne.NewSize(300, 400))
	// w.SetContent(widget.NewLabel("TODOs will go here"))
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	content := container.NewVBox(input, widget.NewButton("Save", func() {
		log.Println("Content was:", input.Text)
		dataStore.SetText(input.Text)
	}))

	lable := widget.NewLabel("")
	dataStoreGetterContent := container.NewVBox(lable, widget.NewButton("Get From Data Store", func() {
		lable.SetText(dataStore.GetText())
		log.Println("Content was:", dataStore.GetText())
	}))
	w.SetContent(container.NewBorder(
		content,
		nil,
		nil,
		dataStoreGetterContent,
	))
	w.ShowAndRun()
}
