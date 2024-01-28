package state

import (
	"fmt"

	"github.com/xprnio/raygo/internal/game/events"
	"github.com/xprnio/raygo/internal/game/ui"
	"github.com/xprnio/raygo/internal/net/client"
)

type ServerState struct {
  Updates chan bool

  ServerName string
  Players []Player

  client *client.Client
}

func NewServerState(client *client.Client) *ServerState {
  return &ServerState{
    Updates: make(chan bool),
    client: client,
  }
}

func (s *ServerState) Init(em *events.EventManager) {
  c := s.client.Connection
  c.WriteString("?info")
  em.AddHandler(func(e events.Event) {
    switch e := e.(type) {
    case ChallengeEvent:
      challenge := fmt.Sprintf("?challenge\n%s", e.Target)
      err := c.WriteString(challenge)
      if err != nil {
        em.Emit(ui.NewToastEvent(err.Error()))
        return
      }
      break
    }
  })

  for {
    res, err := c.ReadString()
    if err != nil {
      em.Emit(ui.NewToastEvent(err.Error()))
      return
    }

    if len(res) == 0 {
      continue
    }

    switch res {
    case "!server-name":
      if s.ServerName, err = c.ReadString(); err != nil {
        em.Emit(ui.NewToastEvent(err.Error()))
        return
      }

      s.Updates <- true
      break
    case "!players":
      if s.Players, err = s.readPlayers(); err != nil {
        em.Emit(ui.NewToastEvent(err.Error()))
        return
      }

      s.Updates <- true
      break
    case "$challenge":
      res, err := c.ReadString()
      if err != nil {
        em.Emit(ui.NewToastEvent(err.Error()))
        break
      }

      if res == "success" {
        em.Emit(ui.NewToastEvent("challenge sent"))
        break
      }

      em.Emit(ui.NewToastEvent(res))
      break
    case "$decline-challenge":
      res, err := c.ReadString()
      if err != nil {
        em.Emit(ui.NewToastEvent(err.Error()))
        break
      }

      if res != "success" {
        em.Emit(ui.NewToastEvent(res))
        break
      }

      em.Emit(ui.NewToastEvent("challenge declined"))
      break
    case "!challenge-accepted":
      name, err := c.ReadString()
      if err != nil {
        em.Emit(ui.NewToastEvent(err.Error()))
        break
      }

      message := fmt.Sprintf("%s accepted the challenge", name)
      em.Emit(ui.NewToastEvent(message))
      break
    case "!challenge-declined":
      name, err := c.ReadString()
      if err != nil {
        em.Emit(ui.NewToastEvent(err.Error()))
        break
      }

      message := fmt.Sprintf("%s declined the challenge", name)
      em.Emit(ui.NewToastEvent(message))
      break
    case "$accept-challenge":
      res, err := c.ReadString()
      if err != nil {
        em.Emit(ui.NewToastEvent(err.Error()))
        break
      }
      
      if res != "success" {
        em.Emit(ui.NewToastEvent(res))
        break
      }

      em.Emit(ui.NewToastEvent("challenge accepted"))
      break
    case "!challenge":
      name, err := c.ReadString()
      if err != nil {
        em.Emit(ui.NewToastEvent(err.Error()))
        break
      }

      em.Emit(ui.ConfirmEvent{
        Message: fmt.Sprintf("%s challenged you", name),

        CancelText: "Decline",
        OnCancel: func() {
          fmt.Print("declining challenge")
          message := fmt.Sprintf("?decline-challenge\n%s", name)
          c.WriteString(message)
        },

        ConfirmText: "Accept",
        OnConfirm: func() {
          fmt.Print("accepting challenge")
          message := fmt.Sprintf("?accept-challenge\n%s", name)
          c.WriteString(message)
        },
      })
      break
    default:
      fmt.Printf("received invalid message: %s\n", res)
      break
    }
  }
}

func (s *ServerState) readPlayers() ([]Player, error) {
  c := s.client.Connection
  players := []Player{}

  for {
    name, err := c.ReadString()
    if err != nil {
      return nil, err
    }

    if len(name) > 0 {
      player := Player{Name: name}
      if player.Name == s.client.Name {
        player.IsLocal = true
      }
      players = append(players, player)
      continue
    }

    break
  }

  return players, nil
}
