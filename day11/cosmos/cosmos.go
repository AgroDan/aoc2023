package cosmos

import "fmt"

type Coord struct {
	X, Y int
}

type Universe struct {
	Galaxies      []Coord
	Length, Width int
}

func (u Universe) GalaxyExists(c Coord) bool {
	// Checks the entire universe looking for
	// the galaxy that exists in the given coord
	// if possible.
	for _, v := range u.Galaxies {
		if v == c {
			return true
		}
	}
	return false
}

func NewUniverse(line []string) Universe {
	u := Universe{}
	// Get the length/width of universe
	// Width = X axis
	// Length = Y axis
	u.Length = len(line)
	u.Width = len(line[0])

	// now find the galaxies
	for y := 0; y < len(line); y++ {
		// This is Y coordinate, up and down

		for x := 0; x < len(line[y]); x++ {
			// This is the X coordinate, left and right
			if line[y][x] == '#' {
				c := Coord{
					X: x,
					Y: y,
				}
				u.Galaxies = append(u.Galaxies, c)
			}
		}
	}
	return u
}

func (u Universe) PrintUniverse() {
	for y := 0; y < u.Length; y++ {
		for x := 0; x < u.Width; x++ {
			tempCoord := Coord{
				X: x,
				Y: y,
			}
			if u.GalaxyExists(tempCoord) {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func (u *Universe) updateRows(from int) {
	// This updates all the galaxies' rows by one
	// starting from the "from" variable. This is
	// the Y axis!
	for i := 0; i < len(u.Galaxies); i++ {
		if u.Galaxies[i].Y >= from {
			u.Galaxies[i].Y++
		}
	}
	u.Length++
}

func (u *Universe) updateCols(from int) {
	// Updates all the galaxies' columns by one
	// starting from the "from" variable. This is
	// the X axis!
	for i := 0; i < len(u.Galaxies); i++ {
		if u.Galaxies[i].X >= from {
			u.Galaxies[i].X++
		}
	}
	u.Width++
}

func (u *Universe) ExpandUniverse() {
	// This is a workhorse function which does two things.
	// First, it determines if the universe can be expanded,
	// Then it updates all the galaxies.

	// First, lets check the columns, we care about the X coordinates
	// here.

	var emptyCols []int
	for x := 0; x < u.Width; x++ {
		empty := true
		for _, v := range u.Galaxies {
			if v.X == x {
				// not empty
				empty = false
				break
			}
		}
		if empty {
			emptyCols = append(emptyCols, x)
		}
	}
	expCols := 0
	for _, v := range emptyCols {
		u.updateCols(v + expCols)
		expCols++
	}

	// Now do the rows
	var emptyRows []int
	for y := 0; y < u.Length; y++ {
		empty := true
		for _, v := range u.Galaxies {
			if v.Y == y {
				// not empty
				empty = false
				break
			}
		}
		if empty {
			emptyRows = append(emptyRows, y)
		}
	}
	expRows := 0
	for _, v := range emptyRows {
		// need to update all further rows too
		u.updateRows(v + expRows)
		expRows++
	}
}

func (u *Universe) arbitrarilyUpdateRows(from, howMuch int) {
	// This updates all the galaxies' rows by one
	// starting from the "from" variable. This is
	// the Y axis!
	for i := 0; i < len(u.Galaxies); i++ {
		if u.Galaxies[i].Y >= from {
			u.Galaxies[i].Y += howMuch
		}
	}
	u.Length += howMuch
}

func (u *Universe) arbitrarilyUpdateCols(from, howMuch int) {
	// Updates all the galaxies' columns by one
	// starting from the "from" variable. This is
	// the X axis!
	for i := 0; i < len(u.Galaxies); i++ {
		if u.Galaxies[i].X >= from {
			u.Galaxies[i].X += howMuch
		}
	}
	u.Width += howMuch
}

func (u *Universe) ArbitrarilyExpandUniverse(howMuch int) {
	// This is a workhorse function which does two things.
	// First, it determines if the universe can be expanded,
	// Then it updates all the galaxies.

	// First, lets check the columns, we care about the X coordinates
	// here.

	howMuch--

	var emptyCols []int
	for x := 0; x < u.Width; x++ {
		empty := true
		for _, v := range u.Galaxies {
			if v.X == x {
				// not empty
				empty = false
				break
			}
		}
		if empty {
			emptyCols = append(emptyCols, x)
		}
	}
	expCols := 0
	for _, v := range emptyCols {
		u.arbitrarilyUpdateCols(v+expCols, howMuch)
		expCols += howMuch
	}

	// Now do the rows
	var emptyRows []int
	for y := 0; y < u.Length; y++ {
		empty := true
		for _, v := range u.Galaxies {
			if v.Y == y {
				// not empty
				empty = false
				break
			}
		}
		if empty {
			emptyRows = append(emptyRows, y)
		}
	}
	expRows := 0
	for _, v := range emptyRows {
		// need to update all further rows too
		u.arbitrarilyUpdateRows(v+expRows, howMuch)
		expRows += howMuch
	}
}
