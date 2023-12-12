package pipes

import "errors"

// ==== Coordinates ==========================//

type Coord struct {
	X, Y int
}

func (c *Coord) Peek(dir int) Coord {
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

// ==== Set =================================//

var exists = struct{}{}

type LocationSet struct {
	s map[Coord]struct{}
}

func NewLocationSet() LocationSet {
	return LocationSet{
		s: make(map[Coord]struct{}),
	}
}

func (l *LocationSet) Add(c Coord) {
	l.s[c] = exists
}

func (l *LocationSet) Delete(c Coord) {
	delete(l.s, c)
}

func (l *LocationSet) Contains(c Coord) bool {
	_, b := l.s[c]
	return b
}

// ==== Queue ===============================//

// Queue for searching
type Queue struct {
	q []Coord
}

func (q *Queue) Push(c Coord) {
	q.q = append(q.q, c)
}

func (q *Queue) Pop() (Coord, error) {
	if len(q.q) == 0 {
		return Coord{}, errors.New("empty queue")
	}
	nc := q.q[0]
	q.q = q.q[1:]
	return nc, nil
}
