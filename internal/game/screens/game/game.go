package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/xprnio/raygo/internal/game/arena"
	"github.com/xprnio/raygo/internal/game/controllers"
	"github.com/xprnio/raygo/internal/game/entities"
	"github.com/xprnio/raygo/internal/game/events"
	"github.com/xprnio/raygo/internal/net/client"
)

type GameScreen struct {
  Width, Height int32

  client *client.Client
  arena *arena.Arena
  entities []entities.Entity
}

func New(width, height int32, client *client.Client) *GameScreen {
  return &GameScreen{
    Width: width,
    Height: height,
    client: client,
  }
}

func (g *GameScreen) createLocalPaddle() *entities.Paddle {
  p := entities.NewLocalPaddle(rl.Red)
  position := p.Controller.GetPosition()

  position.X = (float32(g.Width) - p.Size.X) / 2
  position.Y = (float32(g.Height) - p.Size.Y) - g.arena.WallWidth

  p.Controller.SetPosition(position)

  return p
}

func (g *GameScreen) crteateMirrorPaddle(ctrl controllers.PaddleController) *entities.Paddle {
  p := entities.NewMirrorPaddle(ctrl, rl.Green)
  position := p.Controller.GetPosition()

  position.X = (float32(g.Width) - p.Size.X) / 2
  position.Y = g.arena.WallWidth

  p.Controller.SetPosition(position)
  return p
}

func (g *GameScreen) Init(em *events.EventManager) {
  g.arena = arena.New(float32(g.Width), float32(g.Height))

  local := g.createLocalPaddle()
  mirror := g.crteateMirrorPaddle(local.Controller)

  g.entities = append(g.entities, local)
  g.entities = append(g.entities, mirror)
}

func (g *GameScreen) Update(d float32) {
  for _, e := range g.entities {
    e.Update(d);
    bounds := g.arena.KeepInBounds(e.GetBounds())
    e.SetBounds(bounds)
  }
}

func (g *GameScreen) Draw() {
  g.arena.Draw()
  for _, e := range g.entities {
    e.Draw()
  }
}
