package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func abs(path string) string {
	abs, _ := filepath.Abs(path)
	return abs
}

func TestGetWorkDir(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{[]string{"main"}, abs(".")},
		{[]string{"main", "test"}, abs("./test")},
		{[]string{"main", "../test"}, abs("../test")},
	}

	for _, tt := range tests {
		os.Args = tt.args

		got := GetWorkDir()

		if got != tt.want {
			t.Fatalf("got='%v', want='%v'", got, tt.want)
		}
	}
}

func TestValidateDir(t *testing.T) {
	td := t.TempDir()
	t1 := filepath.Join(td, "exists")
	t2 := filepath.Join(td, "doesnotexist")
	t3 := filepath.Join(td, "notadir")

	_ = os.Mkdir(t1, os.FileMode(0755))
	_ = os.WriteFile(t3, []byte(""), 0755)

	tests := []struct {
		path string
		want error
	}{
		{t1, nil},
		{t2, fmt.Errorf("%v does not exist", t2)},
		{t3, fmt.Errorf("%v is not a directory", t3)},
	}

	for _, tt := range tests {
		got := ValidateDir(tt.path)

		if got != nil && got.Error() != tt.want.Error() {
			t.Fatalf("got='%v', want='%v'", got, tt.want)
		}
	}
}

func TestGetFiles(t *testing.T) {
	td := t.TempDir()
	t1 := filepath.Join(td, "one")
	t2 := filepath.Join(td, "two")

	_ = os.Mkdir(t1, os.FileMode(0755))
	_ = os.WriteFile(t2, []byte(""), 0755)

	got, err := GetFiles(td)

	if err != nil {
		t.Fatalf("err should be nil")
	}

	if got[0] != "one" {
		t.Fatalf("first element should be 'one'")
	}

	if got[1] != "two" {
		t.Fatalf("first element should be 'two'")
	}
}
