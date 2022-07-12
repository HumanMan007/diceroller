/*
	Functions for generating and manipulating probability sets defined as an
	int->float64 map where the value represents the probability of the key with
	the sum of the values being expected to aproximately equal 1.0
*/

package dicelib

import (
	"fmt"
)

type PSet map[int]float64

// Equivalent to die"d"faces
func GenerateDiceRoll(die, faces int) (PSet, error) {
	if faces < 1 || die < 0 {
		return nil, fmt.Errorf("invalid dice expression %dd%d", faces, die)
	}

	// TODO - optimize with bell curves
	dice := PSet{}
	for i := 0; i < faces; i++ {
		dice[i+1] = 1.0 / float64(faces)
	}

	set := map[int]float64{0: 1.0}
	for i := 0; i < die; i++ {
		AddDiceSet(set, dice)
	}

	return set, nil
}

func AddDiceSet(set1, set2 PSet) PSet {
	return opDiceSet(set1, set2, func(i1, i2 int) int { return i1 + i2 })
}

func SubDiceSet(set1, set2 PSet) PSet {
	return opDiceSet(set1, set2, func(i1, i2 int) int { return i1 - i2 })
}

func MulDiceSet(set1, set2 PSet) PSet {
	return opDiceSet(set1, set2, func(i1, i2 int) int { return i1 * i2 })
}

func DivDiceSet(set1, set2 PSet) PSet {
	return opDiceSet(set1, set2, func(i1, i2 int) int { return i1 / i2 })
}

func opDiceSet(set1, set2 PSet, operation func(int, int) int) PSet {
	newSet := PSet{}
	for i, val1 := range set1 {
		for j, val2 := range set2 {
			newSet[operation(i, j)] += val1 * val2
		}
	}
	return newSet
}
