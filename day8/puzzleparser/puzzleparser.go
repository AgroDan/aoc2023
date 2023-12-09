package puzzleparser

import (
	"day8/network"
	"strings"
)

/*
 * Changing the name to puzzleparser because I think there's a
 * parser library already and I don't want to confuse it with
 * anything.
 */

func Parse(lines []string) network.FullMap {
	// First, let's get the first line of directions
	dirs := strings.Trim(lines[0], " ")

	var dirList []rune
	for _, v := range dirs {
		dirList = append(dirList, v)
	}

	var d network.Direction
	for i := 0; i < len(dirList); i++ {
		d.Push(dirList[i])
	}

	nodeHash := make(map[string]*network.Node)

	// Now lets loop through the nodes
	for _, v := range lines[2:] {
		thisNode := network.Node{}
		// First get rid of spaces
		thisLine := strings.ReplaceAll(v, " ", "")
		splitName := strings.Split(thisLine, "=")
		thisName := splitName[0]
		noParens := strings.TrimFunc(splitName[1], func(r rune) bool {
			switch r {
			case '(':
				return true
			case ')':
				return true
			default:
				return false
			}
		})
		pDirs := strings.Split(noParens, ",")

		thisNode.Name, thisNode.LName, thisNode.RName = thisName, pDirs[0], pDirs[1]

		if _, ok := nodeHash[thisName]; !ok {
			// only add stuff if there isn't a hash already
			nodeHash[thisName] = &thisNode
		}
	}

	// Now that we listed all the nodes, let's mark them all.
	for _, v := range nodeHash {
		v.Left = nodeHash[v.LName]
		v.Right = nodeHash[v.RName]
	}

	return network.FullMap{
		Map: nodeHash,
		Dir: d,
	}
}
