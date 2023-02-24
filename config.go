package main

type Config struct {
	WindowWidth  int32
	WindowHeight int32
	WindowTitle  string

	TargetFPS int32
}

func CreateConfig() Config {
	return Config{
		WindowWidth:  1000,
		WindowHeight: 500,
		WindowTitle:  "Pong",

		TargetFPS: 60,
	}
}
