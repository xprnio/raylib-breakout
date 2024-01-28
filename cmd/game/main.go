package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/xprnio/raygo/internal/game"
)

const (
  Width = 800
  Height = 400
)

func main() {
  startGame()
}

func startGame() {
  g := game.New(Width, Height)
	rl.InitWindow(g.Width, g.Height, "malnourished")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
  g.Init()

	for {
    if rl.WindowShouldClose() {
      break
    }

    if g.ShouldExit() {
      break
    }

    delta := rl.GetFrameTime()
    g.Update(delta)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
    g.Draw()
		rl.EndDrawing()
	}
}
