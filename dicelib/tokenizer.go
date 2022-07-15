package dicelib

import (
	"strings"
)

type TokenType int

const (
	NUMBER TokenType = iota
	LEFT_PAREN
	RIGHT_PAREN
	COMA

	ADVANTAGE
	DISADVANTAGE
	REROLL

	IF
	THEN

	D
	PLUS
	MINUS
	TIMES
	DIVIDE

	EOF
	UNEXPECTED
)

type Token struct {
	Type TokenType
	Str  string
}

func nextToken(input string) (Token, string) {
	if len(input) == 0 {
		return Token{EOF, ""}, input
	}

	if input[0] == 0 || input[0] == '\t' || input[0] == '\n' { // Lazy consume whitespace
		return nextToken(input[1:])
	}

	if strings.ToLower(input[0:6]) == "reroll" {
		return Token{REROLL, input[:6]}, input[6:]
	}

	if strings.ToLower(input[0:4]) == "then" {
		return Token{THEN, input[:4]}, input[6:]
	}

	switch strings.ToLower(input[0:3]) {
	case "adv":
		return Token{ADVANTAGE, input[:3]}, input[3:]
	case "dis":
		return Token{DISADVANTAGE, input[:3]}, input[3:]
	}

	if strings.ToLower(input[0:2]) == "if" {
		return Token{IF, input[:2]}, input[2:]
	}

	switch input[0] {
	case '(', '[', '{':
		return Token{LEFT_PAREN, input[:1]}, input[1:]
	case ')', ']', '}':
		return Token{RIGHT_PAREN, input[:1]}, input[1:]
	case '+':
		return Token{PLUS, input[:1]}, input[1:]
	case '-':
		return Token{MINUS, input[:1]}, input[1:]
	case '*':
		return Token{TIMES, input[:1]}, input[1:]
	case '/':
		return Token{DIVIDE, input[:1]}, input[1:]
	case 'd', 'D':
		return Token{D, input[:1]}, input[1:]
	case ',':
		return Token{COMA, input[:1]}, input[1:]
	}

	val, input := getInt(input)
	if val != "" {
		return Token{NUMBER, val}, input
	}

	return Token{UNEXPECTED, input[0:1]}, input[1:]
}

func isDigit(char byte) bool {
	return char >= '0' && char < '9'
}

func getInt(input string) (string, string) {
	curr := 0
	for len(input) > curr && isDigit(input[curr]) {
		curr++
	}

	return input[:curr], input[curr:]
}
