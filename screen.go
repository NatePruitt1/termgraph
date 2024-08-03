package termgraph

import (
	"errors"
  "fmt"
	tsize "github.com/kopoli/go-terminal-size"
)

// Screen is a struct that represents the terminal screen, and holds the
// information required to manage the screen.
// The screen has a private interface that allows one to "request" cell reads
// and writes.
// The screen manages cells, which are the private representation of the characters
// on the screen, and areas, which are the public representation of the
// writeable areas on the screen.
type Screen struct {

	//The characters that are shown on the cells screen.
	//We hold pointers because we want these values to be mutable.
	//This is to avoid tons of memory allocation when creating new screens.
	//In the form [x][y]
	cells [][](*Cell)

	//The areas on the screen
	areas [](*Area)
}

//---------- Public Screen Utilities ----------

// Takes control of the screen represented by Stdout.
// Sets up the private cell arrays, and then creates the first "parent"
// area.
func TakeScreen() (*Screen, error) {
	newScreen := new(Screen)
	size, err := tsize.GetSize()
	if err != nil {
		return nil, err
	} else {
		newScreen.areas = make([]*Area, 1)
		newScreen.areas[0] = newArea(0, 0, 0, 0, newScreen.GetWidth(), newScreen.GetHeight(), "parent")

		//Create the cell arrays
		newScreen.cells = make([][]*Cell, size.Width)

		for i := range newScreen.cells {
			newScreen.cells[i] = make([]*Cell, size.Height)
			for j := range newScreen.cells[i] {
				newScreen.cells[i][j] = newCell(i, j, ' ', "")
				newScreen.cells[i][j].owners[newScreen.areas[0]] = true
			}
		}

		return newScreen, nil
	}
}

//Update the screen
//This function first calculates a string, and then prints it to Stdout, which
//causes the screen to represent all of the changes made to its areas.
func (screen *Screen) UpdateScreen() {
  //TODO: Draw an "Edit tree" which creates the most efficient possible
  //edit to bring the current screen to the new one. Research curses library
  //approach.

  //For now, delete the screen and update it one by one.
  ClearScreen()

  for x := range screen.cells {
    for y := range screen.cells[x] {
      MoveCursor(x,y)
      fmt.Printf("%c", screen.cells[x][y].getValue())
      screen.cells[x][y].updateCell()
    }
  }
}

// ---------- Private Screen Utilities ----------
func (screen *Screen) addArea(area *Area) {
	screen.areas = append(screen.areas, area)
	for x := area.absX; x < area.width+area.absX; x++ {
		for y := area.absY; y < area.height+area.absY; y++ {
			screen.cells[x][y].owners[area] = true
		}
	}
}

func (screen *Screen) setLocation(aX, aY int, c rune, area *Area) error {
	//check that the set is within bounds
	if screen.checkBounds(aX, aY) == false {
		return errors.New("Out of bounds cell")
	}

	//check in that the cell is writeable by this area
	_, ok := screen.cells[aX][aY].owners[area]

	if ok {
		screen.cells[aX][aY].setCell(c, "")
		return nil
	} else {
		return errors.New("Area does not own cell")
	}
}

func (screen *Screen) checkBounds(x, y int) bool {
	return x >= 0 && x < screen.GetWidth() && y >= 0 && y < screen.GetHeight()
}

// ---------- Screen Getters and Setters ----------
func (screen *Screen) GetWidth() int {
	return len(screen.cells)
}

func (screen *Screen) GetHeight() int {
	return len(screen.cells[0])
}

func (screen *Screen) GetCell(x, y int) (*Cell, error) {
	if screen.checkBounds(x, y) {
		return screen.cells[x][y], nil
	} else {
		return nil, errors.New("Cell out of bounds")
	}
}
