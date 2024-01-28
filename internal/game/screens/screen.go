package screens

import "github.com/xprnio/raygo/internal/game/events"

type ScreenEvent struct {
  Screen Screen
}

func NewScreenEvent(screen Screen) ScreenEvent {
  return ScreenEvent{ Screen: screen }
}

type Screen interface {
  Init(em *events.EventManager)
  Update(d float32)
  Draw()
}

type ScreenManager struct {
  Width, Height int32

  Events *events.EventManager
  Current Screen
}

func NewManager(width, height int32) *ScreenManager {
  return &ScreenManager{
    Width: width,
    Height: height,
  }
}

func (m *ScreenManager) Init(em *events.EventManager) {
  m.Events = em
}

func (m *ScreenManager) Set(screen Screen) {
  screen.Init(m.Events)
  m.Current = screen
}

func (m *ScreenManager) Update(d float32) {
  if m.Current != nil {
    m.Current.Update(d)
  }
}

func (m *ScreenManager) Draw() {
  if m.Current != nil {
    m.Current.Draw()
  }
}
