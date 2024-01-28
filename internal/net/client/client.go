package client

import (
	"fmt"
	"net"

	"github.com/xprnio/raygo/internal/net/utils"
)

type Client struct {
  Name string
  Connection *utils.Connection
}

func New(name string) *Client {
  return &Client{ Name: name }
}

func (c *Client) Connect(address string) error {
  addr, err := net.ResolveTCPAddr("tcp", address)
  if err != nil {
    return err
  }

  conn, err := net.DialTCP("tcp", nil, addr)
  if err != nil {
    return err
  }

  c.Connection = utils.NewConnection(conn)
  if err := c.sendConnect(); err != nil {
    return err
  }

  return nil
}

func (c *Client) Close() {
  if c.Connection != nil {
    c.Connection.Conn.Close()
  }
}

func (c *Client) sendConnect() error {
  request := fmt.Sprintf("join\n%s\n", c.Name)
  err := c.Connection.WriteString(request)
  if err != nil {
    return err
  }

  response, err := c.Connection.ReadString()
  if err != nil {
    return err
  }

  if response != "success" {
    return fmt.Errorf(response)
  }

  return nil
}
