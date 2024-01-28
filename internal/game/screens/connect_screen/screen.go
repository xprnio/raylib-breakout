package connect_screen

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/xprnio/raygo/internal/game/events"
	"github.com/xprnio/raygo/internal/game/screens"
	"github.com/xprnio/raygo/internal/game/screens/server_lobby"
	"github.com/xprnio/raygo/internal/game/ui"
	"github.com/xprnio/raygo/internal/net/client"
)

type Screen struct {
  Width, Height int32
  events *events.EventManager

  addressInput *ui.LabelInput
  nameInput *ui.LabelInput
  joinButton ui.Button
}

func New(width, height int32) *Screen {
  return &Screen{
    Width: width,
    Height: height,
  }
}

func (s *Screen) Init(em *events.EventManager) {
  s.events = em

  spacing := float32(24)
  width := float32(420)
  position := rl.NewVector2((float32(s.Width) - width) / 2, 100)
  inputStyle := ui.InputStyle{
    FontSize: 24,
    FontSpacing: 2,
    BorderWidth: 2,
    Cursor: ui.CursorStyle{
      Color: rl.Red,
      ColorBlink: rl.ColorAlpha(rl.Red, 0.2),
    },
    BaseColors: ui.InputColors{
      Background: rl.White,
      BorderColor: rl.Green,
      Foreground: rl.Black,
    },
    HoverColors: ui.InputColors{
      Background: rl.White,
      BorderColor: rl.Green,
      Foreground: rl.Black,
    },
    FocusColors: ui.InputColors{
      Background: rl.White,
      BorderColor: rl.Green,
      Foreground: rl.Black,
    },
  }
  s.addressInput = ui.NewLabelInput(
    "Server address:",
    position, width,
    func(li *ui.LabelInput) {
      li.Input.Value = "127.0.0.1:6969"
      li.Input.Style = inputStyle
    },
  )
  s.nameInput = ui.NewLabelInput(
    "Username:",
    rl.Vector2Add(
      position,
      rl.NewVector2(0, s.addressInput.GetSize().Y + spacing),
    ), width,
    func(li *ui.LabelInput) {
      li.Input.Style = inputStyle
    },
  )
  s.joinButton = ui.NewButton(
    "Connect",
    rl.NewRectangle(
      s.nameInput.GetPosition().X,
      s.nameInput.GetPosition().Y + s.nameInput.GetSize().Y + spacing,
      s.nameInput.GetSize().X,
      s.nameInput.GetSize().Y,
    ),
    func() {
      c := client.New(s.nameInput.Value())
      err := c.Connect(s.addressInput.Value())
      if err != nil {
        em.Emit(ui.NewToastEvent(err.Error()))
        return
      }

      s := server_lobby.New(s.Width, s.Height, c)
      // if err := s.LoadState(); err != nil {
      //   em.Emit(ui.NewToastEvent(err.Error()))
      //   return
      // }

      em.Emit(ui.ClearToasts())
      em.Emit(screens.NewScreenEvent(s))
    },
  )
}

func (s *Screen) Update(d float32) {
  s.addressInput.Update(d)
  s.nameInput.Update(d)
  s.joinButton.Update(d)
}

func (s *Screen) Draw() {
  s.addressInput.Draw()
  s.nameInput.Draw()
  s.joinButton.Draw()
}
