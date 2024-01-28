package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/xprnio/raygo/internal/game/events"
)

type ClearToastsEvent struct{}
func ClearToasts() ClearToastsEvent {
  return ClearToastsEvent{}
}

type Modal interface {
  Element

  Visible() bool
}

type ModalManager struct {
  Width, Height int32
  Size rl.Vector2
  modals []Modal
}

func NewToastManager(width, height int32) *ModalManager {
  return &ModalManager{
    Width: width,
    Height: height,
    Size: rl.NewVector2(float32(width) / 2, 64),
  }
}

func (m *ModalManager) Init(em *events.EventManager) {
  em.AddHandler(func(e events.Event) {
    switch e := e.(type) {
    case ToastEvent:
      toast := NewToast(string(e), rl.Vector2Zero(), m.Size)
      m.modals = append(m.modals, toast)
      m.reflow()
      break
    case ConfirmEvent:
      confirm := NewConfirm(e.Message, func(c *Confirm) {
        c.Size = rl.NewVector2(
          float32(m.Width) / 1.5,
          float32(m.Height) / 1.5,
        )

        center := rl.NewVector2(
          float32(m.Width) / 2,
          float32(m.Height) / 2,
        )
        c.Position = rl.NewVector2(
          center.X - c.Size.X / 2,
          center.Y - c.Size.Y / 2,
        )

        padding := float32(32)
        buttonSize := rl.NewVector2(
          (c.Size.X / 2) - padding * 2,
          48,
        )

        confirmRect := rl.NewRectangle(
          center.X + padding,
          center.Y + padding,
          buttonSize.X, buttonSize.Y,
        )
        c.ConfirmButton = NewButton(e.ConfirmText, confirmRect, e.OnConfirm)

        cancelRect := rl.NewRectangle(
          c.Position.X + padding,
          center.Y + padding,
          buttonSize.X, buttonSize.Y,
        )
        c.CancelButton = NewButton(e.CancelText, cancelRect, e.OnCancel)
      })
      m.modals = append(m.modals, confirm)
      break
    case ClearToastsEvent:
      modals := []Modal{}
      for _, m := range m.modals {
        switch m := m.(type) {
        case *Toast:
          break
        default:
          modals = append(modals, m)
          break
        }
      }
      m.reflow()
      break
    }
  })
}

func (m *ModalManager) Update(d float32) {
  visible := []Modal{}
  for _, m := range m.modals {
    m.Update(d)

    if m.Visible() {
      visible = append(visible, m)
    }
  }

  if len(visible) < len(m.modals) {
    m.modals = visible
    m.reflow()
  }
}

func (m *ModalManager) reflow() {
  position := rl.NewVector2(
    (float32(m.Width) - m.Size.X) / 2,
    float32(m.Height) - m.Size.Y,
  )
  for i, modal := range m.modals {
    switch modal := modal.(type) {
    case *Toast:
      modal.Position = position
      position = rl.Vector2Subtract(
        position,
        rl.NewVector2(0, modal.Size.Y),
      )
      m.modals[i] = modal
      break
    default:
      break
    }
  }
}

func (m *ModalManager) Draw() {
  for _, modal := range m.modals {
    modal.Draw()
  }
}
