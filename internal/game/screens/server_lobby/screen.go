package server_lobby

import (
	"github.com/xprnio/raygo/internal/game/events"
	"github.com/xprnio/raygo/internal/game/state"
	"github.com/xprnio/raygo/internal/net/client"
)

type Screen struct {
  Width, Height int32
  State *state.ServerState

  client *client.Client
  elements *Elements
}

func New(width, height int32, client *client.Client) *Screen {
  return &Screen{
    Width: width,
    Height: height,
    State: state.NewServerState(client),

    client: client,
    elements: NewElements(width, height),
  }
}

func (s *Screen) Init(em *events.EventManager) {
  s.elements.Init(em)

  go s.listenUpdates()
  go s.State.Init(em)
}

func (s Screen) listenUpdates() {
  for <- s.State.Updates {
    s.elements.UpdateState(s.State)
  }
}

func (s *Screen) Update(d float32) {
  s.elements.Update(d)
}

func (s *Screen) Draw() {
  s.elements.Draw()
}
