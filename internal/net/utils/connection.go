package utils

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Connection struct {
  Conn net.Conn
  Reader *bufio.Reader
}

func NewConnection(conn net.Conn) *Connection {
  return &Connection{
    Conn: conn,
    Reader: bufio.NewReader(conn),
  }
}

func (c *Connection) WriteString(data string) (error) {
  data = fmt.Sprintf("%s\n", data)
  _, err := c.Conn.Write([]byte(data))
  return err
}

func (c *Connection) ReadString() (string, error) {
  data, err := c.Reader.ReadString('\n')
  if err != nil {
    return "", err
  }

  return strings.TrimSpace(data), nil
}
