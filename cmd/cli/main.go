package main

import (
	"database/sql"
	"fmt"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	_ "github.com/lib/pq"
	"go-exercise/pkg/openapi3"
	"log"
	"net/http"
	_ "os/exec"
)

type TaskServer struct {
}

func main() {
	s := TaskServer{}

	h := openapi3.Handler(s)

	connStr := "postgresql://changeme:changeme@localhost:58570?sslmode=disable"
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	insertStmt := `insert into "task"("id", "views") values("4aca4462-693d-458f-8688-9e48c6e90609", 1)`
	_, _ = db.Exec(insertStmt)

	http.ListenAndServe(":3000", h)

}

func (t TaskServer) ReadTask(w http.ResponseWriter, r *http.Request, taskId openapi_types.UUID) {

	fmt.Println("Hello World!")

	// our logic to retrieve all todos from a persistent layer
}

func (t TaskServer) CreateTask(w http.ResponseWriter, r *http.Request) {
	// our logic to store the todo into a persistent layer

}

func (t TaskServer) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	// our logic to delete a todo from the persistent layer
}

var db *sql.DB
