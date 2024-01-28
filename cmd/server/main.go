package main

import (
	"fmt"
	"os"

	"github.com/xprnio/raygo/internal/net/server"
)

func main() {
  s, err := server.New("Kohila Surf 24/7", ":6969")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Printf("starting server on %s\n", s.Addr)
  s.Start()
}
