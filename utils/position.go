package utils

import (
	"math/rand"
)

type position [2]int

// generate random coordinates within the given range
func RandomPosition(maxX, maxY int) position {
	x := 2 + rand.Intn(maxX-2)
	y := 2 + rand.Intn(maxY-2)
	return [2]int{x, y}
}

// check if positions overlap
func PositionOverlap(pos1, pos2 position) bool {
	if pos1[0] == pos2[0] && pos1[1] == pos2[1] {
		return true
	}

	return false
}
