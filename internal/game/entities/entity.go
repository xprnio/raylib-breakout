package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity interface {
  Init();
  Update(delta float32);
  Draw();

  GetBounds() rl.Rectangle
  SetBounds(rect rl.Rectangle)
}
