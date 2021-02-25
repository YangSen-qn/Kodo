package util

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Split(s string, separator string) []string {
	return strings.Split(s, separator)
}

func SplitFromFile(filePath string, separator string, partResult chan<- string, errorResult chan<- error) {
	defer func() {
		close(partResult)
		close(errorResult)
	}()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("log error", err)
		errorResult <- err
		return
	}
	defer file.Close()

	leftString := ""
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			errorResult <- err
			break
		}
		if n == 0 {
			break
		}

		leftString = leftString + string(buf)
		stringList := Split(leftString, separator)
		listLength := len(stringList)
		for i, s := range stringList {
			if i < listLength-1 {
				partResult <- s
			} else {
				leftString = s
			}
		}
	}

	if len(leftString) > 0 {
		partResult <- leftString
	}

	return
}
