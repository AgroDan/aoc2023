package rocks

import (
	"fmt"

	"github.com/mitchellh/hashstructure/v2"
)

// This file will deal with the load equations

func (r RockMap) getLoadBearingScore(c Coord) int {
	// This assumes the coordinate passed is a round rock.
	return r.MaxRows - c.Y
}

func (r RockMap) GetTotalLoad() int {
	// This will loop through each round rock present in the
	// map given and return the load bearing score
	var total int = 0
	for k := range r.roundRocks {
		total += r.getLoadBearingScore(k)
	}
	return total
}

// HERE BE METRICS

type Breadcrumb struct {
	b map[uint64]int
}

func NewBreadCrumb() Breadcrumb {
	nb := Breadcrumb{}
	nb.b = make(map[uint64]int)
	return nb
}

func (b *Breadcrumb) Add(r RockMap) bool {
	// hashes the rockmap and stores it
	// returns true if there is a collision!
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		panic("could not hash")
	}
	b.b[hash]++
	return b.b[hash] > 1
}

type LoadMarker struct {
	totalLoads map[int]int
}

func NewLoadMarker() LoadMarker {
	return LoadMarker{
		totalLoads: make(map[int]int),
	}
}

func (l *LoadMarker) Add(loadValue int) {
	l.totalLoads[loadValue]++
}

func (l LoadMarker) PrintVals() {
	// just dumps all the values in a readable format
	for k, v := range l.totalLoads {
		fmt.Printf("Load value %d has repeated %d times.\n", k, v)
	}
}

// func (l LoadMarker) PrintTopVals(howMany int) {
// 	// Prints the top X load values
// 	keys := make([]int, 0, len(l.totalLoads))
// 	for _, v := range l.totalLoads {
// 		keys = append(keys, v)
// 	}
// 	sort.Ints(keys)
// 	for i := 0; i < howMany; i++ {
// 		for k := range l.totalLoads {
// 			if keys[i] == k {
// 				fmt.Printf("Group %d has repeated %d times.\n", i, l.totalLoads[k])
// 			}
// 		}

// 	}
// 	fmt.Printf("Here: %+v\n", l.totalLoads)
// }

/*
 *
 * Here is a better way to record metrics
 *
 */

type RecordsManager struct {
	cycleHashes []string
	idxValue    []int
}

func (r RecordsManager) CycleHashes() []string {
	return r.cycleHashes
}

func (r RecordsManager) IdxValue() []int {
	return r.idxValue
}

func NewRecordsManager() RecordsManager {
	// sets up a records manager object
	return RecordsManager{}
}

func (r *RecordsManager) Add(rm RockMap) {
	// Adds the current rockmap to the struct
	// hash, err := hashstructure.Hash(rm, hashstructure.FormatV2, nil)
	// if err != nil {
	// 	panic("could not get hash")
	// }
	// r.cycleHashes = append(r.cycleHashes, hash)
	hash, err := rm.Hash()
	if err != nil {
		panic(err)
	}
	r.cycleHashes = append(r.cycleHashes, hash)
	total := rm.GetTotalLoad()
	r.idxValue = append(r.idxValue, total)
}

func (r RecordsManager) Get(idx int) (int, string) {
	return r.idxValue[idx], r.cycleHashes[idx]
}

func (r RecordsManager) FindNextMatch(idx int) int {
	// Finds the next index of the next match of the
	// provided index
	for i := idx + 1; i < len(r.idxValue); i++ {
		if r.idxValue[idx] == r.idxValue[i] &&
			r.cycleHashes[idx] == r.cycleHashes[i] {
			return i
		}
	}
	return -1 // didn't find any
}
