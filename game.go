package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	config Config
	p1     Player
	p2     Player
	ball   Ball
}

type Object interface {
	rl.Rectangle
}

func (g *Game) Init() {
	g.config = CreateConfig()
	rl.InitWindow(g.config.WindowWidth, g.config.WindowHeight, g.config.WindowTitle)
	rl.SetTargetFPS(g.config.TargetFPS)

	halfh := g.config.WindowHeight / 2

	var offset float32 = 10
	var paddleWidth float32 = 10
	var paddleHeight float32 = 100
	var paddleSpeed float32 = 200

	g.p1 = Player{
		paddleRect:  rl.NewRectangle(offset, float32(halfh)-paddleHeight/2, paddleWidth, paddleHeight),
		paddleSpeed: paddleSpeed,
		points:      0,
	}

	g.p2 = Player{
		paddleRect:  rl.NewRectangle(float32(g.config.WindowWidth)-offset-paddleWidth, float32(halfh)-paddleHeight/2, paddleWidth, paddleHeight),
		paddleSpeed: paddleSpeed,
		points:      0,
	}

	g.ball = DefaultBall(g, 1)
}

func (g *Game) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	//Draw Players
	g.p1.Draw()
	g.p2.Draw()

	// Draw Ball
	g.ball.Draw()

	// Draw bg lines
	screenMiddle := g.config.WindowWidth / 2
	lineWidth := int32(4)
	for i := 0; i < int(g.config.WindowHeight); i += 20 {
		rl.DrawRectangle(screenMiddle-lineWidth/2, int32(i), lineWidth, 10, rl.White)
	}

	// Draw points

	var offset int32 = 30

	rl.DrawText(strconv.Itoa(int(g.p1.points)), screenMiddle-offset-15, offset, 25, rl.White)
	rl.DrawText(strconv.Itoa(int(g.p2.points)), screenMiddle+offset, offset, 25, rl.White)

	rl.EndDrawing()
}

func (g *Game) Update() {
	dt := rl.GetFrameTime()

	if Colliding(&g.ball, &g.p1) {
		g.ball.dx *= -1
		g.ball.dy += (g.ball.rect.Y + g.ball.rect.Height/2 - g.p1.paddleRect.Y + g.p1.paddleRect.Height/2)
	} else if Colliding(&g.ball, &g.p2) {
		g.ball.dx *= -1
		g.ball.dy += (g.ball.rect.Y + g.ball.rect.Height/2 - g.p2.paddleRect.Y + g.p2.paddleRect.Height/2)
	}

	// Update ball speed
	g.ball.dx *= g.ball.speedupMultiplier
	g.ball.dy *= g.ball.speedupMultiplier

	g.ball.rect.X += g.ball.dx * dt
	g.ball.rect.Y += g.ball.dy * dt

	// Ball collision
	if !InBounds(g, g.ball.rect) {
		g.ball.dy *= -1
	}

	// Handle points
	if g.ball.rect.X < 0 {
		g.p2.points += 1
		g.ball = DefaultBall(g, -1)
	} else if g.ball.rect.X > float32(g.config.WindowWidth) {
		g.p1.points += 1
		g.ball = DefaultBall(g, 1)
	}
}

func (g *Game) HandleInput() {
	dt := rl.GetFrameTime()

	// p1 input
	if rl.IsKeyDown(rl.KeyW) {
		g.p1.paddleRect.Y = Clamp(g.p1.paddleRect.Y-g.p1.paddleSpeed*dt, 0, float32(g.config.WindowHeight)-g.p1.paddleRect.Height)
	} else if rl.IsKeyDown(rl.KeyS) {
		g.p1.paddleRect.Y = Clamp(g.p1.paddleRect.Y+g.p1.paddleSpeed*dt, 0, float32(g.config.WindowHeight)-g.p1.paddleRect.Height)
	}

	// p2 input
	if rl.IsKeyDown(rl.KeyUp) {
		g.p2.paddleRect.Y = Clamp(g.p2.paddleRect.Y-g.p2.paddleSpeed*dt, 0, float32(g.config.WindowHeight)-g.p2.paddleRect.Height)
	} else if rl.IsKeyDown(rl.KeyDown) {
		g.p2.paddleRect.Y = Clamp(g.p2.paddleRect.Y+g.p2.paddleSpeed*dt, 0, float32(g.config.WindowHeight)-g.p2.paddleRect.Height)
	}
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

func InBounds(g *Game, rec rl.Rectangle) bool {
	return rec.Y > 0 && (rec.Y+rec.Height) < float32(g.config.WindowHeight)
}

func Clamp(num, min, max float32) float32 {
	if num < min {
		return min
	} else if num > max {
		return max
	} else {
		return num
	}
}
