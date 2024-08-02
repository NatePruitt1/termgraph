package termgraph;

import "fmt"

//---------- Public Control Functions ----------

//Moves the cursor to a location. Returns an error if the location is out of
//bounds
func MoveCursor(x int, y int) {
	fmt.Print("\x1b[%d;%dH", y, x)
}

//Prints the clear screen control code to completely erase the screen.
func ClearScreen() {
	fmt.Print("\x1b[2J")
}
