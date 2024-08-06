package termgraph

import (
	"fmt"
	"io"
)

//---------- Public Control Functions ----------

//Moves the cursor to a location. Returns an error if the location is out of
//bounds
func MoveCursor(x int, y int, pipe io.Writer) {
	fmt.Fprintf(pipe, "\x1b[%d;%dH", y + 1, x + 1)
}

//Prints the clear screen control code to completely erase the screen.
func ClearScreen(pipe io.Writer) {
	fmt.Fprintf(pipe, "\x1b[2J")
}
