package main

import (
	"fmt"
	"strconv"
)

// The schematic object will contain the entire schematic
type Schematic struct {
	mapping [][]rune
}

// This is just a map position that can be mapped directly
// to the schematic object
type Coord struct {
	X, Y int
}

// The block object will contain a start and end coordinate of
// digits within a schematic.
type Block struct {
	startCoord, endCoord Coord
}

// the Gear object will contain a list of blocks that are actual numbers
type Gear struct {
	gearLoc Coord
	numbers []Block
}

func (b Block) PrintBlock() {
	// just prints the block
	fmt.Printf("Start: %d-%d, End: %d-%d\n", b.startCoord.X, b.startCoord.Y, b.endCoord.X, b.endCoord.Y)
}

func NewSchematic(lines []string) Schematic {
	// this will read in every single line
	// and go over every single character
	// and generate a Schematic, which will
	// be a mapping of each character.
	s := Schematic{}
	for _, line := range lines {
		ycoord := make([]rune, 0)
		for _, char := range line {
			ycoord = append(ycoord, char)
		}
		s.mapping = append(s.mapping, ycoord)
	}
	return s
}

func (s Schematic) PrintSchematic() {
	// This will simply print out the schematic.
	for i, line := range s.mapping {
		fmt.Printf("Line %3d: ", i)
		for _, ch := range line {
			fmt.Printf("%c", ch)
		}
		fmt.Printf("\n")
	}
}

func (s Schematic) GetNumber(b Block) int {
	// This returns the number at the block
	var numStack []int

	for i := b.startCoord.X; i <= b.endCoord.X; i++ {
		num, _ := strconv.Atoi(string(s.mapping[b.startCoord.Y][i]))
		numStack = append(numStack, num)
	}

	// now convert the numStack to an integer
	retval := 0
	for _, v := range numStack {
		retval = (retval * 10) + v
	}
	return retval
}

func (s Schematic) GetAdjacent(b Block) Block {
	// This function will return a block that is technically adjacent (surrounding)
	// the given block. This will ALSO accommodate for edges.

	sCoord := Coord{X: b.startCoord.X - 1, Y: b.startCoord.Y - 1}
	eCoord := Coord{X: b.endCoord.X + 1, Y: b.endCoord.Y + 1}

	// Now let's do some boundary checking
	// top boundary
	if sCoord.Y < 0 {
		sCoord.Y = 0
	}

	// left boundary
	if sCoord.X < 0 {
		sCoord.X = 0
	}

	// bottom boundary
	if eCoord.Y >= len(s.mapping) {
		eCoord.Y = len(s.mapping) - 1
	}

	// right boundary
	if eCoord.X >= len(s.mapping[eCoord.Y]) {
		eCoord.X = len(s.mapping[eCoord.Y]) - 1
	}
	return Block{startCoord: sCoord, endCoord: eCoord}
}

func (s Schematic) IsPartNumber(b Block) bool {
	// This will check around a number block to determine if it is a part number.
	// A part number will be adjacent to a symbol, not a period. So if number or .,
	// then is not adjacent.

	adj := s.GetAdjacent(b)
	for Y := adj.startCoord.Y; Y <= adj.endCoord.Y; Y++ {
		for X := adj.startCoord.X; X <= adj.endCoord.X; X++ {
			thisRune := s.mapping[Y][X]
			if IsRuneNumeric(thisRune) {
				continue
			} else if thisRune == '.' {
				continue
			} else {
				return true
			}
		}
	}
	return false
}

func (s Schematic) IsAdjacentToBlock(srcBlock, checkBlock Block) bool {
	// This, given a srcBlock, will determine if the provided block
	// is adjacent to the checkBlock parameter

	// First, return a block of all things adjacent to it
	adjBlock := s.GetAdjacent(srcBlock)

	// top left
	UL := Coord{X: adjBlock.startCoord.X, Y: adjBlock.startCoord.Y}
	LR := Coord{X: adjBlock.endCoord.X, Y: adjBlock.endCoord.Y}

	// since the checkBlock _should_ not have additional Y coordinates, just
	// loop over the X coord
	for i := checkBlock.startCoord.X; i <= checkBlock.endCoord.X; i++ {
		if checkBlock.startCoord.Y >= UL.Y &&
			checkBlock.startCoord.Y <= LR.Y &&
			i >= UL.X &&
			i <= LR.X {
			return true
		}
	}
	return false
}

func (s Schematic) FindStars() []Block {
	// This function will return a slice of Blocks denoting the stars (gears)
	// in the provided map. The only way to tell if they are gears is through
	// another function, so this only gets the location of the asterisks
	var blockSlice []Block
	for Y := 0; Y < len(s.mapping); Y++ {
		for X := 0; X < len(s.mapping[Y]); X++ {
			// now let's get the asterisks
			if s.mapping[Y][X] == '*' {
				starCoord := Coord{X: X, Y: Y}
				b := Block{startCoord: starCoord, endCoord: starCoord}
				blockSlice = append(blockSlice, b)
			}
		}
	}
	return blockSlice
}

func (s Schematic) FindGears(starList, numList []Block) []Gear {
	// This function will check through the list of potential gears
	// and return a slice of confirmed gears, which will have the
	// associated numbers with them.
	var gearList []Gear
	for _, star := range starList {
		var blockList []Block
		for _, num := range numList {
			if s.IsAdjacentToBlock(star, num) {
				// we have a number adjancent to a gear
				blockList = append(blockList, num)
			}
		}
		if len(blockList) > 1 {
			// numbers next to a gear, we have a gear
			g := Gear{gearLoc: star.startCoord, numbers: blockList}
			gearList = append(gearList, g)
		}
	}
	return gearList
}

func (s Schematic) GetGearRatio(g Gear) int {
	// This, given a gear, will return the gear ratio (the numbers that
	// are adjacent to the gear multiplied together)
	var gearRatio int = 1 // start at 1 because 0 would cancel everything out
	for _, num := range g.numbers {
		gearRatio *= s.GetNumber(num)
	}
	return gearRatio
}

func FindNumbersOnY(Y int, X []rune) []Block {
	// This will return a slice of Blocks on a line with the coordinates of each
	var numbers []int
	for i := 0; i < len(X); i++ {
		if IsRuneNumeric(X[i]) {
			numbers = append(numbers, i)
		}
	}

	blocks := []Block{}
	// now we have a slice of placements where there is a number. We can now loop
	// until we have all the numbers
	var start, end, idx int = 0, 0, 0
	// fmt.Printf("%+v\n", numbers)
	for {
		// fmt.Printf("IDX: %d\n", idx)
		start, end, idx = genblock(idx, numbers)
		if idx == -1 {
			break
		}
		StartCoord := Coord{X: start, Y: Y}
		EndCoord := Coord{X: end, Y: Y}
		b := Block{startCoord: StartCoord, endCoord: EndCoord}
		blocks = append(blocks, b)
	}
	return blocks
}

func genblock(idx int, intList []int) (start, end, newidx int) {
	// this takes an array of integers that are locations of _known numbers_
	// and returns a block object of the beginning and end of each number location
	// returns the place where it left off, and idx is where to start looking

	// first, if we're at the end, bomb out with impossible numbers:
	if idx >= len(intList) {
		return -1, -1, -1
	}

	// otherwise, initialize
	start = intList[idx]
	for i := idx; i < len(intList); i++ {
		// bomb out if end of list
		if i+1 >= len(intList) {
			end = intList[i]
			newidx = i + 1
			break
		}

		// now peek
		if intList[i]+1 < intList[i+1] {
			end = intList[i]
			newidx = i + 1
			break
		}
	}
	return
}

func IsRuneNumeric(r rune) bool {
	_, err := strconv.Atoi(string(r))
	if err != nil {
		return false
	}
	return true
}
