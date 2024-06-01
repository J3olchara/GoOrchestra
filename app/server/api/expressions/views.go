package expressions

import (
	"encoding/json"
	"github.com/J3olchara/GoOrchestra/app/server/api/core"
	"github.com/J3olchara/GoOrchestra/app/server/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
}

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Conn.Query("SELECT id, status, result FROM expressions ORDER BY status")
	if err != nil {
		core.ServerError.WithError(err, w, r)
		return
	}
	exprs := make([]ExpressionUser, 0, 10)
	for rows.Next() {
		var expr ExpressionUser
		err = rows.Scan(&expr.ID, &expr.Status, &expr.Result)
		if err != nil {
			core.ServerError.WithError(err, w, r)
			return
		}
		exprs = append(exprs, expr)
	}
	data, err := json.Marshal(map[string]interface{}{"expressions": exprs})
	if err != nil {
		core.ServerError.WithError(err, w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var expr ExpressionServer
	err := json.NewDecoder(r.Body).Decode(&expr)
	expr.Expr = strings.ReplaceAll(expr.Expr, " ", "")
	if !ValidateExpression(expr.Expr) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		log.Println(err)
		return
	}
	row := db.DB.Conn.QueryRow("INSERT INTO expressions(result, expression) VALUES($1, $1) RETURNING id", expr.Expr)
	if err != nil {
		core.ServerError.WithError(err, w, r)
		return
	}
	err = row.Scan(&expr.ID)
	if err != nil {
		core.ServerError.WithError(err, w, r)
		return
	}
	data, err := json.Marshal(map[string]interface{}{"id": expr.ID})
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func (h Handler) Retrieve(w http.ResponseWriter, r *http.Request) {
	var err error
	var expr ExpressionUser
	expr.ID, err = strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		core.NotFound.ServeHTTP(w, r)
		return
	}
	row := db.DB.Conn.QueryRow("SELECT result, status from expressions WHERE id = $1", expr.ID)
	err = row.Scan(&expr.Result, &expr.Status)
	if err != nil {
		log.Println(err)
		core.NotFound.ServeHTTP(w, r)
		return
	}

	data, err := json.Marshal(map[string]interface{}{"expression": expr})
	if err != nil {
		core.ServerError.WithError(err, w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
