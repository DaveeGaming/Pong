package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	config   Config
	p1       Player
	p2       Player
	ball     Ball
	playing  bool
	powerups []PowerUp

	powerupCooldown     float32
	powerupBaseCooldown float32
}

func (g *Game) Init() {
	g.config = CreateConfig()
	rl.InitWindow(g.config.WindowWidth, g.config.WindowHeight, g.config.WindowTitle)
	rl.SetTargetFPS(g.config.TargetFPS)

	halfh := g.config.WindowHeight / 2

	var offset float32 = 10
	var paddleWidth float32 = 10
	var paddleHeight float32 = 130
	var paddleSpeed float32 = 250

	g.p1 = Player{
		paddleRect:  rl.NewRectangle(offset, float32(halfh)-paddleHeight/2, paddleWidth, paddleHeight),
		paddleSpeed: paddleSpeed,
		points:      0,
		keyUp:       rl.KeyW,
		keyDown:     rl.KeyS,
	}

	g.p2 = Player{
		paddleRect:  rl.NewRectangle(float32(g.config.WindowWidth)-offset-paddleWidth, float32(halfh)-paddleHeight/2, paddleWidth, paddleHeight),
		paddleSpeed: paddleSpeed,
		points:      0,
		keyUp:       rl.KeyUp,
		keyDown:     rl.KeyDown,
	}

	g.powerupBaseCooldown = 10
	g.powerupCooldown = g.powerupBaseCooldown
	g.ball = DefaultBall(g, 1)
	g.playing = true
}

func (g *Game) Draw() {
	rl.BeginDrawing()

	if !g.playing {
		g.GameOver()
		rl.EndDrawing()
		return
	}

	rl.ClearBackground(rl.Black)

	for i := 0; i < len(g.powerups); i++ {
		g.powerups[i].Draw()
	}

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

	rl.DrawText(strconv.Itoa(int(g.p1.points)), screenMiddle-offset-13, offset, 25, rl.White)
	rl.DrawText(strconv.Itoa(int(g.p2.points)), screenMiddle+offset, offset, 25, rl.White)

	rl.EndDrawing()
}

func (g *Game) Update() {
	if !g.playing {
		return
	}
	dt := rl.GetFrameTime()

	if g.powerupCooldown < 0 && len(g.powerups) == 0 {
		RandomPowerUp(g, poverupTypes)
		g.powerupCooldown = g.powerupBaseCooldown
	} else {
		g.powerupCooldown -= dt
	}

	for i := 0; i < len(g.powerups); i++ {
		curr := &g.powerups[i]
		curr.Colliding(g, g.ball)
		curr.Update(g, dt)
	}

	if Colliding(&g.ball, &g.p1) {
		g.ball.dx *= -1
		g.ball.dy += (g.ball.rect.Y + g.ball.rect.Height/2 - g.p1.paddleRect.Y + g.p1.paddleRect.Height/2)
		g.ball.dx += g.ball.speedupAmount
	} else if Colliding(&g.ball, &g.p2) {
		g.ball.dx *= -1
		g.ball.dy += (g.ball.rect.Y + g.ball.rect.Height/2 - g.p2.paddleRect.Y + g.p2.paddleRect.Height/2)
		g.ball.dx -= g.ball.speedupAmount
	}

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

	if g.p1.points == 5 || g.p2.points == 5 {
		g.playing = false
	}

	// Update ball speed
	g.ball.rect.X += g.ball.dx * dt
	g.ball.rect.Y += g.ball.dy * dt
}

func (g *Game) GameOver() {
	var winner string
	if g.p1.points > g.p2.points {
		winner = "Player 1 won the game"
	} else {
		winner = "Player 2 won the game"
	}
	var textSize rl.Vector2 = rl.MeasureTextEx(rl.GetFontDefault(), winner, 50, 5)

	rl.DrawText(winner, g.config.WindowWidth/2-int32(textSize.X)/2, g.config.WindowHeight/2-int32(textSize.Y), 50, rl.White)
}

func (g *Game) HandleInput() {
	if !g.playing {
		return
	}

	dt := rl.GetFrameTime()
	g.p1.HandleInput(g, dt)
	g.p2.HandleInput(g, dt)
}

func InBounds(g *Game, rec rl.Rectangle) bool {
	return rec.Y > 0 && (rec.Y+rec.Height) < float32(g.config.WindowHeight)
}

func RandomPowerUp(g *Game, p []PowerUp) {
	powerup := p[rl.GetRandomValue(0, int32(len(p))-1)]
	powerup.rect.X = float32(rl.GetRandomValue(50, g.config.WindowWidth-50))
	powerup.rect.Y = float32(rl.GetRandomValue(50, g.config.WindowHeight-50))
	g.powerups = append(g.powerups, powerup)
}
