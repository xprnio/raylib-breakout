package server_lobby

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/xprnio/raygo/internal/game/events"
	"github.com/xprnio/raygo/internal/game/screens/server_lobby/player_list"
	"github.com/xprnio/raygo/internal/game/state"
	"github.com/xprnio/raygo/internal/game/ui"
)

type Elements struct {
  Width, Height int32
  ServerName ui.Label
  PlayerList *player_list.PlayerList
}

func NewElements(width, height int32) *Elements {
  e := &Elements{
    Width: width,
    Height: height,
  }

  e.PlayerList = e.initPlayerList(width, height)
  e.ServerName = e.initServerName(width, height)

  return e
}

func (e *Elements) UpdateState(s *state.ServerState) {
  e.PlayerList.UpdatePlayers(s.Players)

  e.ServerName.Text = s.ServerName
  width := float32(e.Width) - e.PlayerList.Size.X
  size := e.ServerName.Size()
  e.ServerName.Position = rl.NewVector2((width - size.X) / 2, 16)
}

func (e *Elements) Init(em *events.EventManager) {
  e.PlayerList.Init(em)
}

func (e *Elements) Update(d float32) {
  e.ServerName.Update(d)
  e.PlayerList.Update(d)
}

func (e *Elements) Draw() {
  e.ServerName.Draw()
  e.PlayerList.Draw()
}

func (e *Elements) initServerName(width, height int32) ui.Label {
  return ui.NewLabel(
    "",
    func(l *ui.Label) {
      width := float32(width) - e.PlayerList.Size.X
      size := l.Size()
      l.Position = rl.NewVector2((width - size.X) / 2, 16)
    },
  )
}

func (e *Elements) initPlayerList(width, height int32) *player_list.PlayerList {
  return player_list.NewPlayerList(func(pl *player_list.PlayerList) {
    pl.Style = player_list.PlayerListStyle{
      Background: rl.ColorAlpha(rl.White, 0.25),
      HeadForeground: rl.White,
      ItemForeground: rl.White,
      ItemLocalForeground: rl.Green,
      Padding: 20,
      HeadFontSize: 32,
      ItemFontSize: 24,
      ItemSpacing: 16,
    }
    pl.Size = rl.NewVector2(float32(width) / 4, float32(height))
    pl.Position = rl.NewVector2(float32(width) - pl.Size.X, 0)
  })
}
