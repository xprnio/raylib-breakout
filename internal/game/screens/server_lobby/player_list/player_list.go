package player_list

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/xprnio/raygo/internal/game/events"
	"github.com/xprnio/raygo/internal/game/state"
)

type PlayerList struct {
  Position rl.Vector2
  Size rl.Vector2

  Style PlayerListStyle
  Players []*Player

  events *events.EventManager
}

type PlayerListOption func(pl *PlayerList)
func NewPlayerList(options ...PlayerListOption) *PlayerList {
  pl := &PlayerList{}

  for _, option := range options {
    option(pl)
  }

  return pl
}

func (pl *PlayerList) Init(em *events.EventManager) {
  pl.events = em
}

func (pl *PlayerList) UpdatePlayers(state []state.Player) {
  players := []*Player{}
  for _, p := range state {
    color := pl.Style.ItemForeground
    if p.IsLocal {
      color = pl.Style.ItemLocalForeground
    }
    player := NewPlayer(p.Name, color)
    player.Init(pl.events)

    players = append(players, player)
  }

  pl.Players = players
  pl.reflow()
}

func (pl *PlayerList) Update(d float32) {
  for _, p := range pl.Players {
    p.Update(d)
  }
}

func (pl *PlayerList) reflow() {
  hb := pl.headBounds()
  pp := rl.Vector2Add(
    pl.Position,
    rl.NewVector2(0, hb.Y + hb.Height + pl.Style.Padding),
  )

  for i, p := range pl.Players {
    size := p.TextSize(pl.Style.ItemFontSize, 2)
    offset := float32(i) * (size.Y + pl.Style.ItemSpacing)
    position := rl.Vector2Add(
      pp,
      rl.NewVector2(pl.Style.Padding, offset),
    )

    p.Bounds = rl.NewRectangle(
      position.X, position.Y,
      size.X, size.Y,
    )
  }

}

func (pl *PlayerList) Draw() {
  rl.DrawRectangleV(pl.Position, pl.Size, pl.Style.Background)
  pl.DrawHead()

  for _, p := range pl.Players {
    p.Draw()
  }
}

func (pl *PlayerList) headSize() rl.Vector2 {
  font := rl.GetFontDefault()
  fontSize := pl.Style.HeadFontSize
  fontSpacing := float32(2)

  return rl.MeasureTextEx(font, "Players", fontSize, fontSpacing)
}

func (pl *PlayerList) headBounds() rl.Rectangle {
  size := pl.headSize()
  position := rl.Vector2Add(
    pl.Position,
    rl.NewVector2((pl.Size.X - size.X) / 2, pl.Style.Padding),
  )
  return rl.NewRectangle(
    position.X, position.Y,
    pl.Size.X, size.Y,
  )
}

func (pl *PlayerList) DrawHead() {
  font := rl.GetFontDefault()
  fontSize := pl.Style.HeadFontSize
  fontSpacing := float32(2)

  size := pl.headSize()
  position := rl.Vector2Add(
    pl.Position,
    rl.NewVector2((pl.Size.X - size.X) / 2, pl.Style.Padding),
  )
  rl.DrawTextEx(font, "Players", position, fontSize, fontSpacing, pl.Style.HeadForeground)
}

