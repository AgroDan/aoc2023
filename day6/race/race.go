package race

import "fmt"

/*
 * This file will define the race object and any functions
 * that work specifically with it.
 */

type Race struct {
	time, distance []int
}

func NewRace(t, d []int) Race {
	// Constructor for Race object
	return Race{
		time:     t,
		distance: d,
	}
}

func (r Race) PrintRace() {
	// just prints the race to confirm we parsed correctly
	fmt.Printf("Time:")
	for _, v := range r.time {
		fmt.Printf("%5d", v)
	}
	fmt.Printf("\nDistance:")
	for _, v := range r.distance {
		fmt.Printf("%5d", v)
	}
	fmt.Printf("\n")
}

func (r Race) PuzzlePartOne() int {
	// this should solve the puzzle for part 1, given the calculations
	var result int = 1
	for i := 0; i < len(r.time); i++ {
		raceTime, record := r.time[i], r.distance[i]
		winTimes := calculateRace(raceTime, record)
		result *= len(winTimes)
	}
	return result
}

func calculateRace(raceTime, record int) []int {
	// given the raceTime and record, calculate each possibility
	// of holding down the button and letting the boat go, returning
	// only possibilities that beat the record time.
	var recordBeaters []int
	for i := 0; i < raceTime; i++ {
		// i is the length of time in milliseconds that the button
		// is held down.
		timeMoving := raceTime - i
		momentum := i * timeMoving
		if momentum > record {
			recordBeaters = append(recordBeaters, momentum)
		}
	}
	return recordBeaters
}
