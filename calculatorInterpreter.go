package main

import (
	"bufio"
	"errors"
	"fmt"
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
	OPERATORS = []rune{'+', '-', '/', '*'}
)

type Token struct {
	value string
	class int
}

func main() {
	input := []rune(readInput(""))

	tokens := InitList[*Token]()

}

func getNextToken(input []rune, pos int) *Token {

}

func getType(input)

func execExpr(tokens *[]Token, operandPos int) {
	leftNum := 0
	rightNum := 0
	var operation func(int, int)
	var err error
	afterOperand := false
	for pos, token := range *tokens {
		fmt.Println(pos, leftNum, rightNum)
		if pos == operandPos {
			if token.value == "+" {
				operation = add
			} else if token.value == "-" {
				operation = subtract
			} else if token.value == "*" {
				operation = multiply
			} else if token.value == "/" {
				operation = divide
			}

		} else if token.class == WHITE_SPACE {
		} else if token.value == "" {
			break
		} else if !afterOperand {
			leftNum, err = strconv.Atoi(token.value)
			afterOperand = true
		} else if afterOperand {
			rightNum, err = strconv.Atoi(token.value)
		}
		if err != nil {
			panic(err)
		}
	}
	operation(leftNum, rightNum)
}

func add(left int, right int) {
	fmt.Println(left, right)
	fmt.Println(left + right)
}

func subtract(left int, right int) {
	fmt.Println(left - right)
}

func multiply(left int, right int) {
	fmt.Println(left * right)
}

func divide(left int, right int) {
	fmt.Println(left / right)
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

func getDigit(input []rune, pos int) (string, int) {
	digit := ""
	length := 0
	for ; isDigit(input[pos]); pos++ {
		digit += string(input[pos])
		length = pos
	}
	return digit, length + 1
}

func findOperand(tokens []Token) (int, error) {
	for pos, token := range tokens {
		if token.class == OPERATOR {
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

func isOperand(char rune) bool {
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
