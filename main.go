package main

import (
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var deathRes, surviveRes fyne.Resource
var runCount *widget.Label
var gameStartBtn *widget.Button
var gameChan chan struct{}
var isRun bool
var bottom *fyne.Container

func main() {
	gameChan = make(chan struct{}, 1)
	defer close(gameChan)
	g := newGame(16, 16, 300*time.Millisecond)
	g.loadRes()
	go g.start()
	app := app.New()
	w := app.NewWindow("game of life")
	obj := g.getGameBody()
	runCount = widget.NewLabel("0")
	gameStartBtn = gameStartButton()
	top := fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		widget.NewLabel("run times: "),
		runCount,
		gameStartBtn,
		restartBtn(g))
	bottom = fyne.NewContainerWithLayout(layout.NewGridLayout(g.width), obj...)
	w.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), top, bottom))
	w.ShowAndRun()
}
