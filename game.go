package main

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

const (
	isDeath   = 0
	isSurvive = 1
	isKeep    = 2
)

type game struct {
	list          map[string]*gameButton
	all           []string // 坐标
	width, height int
	speed         time.Duration
}

type gameButton struct {
	btn  *widget.Button
	x    int
	y    int
	life int // 0:死亡 1:存活
}

//newGame 棋盘大小
func newGame(x, y int, d time.Duration) *game {
	g := &game{width: x, height: y, speed: d}
	g.list = make(map[string]*gameButton)
	return g
}

func (g *game) getGameBody() []fyne.CanvasObject {
	obj := []fyne.CanvasObject{}
	for i := 0; i < g.width; i++ {
		for k := 0; k < g.height; k++ {
			gb := newGameButton(i, k)
			g.setBtn(gb, fmt.Sprintf("%dx%d", i, k))
			obj = append(obj, gb.btn)
		}
	}
	return obj
}

func (g *game) loadRes() {
	deathRes, _ = fyne.LoadResourceFromPath("./res/death.png")
	surviveRes, _ = fyne.LoadResourceFromPath("./res/survive.png")
}

func (g *game) start() {
	count := 0
	for {
		select {
		case <-gameChan:
			if !isRun {
				isRun = true
				count = 0
				gameStartBtn.SetText("Pause")
			} else if isRun {
				isRun = false
				gameStartBtn.SetText("Start")
			}
		case <-time.After(g.speed):
			if isRun {
				for coo, tmpStatus := range g.getNextStatus() {
					gb := g.list[coo]
					gb.setLife(tmpStatus)
				}
				count++
				runCount.SetText(fmt.Sprintf("%d", count))
			}
		}
	}
}

func (g *game) getNextStatus() map[string]int {
	nextStatus := make(map[string]int) //新的结果
	for _, coordinate := range g.all {
		if gb, ok := g.list[coordinate]; ok {
			count := g.checkIsDeath(gb.x, gb.y)
			if result := g.getStatus(count, gb.life); result != isKeep {
				nextStatus[coordinate] = result
			}
		}
	}
	return nextStatus
}

func (g *game) setBtn(gb *gameButton, coo string) {
	g.list[coo] = gb
	g.all = append(g.all, coo)
}

//checkIsDeath 循环检测该坐标周围的8个位置
func (g *game) checkIsDeath(x, y int) int {
	surviveCount := 0
	for i := x - 1; i < x+2; i++ {
		if i < 0 || i >= g.width {
			continue
		}
		for j := y - 1; j < y+2; j++ {
			if (j < 0 || j >= g.height) || (i == x && j == y) { //skip slef
				continue
			}
			if gb, ok := g.list[fmt.Sprintf("%dx%d", i, j)]; ok {
				if gb.life == isSurvive {
					surviveCount++
				}
			}
		}
	}
	return surviveCount
}

func (g *game) getStatus(surviveCount int, selfStatus int) (result int) {
	if selfStatus == isSurvive {
		if surviveCount == 2 || surviveCount == 3 {
			result = isKeep
		} else {
			result = isDeath
		}
	} else {
		if surviveCount == 3 {
			result = isSurvive
		} else {
			result = isKeep
		}
	}
	return
}

func newGameButton(x, y int) *gameButton {
	gb := &gameButton{x: x, y: y}
	life := gb.getLife()
	btn := widget.NewButton("", nil)
	if life == 1 {
		btn.SetIcon(surviveRes)
	} else {
		btn.SetIcon(deathRes)
	}
	gb.life = life
	gb.btn = btn
	return gb
}

func (g *gameButton) setLife(status int) {
	switch status {
	case isDeath:
		g.life = status
		g.btn.SetIcon(deathRes)
		g.btn.Refresh()
	case isSurvive:
		g.life = status
		g.btn.SetIcon(surviveRes)
		g.btn.Refresh()
	case isKeep:
		//do nothing
	}
}

func (g *gameButton) getLife() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(2)
}
