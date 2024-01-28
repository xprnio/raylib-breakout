package player_list

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/xprnio/raygo/internal/game/events"
	"github.com/xprnio/raygo/internal/game/state"
)

type Player struct {
  Name string
  Color rl.Color
  Bounds rl.Rectangle
  events *events.EventManager
}

func NewPlayer(name string, color rl.Color) *Player {
  return &Player{
    Name: name,
    Color: color,
  }
}

func (p *Player) Init(em *events.EventManager) {
  p.events = em
}

func (p Player) Update(d float32) {
  if !p.isHovering() {
    return
  }

  if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
    p.events.Emit(state.NewChallenge(p.Name))
  }
}

func (p Player) isHovering() bool {
  m := rl.GetMousePosition()
  return rl.CheckCollisionPointRec(m, p.Bounds)
}

func (p Player) Position() rl.Vector2 {
  return rl.NewVector2(p.Bounds.X, p.Bounds.Y)
}

func (p Player) Size() rl.Vector2 {
  return rl.NewVector2(p.Bounds.Width, p.Bounds.Height)
}

func (p Player) TextSize(fontSize, fontSpacing float32) rl.Vector2 {
  font := rl.GetFontDefault()
  return rl.MeasureTextEx(font, p.Name, fontSize, fontSpacing)
}

func (p Player) Draw() {
  if p.isHovering() {
    rl.DrawRectangleRec(p.Bounds, rl.ColorAlpha(rl.White, 0.1))
  }

  font := rl.GetFontDefault()
  fontSize := 24
  fontSpacing := 2
  rl.DrawTextEx(font, p.Name, p.Position(), float32(fontSize), float32(fontSpacing), p.Color)
}
