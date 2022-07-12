/*
	Artificial distinction bewtween boolean and interger rolls, any value
	superior to zero is considered true
*/

package dicelib

import "fmt"

func binarize(p PSet) PSet {
	return opDiceSet(p, PSet{1: 1.0}, func(i1, i2 int) int { return Bool2Int(i1 >= i2) })
}

// If Then
type ifThen struct {
	node, bnode Node
}

func (it ifThen) calculate() PSet {
	return MulDiceSet(binarize(it.bnode.calculate()), it.node.calculate())
}

func (it ifThen) roll() (int, string) {
	val, expr := it.bnode.roll()
	if val == 0 {
		return 0, "if " + expr + " failed "
	} else {
		val, expr2 := it.node.roll()
		return val, "if " + expr + " then " + expr2
	}
}

func (it ifThen) toString() string {
	return "if " + it.bnode.toString() + " then " + it.node.toString()
}

// Superior, inferior, equals unequals
type comparison struct {
	left, right Node
	operation   func(int, int) bool
	char        string
}

func (cp comparison) calculate() PSet {
	lSet := cp.left.calculate()
	rSet := cp.right.calculate()
	return opDiceSet(lSet, rSet, func(i1, i2 int) int { return Bool2Int(cp.operation(i1, i2)) })
}

func (cp comparison) roll() (int, string) {
	leftVal, leftExpr := cp.left.roll()
	rightVal, rightExpr := cp.left.roll()
	return Bool2Int(cp.operation(leftVal, rightVal)), leftExpr + cp.char + rightExpr
}

func (cp comparison) toString() string {
	return cp.left.toString() + cp.char + cp.right.toString()
}

// belongs to a set
type in struct {
	Node
	array []int
}

func (in in) calculate() PSet {
	return opDiceSet(in.calculate(), PSet{1: 1.0},
		func(i1, _ int) int { return Bool2Int(Contains(i1, in.array)) },
	)
}

func (in in) roll() (bool, string) {
	val, expr := in.Node.roll()
	expr += " in"
	for _, val := range in.array {
		expr += " " + fmt.Sprint(val)
	}
	return Contains(val, in.array), expr
}

func (in in) toString() string {
	nums := ""
	for _, val := range in.array {
		nums += " " + fmt.Sprint(val)
	}
	return in.Node.toString() + "in " + nums
}

// and or
type binaryBool struct {
	left, right Node
	operation   func(int, int) int
	char        string
}

func (bb binaryBool) calculate() PSet {
	left, right := binarize(bb.left.calculate()), binarize(bb.right.calculate())
	return opDiceSet(left, right, bb.operation)
}

func (bb binaryBool) roll() (int, string) {
	val, expr := bb.left.roll()
	val2, expr2 := bb.right.roll()
	val, val2 = Bool2Int(val > 1), Bool2Int(val2 > 1)
	return bb.operation(val, val2), expr + bb.char + expr2
}

func (bb binaryBool) toString() string {
	expr := bb.left.toString()
	expr2 := bb.right.toString()
	return expr + bb.char + expr2
}
