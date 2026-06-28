package engine

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"searchengine.com/src/internal/store"
)

type Engine struct {
	store *store.Store
}

func NewEngine(store *store.Store) *Engine {
	return &Engine{
		store: store,
	}
}

func (e *Engine) Name() string {
	return "Search Engine"
}

func (e *Engine) Init() error {
	fmt.Printf("\n [%s] Initializing", e.Name())
	return nil
}

func (e *Engine) Start() error {
	fmt.Printf("\n [%s] Started", e.Name())

	userQuery := e.prompt()

	e.handleQuery(userQuery)

	return nil
}

func (e *Engine) Stop() error {
	fmt.Printf("\n [%s] Shutting down", e.Name())
	return nil
}

func (e *Engine) prompt() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprintf(os.Stderr, "\n [%s] Search: ", e.Name())
		s, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input", err)
			break
		}

		s = strings.TrimSpace(s)

		if s != "" {
			return s
		}
	}

	return ""
}

func (e *Engine) handleQuery(query string) {
	// e.v1(query)
	e.v2(strings.ToLower(query))
}

func (e *Engine) v1(query string) {

	for _, file := range e.store.GetAllDocs() {
		if strings.Contains(strings.ToLower(file), strings.ToLower(query)) {
			fmt.Printf("Search results... %s", file)
			break
		}
	}
}

func (e *Engine) v2(query string) {
	searchIndex := e.store.GetSearchIndex()

	// e.store.PrintMap(searchIndex)
	if values, ok := searchIndex[query]; ok {
		fmt.Printf("Found %d result(s): %v\n", len(values), values)
	} else {
		fmt.Printf("No results for '%s' \n", query)
	}
}
