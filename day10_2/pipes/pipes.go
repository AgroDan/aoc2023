package pipes

import (
	"errors"
	"fmt"
)

const (
	N = iota
	S
	E
	W
)

func Surrounding(c Coord) [4]Coord {
	// This will return all 4 directions around the provided
	// coordinate. This does NO bounds checking.
	return [4]Coord{c.Peek(N), c.Peek(S), c.Peek(E), c.Peek(W)}
}

func (p PipesMap) SafeSurrounding(c Coord) []Coord {
	// This will return a slice of coordinates that will have
	// bounds checking on it. If the coordinate it checks
	// is out of bounds, it simply won't add it to the
	// return slice
	var retval []Coord
	dirs := [4]int{N, S, E, W}
	for _, v := range dirs {
		thisCoord := c.Peek(v)
		if thisCoord.Y >= len(p.m) || thisCoord.Y < 0 ||
			thisCoord.X >= len(p.m[0]) || thisCoord.X < 0 {
			continue
		}
		retval = append(retval, thisCoord)
	}
	return retval
}

type PipesMap struct {
	m        [][]rune
	starting Coord
}

func (p PipesMap) PrintMap() {
	for y := 0; y < len(p.m); y++ {
		for x := 0; x < len(p.m[y]); x++ {
			fmt.Printf("%c", p.m[y][x])
		}
		fmt.Printf("\n")
	}
}

func (p PipesMap) IsConnected(from, to Coord) bool {
	// This function will determine if the _from_ coord is
	// connected to the _to_ coord.

	toRune, err1 := p.GetRune(to)
	if err1 != nil {
		// out of bounds
		return false
	}

	dir1, dir2, err2 := GetTileDirection(toRune)
	if err2 != nil {
		// could be an open tile, in which case this isn't connected
		return false
	}

	if to.Peek(dir1) == from || to.Peek(dir2) == from {
		return true
	}
	return false
}

func NewPipesMap(lines []string) PipesMap {
	// Constructor for the PipesMap object
	p := PipesMap{}
	for y := 0; y < len(lines); y++ {
		var line []rune
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == 'S' {
				p.starting.X = x
				p.starting.Y = y
			}
			line = append(line, rune(lines[y][x]))
		}
		p.m = append(p.m, line)
	}

	// Now, lets figure out what tile is the starting tile
	var connectedDirs []int
	directions := [4]int{N, S, E, W}
	for _, d := range directions {
		if p.IsConnected(p.starting, p.starting.Peek(d)) {
			connectedDirs = append(connectedDirs, d)
		}
	}

	// should only be 2 connected.
	switch connectedDirs[0] {
	case N:
		switch connectedDirs[1] {
		case S:
			p.m[p.starting.Y][p.starting.X] = '|'
		case W:
			p.m[p.starting.Y][p.starting.X] = 'J'
		case E:
			p.m[p.starting.Y][p.starting.X] = 'L'
		}
	case S:
		switch connectedDirs[1] {
		case W:
			p.m[p.starting.Y][p.starting.X] = '7'
		case N:
			p.m[p.starting.Y][p.starting.X] = '|'
		case E:
			p.m[p.starting.Y][p.starting.X] = 'F'
		}
	case E:
		switch connectedDirs[1] {
		case S:
			p.m[p.starting.Y][p.starting.X] = 'F'
		case W:
			p.m[p.starting.Y][p.starting.X] = '-'
		case N:
			p.m[p.starting.Y][p.starting.X] = 'L'
		}
	case W:
		switch connectedDirs[1] {
		case N:
			p.m[p.starting.Y][p.starting.X] = 'J'
		case E:
			p.m[p.starting.Y][p.starting.X] = '-'
		case S:
			p.m[p.starting.Y][p.starting.X] = '7'
		}
	}

	return p
}

func (p PipesMap) GetRune(c Coord) (rune, error) {
	// returns the rune at the listed coordinate.
	if c.Y >= len(p.m) || c.Y < 0 ||
		c.X >= len(p.m[0]) || c.X < 0 {
		return '?', errors.New("out of bounds")
	}
	return p.m[c.Y][c.X], nil
}

func (p PipesMap) IsAnEdge(c Coord) bool {
	// This function determines if we are at an edge (or beyond)
	if c.Y >= len(p.m)-1 || c.Y <= 0 ||
		c.X >= len(p.m[0])-1 || c.X <= 0 {
		return true
	}
	return false
}

func GetTileDirection(tile rune) (int, int, error) {
	// This, given the appropriate tile, will return the
	// two directions the pipe on the tile should throw.
	switch tile {
	case '|':
		return N, S, nil
	case '-':
		return E, W, nil
	case 'J':
		return W, N, nil
	case '7':
		return W, S, nil
	case 'L':
		return N, E, nil
	case 'F':
		return E, S, nil
	case '.':
		return N, N, errors.New("open ground tile")
	case 'S':
		return N, N, errors.New("starting point tile")
	default:
		return N, N, errors.New("unknown tile")
	}
}

func GetOppositeDirection(dir int) int {
	// Given a direction, simply returns its opposite
	switch dir {
	case N:
		return S
	case S:
		return N
	case E:
		return W
	case W:
		return E
	default:
		return -1
	}
}

func (p PipesMap)