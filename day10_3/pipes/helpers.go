package pipes

import "errors"

const (
	N = iota
	S
	E
	W
)

const (
	NorthSouth = iota // |
	WestEast          // -
	NorthEast         // L
	NorthWest         // J
	SouthWest         // 7
	SouthEast         // F
	Ground            // .
	Starting          // S
)

type Coord struct {
	X, Y int
}

func GetCoordDirection(c Coord, dir int) Coord {
	// Returns a coordinate given a direction
	switch dir {
	case N:
		return Coord{
			X: c.X,
			Y: c.Y - 1,
		}
	case S:
		return Coord{
			X: c.X,
			Y: c.Y + 1,
		}
	case E:
		return Coord{
			X: c.X + 1,
			Y: c.Y,
		}
	case W:
		return Coord{
			X: c.X - 1,
			Y: c.Y,
		}
	default:
		return c
	}
}

func GetTileType(t rune) int {
	// Given the rune, returns the tile type.
	switch t {
	case '|':
		return NorthSouth
	case '-':
		return WestEast
	case 'L':
		return NorthEast
	case 'J':
		return NorthWest
	case '7':
		return SouthWest
	case 'F':
		return SouthEast
	case '.':
		return Ground
	case 'S':
		return Starting
	default:
		return -1
	}
}

func GetNeighborDirections(dir int) (int, int) {
	switch dir {
	case NorthSouth:
		return N, S
	case WestEast:
		return W, E
	case NorthEast:
		return N, E
	case NorthWest:
		return N, W
	case SouthWest:
		return S, W
	case SouthEast:
		return S, E
	default:
		return -1, -1
	}
}

/* *************************************************
 * The below section will involve the usage of a set
 * and how to create a visited set
 * *************************************************/

var exists = struct{}{}

type Set struct {
	m map[Coord]struct{}
}

func NewSet() *Set {
	s := &Set{}
	s.m = make(map[Coord]struct{})
	return s
}

func (s *Set) Add(value Coord) {
	s.m[value] = exists
}

func (s *Set) Remove(value Coord) {
	delete(s.m, value)
}

func (s *Set) Contains(value Coord) bool {
	_, c := s.m[value]
	return c
}

func (s Set) Len() int {
	// returns the length
	return len(s.m)
}

/* *************************************************
 * The below section will involve the usage of a queue
 * and how to create a visited set
 * *************************************************/

type Queue struct {
	q []Coord
}

func (q *Queue) Push(c Coord) {
	// adds a coordinate to the queue
	q.q = append(q.q, c)
}

func (q *Queue) Pop() (Coord, error) {
	// pops the top value off the queue
	if len(q.q) == 0 {
		return Coord{}, errors.New("empty queue")
	}
	ne := q.q[0]
	q.q = q.q[1:]
	return ne, nil
}

// Because I couldn't find a better place for this function

func IsOdd(n int) bool {
	return n%2 == 1
}
