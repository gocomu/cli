package main

import (
	"test/cmd"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	gocomu "github.com/gocomu/cli"
)

var playButton *widget.Button

func main() {
	// load project info details from gocomu.yml located at proejct root
	yamlData, _ := gocomu.Yaml()

	a := app.New()

	playButton = widget.NewButton("Play", func() {
		//a.Quit()
		play()
	})

	w := a.NewWindow(yamlData.Name)
	w.SetContent(widget.NewVBox(
		widget.NewLabel(yamlData.Description),
		playButton,
	))

	w.ShowAndRun()
}

func play() {
	playButton.Disable()
	cmd.GOCOMU()
	playButton.Enable()
}
