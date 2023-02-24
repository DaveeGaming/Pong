package main

type Config struct {
	WindowWidth  int32
	WindowHeight int32
	WindowTitle  string

	TargetFPS int32
}

func CreateConfig() Config {
	return Config{
		WindowWidth:  800,
		WindowHeight: 450,
		WindowTitle:  "Pong",

		TargetFPS: 60,
	}
}
