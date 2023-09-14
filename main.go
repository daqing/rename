package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

// return newName from baseName
// e.g. xyz.tar.gz -> xyz.tar
func newName(baseName string) string {
	parts := strings.Split(baseName, ".")
	sz := len(parts)

	return strings.Join(parts[0:sz-1], ".")
}

func handleRenameToSuffix(from string, suffix string) error {
	if suffix[0] != '.' {
		suffix = "." + suffix
	}

	dir := path.Dir(from)
	baseName := path.Base(from)
	newName := newName(baseName)
	ext := suffix

	fullNewName := dir + "/" + newName + ext

	if err := os.Rename(from, fullNewName); err != nil {
		fmt.Printf("Rename error: %s\n", err)
		return err
	}

	fmt.Printf("File [%s] renamed to [%s]\n", from, fullNewName)
	return nil
}

func main() {
	if len(os.Args) < 3 {
		showHelp()

		os.Exit(1)
	}

	var from = os.Args[1]
	var to = os.Args[2]

	if to[0] == '.' {
		if err := handleRenameToSuffix(from, to); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}

	dir := path.Dir(from)
	ext := path.Ext(from)

	fullNewName := dir + "/" + to + ext

	if err := os.Rename(from, fullNewName); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Printf("File [%s] renamed to [%s]\n", from, fullNewName)

	os.Exit(0)
}

func showHelp() {
	fmt.Printf("Usage: rename [from] [to]\n\n")
	fmt.Printf("[examples]\n\n")
	fmt.Println("$ rename a.txt b")
	fmt.Printf("  a.txt renamed to b.txt\n\n")
	fmt.Println("$ rename a.txt .mp3")
	fmt.Println("  a.txt renamed to a.mp3")
}
