package expressions

import (
	"regexp"
)

type ExpressionUser struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	Result string `json:"result"`
}

type ExpressionServer struct {
	ExpressionUser
	Expr string `json:"expression"`
}

type ExpressionDB struct {
	ExpressionServer
	Blocked string `json:"blocked"`
}

func (e *ExpressionServer) NormalizeExpression() {
	pattern := regexp.MustCompile(`[(]+([0-9]+)[)]+`)
	e.Expr = pattern.ReplaceAllStringFunc(e.Expr, func(str string) string {
		return pattern.FindStringSubmatch(str)[1]
	})
}
