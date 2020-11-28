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

func NewStore() *Store {
	return &Store{}
}

func (s *Store) FindAll() []Todo {
	return s.db
}

func (s *Store) FindIndex(id int) (int, bool) {
	var foundIdx int

	for i, todo := range s.db {
		if todo.ID == id {
			foundIdx = i
		}
	}
	return foundIdx, foundIdx > -1
}

func (s *Store) Create(t Todo) {
	s.db = append(s.db, t)
}

func (s *Store) Update(t Todo) {
	i, found := s.FindIndex(t.ID)
	if found {
		s.db[i].Completed = t.Completed
	}
}

func (s *Store) Destroy(id int) {
	i, found := s.FindIndex(id)
	if found {
		s.db = append(s.db[:i], s.db[i+1:]...)
	}
}
