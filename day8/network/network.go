package network

import "fmt"

/*
 * This is a network map as defined in the day 8 challenge.
 * It will be used by the parser to create a full map. Each
 * node in the network will contain two pointers which point
 * to another node.
 */

type Node struct {
	Name         string
	Left, Right  *Node
	LName, RName string // This is to prevent a chicken-and-egg situation, hard to explain on one line
}

type Direction struct {
	d []rune
}

func (d *Direction) Push(dir rune) {
	// pushes a direction onto the direction stack
	d.d = append(d.d, dir)
}

func (d *Direction) Pop(dir rune) rune {
	// pops a value off the stack
	retval := d.d[len(d.d)-1]
	d.d = d.d[:len(d.d)-1]
	return retval
}

type FullMap struct {
	Map    map[string]*Node
	Dir    Direction
	dirIdx int
}

func (f *FullMap) NextDirection() rune {
	i := f.dirIdx
	if f.dirIdx+1 >= len(f.Dir.d) {
		f.dirIdx = 0
	} else {
		f.dirIdx++
	}
	return f.Dir.d[i]

}

func (f FullMap) PrintMap() {
	// Just confirms we parsed correctly
	fmt.Printf("Directions: ")
	for _, v := range f.Dir.d {
		fmt.Printf("%c ", v)
	}
	fmt.Printf("\n\n")
	fmt.Printf("Maps:\n")
	for _, v := range f.Map {
		fmt.Printf("%s: (%s, %s)\n", v.Name, v.LName, v.RName)
	}
}

func (f *FullMap) DeepCopy() FullMap {
	// As the name implies, this will make a full deep copy of
	// a fullmap object.
	dirCopy := make([]rune, len(f.Dir.d))
	copy(dirCopy, f.Dir.d)
	d := Direction{
		d: dirCopy,
	}

	mapCopy := make(map[string]*Node)
	for k, v := range f.Map {
		temp := Node{
			Name:  v.Name,
			LName: v.LName,
			RName: v.RName,
		}
		mapCopy[k] = &temp
	}

	// Now build the new addresses
	for _, v := range f.Map {
		mapCopy[v.Name].Left = mapCopy[v.LName]
		mapCopy[v.Name].Right = mapCopy[v.RName]
	}

	return FullMap{
		Map: mapCopy,
		Dir: d,
	}
}

// For the ghosts
type Traveller struct {
	Ptr *Node
	FM  FullMap
}
