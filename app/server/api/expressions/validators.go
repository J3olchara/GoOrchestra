package expressions

import (
	"slices"
	"strings"
)

func ValidateExpression(expr string) bool {
	expr = strings.ReplaceAll(expr, " ", "")
	signs := []uint8{'+', '-', '*', '/'}
	if !ContainsOnly(expr, "123456789+-*/()") || len(expr) == 0 || slices.Contains(signs, expr[0]) || slices.Contains(signs, expr[len(expr)-1]) {
		return false
	}
	for i := 1; i < len(expr); i++ {
		if slices.Contains(signs, expr[i-1]) && slices.Contains(signs, expr[i]) {
			return false
		}
	}
	return true
}
