package kernal

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"searchengine.com/src/internal/store"
)

type FileReader struct {
	dir   string
	store *store.Store
}

var stopWords = map[string]struct{}{
	// Articles & Determiners
	"the": {}, "a": {}, "an": {}, "this": {}, "that": {},
	"these": {}, "those": {}, "my": {}, "your": {}, "his": {},
	"her": {}, "its": {}, "our": {}, "their": {}, "each": {},
	"every": {}, "some": {}, "any": {}, "no": {}, "all": {},

	// Prepositions
	"in": {}, "on": {}, "at": {}, "by": {}, "for": {},
	"with": {}, "about": {}, "against": {}, "between": {}, "into": {},
	"through": {}, "during": {}, "before": {}, "after": {}, "above": {},
	"below": {}, "to": {}, "from": {}, "up": {}, "down": {},

	// Conjunctions
	"and": {}, "but": {}, "or": {}, "nor": {}, "for_conj": {},
	"yet": {}, "so": {}, "although": {}, "because": {}, "since": {},
	"unless": {}, "while": {}, "whereas": {}, "if": {}, "whether": {},

	// Pronouns
	"i": {}, "you": {}, "he": {}, "she": {}, "it": {},
	"we": {}, "they": {}, "me": {}, "him": {}, "her_pron": {},
	"us": {}, "them": {}, "myself": {}, "yourself": {}, "himself": {},
	"herself": {}, "itself": {}, "ourselves": {}, "themselves": {}, "who": {},
	"whom": {}, "whose": {}, "which": {}, "what": {},

	// Auxiliary & Common Verbs
	"am": {}, "is": {}, "are": {}, "was": {}, "were": {},
	"be": {}, "been": {}, "being": {}, "have": {}, "has": {},
	"had": {}, "do": {}, "does": {}, "did": {}, "can": {},
	"could": {}, "will": {}, "would": {}, "shall": {}, "should": {},
	"may": {}, "might": {}, "must": {}, "go": {}, "get": {},
	"make": {}, "say": {}, "know": {}, "take": {}, "see": {},
}

func NewFileReader(dir string, store *store.Store) *FileReader {
	return &FileReader{
		dir:   dir,
		store: store,
	}
}

func (f *FileReader) Name() string {
	return "FileReader"
}

func (f *FileReader) Init() error {
	fmt.Printf("\n [%s] Initializing file reader", f.Name())

	fmt.Printf("\n [%s] Getting all docuement names", f.Name())

	docNames, err := f.GetAllFileNames()

	if err != nil {
		fmt.Printf("\n [%s] Failed getting doc names", f.Name())
	}

	f.store.UpdateDocuments(docNames)

	return nil
}

func (f *FileReader) Start() error {
	fmt.Printf("[%s] Started file reader", f.Name())
	return nil
}

func (f *FileReader) Stop() error {
	fmt.Printf("\n [%s] Shuttin down", f.Name())
	return nil
}

func (f *FileReader) GetAllFileNames() ([]string, error) {
	var files []string

	wd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(filepath.Join(wd, f.dir))

	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
			f.Tokenize(entry.Name())
		}
	}

	return files, nil
}

func (f *FileReader) Tokenize(docFile string) {
	wd, err := os.Getwd()

	fmt.Printf("\n [%s] Tokenzing: %s", f.Name(), docFile)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting working directory: ", err)
		return
	}

	path := filepath.Join(wd, f.dir, docFile)

	file, err := os.Open(path)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		for _, word := range words {
			cleanedWord := removePunctuation(word)

			if !isStopWord(cleanedWord) {
				f.store.UpdateSearchIndex(cleanedWord, docFile)
			}
		}
	}
}

func removePunctuation(word string) string {
	const punctuation = ".,!?;\"()[]{}"

	cleanedWord := strings.Trim(strings.ToLower(word), punctuation)
	return cleanedWord
}

func isStopWord(word string) bool {
	_, exist := stopWords[word]

	return exist
}
