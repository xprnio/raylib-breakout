package server_lobby

import (
	"github.com/xprnio/raygo/internal/game/screens/server_lobby/player_list"
	"github.com/xprnio/raygo/internal/net/client"
)

type State struct {
  ServerName string
  Players []player_list.Player

  client *client.Client
}

func NewState(client *client.Client) (*State, error) {
  s := &State{ client: client }

  if err := s.loadServerInfo(); err != nil {
    return nil, err
  }

  return s, nil
}

func (s *State) loadServerInfo() error {
  c := s.client.Connection
  err := c.WriteString("?info")
  if err != nil {
    return err
  }

  for {
    res, err := c.ReadString()
    if err != nil {
      return err
    }

    switch res {
    case "!server-name":
      if s.ServerName, err = c.ReadString(); err != nil {
        return err
      }

      break
    case "!players":
      if s.Players, err = s.readPlayers(); err != nil {
        return err
      }
      break
    }

    if len(s.ServerName) == 0 {
      continue
    }

    if len(s.Players) == 0 {
      continue
    }

    break
  }

  return nil
}

func (s *State) readPlayers() ([]player_list.Player, error) {
  players := []player_list.Player{}
  c := s.client.Connection
  for {
    data, err := c.ReadString()
    if err != nil {
      return nil, err
    }

    if len(data) == 0 {
      return players, nil
    }

    players = append(players, player_list.Player{
      Name: data,
    })
  }
}
