package main_menu

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/xprnio/raygo/internal/game/events"
	"github.com/xprnio/raygo/internal/game/screens"
	"github.com/xprnio/raygo/internal/game/screens/connect_screen"
	"github.com/xprnio/raygo/internal/game/ui"
)

type MainMenuScreen struct {
  Width, Height int32

  elements []ui.Element
  events *events.EventManager
}

func New(width, height int32) *MainMenuScreen {
  return &MainMenuScreen{
    Width: width,
    Height: height,
  }
}

func (s *MainMenuScreen) initButtons() *ui.ButtonList {
  size := rl.NewVector2(320, 48)
  position := rl.NewVector2(
    (float32(s.Width) - size.X) / 2,
    float32(s.Height) / 2,
  )
  buttons := ui.NewButtonList(position, size, 32)
  buttons.Add("Connect", func(){
    cs := connect_screen.New(s.Width, s.Height)
    s.events.Emit(screens.NewScreenEvent(cs))
  })
  buttons.Add("Quit", func(){
    s.events.Emit(events.NewExitEvent())
  })

  return buttons
}

func (s *MainMenuScreen) Init(em *events.EventManager) {
  s.events = em

  s.elements = []ui.Element{
    s.initButtons(),
    ui.NewLabel(
      "turn my lights on",
      func(l *ui.Label) {
        size := l.Size()
        l.Position = rl.NewVector2(
          (float32(s.Width) - size.X) / 2,
          float32(s.Height) / 3,
        ) 
      },
    ),
  }
}

func (m *MainMenuScreen) Update(d float32) {
  for _, el := range m.elements {
    el.Update(d)
  }
}

func (m *MainMenuScreen) Draw() {
  for _, el := range m.elements {
    el.Draw()
  }
}
