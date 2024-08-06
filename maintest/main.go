package main

import (
    termgraph "github.com/NatePruitt1/termgraph"
    "fmt"
  )

func main() {
  screen, err := termgraph.TakeScreen()
  if err != nil {
    return;
  } else {   
    fmt.Print(screen)
    area := screen.GetArea(0)
    area.SetLocation(50, 50, 'c')
  }
}
