package dicelib

import (
	"fmt"
	"math/rand"
)

type Node interface {
	calculate() PSet
	roll() (int, string)
	toString() string
}

// Die and number(d1)
type die struct {
	faces, die int
	PSet
}

func (d die) calculate() PSet {
	return d.PSet
}

func (d die) roll() (int, string) {
	randProb := rand.Float64()
	var val int
	var prob float64
	for val, prob = range d.PSet {
		randProb -= prob
		if randProb > 0 {
			return val + 1, fmt.Sprint(val + 1)
		}
	}
	// Can tecnically happen due to floating point error
	return val + 1, fmt.Sprint(val + 1)
}

// Adition, subtraction, division, multiplication
type binary struct {
	left, right Node
	char        string
	operation   func(int, int) int
}

func (bn binary) calculate() PSet {
	return opDiceSet(bn.left.calculate(), bn.right.calculate(), bn.operation)
}

func (bn binary) roll() (int, string) {
	result, expr := bn.left.roll()
	result2, expr2 := bn.right.roll()
	return bn.operation(result, result2), expr + bn.char + expr2
}

func (bn binary) toString() string {
	return bn.left.toString() + bn.char + bn.right.toString()
}

// Parenthesis
type paren struct {
	Node
}

func (pr paren) calculate() PSet {
	return pr.Node.calculate()
}

func (pr paren) roll() (int, string) {
	val, expr := pr.Node.roll()
	return val, "(" + expr + ")"
}

func (pr paren) toString() string {
	return "(" + pr.Node.toString() + ")"
}

// Advantage, disadvantage
type unFunction struct {
	string
	Node
	function func(int, int) int
}

func (un unFunction) calculate() PSet {
	val1 := un.Node.calculate()
	val2 := un.Node.calculate()
	return opDiceSet(val1, val2, un.function)
}

func (un unFunction) roll() (int, string) {
	val1, expr1 := un.Node.roll()
	val2, expr2 := un.Node.roll()
	return un.function(val1, val2), un.string + "(" + expr1 + "," + expr2 + ")"
}

func (un unFunction) toString() string {
	return un.string + "(" + un.Node.toString() + ")"
}

// Reroll (separated from unFunction for the toString)
type reroll struct {
	Node
	array []int
}

func (re reroll) calculate() PSet {
	ret := PSet{}
	set := re.Node.calculate()
	for val, prob := range set {
		if Contains(val, re.array) {
			for valInner, probInner := range set {
				ret[valInner] += prob * probInner
			}
		} else {
			ret[val] += prob
		}
	}
	return ret
}

func (re reroll) roll() (int, string) {
	val, expr := re.Node.roll()
	if Contains(val, re.array) {
		val, expr2 := re.Node.roll()
		return val, "(" + expr + " -> " + expr2 + ")"
	}
	return val, "(" + expr + ")"
}

func (re reroll) toString() string {
	ints := ""
	for _, val := range re.array {
		ints += " " + string(val)
	}
	return "reroll(" + re.Node.toString() + ", " + ints + ")"
}

// Prefixed minus
type pMinus struct {
	Node
}

func (p pMinus) calculate() PSet {
	return SubDiceSet(PSet{0: 1.0}, p.calculate())
}

func (p pMinus) roll() (int, string) {
	val, expr := p.roll()
	return -val, "-" + expr
}

func (p pMinus) toString() string {
	return "-" + p.Node.toString()
}
