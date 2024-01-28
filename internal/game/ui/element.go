package ui

type Element interface {
  Update(d float32)
  Draw()
}
