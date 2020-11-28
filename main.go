package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	baseFrontDir := "examples/vanillajs"
	db := []Todo{}
	db = append(db, Todo{1, "todo1", false})
	http.HandleFunc("/api/todos", func(rw http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(rw)
		enc.Encode(&db)
	})
	http.Handle("/", http.FileServer(http.Dir(baseFrontDir)))
	fmt.Println("server is running http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
