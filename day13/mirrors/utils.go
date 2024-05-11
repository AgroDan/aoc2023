package mirrors

import "fmt"

// First we're not going to break up the mirror sets in each group,
// we'll go line-by-line and let the full input parser break up into
// each mirror set.

// good ol' coord object
type Coord struct {
	X, Y int
}

type MirrorSet struct {
	M [][]rune
}

func NewMirrorSet() MirrorSet {
	return MirrorSet{}
}

func (mirror *MirrorSet) AddRow(line string) {
	// This will simply append a line to the mirrorset
	runes := []rune(line)
	mirror.M = append(mirror.M, runes)
}

// Now let's parse the whole file

func GenerateMirrorSets(lines []string) []MirrorSet {
	var ms []MirrorSet
	for i := 0; i < len(lines); i++ {
		var j int = i
		thisMS := NewMirrorSet()
		for {
			if j >= len(lines) || len(lines[j]) < 1 {
				ms = append(ms, thisMS)
				i = j
				break
			}
			thisMS.AddRow(lines[j])
			j++
		}
	}
	return ms
}

func (mirror MirrorSet) PrintMirror() {
	// This just prints the mirror
	for y := 0; y < len(mirror.M); y++ {
		for x := 0; x < len(mirror.M[y]); x++ {
			fmt.Printf("%c", mirror.M[y][x])
		}
		fmt.Printf("\n")
	}
}
