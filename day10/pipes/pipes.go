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

type Coord struct {
	// standard coord struct
	X, Y int
}

func (c *Coord) Move(dir int) {
	// There is no error checking here. It will move the coordinate
	// in any direction told. It is not the Coord's responsibility to
	// perform any sort of bounds-checking.
	switch dir {
	case N:
		c.Y--
	case S:
		c.Y++
	case E:
		c.X++
	case W:
		c.X--
	}
}

func (c Coord) Peek(dir int) Coord {
	// This won't change anything, just return the coordinate of
	// the given direction. Again, no bounds checking here.
	r := Coord{
		X: c.X,
		Y: c.Y,
	}
	switch dir {
	case N:
		r.Y--
	case S:
		r.Y++
	case E:
		r.X++
	case W:
		r.X--
	}
	return r
}

type PipesMap struct {
	m        [][]rune
	starting Coord
}

func NewPipesMap(lines []string) PipesMap {
	// Constructor for the PipesMap object
	p := PipesMap{}
	for i := 0; i < len(lines); i++ {
		var line []rune
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == 'S' {
				p.starting.X = j
				p.starting.Y = i
			}
			line = append(line, rune(lines[i][j]))
		}
		p.m = append(p.m, line)
	}
	// Now replace the starting tile
	p.replaceStartTile()
	return p
}

func (p PipesMap) PrintMap() {
	for i := 0; i < len(p.m); i++ {
		for j := 0; j < len(p.m[i]); j++ {
			fmt.Printf("%c", p.m[i][j])
		}
		fmt.Printf("\n")
	}
}

func (p PipesMap) PrintVisitedMap(bc Circuit) {
	for y := 0; y < len(p.m); y++ {
		for x := 0; x < len(p.m[y]); x++ {
			thisCoord := Coord{X: x, Y: y}
			if bc.Contains(thisCoord) {
				fmt.Printf("#")
			} else {
				fmt.Printf("%c", p.m[y][x])
			}
		}
		fmt.Printf("\n")
	}
}

func (p PipesMap) IdentifyTile(c Coord) (Coord, Coord, error) {
	// This function, given a specific coordinate, will identify
	// the two places you can go given a specific tile coordinate.
	// If this is a pipe, it will return the two coordinates that
	// are possible to traverse through. If this is a dirt block
	// (meaning not inside a pipe) it will throw an error.
	thisTile := p.m[c.Y][c.X]
	// fmt.Printf("This tile: %c\n", thisTile)
	switch thisTile {
	case '|':
		// vertical pipe
		return c.Peek(N), c.Peek(S), nil
	case '-':
		// horizontal pipe
		return c.Peek(W), c.Peek(E), nil
	case 'L':
		// 90 degree N/E bend
		return c.Peek(N), c.Peek(E), nil
	case 'J':
		// 90 degree N/W bend
		return c.Peek(N), c.Peek(W), nil
	case '7':
		// 90 degree W/S bend
		return c.Peek(W), c.Peek(S), nil
	case 'F':
		// 90 degree S/E bend
		return c.Peek(S), c.Peek(E), nil
	case '.':
		// ground, throw error
		return c, c, errors.New("ground detected, no known direction can be supplied")
	case 'S':
		// starting position, this should have been overwritten
		return c, c, errors.New("starting position found, should have been replaced with valid tile")
	default:
		return c, c, errors.New("unknown tile found")
	}
}

func (p PipesMap) identifyTileDirs(c Coord) (int, int, error) {
	// similar to IdentifyTile(), this instead returns the directions the provided space
	// will route to. This can potentially be used to determine if we can squeeze between pipes.
	thisTile := p.m[c.Y][c.X]
	switch thisTile {
	case '|':
		return N, S, nil
	case '-':
		return W, E, nil
	case 'L':
		return N, E, nil
	case 'J':
		return N, W, nil
	case '7':
		return W, S, nil
	case 'F':
		return S, E, nil
	case '.':
		return N, N, errors.New("ground detected")
	case 'S':
		return N, N, errors.New("starting point detected")
	default:
		return N, N, errors.New("unknown tile")
	}
}

func (p PipesMap) IsInBounds(c Coord) bool {
	// Simply does some bounds checking. If we're outside of the pipesmap, return false
	if c.X < 0 || c.Y < 0 {
		// easy win
		return false
	}

	if c.X >= len(p.m[0]) || c.Y >= len(p.m) {
		return false
	}
	return true
}

func (p *PipesMap) replaceStartTile() error {
	// This function will attempt to replace the starting tile with the proper tile that should
	// be under it. It does this by looking at the 4 coorindates around the tile and determining
	// if they are valid in an attempt to be circuitous. Will probably refine this later, but
	// this will assume that at most two of the tiles are valid.
	dTiles := [4]Coord{p.starting.Peek(N), p.starting.Peek(S), p.starting.Peek(W), p.starting.Peek(E)}

	fmt.Printf("Starting pos: %+v\n", p.starting)
	var validCoords []Coord

	for _, v := range dTiles {
		fmt.Printf("Checking: %+v\n", v)
		if p.IsInBounds(v) {
			if aCoord, bCoord, err := p.IdentifyTile(v); err == nil {
				fmt.Printf("aCoord: %+v, bCoord: %+v\n", aCoord, bCoord)
				if aCoord == p.starting || bCoord == p.starting {
					validCoords = append(validCoords, v)
				}
			} else {
				fmt.Printf("Error!: %s\n", err)
			}
		} else {
			fmt.Printf("Out of bounds??? %+v\n", v)
		}
	}
	fmt.Printf("Valid coords: %+v\n", validCoords)

	if len(validCoords) != 2 {
		// who knows what this is
		return errors.New("more than 2 directions detected, will have to figure this out later")
	}

	// Prepare for garbage logic but I'm on a time crunch
	switch validCoords[0] {
	case p.starting.Peek(N):
		switch validCoords[1] {
		case p.starting.Peek(S):
			// North and South valid, must be a vertical pipe
			p.m[p.starting.Y][p.starting.X] = '|'
		case p.starting.Peek(W):
			// N and W, must be J
			p.m[p.starting.Y][p.starting.X] = 'J'
		case p.starting.Peek(E):
			// N and E, must be L
			p.m[p.starting.Y][p.starting.X] = 'L'
		}
	case p.starting.Peek(E):
		switch validCoords[1] {
		case p.starting.Peek(W):
			// must be a horizontal pipe, E/W
			p.m[p.starting.Y][p.starting.X] = '-'
		case p.starting.Peek(N):
			// must be L, E and N
			p.m[p.starting.Y][p.starting.X] = 'L'
		case p.starting.Peek(S):
			// E and S, must be F
			p.m[p.starting.Y][p.starting.X] = 'F'
		}
	case p.starting.Peek(S):
		switch validCoords[1] {
		case p.starting.Peek(W):
			// S and W, must be 7
			p.m[p.starting.Y][p.starting.X] = '7'
		case p.starting.Peek(N):
			// N and S, must be |
			p.m[p.starting.Y][p.starting.X] = '|'
		case p.starting.Peek(E):
			// S and E, must be F
			p.m[p.starting.Y][p.starting.X] = 'F'
		}
	case p.starting.Peek(W):
		switch validCoords[1] {
		case p.starting.Peek(N):
			// N and W, must be J
			p.m[p.starting.Y][p.starting.X] = 'J'
		case p.starting.Peek(E):
			// W and E, must be -
			p.m[p.starting.Y][p.starting.X] = '-'
		case p.starting.Peek(S):
			// W and S, must be 7
			p.m[p.starting.Y][p.starting.X] = '7'
		}
	}
	return nil
}

func (p PipesMap) IsEdge(c Coord) bool {
	// simply returns true or false if we are at the edge
	// of a map, or beyond
	if c.X >= len(p.m[0])-1 || c.X <= 0 ||
		c.Y >= len(p.m)-1 || c.Y <= 0 {
		return true
	}
	return false
}

func (p *PipesMap) PrintMarkedSpaces(c []Coord, bc Circuit) {
	// This will print the map complete with marked spaces
	// and pipes
	for y := 0; y < len(p.m); y++ {
		for x := 0; x < len(p.m[y]); x++ {
			thisCoord := Coord{X: x, Y: y}
			var printable rune
			if bc.Contains(thisCoord) {
				printable = '#'
			} else {
				printable = p.m[y][x]
			}
			for _, v := range c {
				if thisCoord == v {
					printable = 'X'
				}
			}
			fmt.Printf("%c", printable)
		}
		fmt.Printf("\n")
	}
}
