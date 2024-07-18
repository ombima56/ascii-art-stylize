package ascii

import (
	"errors"
	"fmt"
	"os"
)

func FileCheck(fileName string) (string, error) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return fileName, fmt.Errorf("error stating file: %v", err)
	}

	fileSize := fileInfo.Size()
	expectedSizes := map[string]int64{
		"bannerfiles/standard.txt":   6623,
		"bannerfiles/thinkertoy.txt": 5558,
		"bannerfiles/shadow.txt":     7463,
	}

	expectedSize, exists := expectedSizes[fileName]
	if !exists {
		return fileName, errors.New("unknown banner file")
	}

	if fileSize != expectedSize {
		return fileName, errors.New("the banner file has been altered")
	}

	return fileName, nil
}
