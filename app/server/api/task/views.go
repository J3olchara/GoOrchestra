package task

import (
	"encoding/json"
	"github.com/J3olchara/GoOrchestra/app/server/api/core"
	"github.com/J3olchara/GoOrchestra/app/server/api/expressions"
	"github.com/J3olchara/GoOrchestra/app/server/db"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
}

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	var expr expressions.ExpressionServer
	row := db.DB.Conn.QueryRow("UPDATE expressions SET blocked = true WHERE id=(SELECT id FROM expressions WHERE blocked=false AND status='waiting' LIMIT 1) RETURNING id, result")
	err := row.Scan(&expr.ID, &expr.Expr)
	if err != nil {
		core.NotFoundHandler{}.ServeHTTP(w, r)
		return
	}
	task := GetTaskWithBrackets(expr)
	if err = task.Create(); err != nil {
		db.DB.Conn.Exec("UPDATE expressions SET blocked = false WHERE id=$1", expr.ID)
		core.ServerError.WithError(err, w, r)
		return
	}
	data, err := json.Marshal(map[string]interface{}{"task": task})
	if err != nil {
		db.DB.Conn.Exec("UPDATE expressions SET blocked = false WHERE id=$1", expr.ID)
		core.ServerError.WithError(err, w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var task *TaskServer
	err := json.NewDecoder(r.Body).Decode(&task)
	result := strconv.FormatFloat(task.Result, 'g', -1, 64)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	expr, task := GetExpressionByTask(task.ID)
	if expr == nil {
		core.NotFound.ServeHTTP(w, r)
		return
	}
	arg1 := strconv.FormatFloat(task.Arg1, 'g', -1, 64)
	arg2 := strconv.FormatFloat(task.Arg2, 'g', -1, 64)
	expr.Expr = strings.ReplaceAll(expr.Expr, arg1+task.Operation+arg2, result)
	expr.NormalizeExpression()
	if _, err = strconv.ParseFloat(expr.Expr, 8); err == nil {
		_, err = db.DB.Conn.Exec(
			"UPDATE expressions SET result=$1, blocked=false, status='done' where id=$2;",
			expr.Expr, expr.ID,
		)

	} else {
		_, err = db.DB.Conn.Exec(
			"UPDATE expressions SET result=$1, blocked=false where id=$2;",
			expr.Expr, expr.ID,
		)
	}
	if err != nil {
		core.ServerError.WithError(err, w, r)
		return
	}
	_, err = db.DB.Conn.Exec(
		"DELETE FROM task WHERE id=$1",
		task.ID,
	)
	if err != nil {
		core.ServerError.WithError(err, w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
}
