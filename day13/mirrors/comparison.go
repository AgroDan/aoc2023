package mirrors

func (m MirrorSet) CompareRows(row1, row2 int) bool {
	if row1 >= len(m.M) || row2 >= len(m.M) {
		return false
	}

	if row1 < 0 || row2 < 0 {
		return false
	}

	for i := 0; i < len(m.M[row1]); i++ {
		if m.M[row1][i] != m.M[row2][i] {
			return false
		}
	}

	return true
}

func (m MirrorSet) CompareDiffRows(row1, row2 int) int {
	// This function, similar to the above, compares each
	// row and determines how many entries in the row are
	// DIFFERENT. This is useful for the challenge since
	// technically we care only about two rows that are
	// off by exactly 1.
	if row1 >= len(m.M) || row2 >= len(m.M) {
		return -1
	}

	if row1 < 0 || row2 < 0 {
		return -1
	}

	var diff int = 0
	for i := 0; i < len(m.M[row1]); i++ {
		if m.M[row1][i] != m.M[row2][i] {
			diff++
		}
	}
	return diff
}

func (m MirrorSet) CompareCols(col1, col2 int) bool {
	if col1 >= len(m.M[0]) || col2 >= len(m.M[0]) {
		return false
	}

	if col1 < 0 || col2 < 0 {
		return false
	}

	for i := 0; i < len(m.M); i++ {
		if m.M[i][col1] != m.M[i][col2] {
			return false
		}
	}

	return true
}

func (m MirrorSet) CompareDiffCols(col1, col2 int) int {
	// same as above but diff
	if col1 >= len(m.M[0]) || col2 >= len(m.M[0]) {
		return -1
	}

	if col1 < 0 || col2 < 0 {
		return -1
	}

	var diff int = 0
	for i := 0; i < len(m.M); i++ {
		if m.M[i][col1] != m.M[i][col2] {
			diff++
		}
	}
	return diff
}

func (m MirrorSet) IsColRefraction(col int) bool {
	// This function, given the _left-most_ column, will
	// branch out recursively from each supposed
	// refraction point and just confirm whether or
	// not the given point is an official refraction
	// point.
	leftCol := col
	rightCol := col + 1
	// if we're starting here, then this doesn't match
	if leftCol >= len(m.M[0]) || rightCol >= len(m.M[0]) {
		return false
	}
	if leftCol < 0 || rightCol < 0 {
		return false
	}

	for {
		if leftCol >= len(m.M[0]) || rightCol >= len(m.M[0]) {
			break
		}
		if leftCol < 0 || rightCol < 0 {
			break
		}
		if !m.CompareCols(leftCol, rightCol) {
			return false
		}
		leftCol--
		rightCol++
	}
	return true
}

func (m MirrorSet) IsColRefractionOffByOne(col int) bool {
	// this function is similar, but it will attempt to determine
	// the above with the puzzle requirements, where there is one
	// AND ONLY ONE difference. If there is any more, then it can't
	// be a legit difference.
	leftCol := col
	rightCol := col + 1

	// sanity checking
	if leftCol >= len(m.M[0]) || rightCol >= len(m.M[0]) {
		return false
	}
	if leftCol < 0 || rightCol < 0 {
		return false
	}
	var diff int = 0
	for {
		if diff < 0 || diff > 1 {
			return false
		}
		if leftCol >= len(m.M[0]) || rightCol >= len(m.M[0]) {
			break
		}
		if leftCol < 0 || rightCol < 0 {
			break
		}
		diff += m.CompareDiffCols(leftCol, rightCol)

		leftCol--
		rightCol++
	}

	// Now we only care if we're off by one
	return diff == 1
}

func (m MirrorSet) IsRowRefraction(row int) bool {
	// This is the same as the above, but for rows. This takes
	// the TOP-most row.
	upRow := row
	downRow := row + 1
	if upRow >= len(m.M) || downRow >= len(m.M) {
		return false
	}
	if upRow < 0 || downRow < 0 {
		return false
	}

	for {
		if upRow >= len(m.M) || downRow >= len(m.M) {
			break
		}
		if upRow < 0 || downRow < 0 {
			break
		}
		if !m.CompareRows(upRow, downRow) {
			return false
		}
		upRow--
		downRow++
	}
	return true
}

func (m MirrorSet) IsRowRefractionOffByOne(row int) bool {
	// This is the same as the above, but for rows. This takes
	// the TOP-most row and only counts if we're off by one
	upRow := row
	downRow := row + 1
	if upRow >= len(m.M) || downRow >= len(m.M) {
		return false
	}
	if upRow < 0 || downRow < 0 {
		return false
	}

	var diff int = 0
	for {
		if diff < 0 || diff > 1 {
			return false
		}
		if upRow >= len(m.M) || downRow >= len(m.M) {
			break
		}
		if upRow < 0 || downRow < 0 {
			break
		}
		diff += m.CompareDiffRows(upRow, downRow)
		upRow--
		downRow++
	}
	return diff == 1
}
