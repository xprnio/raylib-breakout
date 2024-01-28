package server

import (
	"fmt"
	"io"
	"net"

	"github.com/xprnio/raygo/internal/net/utils"
)

type Server struct {
  Name string
  Addr *net.TCPAddr

  Players map[string]*Player
  Challenges map[*Player]*Player
}


func New(name, address string) (*Server, error) {
  addr, err := net.ResolveTCPAddr("tcp", address)
  if err != nil {
    return nil, err
  }

  s := &Server{
    Name: name,
    Addr: addr,
    Players: make(map[string]*Player),
    Challenges: make(map[*Player]*Player),
  }
  return s, nil
}

func (s *Server) Start() error {
  l, err := net.ListenTCP("tcp", s.Addr)
  if err != nil {
    return err
  }

  for {
    conn, err := l.Accept()
    if err != nil {
      return err
    }

    go s.handleConnection(conn)
  }
}

func (s *Server) findPlayer(name string) *Player {
  for _, p := range s.Players {
    if p.Name == name {
      return p
    }
  }

  return nil
}

func (s *Server) handleConnection(conn net.Conn) {
  defer conn.Close()

  c := utils.NewConnection(conn)

  fmt.Println("client connected")
  for {
    err := s.handleCommand(c)
    if err == io.EOF {
      fmt.Println("client disconnected")
      return
    }

    if err != nil {
      fmt.Println("error:", err)
      response := fmt.Sprintf("error: %s", err.Error())
      c.WriteString(response)
      break
    }
  }
}

func (s *Server) handleDisconnect(p *Player) {
  fmt.Printf("player disconnected: %s\n", p.Name)
  s.removePlayer(p)
}

func (s *Server) handleCommand(c *utils.Connection) error {
  command, err := c.ReadString()
  if err != nil {
    return err
  }

  if len(command) == 0 {
    return nil
  }

  switch command {
  case "join":
    player, err := s.handleJoin(c)
    if err != nil {
      return err
    }

    s.addPlayer(player)

    defer s.handleDisconnect(player)
    fmt.Println("player joined:", player.Name)

    return player.Handle()
  default:
    return fmt.Errorf("invalid command: %s", command)
  }
}

func (s *Server) addPlayer(player *Player) {
  s.Players[player.Name] = player;
  s.notifyPlayers()
}

func (s *Server) removePlayer(player *Player) {
  delete(s.Players, player.Name)
  s.notifyPlayers()
}

func (s *Server) notifyPlayers() {
  players := "!players"
  for _, p := range s.Players {
    players = fmt.Sprintf("%s\n%s", players, p.Name)
  }

  for _, p := range s.Players {
    err := p.conn.WriteString(fmt.Sprintf("%s\n", players))
    if err != nil {
      err := fmt.Errorf("error sending player list to %s", p.Name)
      fmt.Println(err)
    }
  }
}

func (s *Server) handleJoin(c *utils.Connection) (*Player, error) {
  name, err := c.ReadString()
  if err != nil {
    return nil, err
  }

  if len(name) < 3 {
    return nil, fmt.Errorf("len(name) must be at least 3")
  }

  if len(name) > 16 {
    return nil, fmt.Errorf("len(name) must be less than 16")
  }

  if _, exists := s.Players[name]; exists {
    return nil, fmt.Errorf("name already used")
  }

  player := NewPlayer(name, s, c)
  player.conn.WriteString("success")
  return player, nil
}
