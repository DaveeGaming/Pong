package main

import rl "github.com/gen2brain/raylib-go/raylib"

var Pong Game = Game{playing: true}

func main() {
	Pong.Init()

	for !rl.WindowShouldClose() {
		Pong.HandleInput()
		Pong.Update()
		Pong.Draw()
	}

	rl.CloseWindow()
}
