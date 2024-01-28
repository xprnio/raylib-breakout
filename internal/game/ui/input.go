package ui

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type InputColors struct {
  Background rl.Color
  Foreground rl.Color

  BorderColor rl.Color
}

type InputStyle struct {
  FontSize float32
  FontSpacing float32
  BorderWidth float32

  Cursor CursorStyle

  BaseColors InputColors
  HoverColors InputColors
  FocusColors InputColors
}

type CursorStyle struct {
  Color rl.Color
  ColorBlink rl.Color
}

type Input struct {
  Position rl.Vector2
  Width float32

  Style InputStyle

  Value string
  focused bool
  focusedCounter float64
}

type InputOption func(i *Input)

func NewInput(position rl.Vector2, width float32, options ...InputOption) *Input {
  i := &Input{
    Position: position,
    Width: width,
    Value: "",
  }

  for _, option := range options {
    option(i)
  }

  return i
}

func (i *Input) GetTextSize() rl.Vector2 {
  font := rl.GetFontDefault()
  fontSize := i.Style.FontSize
  fontSpacing := i.Style.FontSpacing
  return rl.MeasureTextEx(font, i.Value, fontSize, fontSpacing)
}

func (i *Input) GetSize() rl.Vector2 {
  textSize := i.GetTextSize()
  content := rl.Vector2Add(
    rl.NewVector2(i.Width, textSize.Y),
    i.padding(),
  )
  return rl.Vector2AddValue(content, i.Style.BorderWidth * 2)
}

func (i *Input) GetBounds() rl.Rectangle {
  size := i.GetSize()
  return rl.NewRectangle(
    i.Position.X, i.Position.Y,
    size.X, size.Y,
  )
}

func (i *Input) isHovering() bool {
  m := rl.GetMousePosition()
  return rl.CheckCollisionPointRec(m, i.GetBounds())
}

func (i *Input) updateInput() {
  if !i.focused {
    return
  }

  key := rl.GetCharPressed()
  for key > 0 {
    if key >= 32 && key <= 125 {
      i.Value = fmt.Sprintf("%s%c", i.Value, rune(key))
    }

    key = rl.GetCharPressed()
  }

  if rl.IsKeyPressed(rl.KeyBackspace) {
    length := len(i.Value)
    if length > 0 {
      i.Value = i.Value[0:length - 1]
    }
  }
}

func (i *Input) updateCounter(d float32) {
  if !i.focused {
    i.focusedCounter = 0
    return
  }

  i.focusedCounter += float64(d)
}

func (i *Input) updateFocus() {
  if !rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
    return
  }

  if i.isHovering() {
    i.focused = true
    return
  }

  i.focused = false
}

func (i *Input) Update(d float32) {
  i.updateFocus()
  i.updateInput()
  i.updateCounter(d)
}

func (i *Input) backgroundColor() rl.Color {
  if i.focused {
    return i.Style.FocusColors.Background
  }

  if i.isHovering() {
    return i.Style.HoverColors.Background
  }

  return i.Style.BaseColors.Background
}

func (i *Input) borderColor() rl.Color {
  if i.focused {
    return i.Style.FocusColors.BorderColor
  }

  if i.isHovering() {
    return i.Style.HoverColors.BorderColor
  }

  return i.Style.BaseColors.BorderColor
}

func (i *Input) Draw() {
  bounds := i.GetBounds()
  rl.DrawRectangleRec(bounds, i.backgroundColor())
  rl.DrawRectangleLinesEx(bounds, i.Style.BorderWidth, i.borderColor())

  i.drawValue()
  i.drawCursor()
}

func (i *Input) padding() rl.Vector2 {
  return rl.Vector2Scale(
    rl.NewVector2(i.Style.BorderWidth, i.Style.BorderWidth),
    2,
  )
}

func (i *Input) valueColor() rl.Color {
  if i.focused {
    return i.Style.FocusColors.Foreground
  }

  if i.isHovering() {
    return i.Style.HoverColors.Foreground
  }

  return i.Style.BaseColors.Foreground
}

func (i *Input) drawValue() {
  font := rl.GetFontDefault()
  fontSize := i.Style.FontSize
  fontSpacing := i.Style.FontSpacing

  position := rl.Vector2Add(i.Position, i.padding())
  // position = rl.Vector2AddValue(
  //   position,
  //   i.Style.BorderWidth,
  // )

  rl.DrawTextEx(font, i.Value, position, fontSize, fontSpacing, i.valueColor())
}

func (i *Input) cursorColor() rl.Color {
  if math.Mod(i.focusedCounter, 1) > 0.5 {
    return i.Style.Cursor.ColorBlink
  }

  return i.Style.Cursor.Color
}

func (i *Input) drawCursor() {
  if i.focused != true {
    return
  }

  textSize := i.GetTextSize()

  size := rl.NewVector2(16, i.Style.BorderWidth * 2)
  position := rl.Vector2AddValue(i.Position, i.Style.BorderWidth)
  position = rl.Vector2Add(position, i.padding())
  position = rl.Vector2Add(position, textSize)
  position = rl.Vector2Subtract(position, rl.NewVector2(0, size.Y))

  rl.DrawRectangleV(position, size, i.cursorColor())
}

