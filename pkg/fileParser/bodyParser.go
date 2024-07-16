package fileParser

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

// return true if the function encounters a request tag and the main thread should continue after bodyParser exits
// otherwise if we get to EOF then return false to indicate that the main thread can wrap up
func bodyParser(preserveNewlines bool, fileScanner *bufio.Scanner, requests map[string]interface{}) (bool, error) {
	fileScanner.Scan()

	for {
		line := strings.TrimSpace(fileScanner.Text())
		if line == "request:" {
			return true, nil
		}

		if bodyString, ok := requests["body"].(string); ok {
			line = strings.TrimSpace(line)
			bodyString = fmt.Sprintf("%s%s", bodyString, line)
			requests["body"] = bodyString
		} else {
			return false, errors.New("improperly formatted map object. map['body'] should have type string")
		}

		if !fileScanner.Scan() {
			return false, nil
		}
	}
}
