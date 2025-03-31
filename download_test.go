package main

import (
	"os"
	"testing"
)

func TestCreateDirectory(t *testing.T) {
	path := createDirectory()

	if path == "" {
		t.Error(`createDirectory() = "", want "[loadimgs dir]/download/[current date]"`)
	}
}

func TestGetValidInteger(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Valid input", "5\n", 5},
		{"Zero input", "0\n", 0},
		{"Negative number", "-42\n", -42},
		{"Large number", "1000000\n", 1000000},
		{"Invalid input before valid", "hello\n8\n", 8},
		{"Floating point before valid", "12.5\n7\n", 7},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()

			r, w, _ := os.Pipe()
			w.WriteString(test.input)
			w.Close()
			os.Stdin = r

			result := getValidInteger("Enter a number: ")
			if result != test.expected {
				t.Errorf(`getValidInteger("Enter a number: ") = %d, want %d`, result, test.expected)
			}
		})
	}
}