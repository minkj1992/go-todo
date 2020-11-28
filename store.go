package main

// Todo is ..
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Store saves todo
type Store struct {
	db []Todo
}
