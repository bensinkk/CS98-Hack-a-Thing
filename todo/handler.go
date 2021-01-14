package function

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"net/http"
)

var db *sql.DB

type Todo struct {
	Description string `json:"description"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{}
	todos = append(todos, Todo{Description: "Run faas-cli up"})
	res, _ := json.Marshal(todos)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}
