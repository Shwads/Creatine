package fileParser

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

// return true if the function encounters a request tag and the main thread should continue after bodyParser exits
// otherwise if we get to EOF then return false to indicate that the main thread can wrap up
func bodyParser(fileScanner *bufio.Scanner, requests map[string]interface{}) (bool, error) {
	fmt.Printf("\nSkipping the line: %s\n\n", fileScanner.Text())
	fileScanner.Scan()

	for {
		line := strings.TrimSpace(fileScanner.Text())
		if line == "request:" {
			fmt.Printf("Exiting body with a new request being found\n")
			return true, nil
		}

		if bodyString, ok := requests["body"].(string); ok {
			line = strings.TrimSpace(line)
			bodyString = fmt.Sprintf("%s%s", bodyString, line)
			requests["body"] = bodyString
		} else {
			fmt.Printf("improper formatting.  object had type %T\n", requests["body"])
			return false, errors.New("improperly formatted map object. map['body'] should have type string")
		}

		if !fileScanner.Scan() {
			fmt.Printf("Exiting the body parser with an EOF\n")
			return false, nil
		}
	}
}
