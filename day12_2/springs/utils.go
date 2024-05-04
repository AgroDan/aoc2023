package springs

type SpringSet struct {
	Springs []rune
	Inst    []int
}

func sumArray(a []int) int {
	// Sums an array of ints
	var total = 0

	for _, v := range a {
		total += v
	}
	return total
}

func contiguousBroken(s []rune, howMany int) bool {
	// this function, given a rune array, will start at
	// the beginning and determine if any '.''s exist in
	// the array up to the howMany mark. if ANY exist,
	// return false. otherwise return true
	// also if howMany is greater than the size of the
	// rune array, return false
	if len(s) < howMany {
		return false
	}
	for i := 0; i < howMany; i++ {
		if s[i] == '.' {
			return false
		}
	}
	return true
}

func GetCombinations(springs []rune, nums []int) int {
	// tally counter
	var total int = 0
	// No more springs, and no more numbers? Validate!
	if len(springs) == 0 {
		if len(nums) == 0 {
			return 1
		}
		// Otherwise this is invalid, ret 0
		return 0
	}

	// no more numbers, but any broken springs? If so invalidate!
	if len(nums) == 0 {
		for _, v := range springs {
			if v == '#' {
				return 0 // no possibility
			}
		}
		return 1
	}

	// Check if valid combination of potential broken springs
	// plus amount of springs left
	if len(springs) < (sumArray(nums) + (len(nums) - 1)) {
		// can't have any more possibilities
		return 0
	}

	// Now let's perform some checking of the actual springs
	if springs[0] == '.' || springs[0] == '?' {
		total += GetCombinations(springs[1:], nums)
	}

	// I'm going to follow 0xdf's logic because wtf
	if (springs[0] == '?' || springs[0] == '#') &&
		contiguousBroken(springs[:nums[0]], nums[0]) &&
		(len(springs) == nums[0] || (springs[nums[0]] == '.' || springs[nums[0]] == '?')) {
		// Do some out of bounds checking
		if nums[0]+1 > len(springs) {
			total += GetCombinations([]rune{}, nums[1:])
		} else {
			total += GetCombinations(springs[nums[0]+1:], nums[1:])
		}
	}

	return total
}

func FoldSpringSet(s SpringSet, foldAmt int) SpringSet {
	// This will multiply the amount of springs and adjustments
	// by the folded amount.
	ns := SpringSet{}

	for i := 0; i < foldAmt; i++ {
		if i > 0 {
			ns.Springs = append(ns.Springs, '?')
		}
		ns.Springs = append(ns.Springs, s.Springs...)
		ns.Inst = append(ns.Inst, s.Inst...)
	}
	return ns
}
