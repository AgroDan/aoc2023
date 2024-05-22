package rocks

import "fmt"

func NewRockMap(lines []string) RockMap {
	r := RockMap{}
	r.cubeRocks = make(map[Coord]bool)
	r.roundRocks = make(map[Coord]bool)
	var length, width int = 0, 0
	for Yi := 0; Yi < len(lines); Yi++ {
		width = 0
		for Xi := 0; Xi < len(lines[Yi]); Xi++ {
			width++
			thisCoord := Coord{
				X: Xi,
				Y: Yi,
			}
			switch lines[Yi][Xi] {
			case 'O':
				// round rock
				r.roundRocks[thisCoord] = true
			case '#':
				r.cubeRocks[thisCoord] = true
			default:
				continue
			}
		}
		length++
	}
	r.MaxRows = length
	r.MaxCols = width
	return r
}

func (r RockMap) PrintRockMap() {
	// simply prints the rockmap as well as some additional info
	fmt.Printf("Max cols: %d, Max Rows: %d\n", r.MaxCols, r.MaxRows)
	for Yi := 0; Yi < r.MaxRows; Yi++ {
		for Xi := 0; Xi < r.MaxCols; Xi++ {
			thisCoord := Coord{
				X: Xi,
				Y: Yi,
			}
			if r.DoesRoundRockExist(thisCoord) {
				fmt.Printf("O")
				continue
			}
			if r.DoesStaticRockExist(thisCoord) {
				fmt.Printf("#")
				continue
			}
			fmt.Printf(".")
		}
		fmt.Printf("\n")
	}
}
