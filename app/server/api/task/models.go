package task

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/J3olchara/GoOrchestra/app/server/api/expressions"
	"github.com/J3olchara/GoOrchestra/app/server/db"
	"time"
)

type Timestamp time.Duration

type Task struct {
	ID            int       `json:"id"`
	Arg1          float64   `json:"arg1"`
	Arg2          float64   `json:"arg2"`
	Operation     string    `json:"operation"`
	OperationTime Timestamp `json:"operation_time"`
	ExpressionId  int       `json:"-"`
}

type TaskServer struct {
	Task
	Result float64 `json:"result"`
}

func (t *Timestamp) Scan(src interface{}) error {
	if time, ok := src.(time.Duration); ok {
		*t = Timestamp(time)
	}
	return nil
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(t).Milliseconds())
}

func (t Timestamp) Value() (driver.Value, error) {
	return time.Duration(t).Milliseconds(), nil
}

func (t *Task) Create() error {
	row := db.DB.Conn.QueryRow(
		"INSERT INTO Task(arg1, arg2, operation, durationMS, expression_id) VALUES($1, $2, $3, $4, $5) returning id",
		t.Arg1, t.Arg2, t.Operation, t.OperationTime, t.ExpressionId,
	)
	return row.Scan(&t.ID)
}

func GetExpressionByTask(id int) (*expressions.ExpressionServer, *TaskServer) {
	var expr expressions.ExpressionServer
	var task TaskServer
	row := db.DB.Conn.QueryRow(
		"SELECT E.id, E.result, T.arg1, T.arg2, T.operation from task as T LEFT JOIN public.expressions E on T.expression_id = E.id WHERE T.id = $1",
		id,
	)
	task.ID = id
	err := row.Scan(&expr.ID, &expr.Expr, &task.Arg1, &task.Arg2, &task.Operation)
	if err != nil {
		return nil, nil
	}
	return &expr, &task
}
