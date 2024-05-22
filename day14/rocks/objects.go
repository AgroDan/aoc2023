package rocks

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"fmt"
)

const (
	North = iota
	South
	East
	West
)

// standard coord obj
type Coord struct {
	X, Y int
}

// // this object will represent a rock
// type RoundRock struct {
// 	pos Coord
// }

// func NewRock(c Coord) RoundRock {
// 	return RoundRock{pos: c}
// }

// func (r RoundRock) PeekDir(dir int) Coord {
// 	// This function just returns the coordinate results
// 	// of this direction, it does not change the roundrock
// 	// position

// }

func (r RockMap) PeekDir(start Coord, dir int) Coord {
	// This will return the coordinate of the direction
	// from the starting point. This will NOT perform
	// error checking!
	switch dir {
	case North:
		return Coord{
			X: start.X,
			Y: start.Y - 1,
		}
	case South:
		return Coord{
			X: start.X,
			Y: start.Y + 1,
		}
	case East:
		return Coord{
			X: start.X + 1,
			Y: start.Y,
		}
	case West:
		return Coord{
			X: start.X - 1,
			Y: start.Y,
		}
	default:
		return Coord{
			X: -1,
			Y: -1,
		}
	}
}

func (r RockMap) IsCoordInBounds(c Coord) bool {
	// Will return a true or false as to whether
	// the supplied coordinate is actually in bounds.
	// This will not check if something exists at
	// this coordinate.
	if c.X < 0 || c.X >= r.MaxCols {
		return false
	}

	if c.Y < 0 || c.Y >= r.MaxRows {
		return false
	}
	return true
}

// a map of all coordinates of stationary rocks
type RockMap struct {
	cubeRocks, roundRocks map[Coord]bool
	MaxRows, MaxCols      int
}

func (r RockMap) DeepCopy() RockMap {
	// Creates a deep copy of the struct object
	nr := RockMap{}
	nr.cubeRocks = make(map[Coord]bool)
	nr.roundRocks = make(map[Coord]bool)
	nr.MaxRows, nr.MaxCols = r.MaxRows, r.MaxCols

	for k, v := range r.cubeRocks {
		nr.cubeRocks[k] = v
	}
	for k, v := range r.roundRocks {
		nr.roundRocks[k] = v
	}
	return nr
}

func (r *RockMap) MoveRocks(dir int) bool {
	// This function, given a direction, will move all the
	// round rocks once in that direction.
	//
	// This function will return a true if movement happened
	// otherwise false if nothing changed
	var state bool = false
	switch dir {
	case North:
		// we will iterate over each column from top to bottom
		for Xi := 0; Xi < r.MaxCols; Xi++ {
			for Yi := 0; Yi < r.MaxRows; Yi++ {
				thisCoord := Coord{
					X: Xi,
					Y: Yi,
				}

				if r.DoesRoundRockExist(thisCoord) {
					// Move it up if possible
					moveState := r.MoveOneRock(thisCoord, dir)
					if moveState == nil {
						state = true
					}
				}
			}
		}
	case South:
		// Do the above in reverse basically, iterate over rows
		// in reverse that is
		for Xi := 0; Xi < r.MaxCols; Xi++ {
			for Yi := r.MaxRows - 1; Yi >= 0; Yi-- {
				thisCoord := Coord{
					X: Xi,
					Y: Yi,
				}
				if r.DoesRoundRockExist(thisCoord) {
					// Move it down
					moveState := r.MoveOneRock(thisCoord, dir)
					if moveState == nil {
						state = true
					}
				}
			}
		}
	case East:
		// now we play with the outer loop
		for Yi := 0; Yi < r.MaxRows; Yi++ {
			for Xi := r.MaxCols - 1; Xi >= 0; Xi-- {
				thisCoord := Coord{
					X: Xi,
					Y: Yi,
				}
				if r.DoesRoundRockExist(thisCoord) {
					// Move it east
					moveState := r.MoveOneRock(thisCoord, dir)
					if moveState == nil {
						state = true
					}
				}
			}
		}
	case West:
		for Yi := 0; Yi < r.MaxRows; Yi++ {
			for Xi := 0; Xi < r.MaxCols; Xi++ {
				thisCoord := Coord{
					X: Xi,
					Y: Yi,
				}
				if r.DoesRoundRockExist(thisCoord) {
					// move it west
					moveState := r.MoveOneRock(thisCoord, dir)
					if moveState == nil {
						state = true
					}
				}
			}
		}

	default:
		panic("WHAT???")
	}
	return state
}

func (r *RockMap) MoveOneRock(c Coord, dir int) error {
	// This will move a rock from the particular coordinate
	// one block towards the supplied direction. Will error
	// if something exists in that direction, or if we hit a
	// boundary, or if this is an invalid rock
	if !r.DoesRoundRockExist(c) {
		return errors.New("no rock exists on provided coord")
	}

	if r.DoesStaticRockExist(c) {
		return errors.New("cannot move static rock")
	}

	if !r.IsCoordInBounds(c) {
		return errors.New("source rock is not in bounds")
	}

	peek := r.PeekDir(c, dir)
	if !r.IsCoordInBounds(peek) {
		return errors.New("dir is off the map")
	}

	if r.DoesRoundRockExist(peek) || r.DoesStaticRockExist(peek) {
		return errors.New("rock exists in intended direction")
	}

	// all tests pass otherwise, let's move the rock
	r.AddRoundRock(peek)
	r.DeleteRoundRock(c)
	return nil
}

func (r *RockMap) AddStaticRock(c Coord) error {
	// Adds a stationary rock to the rockmap
	// First, does the rock exist already
	if !r.DoesStaticRockExist(c) {
		r.cubeRocks[c] = true
		return nil
	} else {
		return errors.New("static rock already exists")
	}
}

func (r *RockMap) AddRoundRock(c Coord) {
	// adds a round rock. NO CHECKING! JUST ADDS IT!
	r.roundRocks[c] = true
}

func (r *RockMap) DeleteRoundRock(c Coord) error {
	if r.DoesRoundRockExist(c) {
		delete(r.roundRocks, c)
		return nil
	} else {
		return errors.New("round rock does not exist at coord")
	}
}

func (r RockMap) DoesStaticRockExist(c Coord) bool {
	if _, ok := r.cubeRocks[c]; ok {
		return true
	}
	return false
}

func (r RockMap) DoesRoundRockExist(c Coord) bool {
	if _, ok := r.roundRocks[c]; ok {
		return true
	}
	return false
}

func (r RockMap) IsValidEmptySpace(c Coord) bool {
	// First, are there any existing rocks here
	if r.DoesRoundRockExist(c) {
		return false
	}
	if r.DoesRoundRockExist(c) {
		return false
	}

	// now are these within the boundaries of the map?
	if c.X < 0 || c.X >= r.MaxCols {
		return false
	}
	if c.Y < 0 || c.Y >= r.MaxRows {
		return false
	}

	return true
}

func (r RockMap) Hash() (string, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	// serialize everything
	err := encoder.Encode(r)
	if err != nil {
		return "", err
	}

	// Now let's try to encode the arrays
	for k := range r.roundRocks {
		err = encoder.Encode(k)
		if err != nil {
			return "", err
		}
	}

	// Compute the hash
	hash := sha256.Sum256(buf.Bytes())
	return fmt.Sprintf("%x", hash), nil
}
