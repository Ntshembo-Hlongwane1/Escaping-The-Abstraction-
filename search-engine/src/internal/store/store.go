package store

import "fmt"

type Store struct {
	documents   []string // This all the file names
	searchIndex map[string][]string
}

func NewStore() *Store {
	return &Store{
		documents:   make([]string, 0),
		searchIndex: make(map[string][]string, 0),
	}
}

func (s *Store) Name() string {
	return "\n Store Subsystem"
}

func (s *Store) Init() error {
	fmt.Printf("\n [%s] Initializing --> document store", s.Name())
	return nil
}

func (s *Store) Start() error {
	fmt.Printf("\n [%s] Started", s.Name())
	return nil
}

func (s *Store) Stop() error {
	fmt.Printf("\n [%s] Shutting down", s.Name())
	return nil
}

func (s *Store) UpdateDocuments(newDocs []string) {
	s.documents = newDocs
}

func (s *Store) GetAllDocs() []string {
	return s.documents
}

func (s *Store) UpdateSearchIndex(key, value string) {
	values := s.searchIndex[key]

	if len(values) > 0 && values[len(values)-1] == value {
		return
	}

	s.searchIndex[key] = append(s.searchIndex[key], value)
}

func (s *Store) GetSearchIndex() map[string][]string {
	return s.searchIndex
}

func (s *Store) PrintMap(m map[string][]string) {
	fmt.Printf("\n [%s] Search Index: ", s.Name())

	for key, values := range m {
		fmt.Printf(" %s: %v\n", key, values)
	}
}
