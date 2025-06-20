package main

import "fmt"

func main() {
	words := []string{"go", "java", "go", "python", "go", "java"}
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	fmt.Println("Word counts:")
	for k, v := range wordCount {
		fmt.Printf("%v: %v\n", k, v)
	}

	// Accessing, adding, deleting
	fmt.Println("Count for 'go':", wordCount["go"])
	wordCount["rust"] = 1
	fmt.Println("Added rust:", wordCount)
	delete(wordCount, "python")
	fmt.Println("After deleting python:", wordCount)
}
