package task

import (
	"github.com/J3olchara/GoOrchestra/app/server/api/core"
	"github.com/J3olchara/GoOrchestra/app/server/api/expressions"
	"log"
	"os"
	"strconv"
	"time"
)

func GetTaskWithBrackets(expr expressions.ExpressionServer) *Task {
	l, r := 0, len(expr.Expr)-1
	for r-l > 0 {
		if expr.Expr[l] != '(' {
			l += 1
		}
		if expr.Expr[r] != ')' {
			r -= 1
		}
		if expr.Expr[l] == '(' && expr.Expr[r] == ')' {
			expr.Expr = expr.Expr[l+1 : r]
			return GetTaskWithBrackets(expr)
		}
	}
	return GetTaskWithoutBrackets(expr)
}

func IsFloatChar(s uint8) bool {
	return '0' <= s && s <= '9' || s == '.'
}

func GetTaskWithoutBrackets(expr expressions.ExpressionServer) *Task {
	// expects no brackets inside expr
	var err error
	ChooseTask := map[uint8][][]float64{
		'/': {},
		'*': {},
		'-': {},
		'+': {},
	}
	i := 0
	for i < len(expr.Expr) {
		// arg1
		l, r := i, i
		for r < len(expr.Expr) && IsFloatChar(expr.Expr[r]) {
			r++
		}
		arg1, _ := strconv.ParseFloat(expr.Expr[l:r], 64)

		operation := expr.Expr[r]

		// arg2
		l, r = r+1, r+1
		for r < len(expr.Expr) && IsFloatChar(expr.Expr[r]) {
			r++
		}
		arg2, _ := strconv.ParseFloat(expr.Expr[l:r], 64)

		ChooseTask[operation] = append(ChooseTask[operation], []float64{arg1, arg2})
		i = r
	}

	var args []float64
	var operationTime int
	var Operation uint8
	for _, op := range []uint8{'/', '*', '-', '+'} {
		if allArgs := ChooseTask[op]; len(allArgs) != 0 {
			args = allArgs[0]
			operationTime, err = strconv.Atoi(os.Getenv(core.ExpressionConfigNames[op]))
			Operation = op
			break
		}
	}

	if err != nil {
		log.Fatal(err, "some troubles with times environs")
	}
	if len(args) == 0 {
		return nil
	}
	return &Task{
		Arg1:          args[0],
		Arg2:          args[1],
		Operation:     string(Operation),
		OperationTime: Timestamp(time.Duration(operationTime) * time.Millisecond),
		ExpressionId:  expr.ID,
	}
}
