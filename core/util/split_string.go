package util

import (
	"bufio"
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

	leftString := "" // 上一行遗留
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				errorResult <- err
			}
			break
		}

		line = strings.TrimSpace(line)
		leftString = leftString + line
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
