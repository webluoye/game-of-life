package main

import (
	"fyne.io/fyne/widget"
)

func restartBtn(g *game) *widget.Button {
	return widget.NewButton("Restart", func() {
		if isRun {
			gameChan <- struct{}{}
		}
		bottom.Objects = g.getGameBody()
		bottom.Refresh()
	})
}

func gameStartButton() *widget.Button {
	return widget.NewButton("Start", func() {
		gameChan <- struct{}{}
	})
}
