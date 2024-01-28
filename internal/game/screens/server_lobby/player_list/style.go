package player_list

import rl "github.com/gen2brain/raylib-go/raylib"

type PlayerListStyle struct {
  Background rl.Color

  Padding float32

  // List Head
  HeadForeground rl.Color
  HeadFontSize float32

  // List Items
  ItemForeground rl.Color
  ItemLocalForeground rl.Color
  ItemFontSize float32
  ItemSpacing float32
}
