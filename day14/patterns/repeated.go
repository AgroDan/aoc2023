package patterns

// type RecordsManager struct {
// 	cycleHashes []string
// 	idxValue    []int
// }

type Record interface {
	CycleHashes() []string
	IdxValue() []int
}

func compareArrays(a, b []uint64) bool {
	// This function will take two arrays and just compare them.
	// if they are different in any way, return false
	// also these arrays had better be the same size! I won't
	// check them in the interest of speed but just make sure
	// they're the same size.
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func compareArraysV2(a, b []string, c, d []int) bool {
	// same as above but for next Variation, also checks load bearing vals
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] && c[i] != d[i] {
			return false
		}
	}
	return true
}

func FindPattern(arr []uint64, windowSize int) []int {
	// This will, given the array to search for and the window size
	// to search with, will return a list of times we see a repeated pattern.
	var total []int
	if windowSize > len(arr)/2 {
		return total // can't compare if the window size is bigger than half the
		// size of the array
	}

	// set up for a finite set
	exists := struct{}{}
	mySet := make(map[int]struct{})

	for i := 0; i < len(arr)-windowSize; i++ {
		thisWindow := arr[i : i+windowSize]
		for j := i + windowSize; j < len(arr)-windowSize; j++ {
			compareWindow := arr[j : j+windowSize]
			if compareArrays(thisWindow, compareWindow) {
				mySet[j] = exists
				j += (windowSize - 1) // for efficiency's sake
			}
		}
	}
	for k := range mySet {
		total = append(total, k)
	}
	return total
}

func FindPatternV2(r Record, windowSize, startingIdx int) []int {
	// This will, given the array to search for and the window size
	// to search with, will return a list of times we see a repeated pattern
	// that includes the starting position+windowSize.
	var total []int
	if windowSize > len(r.CycleHashes())/2 {
		return total // can't compare if the window size is bigger than half the
		// size of the array
	}

	if startingIdx >= len(r.CycleHashes()) {
		return total
	}

	// // set up for a finite set
	// exists := struct{}{}
	// mySet := make(map[int]struct{})

	compareHashWindow := r.CycleHashes()[startingIdx : startingIdx+windowSize]
	compareIdxWindow := r.IdxValue()[startingIdx : startingIdx+windowSize]
	for i := startingIdx + windowSize; i < len(r.CycleHashes())-windowSize; i++ {
		thisCycleWindow := r.CycleHashes()[i : i+windowSize]
		thisIdxWindow := r.IdxValue()[i : i+windowSize]

		if compareArraysV2(compareHashWindow, thisCycleWindow, compareIdxWindow, thisIdxWindow) {
			total = append(total, i)
			i += (windowSize - 1)
		}
	}
	return total
}

// func FindOnePattern(arr []uint64, windowSize, idx int) []int {
// 	// given a starting index AND the window size, it will check
// 	// the rest of the array for exactly that pattern. It will
// 	// return every instace of the array that it finds it.
// 	compareArray := arr[idx:idx+windowSize]
// 	for i := idx+windowSize; i < len(arr)-windowSize; i++ {
// 		if compareArrays(comp)
// 	}
// }
