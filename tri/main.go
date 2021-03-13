package main

import (
	"bufio"
	"fmt"
	"maw/trie"
	"os"
	"strings"
)

func getTrie(filePath string) trie.Trie {
	disctionaryFile, err := os.Open(filePath)
	if nil != err {
		panic(err)
	}
	scanner := bufio.NewScanner(disctionaryFile)
	scanner.Split(bufio.ScanLines)

	fmt.Println("Reading the dictionary")
	wordCount := 0
	trie := trie.NewTrie()
	for scanner.Scan() {
		newWord := scanner.Text()
		trie.Insert(newWord)
		wordCount++
	}
	fmt.Printf("Found %d words\n", wordCount)
	return trie
}

func main() {
	args := os.Args[1:]
	if 1 != len(args) {
		fmt.Println("Rquired one argument with the doctionary")
	}
	dictionaryPath := args[0]
	trie := getTrie(dictionaryPath)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a new word to complete: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		completions := trie.AutoComplete(text)
		for _, word := range completions {
			fmt.Println(word)
		}
	}
}
