package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_game_getStatus(t *testing.T) {
	type args struct {
		surviveCount int
		selfStatus   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"活跃--周围=0", args{0, isSurvive}, isDeath},
		{"活跃--周围=1", args{1, isSurvive}, isDeath},
		{"活跃--周围=2", args{2, isSurvive}, isKeep},
		{"活跃--周围=3", args{3, isSurvive}, isKeep},
		{"活跃--周围=4", args{4, isSurvive}, isDeath},
		{"死亡--周围=2", args{2, isDeath}, isKeep},
		{"死亡--周围=3", args{3, isDeath}, isSurvive},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &game{}
			if got := g.getStatus(tt.args.surviveCount, tt.args.selfStatus); got != tt.want {
				t.Errorf("game.getStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getList(x, y int) int {
	mockCoo := make(map[string]int, 16)
	mockCoo["0x0"] = isDeath
	mockCoo["0x1"] = isSurvive
	mockCoo["0x2"] = isDeath
	mockCoo["0x3"] = isDeath
	mockCoo["1x0"] = isSurvive //
	mockCoo["1x1"] = isDeath
	mockCoo["1x2"] = isDeath
	mockCoo["1x3"] = isSurvive
	mockCoo["2x0"] = isSurvive //
	mockCoo["2x1"] = isSurvive
	mockCoo["2x2"] = isSurvive
	mockCoo["2x3"] = isDeath
	mockCoo["3x0"] = isDeath //
	mockCoo["3x1"] = isSurvive
	mockCoo["3x2"] = isDeath
	mockCoo["3x3"] = isSurvive
	if res, ok := mockCoo[fmt.Sprintf("%dx%d", x, y)]; ok {
		return res
	}
	return -1
}

func Test_game_checkIsDeath(t *testing.T) {
	g := newGame(4, 4, 1*time.Second)
	/*
		0 1 0 0
		1 0 0 1
		1 1 1 0
		0 1 0 1
	*/
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			gb := &gameButton{x: i, y: j, life: getList(i, j)}
			coo := fmt.Sprintf("%dx%d", i, j)
			g.all = append(g.all, coo)
			g.list[coo] = gb
		}
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"坐标0,0", args{0, 0}, 2},
		{"坐标0,1", args{0, 1}, 1},
		{"坐标0,2", args{0, 2}, 2},
		{"坐标0,3", args{0, 3}, 1},
		{"坐标1,0", args{1, 0}, 3},
		{"坐标1,1", args{1, 1}, 5},
		{"坐标1,2", args{1, 2}, 4},
		{"坐标3,0", args{3, 0}, 3},
		{"坐标3,3", args{3, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.checkIsDeath(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("game.checkIsDeath() = %v, want %v", got, tt.want)
			}
		})
	}
}
