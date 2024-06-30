package Ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LoadBanner loads the banner characters from a file into a map
func LoadBanner(name string) (map[rune]string, error) {
	var height int                  // Tracks the current height of the character being read
	Banner := make(map[rune]string) // Map to store the banner characters
	currentChar := rune(32)
	charLine := []string{}                     // Slice to store lines of the current character
	filePath := "bannerfiles/" + name + ".txt" // Construct the file path
	_, err := FileCheck(filePath)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Skip the first line (assuming it's a header or empty)
	for scanner.Scan() {
		line := scanner.Text()

		if height == 8 {
			Banner[currentChar] = strings.Join(charLine, "\n")

			currentChar++
			height = 0
			charLine = []string{}
		} else {
			charLine = append(charLine, line)
			height++
		}
	}
	// After the loop, check if there's a partially read character and store it
	if height > 0 {
		Banner[currentChar] = strings.Join(charLine, "\n")
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return Banner, nil
}
