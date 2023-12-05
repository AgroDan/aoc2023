package seed

import "fmt"

// type Seed struct {
// 	// This seed will be the first thing parsed.
// 	seed []int
// }

// This is the ratio of each X to Y instruction
type Ratio struct {
	Dst    int
	Src    int
	Length int
}

type Instruction struct {
	From, To string
	R        []Ratio
}

type Almanac struct {
	// This is putting it all together.
	Seeds []int
	Inst  []Instruction
}

func (a Almanac) PrintAlmanac() {
	// Prints the almanac in a way that only it can
	fmt.Printf("Seeds: ")
	for _, v := range a.Seeds {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n\n")

	// Now print the targets:
	for _, v := range a.Inst {
		fmt.Printf("%s-to-%s map:\n", v.From, v.To)
		for _, k := range v.R {
			fmt.Printf("%d %d %d\n", k.Dst, k.Src, k.Length)
		}
		fmt.Printf("\n")
	}
}

func (a Almanac) FindDestination(point int) int {
	// this function will determine the destination number, given
	// an arbitrary number to start.
	transitive := point // so we don't change the point
	// fmt.Printf("Start -> %d", point)
	for _, v := range a.Inst {
		for _, k := range v.R {
			// now we go through all the ratios
			srcLow := k.Src
			srcHigh := k.Src + k.Length

			// Are we inside this ratio?
			if transitive >= srcLow && transitive <= srcHigh {
				diff := transitive - srcLow
				transitive = diff + k.Dst
				// fmt.Printf(" -> %d (l: %d, h: %d)", transitive, srcLow, srcHigh)
				break
			}
		}
	}
	// fmt.Printf("\n")
	return transitive
}

func (a Almanac) GenerateSeedRanges() [][]int {
	// I'm almost positive this will take forever to calculate but
	// lets just do this I guess
	var retVal [][]int
	for i := 0; i < len(a.Seeds); i += 2 {
		var v []int
		v = append(v, a.Seeds[i])
		v = append(v, a.Seeds[i+1])
		retVal = append(retVal, v)
	}
	return retVal
}
