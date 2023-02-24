package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PowerUp struct {
	duration     float32
	sprite       *rl.Image
	rect         rl.Rectangle
	OnActivate   func(*Game)
	OnDeactivate func(*Game)

	active bool
	timer  float32
}

func (p *PowerUp) String() string {
	return strconv.Itoa(int(p.rect.X)) + " " + strconv.Itoa(int(p.rect.Y))
}

func (p *PowerUp) Draw() {
	//rl.DrawTexture(rl.LoadTextureFromImage(p.sprite), p.rect.ToInt32().X, p.rect.ToInt32().Y, rl.White)
	if p.active {
		return
	}
	rl.DrawRectangleRec(p.rect, rl.White)
}

func (p *PowerUp) Colliding(g *Game, b Ball) {
	if rl.CheckCollisionRecs(b.rect, p.rect) {
		if !p.active {
			print("collided")
			p.OnActivate(g)
			p.timer = p.duration
			p.active = true
		}
	}
}

func (p *PowerUp) Update(g *Game, dt float32) {
	if p.active {
		print(p.timer)
		if p.timer < 0 {
			p.OnDeactivate(g)
			g.powerups = []PowerUp{}
		} else {
			p.timer -= dt
		}
	}
}

func NewPowerUp(text string, duration float32, on func(*Game), off func(*Game)) PowerUp {
	//sprite := rl.ImageText("T", 20, rl.White)
	return PowerUp{
		//sprite:       sprite,
		rect:         rl.NewRectangle(-100, -100, float32(60), float32(60)),
		duration:     duration,
		OnActivate:   on,
		OnDeactivate: off,
	}
}

// POWERUP TYPES
var poverupTypes []PowerUp = []PowerUp{
	NewPowerUp("B", 5,
		func(g *Game) { g.p1.paddleRect.Height += 50; g.p2.paddleRect.Height += 50 },
		func(g *Game) { g.p1.paddleRect.Height -= 50; g.p2.paddleRect.Height -= 50 },
	),
}
