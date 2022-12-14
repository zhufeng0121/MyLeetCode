package main

import "strconv"

func evalRPN(tokens []string) int {
	symbol := make([]string, 0) // 符号栈
	operand := make([]int, 0)   // 操作数栈

	for i := 0; i < len(tokens); i++ {
		if IsSymbol(tokens[i]) {
			symbol = append(symbol, tokens[i])

		} else {
			v, _ := strconv.Atoi(tokens[i])
			operand = append(operand, v)

		}
		for len(operand) >= 2 && len(symbol) >= 1 {
			op1, op2 := operand[len(operand)-1], operand[len(operand)-2]
			operand = operand[:len(operand)-2]
			operator := symbol[len(symbol)-1]
			symbol = symbol[:len(symbol)-1]
			var temp int

			if operator == "+" {
				temp = op1 + op2
			} else if operator == "-" {
				temp = op2 - op1
			} else if operator == "*" {
				temp = op2 * op1
			} else if operator == "/" {
				temp = op2 / op1
			}
			operand = append(operand, temp)

		}
	}

	return operand[0]

}

func IsSymbol(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}
	return false
}
