package termgraph

import "errors"

/*An Area is a part of the screen. It can be bordered or unbordered.
  If an Area is not bordered, then any area which owns the cell can print too
  it.
  If an Area is bordered, then that cell will be marked a "border" cell. This
  means that no sub-area may own it, and no parent area may modify it.

  Areas have children. Child areas may be less than or equal to in size to the
  parent area, and must be fully contained by the parent area.
*/

/*Areas do not hold a reference to their children, instead, children hold
  A map of areas allowed to print to them. When an area attempts to print,
  it will "request" the screen to print at the location, and then receive
  an error or a confirmation of success. */
type Area struct {
  //An areas location on the screen.
  //Areas hold references to their parents and children, since it runs in a
  //tree structure (like html)
  screen *Screen
  parent *Area;
  children []*Area;

  localX int
  localY int
  
  absX int
  absY int

  width int
  height int

  id string
}

//Set Area Location
//One of the pivotal interface functions. Sets a cell within the areas newVal
//value (assuming it is a valid modification).
//Changes will not be represented until the screen in cleared
func (area *Area) SetLocation(lX, lY int, c rune) error {
  aX, aY := area.localToAbsolute(lX, lY)
  err := area.screen.setLocation(aX, aY, c, area)
  return err
}

//Todo: Move this functionality to the screen
//Allocate a new area, assume that this has already been checked and everything
func newArea(lX, lY, aX, aY, width, height int, name string) Area{
  newArea := Area{}
  newArea.parent = nil
  newArea.children = make([]*Area, 0)
  newArea.localX = lX
  newArea.localY = lY
  newArea.absX = aX
  newArea.absY = aY
  newArea.width = width
  newArea.height = height
  newArea.id = name
  return newArea
}

//Creates a child area. The coordinates given are local to the parent area.
//The parent area is checked for violations first.
//IF an error is returned, an area will also be returned. This is the area that
//was offended by the error
func (parent *Area) makeChild(lX, lY, width, height int, name string) (*Area, error) {
  cornerX := lX + width
  cornerY := lY + height

  if lX < 0 || lY < 0 || cornerX > parent.width || cornerY > parent.height {
    return parent, errors.New("Area out of bounds of parent")
  }

  //go through children, making sure none are "intruded" by this.
  for _, c := range parent.children {
    childCornerX := c.localX + c.width
    childCornerY := c.localY + c.height
    //case 1: the top corner is within the other child
    if lX >= c.localX && lY >= c.localY && lX <= childCornerX && lY <= childCornerY {
      return c, errors.New("Area overlaps existing child area")
    } else if cornerX >= c.localX && cornerY >= c.localY &&cornerX <= childCornerX && cornerY <= childCornerY {
      //case 2: the bottom corner is within the other child
      return c, errors.New("Area overlaps existing child area")
    } else if lX <= c.localX && lY <= c.localY && cornerX >= childCornerX && cornerY >= childCornerY {
      //case 3: This area contains the child (and therefore should be its parent)
      return c, errors.New("Area contains existing child area")
    }
  }
  
  //Parent and children have not objected, create new area.
  aX, aY := parent.localToAbsolute(lX, lY)
  newArea := parent.screen.newScreenArea(newArea(lX, lY, aX, aY ,width, height, name))
  newArea.screen = parent.screen
  newArea.parent = parent
  parent.children = append(parent.children, newArea)
  return newArea, nil
}

//Converts the local coordinates x and y to absolute screen coordinates.
func (area *Area) localToAbsolute(x, y int) (int, int) {
  return area.absX + x, area.absY + y
}
