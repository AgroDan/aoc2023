package cosmos

import (
	"errors"
	"math"
)

/*
 * This will perform the calculations necessary on all the
 * galaxies etc
 */

type Pairs struct {
	// This struct will only accept unique pairs
	// if you use the AddPair() function.
	pairSet map[Coord]map[Coord]int
}

func newPairs() Pairs {
	return Pairs{
		pairSet: make(map[Coord]map[Coord]int),
	}
}

func (p *Pairs) AddPair(from, to Coord) error {
	// First, check to see if this exists already
	if _, exists := p.pairSet[from][to]; exists {
		return errors.New("pair already added")
	}
	if _, exists := p.pairSet[to][from]; exists {
		return errors.New("pair already added")
	}

	// otherwise, add the pair with the distance
	if _, exists := p.pairSet[from]; !exists {
		p.pairSet[from] = make(map[Coord]int)
	}
	p.pairSet[from][to] = GetDistance(from, to)
	return nil
}

func GetDistance(from, to Coord) int {
	// This will return a distance
	// between two coordinates.
	yDistance := math.Abs(float64(from.Y - to.Y))
	xDistance := math.Abs(float64(from.X - to.X))
	return int(yDistance + xDistance)
}

func GeneratePairs(u Universe) Pairs {
	p := newPairs()
	for _, v := range u.Galaxies {
		for i := 0; i < len(u.Galaxies); i++ {
			if v == u.Galaxies[i] {
				continue
			}
			_ = p.AddPair(v, u.Galaxies[i])
		}
	}
	return p
}

func (p Pairs) GetTotal() int {
	// Gets the total amount of distance points for each pair.
	total := 0
	for _, v := range p.pairSet {
		for _, v2 := range v {
			total += v2
		}
	}
	return total
}
