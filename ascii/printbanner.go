package Ascii

import (
	"fmt"
	"strings"
)

// PrintBanner prints the input string using the loaded banner characters.
func PrintBanner(line, filename string) string {
	outPut := make([][]string, 8) // Output slice to store the banner lines.

	banner, err := LoadBanner(filename)
	if err != nil {
		return fmt.Sprintf("Error loading banner: %v", err)
	}

	for _, char := range line {
		if char < 32 || char > 126 {
			return fmt.Sprintf("Character out of range: %q\n", char)
		}
		if ascii, ok := banner[char]; ok {
			asciiLines := strings.Split(ascii, "\n")
			for i := 0; i < len(asciiLines); i++ {
				outPut[i] = append(outPut[i], asciiLines[i])
			}
		} else {
			return fmt.Sprintf("Character not found: %q\n", char)
		}
	}
	result := string('\n')
	// Print the assembled output lines
	for _, line := range outPut {
		result += (strings.Join(line, "")) + string('\n')
	}

	return result
}
