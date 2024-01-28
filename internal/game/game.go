package game

import (
	"github.com/xprnio/raygo/internal/game/events"
	"github.com/xprnio/raygo/internal/game/screens"
	"github.com/xprnio/raygo/internal/game/screens/main_menu"
	"github.com/xprnio/raygo/internal/game/ui"
)

type Game struct {
  Width, Height int32
  Events *events.EventManager
  Screens *screens.ScreenManager
  Toasts *ui.ModalManager
  shouldExit bool
}

func New(width, height int32) *Game {
  return &Game{
    Width: width,
    Height: height,
    Events: events.NewEventManager(),
    Screens: screens.NewManager(width, height),
    Toasts: ui.NewToastManager(width, height),
    shouldExit: false,
  }
}

func (g *Game) Exit() {
  g.shouldExit = true
}

func (g *Game) Init() {
  g.Events.AddHandler(func(e events.Event) {
    switch e := e.(type) {
    case events.ExitEvent:
      g.shouldExit = true
      break
    case screens.ScreenEvent:
      g.Screens.Set(e.Screen)
      break
    }
  })

  g.Screens.Init(g.Events)
  g.Toasts.Init(g.Events)

  main := main_menu.New(g.Width, g.Height)
  g.Screens.Set(main)
}

func (g *Game) Update(d float32) {
  g.Screens.Update(d)
  g.Toasts.Update(d)
}

func (g *Game) Draw() {
  g.Screens.Draw()
  g.Toasts.Draw()
}

func (g *Game) ShouldExit() bool {
  return g.shouldExit
}
