package main

import (
	"fmt"
	"os"
	"strings"
)

const notesFile = "notes.txt"

func addNote(note string) {
	f, err := os.OpenFile(notesFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error writing note:", err)
		return
	}
	defer f.Close()
	_, err = f.WriteString(note + "\n")
	if err != nil {
		fmt.Println("Failed to save note:", err)
	}
	fmt.Println("Note added!")
}

func listNotes() {
	data, err := os.ReadFile(notesFile)
	if err != nil {
		fmt.Println("No notes found.")
		return

	}
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if line != "" {
			fmt.Printf("%d, %s\n", i+1, line)
		}
	}
}

func deleteNote(index int) {
	data, err := os.ReadFile(notesFile)
	if err != nil {
		fmt.Println("No notes found.")
		return
	}
	lines := strings.Split(string(data), "\n")
	if index <= 0 || index > len(lines)-1 {
		fmt.Println("invalid index")
		return
	}
	lines = append(lines[:index-1], lines[index:]...)
	err = os.WriteFile(notesFile, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		fmt.Println("Error deleting note:", err)
		return
	}
	fmt.Println("Note deleted.")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: quicknotes add|list|delete")
		return
	}
	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a note to add.")
			return
		}
		addNote(strings.Join(os.Args[2:], ""))
	case "list":
		listNotes()
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the index of the note to delete.")
			return
		}
		var index int
		fmt.Sscanf(os.Args[2], "%d", &index)
		deleteNote(index)
	default:
		fmt.Println("Unknown command.")
	}
}
