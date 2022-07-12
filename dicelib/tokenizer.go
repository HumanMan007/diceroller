package dicelib

import (
	"fmt"
	"strings"
)

type TokenType int

const (
	ROLL TokenType = iota
	LEFT_PAREN
	RIGHT_PAREN

	ADVANTAGE
	DISADVANTAGE
	REROLL

	IF
	THEN

	PLUS
	MINUS
	TIMES
	DIVIDE

	NEWLINE
	EOF
	UNEXPECTED
)

type Token struct {
	Type TokenType
	Str  string
}

func nextToken(input string) (Token, string, error) {
	if len(input) == 0 {
		return Token{EOF, ""}, input, nil
	}

	if input[0] == 0 || input[0] == '\t' { // Lazy consume whitespace
		return nextToken(input[1:])
	}

	switch input[0] {
	case '(':
		return Token{LEFT_PAREN, "("}, input[1:], nil
	case ')':
		return Token{RIGHT_PAREN, ")"}, input[1:], nil
	case '+':
		return Token{PLUS, "+"}, input[1:], nil
	case '-':
		return Token{MINUS, "-"}, input[1:], nil
	case '*':
		return Token{TIMES, "*"}, input[1:], nil
	case '/':
		return Token{DIVIDE, "/"}, input[1:], nil
	case '\n':
		return Token{NEWLINE, "\n"}, input[1:], nil
	}

	if strings.ToLower(input[0:2]) == "if" {
		return Token{IF, "if"}, input[2:], nil
	}

	switch strings.ToLower(input[0:3]) {
	case "adv":
		return Token{ADVANTAGE, "adv"}, input[3:], nil
	case "dis":
		return Token{DISADVANTAGE, "dis"}, input[3:], nil
	}

	if strings.ToLower(input[0:4]) == "then" {
		return Token{THEN, "then"}, input[6:], nil
	}

	if strings.ToLower(input[0:6]) == "reroll" {
		return Token{REROLL, "reroll"}, input[6:], nil
	}

	val, input, err := getProb(input)
	if val != "" {
		return Token{ROLL, val}, input, err
	}

	return Token{UNEXPECTED, input[0:1]}, input[1:], fmt.Errorf("unexpected carater %c", input[0])
}

func isDigit(char byte) bool {
	return char >= '0' && char < '9'
}

func getInt(input string) (string, string) {
	if !isDigit(input[0]) {
		return "", input
	}

	curr := 0
	for len(input) > curr && isDigit(input[curr]) {
		curr++
	}

	return input[:curr], input[curr:]
}

func getProb(input string) (string, string, error) {
	var val string
	if val, input = getInt(input); val == "" {
		return val, input, nil
	}

	if input[0] == 'd' {
		val2 := ""
		val2, input = getInt(input[1:])
		if val2 == "" {
			return val, input, fmt.Errorf("expected number after 'd', got :%c", input[0])
		}
		val = val + "d" + val2
	}

	return val, input, nil
}
