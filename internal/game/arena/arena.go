package arena

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WallWidth = 32

type Arena struct {
  WallWidth float32
  Bounds rl.Rectangle
  Walls [4]rl.Rectangle
}

func New(width, height float32) *Arena {
  return &Arena{
    WallWidth: WallWidth,
    Bounds: rl.NewRectangle(0, 0, width, height),
    Walls: [4]rl.Rectangle{
      rl.NewRectangle(0, 0, WallWidth, height),
      rl.NewRectangle(width - WallWidth, 0, WallWidth, height),
    },
  }
}

func (a *Arena) Draw() {
  for _, w := range a.Walls {
    rl.DrawRectangleRec(w, rl.RayWhite)
  }
}

func (a *Arena) KeepInBounds(rect rl.Rectangle) rl.Rectangle {
  for _, w := range a.Walls {
    collision := rl.GetCollisionRec(w, rect)
    if collision.Width == 0 {
      continue
    }

    fmt.Printf("%v\n", collision)

    if collision.X == rect.X {
      rl.DrawRectangleRec(collision, rl.Green)
      rect.X += collision.Width
      return rect
    }
    
    if collision.X + collision.Width == rect.X + rect.Width {
      // Collided on the right
      rl.DrawRectangleRec(collision, rl.Blue)
      rect.X -= collision.Width
      return rect
    }
  }
  
  return rect
}
