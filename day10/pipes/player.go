package pipes

import (
	"errors"
	"fmt"
)

/*
 * This puts everything together. This shows the location of the player by appropriating
 * the map and traversing where appropriate.
 */

type Player struct {
	current, prev Coord
	steps         int
	m             *PipesMap
	bc            Circuit
}

func NewPlayer(p *PipesMap) Player {
	nc := NewCircuit()
	nc.Add(p.starting)
	return Player{
		current: p.starting,
		prev:    p.starting,
		steps:   0,
		m:       p,
		bc:      nc,
	}
}

func (p Player) GetBreadCrumbs() Circuit {
	// Just returns the breadcrumb object because
	// I'm just too lazy to refactor my code. oh
	// go pound sand
	return p.bc
}

func (p Player) GetPossibleDirections() (Coord, Coord) {
	// This will return two coordinates of where the player
	// can actually go.
	if aCoord, bCoord, err := p.m.IdentifyTile(p.current); err == nil {
		return aCoord, bCoord
	} else {
		panic(fmt.Sprintf("Hit error: %s", err))
	}
}

func (p *Player) stepOnce() error {
	// NOTE: This is not meant to be called individually. This doesn't decide on a direction.
	// It goes in the direction that we HAVEN'T been in. Throws an error if we can't determine
	// that (ie, we're in the starting position and haven't decided on a direction)
	if p.steps <= 0 {
		return errors.New("this function needs a direction, not to be manually called")
	}
	aCoord, bCoord := p.GetPossibleDirections()
	if aCoord == p.prev {
		// we're going to bCoord!
		// fmt.Printf("Moving to %+v\n", bCoord)
		p.prev = p.current
		p.current = bCoord
		p.bc.Add(p.current)
		p.steps++

	} else if bCoord == p.prev {
		// fmt.Printf("Moving to %+v\n", aCoord)
		// else we're giong to aCoord
		p.prev = p.current
		p.current = aCoord
		p.bc.Add(p.current)
		p.steps++
	} else {
		return errors.New("neither aCoord or bCoord were last visited")
	}
	return nil
}

func (p *Player) StartWalking() {
	// fmt.Printf("Started at X: %d, Y: %d\n", p.current.X, p.current.Y)
	// fmt.Printf("Player: %+v\n", p)
	// fmt.Printf("Starting coords: X: %d, Y: %d\n", p.m.starting.X, p.m.starting.Y)
	// This function will loop until we reach the beginning again. This should be circuitous after all.
	if p.steps <= 0 {
		// we should start at the first direction, doesn't really matter.
		firstStep, _ := p.GetPossibleDirections()
		// fmt.Printf("First step: %+v, second: %+v\n", firstStep, secondStep)
		p.prev = p.current
		p.current = firstStep
		p.bc.Add(p.current)
		p.steps++
		// fmt.Printf("Moved to X: %d, Y: %d\n", p.current.X, p.current.Y)
	}

	// fmt.Printf("Comparing %+v to %+v\n", p.current, p.m.starting)
	// Now let's start walking
	for {
		if p.current == p.m.starting {
			// we looped around. Yay
			// fmt.Printf("Looped around\n")
			break
		}
		if err := p.stepOnce(); err != nil {
			fmt.Printf("Something terrible happened: ")
			fmt.Println(err)
			fmt.Printf("Errored out\n")
			break
		}
	}
}

func (p Player) GetSteps() int {
	return p.steps
}
