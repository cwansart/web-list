package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// checks program args and uses the first argument as working dir
// if no argument is set, use the current dir
func GetWorkDir() string {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	abs, e := filepath.Abs(dir)
	if e != nil {
		return dir
	}
	return abs
}

// checks if dir exists and is a directory
func ValidateDir(dir string) error {
	s, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%v does not exist", dir)
		}
		// unchecked error
		return err
	}
	if !s.IsDir() {
		return fmt.Errorf("%v is not a directory", dir)
	}
	return nil
}

// return all file names in given path, fails if os.ReadDir fails
func GetFiles(path string) ([]string, error) {
	de, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var f []string
	for _, v := range de {
		f = append(f, v.Name())
	}

	return f, nil
}

func main() {
	d := GetWorkDir()
	log.Printf("using dir: %v\n", d)

	err := ValidateDir(d)
	if err != nil {
		log.Fatalf("invalid dir: %v\n", err)
	}

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := GetFiles(d)
		if err != nil {
			log.Fatalf("failed to get files: %v\n", err)
		}

		l := list(d, f)
		l.Render(r.Context(), w)
	}))

	fmt.Println("listening on http://localhost:3000")
	if err := http.ListenAndServe("localhost:3000", nil); err != nil {
		log.Printf("error listening: %v", err)
	}
}
