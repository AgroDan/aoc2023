package pipes

import (
	"errors"
)

/*
 * This file has to do with the _cicuitousness_ (if that's a word) of the pipe
 * in question. The pipe can be a full loop, so we have to map every single inch
 * of the pipe to determine whether or not we hit a wall.
 */

// the Circuit struct datatype will serve as a database of coordinates attached
// to the pipe. If we are travelling the pipe, we want to make breadcrumbs to
// determine the coordinates of everywhere in the pipe.
var exists = struct{}{}

type Circuit struct {
	breadcrumb map[Coord]struct{}
}

func NewCircuit() Circuit {
	// constructs the circuit
	return Circuit{
		breadcrumb: make(map[Coord]struct{}),
	}
}

func (c *Circuit) Add(co Coord) {
	c.breadcrumb[co] = exists
}

func (c *Circuit) Delete(co Coord) {
	delete(c.breadcrumb, co)
}

func (c *Circuit) Contains(co Coord) bool {
	_, b := c.breadcrumb[co]
	return b
}

func (c Circuit) GetListOfPipes() []Coord {
	// This, for my own convenience, simply
	// returns a slice of every tile that a
	// known pipe within the circuit is on.
	var retval []Coord
	for k := range c.breadcrumb {
		retval = append(retval, k)
	}
	return retval
}

// Queue for DFS
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
