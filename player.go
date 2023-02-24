package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	points int32

	paddleRect rl.Rectangle

	paddleSpeed float32
}

func (p *Player) Draw() {
	rl.DrawRectangleRec(
		p.paddleRect,
		rl.White,
	)
}
