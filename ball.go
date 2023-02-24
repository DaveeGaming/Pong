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
