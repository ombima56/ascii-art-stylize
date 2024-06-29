package Ascii

import (
	"errors"
	"os"
)

func FileCheck(fileName string) (string, error) {
	var fileSize int64

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return fileName, err
	}
	fileSize = fileInfo.Size()

	if (fileName == "bannerfile/standard.txt" && fileSize != 6623) ||
		(fileName == "bannerfile/thinkertoy.txt" && fileSize != 5558) ||
		(fileName == "bannerfile/shadow.txt" && fileSize != 7463) {
		return fileName, errors.New("the Banner file has been altered: ")
	}

	return fileName, nil
}
