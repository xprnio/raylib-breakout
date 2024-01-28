package server

import (
	"fmt"

	"github.com/xprnio/raygo/internal/net/utils"
)

type Player struct {
  Name string

  server *Server
  conn *utils.Connection
}

func NewPlayer(name string, server *Server, conn *utils.Connection) *Player {
  return &Player{
    Name: name,
    server: server,
    conn: conn,
  }
}

func (p *Player) Handle() error {
  c := p.conn
  for {
    command, err := c.ReadString()
    if err != nil {
      return err
    }

    if len(command) == 0 {
      continue
    }

    switch command {
      case "?info":
        if err := p.sendServerName(); err != nil {
          return err
        }
        if err := p.sendPlayerList(); err != nil {
          return err
        }
        break
      case "?accept-challenge":
        name, err := c.ReadString()
        if err != nil {
          return err
        }

        target := p.server.findPlayer(name)
        if target == nil {
          c.WriteString("$accept-challenge\nerror: player not found")
          break
        }

        challenge, exists := p.server.Challenges[target]
        if !exists {
          c.WriteString("$accept-challenge\nerror: challenge not found")
          break
        }

        if challenge.Name != p.Name {
          c.WriteString("$accept-challenge\nerror: invalid challenge")
          break
        }

        target.conn.WriteString(fmt.Sprintf("!challenge-accepted\n%s", p.Name))
        c.WriteString("$accept-challenge\nsuccess")
        // TODO: Start game
        break
      case "?decline-challenge":
        name, err := c.ReadString()
        if err != nil {
          return err
        }

        target := p.server.findPlayer(name)
        if target == nil {
          c.WriteString("$accept-challenge\nerror: player not found")
          break
        }

        challenge, exists := p.server.Challenges[target]
        if !exists {
          c.WriteString("$accept-challenge\nerror: challenge not found")
          break
        }

        if challenge.Name != p.Name {
          c.WriteString("$accept-challenge\nerror: invalid challenge")
          break
        }

        target.conn.WriteString(fmt.Sprintf("!challenge-declined\n%s", p.Name))

        delete(p.server.Challenges, target)
        c.WriteString("$decline-challenge\nsuccess")
        break
      case "?cancel-challenge":
        name, err := c.ReadString()
        if err != nil {
          return err
        }

        challenge, exists := p.server.Challenges[p]
        if !exists {
          c.WriteString("$cancel-challenge\nerror: invalid challenge")
          break
        }

        if name != challenge.Name {
          c.WriteString("$cancel-challenge\nerror: invalid challenge")
          break
        }

        delete(p.server.Challenges, p)
        break
      case "?challenge":
        name, err := c.ReadString()
        if err != nil {
          return err
        }

        if name == p.Name {
          c.WriteString("$challenge\nerror: can't challenge yourself")
          break
        }

        player := p.server.findPlayer(name)
        if player == nil {
          c.WriteString("$challenge\nerror: invalid player")
          break
        }

        challenge := fmt.Sprintf("!challenge\n%s", p.Name)
        player.conn.WriteString(challenge)
        c.WriteString("$challenge\nsuccess")

        p.server.Challenges[p] = player
        break
      default:
        return fmt.Errorf("invalid command: %s", command)
    }
  }
}

func (p *Player) sendServerName() error {
  response := fmt.Sprintf("!server-name\n%s", p.server.Name)
  return p.conn.WriteString(response)
}

func (p *Player) sendPlayerList() error {
  players := "!players"
  for _, p := range p.server.Players {
    players = fmt.Sprintf("%s\n%s", players, p.Name)
  }

  err := p.conn.WriteString(fmt.Sprintf("%s\n", players))
  if err != nil {
    return err
  }

  return nil
}
