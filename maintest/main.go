package main

import (
	"fmt"
	"os"
	"time"

	termgraph "github.com/NatePruitt1/termgraph"
)

func main() {
  screen, err := termgraph.TakeScreen()
  if err != nil {
    return;
  } else {   
    root := screen.GetArea(0)
    child, err := root.NewChild(50, 0, 50, 10, "ChildArea")
    if err != nil {
      fmt.Fprintln(os.Stdout, "Error making child")
      fmt.Println(err)
      return
    } else {
      err := child.Put(25, 9, 'c')
      if err != nil {
        fmt.Fprintln(os.Stdout, "Error putting")
        return
      }
      screen.UpdateScreen()
      time.Sleep(time.Second * 10)
    }
  }
}
