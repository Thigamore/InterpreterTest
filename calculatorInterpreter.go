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
	Val   string
	Class int
}

func main() {
	input := []rune(readInput(""))
	pos := 0

	tokenList := InitList(getNextToken(input, &pos))
	currentToken := tokenList.Head

	fmt.Println(currentToken.Val)
	for !checkType(currentToken.Val, EOF) {
		currentToken.Next = &Node[*Token]{
			Val:    getNextToken(input, &pos),
			Before: currentToken,
		}
		currentToken = currentToken.Next
		tokenList.Tail = currentToken
		fmt.Println(*currentToken.Val)
	}

	fmt.Printf("Length: %d\n", pos-1)

	currentToken = tokenList.Head
	//Goes through and does multiplication and division first
	for !checkType(currentToken.Val, EOF) {
		if checkType(currentToken.Val, OPERATOR) {
			if currentToken.Val.Val == "*" || currentToken.Val.Val == "/" {
				currentToken.Val = &Token{
					Val:   fmt.Sprint(execExpr(currentToken)),
					Class: NUMBER,
				}
				if currentToken.Before.Before != nil {
					currentToken.Before.Before.Next = currentToken
					currentToken.Before = currentToken.Before.Before
				} else {
					currentToken.Before = nil
					tokenList.Head = currentToken
				}
				if currentToken.Next.Next != nil {
					currentToken.Next.Next.Before = currentToken
					currentToken.Next = currentToken.Next.Next
				} else {
					currentToken.Next = nil
					tokenList.Tail = currentToken
				}
			}
		}
		currentToken = currentToken.Next
	}
	currentToken = tokenList.Head
	//Goes through and does addition and subtraction second
	for !checkType(currentToken.Val, EOF) {
		if checkType(currentToken.Val, OPERATOR) {
			if currentToken.Val.Val == "+" || currentToken.Val.Val == "-" {
				currentToken.Val = &Token{
					Val:   fmt.Sprint(execExpr(currentToken)),
					Class: NUMBER,
				}
			}
			if currentToken.Before.Before != nil {
				currentToken.Before = currentToken.Before.Before
			} else {
				currentToken.Before = nil
				tokenList.Head = currentToken
			}
			if currentToken.Next.Next != nil {
				currentToken.Next = currentToken.Next.Next
			} else {
				currentToken.Next = nil
				tokenList.Tail = currentToken
			}
		}
		currentToken = currentToken.Next
	}
	fmt.Println()
	fmt.Println(tokenList.Head.Val.Val)
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

func execExpr(tokens *Node[*Token]) int {
	var operation func(int, int) int
	switch tokens.Val.Val {
	case "*":
		operation = multiply
	case "/":
		operation = divide
	case "+":
		operation = add
	case "-":
		operation = subtract
	}
	left, err := strconv.Atoi(tokens.Before.Val.Val)
	if err != nil {
		panic(err)
	}
	right, err := strconv.Atoi(tokens.Next.Val.Val)
	if err != nil {
		panic(err)
	}
	return operation(left, right)
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

func isOperator(char rune) bool {
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
