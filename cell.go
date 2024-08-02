package termgraph;

//A cell is a character on the screen, which can have color codes that come
//before and after it.
type Cell struct {
  //The x and y location on the screen. This should only be set once, and 
  //therefore is readonly. 
  x int
  y int

  //The color code for the cell. If this is null, then this cell has no color,
  //and should be reset before it is written.
  colorCode string
  
  //The current value of this cell.
  currentVal rune
  
  //The value of this cell once it is updated.
  newVal rune

  //owners
  owners map[*Area]bool
}

//Updates cell
//This is to be called after the screen is updated. This is only to be done
//*After* the screen reflects the new value.
func (cell *Cell) updateCell() {
  cell.currentVal = cell.newVal
}

//---------- Cell getters and setters --------------------
func (c *Cell) getLocation() (x, y int) {
  return c.x, c.y
}

func (c *Cell) getValue() rune {
  return c.currentVal
}

func (c *Cell) setValue(r rune) {
  c.newVal = r
}

func (c *Cell) setCell(r rune, color string) {
  c.newVal = r
}

func newCell(x, y int, val rune, color string) *Cell {
  return &Cell{
    x: x,
    y: y,
    currentVal: val,
    newVal: val,
    colorCode: color,
  }
}
