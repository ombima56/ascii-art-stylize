package ascii

import (
	"fmt"
	"strings"
)

// PrintBanner prints the input string using the loaded banner characters.
func PrintBanner(line, filename string) (string, error) {
	outPut := make([][]string, 8) // Output slice to store the banner lines.

	banner, err := LoadBanner(filename)
	if err != nil {
		return "", fmt.Errorf("error loading banner: %w", err)
	}

	for _, char := range line {
		if char < 32 || char > 126 {
			return "", fmt.Errorf("character out of range: %q", char)
		}
		if ascii, ok := banner[char]; ok {
			asciiLines := strings.Split(ascii, "\n")
			for i := 0; i < len(asciiLines); i++ {
				outPut[i] = append(outPut[i], asciiLines[i])
			}
		} else {
			return "", fmt.Errorf("character not found: %q", char)
		}
	}

	var result strings.Builder
	// Print the assembled output lines
	for _, line := range outPut {
		result.WriteString(strings.Join(line, ""))
		result.WriteString("\n")
	}

	return result.String(), nil
}
