package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	points int32

	keyUp   int
	keyDown int

	paddleRect  rl.Rectangle
	paddleSpeed float32
}

func (p *Player) Draw() {
	rl.DrawRectangleRec(
		p.paddleRect,
		rl.White,
	)
}

func (p *Player) HandleInput(g *Game, dt float32) {
	if rl.IsKeyDown(int32(p.keyUp)) {
		p.paddleRect.Y = Clamp(p.paddleRect.Y-p.paddleSpeed*dt, 0, float32(g.config.WindowHeight)-p.paddleRect.Height)
	} else if rl.IsKeyDown(int32(p.keyDown)) {
		p.paddleRect.Y = Clamp(p.paddleRect.Y+p.paddleSpeed*dt, 0, float32(g.config.WindowHeight)-p.paddleRect.Height)
	}
}
