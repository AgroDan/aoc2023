package pipes

import (
	"errors"
	"fmt"
)

type PipeMap struct {
	m           [][]rune
	start       Coord
	PipeCircuit *Set
}

func NewPipeMap(lines []string) PipeMap {
	p := PipeMap{}
	for _, line := range lines {
		runes := []rune(line)
		p.m = append(p.m, runes)
	}

	for i := 0; i < len(p.m); i++ {
		for j := 0; j < len(p.m[i]); j++ {
			if p.m[i][j] == 'S' {
				p.start = Coord{
					Y: i,
					X: j,
				}
			}
		}
	}
	fmt.Printf("Start: %+v\n", p.start)
	p.PipeCircuit = NewSet()
	return p
}

func (p PipeMap) PrintMap() {
	// prints the map
	for y := 0; y < len(p.m); y++ {
		for x := 0; x < len(p.m[y]); x++ {
			fmt.Printf("%c", p.m[y][x])
		}
		fmt.Printf("\n")
	}
}

// Now let's write a function to clean the map!
// Note that this can ONLY be done after the PipeCircuit has
// been filled out! This will basically go over the entire map
// and change anything that isn't in the circuit into dirt.
// Additionally, it will replace the starting coord with the
// proper pipe.

func (p *PipeMap) CleanMyPipes() {
	// this is fairly destructive!
	for y := 0; y < len(p.m); y++ {
		for x := 0; x < len(p.m[y]); x++ {
			thisCoord := Coord{
				X: x,
				Y: y,
			}

			if !p.PipeCircuit.Contains(thisCoord) {
				p.m[y][x] = '.'
			}
		}
	}

	// Now let's alter that pesky start tile
	dirs := p.GetValidNeighbors(p.start)
	// these only get the coords
	// we need to find out the directions to make a valid choice
	var validDirs []int

	for _, v := range dirs {
		var checkDirs = [4]int{N, S, E, W}
		for _, d := range checkDirs {
			peekDirection := GetCoordDirection(p.start, d)
			if peekDirection == v {
				validDirs = append(validDirs, d)
			}
		}
	}

	// should only be 2 dirs here, but panic if otherwise
	if len(validDirs) != 2 {
		panic("more than 2 valid directions shown")
	}
	var startRune rune
	switch validDirs[0] {
	case N:
		switch validDirs[1] {
		case E:
			startRune = 'L'
		case W:
			startRune = 'J'
		case S:
			startRune = '|'
		default:
			panic("not a valid tile from N")
		}
	case S:
		switch validDirs[1] {
		case N:
			startRune = '|'
		case E:
			startRune = 'F'
		case W:
			startRune = '7'
		default:
			panic("not a valid tile from S")
		}
	case E:
		switch validDirs[1] {
		case N:
			startRune = 'L'
		case W:
			startRune = '-'
		case S:
			startRune = 'F'
		default:
			panic("not a valid tile from E")
		}
	case W:
		switch validDirs[1] {
		case N:
			startRune = 'J'
		case E:
			startRune = '-'
		case S:
			startRune = '7'
		default:
			panic("not a valid tile from W")
		}
	default:
		panic("not a valid start tile")
	}
	p.m[p.start.Y][p.start.X] = startRune
}

func (p PipeMap) GetMapLocation(c Coord) rune {
	// Returns the actual rune at the coordinate
	return p.m[c.Y][c.X]
}

func (p PipeMap) GetStart() Coord {
	// returns the start coordinates
	return p.start
}

func (p PipeMap) IsValid(c Coord) bool {
	// This will do bounds checking
	if c.Y < 0 || c.Y >= len(p.m) {
		return false
	}

	if c.X < 0 || c.X >= len(p.m[c.Y]) {
		return false
	}
	return true
}

func (p PipeMap) GetValidNeighbors(c Coord) []Coord {
	var retval []Coord

	// The tile
	// fmt.Printf("X: %d, Y: %d\n", c.X, c.Y)
	currentTile := p.GetMapLocation(c)          // remember this returns the actual rune!
	currentTileType := GetTileType(currentTile) // this returns the enum type!

	var directions []int

	// First, let's see what kind of tile this is so we can infer directions
	switch currentTileType {
	case NorthSouth:
		directions = append(directions, N, S)
	case WestEast:
		directions = append(directions, W, E)
	case NorthEast:
		directions = append(directions, N, E)
	case NorthWest:
		directions = append(directions, N, W)
	case SouthWest:
		directions = append(directions, S, W)
	case SouthEast:
		directions = append(directions, S, E)
	case Starting:
		directions = append(directions, N, S, E, W)
		// default:
		// What could this be then?
	}

	for _, d := range directions {
		checkCoord := GetCoordDirection(c, d)

		if !p.IsValid(checkCoord) {
			continue
		}
		thisTile := p.GetMapLocation(checkCoord)
		tileType := GetTileType(thisTile)

		switch d {
		case N:
			// we are checking north, what applies are things connecting south
			if tileType == NorthSouth || tileType == SouthEast || tileType == SouthWest {
				retval = append(retval, checkCoord)
			}
		case S:
			// checking south, need things connecting north
			if tileType == NorthSouth || tileType == NorthEast || tileType == NorthWest {
				retval = append(retval, checkCoord)
			}
		case E:
			// checking east, need things connecting west
			if tileType == WestEast || tileType == NorthWest || tileType == SouthWest {
				retval = append(retval, checkCoord)
			}
		case W:
			// checking west, need things connecting east
			if tileType == WestEast || tileType == NorthEast || tileType == SouthEast {
				retval = append(retval, checkCoord)
			}
		}
	}

	return retval
}

func (p PipeMap) CountLayers(c Coord, dir int) (int, error) {
	// Given a coordinate, will count how many transitions until it hits a boundary.
	// errors out if given wrong info.

	// direction must be E or W!
	if dir != E && dir != W {
		return -1, errors.New("direction must be E or W")
	}

	// make sure this is a "Ground" object
	checkValidRune := GetTileType(p.GetMapLocation(c))
	if checkValidRune != Ground {
		return -1, errors.New("coordinate must be a Ground TileType")
	}

	// Just for yuks lets make sure the coord is valid
	if !p.IsValid(c) {
		return -1, errors.New("coordinate must be valid within the confines of the map")
	}

	// Now we loop
	var counter int = 0
	currentCoord := c
	for {
		currentCoord = GetCoordDirection(currentCoord, dir)
		// First, let's see if this is valid. Because if it isn't, we failed bounds checking
		// and we're at the border
		// fmt.Printf("Current coord: %+v\n", currentCoord)
		if !p.IsValid(currentCoord) {
			break
		}
		currentTile := GetTileType(p.GetMapLocation(currentCoord))

		// Now with the currentTile, we can begin counting...except for certain circumstances.
		switch currentTile {
		case Ground, WestEast:
			// Don't care about . or -
			continue
		case NorthSouth:
			// Tile is |
			counter++
		case SouthWest:
			// Tile is 7
			// Traverse until . or L is seen.
			// If L, count as 1.
			// if F, count as 2.
			if dir == W {
			swloop:
				for {
					currentCoord = GetCoordDirection(currentCoord, dir)
					// fmt.Printf("Current coord: %+v\n", currentCoord)
					if !p.IsValid(currentCoord) {
						break swloop
					}

					currentTile = GetTileType(p.GetMapLocation(currentCoord))
					switch currentTile {
					case NorthEast:
						// tile is L, making it a swoop and definitely a boundary cross
						// fmt.Printf("Found one layer\n")
						counter++
						break swloop
					case SouthEast:
						// tile is an F, so it's not technically a boundary cross
						// but I'll count it as 2 because it's like going over another layer
						// fmt.Printf("Found two layers\n")
						counter += 2
						break swloop
					default:
						// any other tile we don't care about
						continue
					}
				}
			}
		case NorthEast:
			// Tile is starting with L if we go E
			// If 7, count as 1.
			// if J, count as 2.
			if dir == E {
			neloop:
				for {
					currentCoord = GetCoordDirection(currentCoord, dir)
					if !p.IsValid(currentCoord) {
						break
					}
					currentTile = GetTileType(p.GetMapLocation(currentCoord))
					switch currentTile {
					case SouthWest:
						counter++
						break neloop
					case NorthWest:
						counter += 2
						break neloop
					default:
						continue
					}
				}
			}
			counter++
		case NorthWest:
			// Tile is a J, if going W
			// if F, count as 1
			// if L, count as 2
			if dir == W {
			nwloop:
				for {
					currentCoord = GetCoordDirection(currentCoord, dir)
					if !p.IsValid(currentCoord) {
						break
					}
					currentTile = GetTileType(p.GetMapLocation(currentCoord))
					switch currentTile {
					case SouthEast:
						// F
						// fmt.Printf("Found one layer <- FJ\n")
						counter++
						break nwloop
					case NorthEast:
						// L
						// fmt.Printf("Found two layers <- LJ\n")
						counter += 2
						break nwloop
					default:
						continue
					}
				}
			}
		case SouthEast:
			// Tile is F, if going E
			// if J, count as 1
			// if 7, count as 2
			if dir == E {
			seloop:
				for {
					currentCoord = GetCoordDirection(currentCoord, dir)
					if !p.IsValid(currentCoord) {
						break
					}
					currentTile = GetTileType(p.GetMapLocation(currentCoord))
					switch currentTile {
					case NorthWest:
						// J
						counter++
						break seloop
					case SouthWest:
						// 7
						counter += 2
						break seloop
					default:
						continue
					}
				}
			}
		}

	}

	return counter, nil
}

func (p PipeMap) GetInternal() int {
	// MAKE SURE THE MAP IS CLEANED FIRST!!!
	var retval int = 0
	for y := 0; y < len(p.m); y++ {
		for x := 0; x < len(p.m[y]); x++ {
			if p.m[y][x] == '.' {
				thisCoord := Coord{
					X: x,
					Y: y,
				}
				layers, err := p.CountLayers(thisCoord, W)
				if err != nil {
					panic(err)
				}

				if IsOdd(layers) {
					retval++
				}
			}
		}
	}
	return retval
}
