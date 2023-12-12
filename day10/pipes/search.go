package pipes

import "fmt"

/*
 * I started writing a lot of this in the main package, and since I was
 * referring to a lot of un-exported variables I figured I might as well
 * put it here in the pipes package. This is getting messy but I'm getting
 * tired now.
 */

func GetNotInCircuit(pMap PipesMap, bc Circuit) []Coord {
	// This loops through the entire map and just
	// gets everything that's not in the circuit.
	// It accepts the full map and the breadcrumb
	// list that the Player object should have after
	// completing a full circuit.
	var retval []Coord
	for y := 0; y < len(pMap.m); y++ {
		for x := 0; x < len(pMap.m[y]); x++ {
			thisCoord := Coord{X: x, Y: y}
			// First of all, if we're at an edge, bomb out.
			if pMap.IsEdge(thisCoord) {
				// fmt.Printf("%+v is an edge, yes\n", thisCoord)
				continue
			}

			// Now if we're not at an edge, are we in the
			// circuitous pipe?
			if bc.Contains(thisCoord) {
				continue
			}

			// Otherwise, we're in a potential spot for the nest.
			retval = append(retval, thisCoord)
		}
	}
	// fmt.Printf("Here's what we did: %+v\n", retval)
	// fmt.Printf("How many: %d\n", len(retval))
	return retval
}

/*
 * DEPTH FIRST SEARCH, ONLY 10 DAYS IN BEFORE I NEEDED IT LETS GO
 */

func (p PipesMap) DetectSpaces(startSpace Coord, c Circuit) []Coord {
	// Given a point to check, it will spread out and return a list of
	// coordinates that are enclosed within the circuit. If it hits the
	// map edge, then it is NOT enclosed in the circuit so therefore it
	// is NOT enclosed by the loop.

	// first, is the point an edge? If so bomb out immediately ya turkey
	//
	// for your code
	if p.IsEdge(startSpace) {
		return []Coord{}
	}

	retval := make([]Coord, 0)

	q := Queue{}
	q.Push(startSpace)

	// I'll re-use the Circuit object to store visited
	visited := NewCircuit()

	// // starting point has been visited
	// visited.Add(startSpace)

	// anon function that I'll use here
	oppositeDir := func(thisDir int) int {
		switch thisDir {
		case N:
			return S
		case S:
			return N
		case E:
			return W
		case W:
			return E
		default:
			return 5
		}
	}

	// check it out
	for {
		workingCoord, err := q.Pop()
		if err != nil {
			// empty queue, drop out
			break
		}

		// First, have we visited here? If so ignore and go onto the next
		if visited.Contains(workingCoord) {
			continue
		}

		// If we got here, we can continue to work, but make sure
		// we mark the coord as visited
		visited.Add(workingCoord)

		// fmt.Printf("New DFS\n")
		// fmt.Printf("Size of queue: %d\n", len(q.q))

		// first, is this an edge of the map?
		if p.IsEdge(workingCoord) {
			// It's an edge, which means it can't be enclosed. Return zilch
			return []Coord{}
		}

		if !c.Contains(workingCoord) {
			// Just check and see if this is in the pipe. If not, it's possibly
			// a potential nest space so we'll add it to the possible list.
			retval = append(retval, workingCoord)
		} else {
			// If we got here, it's a pipe that's part of the circuit. But
			// can we go around it?
			dir1, dir2, err := p.identifyTileDirs(workingCoord)
			if err != nil {
				// what is this?
				fmt.Println(err)
				break
			}
			if !visited.Contains(workingCoord.Peek(dir1)) {
				q.Push(workingCoord.Peek(dir1))
			}
			if !visited.Contains(workingCoord.Peek(dir2)) {
				q.Push(workingCoord.Peek(dir2))
			}
			// don't need to get other directions here either
			continue

		}

		// Get all the directions
		dTiles := [4]int{N, S, W, E}
		// dTiles := [4]Coord{workingCoord.Peek(N), workingCoord.Peek(S), workingCoord.Peek(W), workingCoord.Peek(E)}
		for _, v := range dTiles {
			// First, have we visited this yet?
			checkCoord := workingCoord.Peek(v) // Look in this direction

			// Let's first check the directions
			if visited.Contains(checkCoord) {
				// been here, next
				continue
			}

			// Second, is this a pipe space within the circuit?
			if c.Contains(checkCoord) {
				// we are a pipe in the circuit, find the directions the pipe goes in
				dir1, dir2, err := p.identifyTileDirs(checkCoord)
				if err != nil {
					fmt.Println(err)
					break
				}

				// if I went north, then i came from the south. If the pipe has a south
				// direction, then i can technically squeeze through it.
				iCameFrom := oppositeDir(v)
				switch iCameFrom {
				case dir1:
					q.Push(checkCoord)
				case dir2:
					q.Push(checkCoord)
				default:
					// brick wall, ignore
					continue
				}
			}

			// Third, is this tile at the edge of the map? If so just bomb out here, the provided coordinate
			// is definitely not enclosed within the circuit
			if p.IsEdge(checkCoord) {
				return []Coord{}
			}

			// Finally, it might be a space. Let's add this to the potential nest spaces and stick it into the queue.
			retval = append(retval, checkCoord)
			q.Push(checkCoord)
		}
	}
	return retval
}
