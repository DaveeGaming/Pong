package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Ball struct {
	size int32

	rect rl.Rectangle
	dx   float32
	dy   float32

	speedupMultiplier float32
}

func (b *Ball) Draw() {
	rl.DrawRectangleRec(b.rect, rl.White)
}

func Colliding(b *Ball, p *Player) bool {
	return rl.CheckCollisionRecs(b.rect, p.paddleRect)
}

func DefaultBall(g *Game, side float32) Ball {
	var ballSize float32 = 15

	return Ball{
		rect: rl.NewRectangle(float32(g.config.WindowWidth)/2-ballSize/2, float32(g.config.WindowHeight/2)-ballSize/2, ballSize, ballSize),

		dx:                150 * side,
		dy:                0,
		speedupMultiplier: 1.001,
	}
}
