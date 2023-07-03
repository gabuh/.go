package backend

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func Run(addr string) {
	http.HandleFunc("/", greet)
	log.Fatal(http.ListenAndServe(addr, nil))
}
