package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

const (
	NUMBER = iota
	OPERATOR
	WHITE_SPACE
	EOF
)

var (
	OPERATORS = []string{"+", "-", "/", "*"}
)

type Token struct {
	Val   string
	Class int
}

func main() {
	input := []rune(readInput(""))
	pos := 0
	currentToken := &Token{Class: -1}

	for checkType(currentToken, EOF) {
		currentToken = getNextToken(input, &pos)
		if isOperator(currentToken.Val) {
			if currentToken.Val == "+" || currentToken.Val == "-" {
				//TODO ------------------------
				addSubtract()
			} else if currentToken.Val == "*" || currentToken.Val == "/" {
				//TODO------------------------------
				multiplyDivide()
			}
		}
	}

}

func getNextToken(input []rune, pos *int) *Token {
	switch getType(input[*pos]) {
	case NUMBER:
		return &Token{
			Val:   getNumber(input, pos),
			Class: NUMBER,
		}
	case OPERATOR:
		*pos++
		return &Token{
			Val:   string(input[*pos-1]),
			Class: OPERATOR,
		}
	case WHITE_SPACE:
		*pos++
		return getNextToken(input, pos)
	case EOF:
		*pos++
		return &Token{
			Val:   "",
			Class: EOF,
		}
	}
	return nil
}

func getType(input rune) int {
	if isDigit(input) {
		return NUMBER
	} else if isWhiteSpace(input) {
		return WHITE_SPACE
	} else if isOperator(input) {
		return OPERATOR
	} else {
		return EOF
	}
}

func checkType(token *Token, class int) bool {
	if token.Class == class {
		return true
	}
	return false
}

//Split up add subtract and divide multiply for order of operations
func addSubtract(token *Token, before int, input []rune, pos *int) int {
	result := 0
	next := getNextToken(input, pos)
	nextInt, err := strconv.Atoi(next.Val)

	if err != nil {
		panic(err)
	}
	if token.Val == "+" {
		result = add(before, nextInt)
	} else {
		result = subtract(before, nextInt)
	}
}

func multiplyDivide(previous Stack[*Token]) int {

}

func add(left int, right int) int {
	return left + right
}

func subtract(left int, right int) int {
	return left - right
}

func multiply(left int, right int) int {
	return left * right
}

func divide(left int, right int) int {
	return left / right
}

//Put in a file to read from that file
//Leave blank for terminal to be read
func readInput(path string) string {
	var reader *bufio.Reader
	if path == "" {
		reader = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.OpenFile(path, os.O_RDONLY, 0600)
		if err != nil {
			panic(err)
		}
		reader = bufio.NewReader(file)
	}

	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return text
}

func getNumber(input []rune, pos *int) string {
	digit := ""
	for ; isDigit(input[*pos]); (*pos)++ {
		digit += string(input[*pos])
	}
	return digit
}

func findOperand(tokens []Token) (int, error) {
	for pos, token := range tokens {
		if token.Class == OPERATOR {
			return pos, nil
		}
	}
	return 0, errors.New("Failed to find operand")
}

func isDigit(char rune) bool {
	if 48 <= int(char) && int(char) <= 57 {
		return true
	}
	return false
}

func isOperator(char string) bool {
	for _, operator := range OPERATORS {
		if operator == char {
			return true
		}
	}
	return false
}

func isWhiteSpace(char rune) bool {
	if char == ' ' {
		return true
	}
	return false
}
