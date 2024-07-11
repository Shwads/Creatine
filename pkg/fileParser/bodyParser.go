package fileParser

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)


func bodyParser(fileScanner *bufio.Scanner, requests map[string]interface{}) error {
	fmt.Printf("\nEntered bodyParser\n\n")
	fileScanner.Scan()

	for {
		line := strings.TrimSpace(fileScanner.Text())
		if bodyString, ok := requests["body"].(string); ok {
			line = strings.TrimSpace(line)
			bodyString = fmt.Sprintf("%s%s", bodyString, line)
			requests["body"] = bodyString
		} else {
			fmt.Printf("improper formatting.  object had type %T\n", requests["body"])
			return errors.New("improperly formatted map object. map['body'] should have type string")
		}

		if !fileScanner.Scan() {
			fmt.Printf("\nExiting bodyParser\n\n")
			return nil
		}
	}
}
