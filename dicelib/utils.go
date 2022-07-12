package dicelib

func Contains[E comparable](needle E, haysack []E) bool {
	for _, val := range haysack {
		if needle == val {
			return true
		}
	}
	return false
}
