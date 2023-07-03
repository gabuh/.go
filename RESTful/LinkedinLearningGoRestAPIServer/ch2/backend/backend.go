package backend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB   *sql.DB
	Port string
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func (a *App) Initialize() {
	DB, err := sql.Open("sqlite3", "../../practice.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	a.DB = DB
}

func (a *App) Run() {
	http.HandleFunc("/", greet)
	fmt.Println("Server started and listening on the port: ", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, nil))
}
