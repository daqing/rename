package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

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
	} else if to[0] == ':' {
		if err := handleTransform(from, to); err != nil {
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
	fmt.Printf("  a.txt renamed to a.mp3\n\n")
	fmt.Println("$ rename a.txt :upper")
	fmt.Printf("  a.txt renamed to A.txt\n\n")
	fmt.Println("$ rename LONG-NAME.txt :lower")
	fmt.Printf("  LONG-NAME.txt renamed to long-name.txt\n\n")
	fmt.Println("$ rename a.txt :plus -super")
	fmt.Printf("  a.txt renamed to a-super.txt\n\n")
}

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

// we would do some transformation to the file name
// e.g. rename a.txt :upper -> A.txt
// e.g. rename A.txt :lower -> a.txt
func handleTransform(from string, to string) error {
	fmt.Println(from)
	fmt.Printf("transform to: %s\n", to)

	// get basename
	dir := path.Dir(from)
	baseName := path.Base(from)
	newName := newName(baseName)

	fmt.Println("newName: ", newName)
	fmt.Println("dir: ", dir)

	switch to {
	case ":upper":
		newName = strings.ToUpper(newName)
	case ":lower":
		newName = strings.ToLower(newName)
	case ":plus":
		if len(os.Args) > 3 {
			newName = newName + os.Args[3]
		} else {
			return fmt.Errorf("missing argument for :plus")
		}
	default:
		// do nothing
	}

	ext := path.Ext(from)

	return os.Rename(from, dir+"/"+newName+ext)
}
