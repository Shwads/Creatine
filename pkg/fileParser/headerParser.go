package fileParser

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

func headerParser(fileScanner *bufio.Scanner, requests map[string]interface{}) error {
	fmt.Printf("\nEntered headerParser\n\n")
	var currTag string

	fileScanner.Scan()

	for {
		isTag := false
		line := strings.Split(strings.TrimSpace(fileScanner.Text()), ":")
		line[0] = strings.TrimSpace(line[0])

		if len(line) > 1 {
			if line[1] == "" {
				isTag = true
				line = []string{line[0]}
			}
		}

		if line[0] == "body" {
			if len(line) > 1 {
				line[1] = strings.TrimSpace(line[1])
				if line[1] == "|" {
					fmt.Printf("calling bodyParser from headerParser\n")
					requests["body"] = ""
					bodyParser(fileScanner, requests)
					return nil
				} else {
					requests["body"] = line[1]
					return nil
				}
			} else {
				return errors.New("parser expects a body key to have a value, use a pipe | to signify a multi line value")
			}
		}

		if len(line) > 1 {
			fmt.Printf("length of line: %s was >1\n", fileScanner.Text())
			line[1] = strings.TrimSpace(line[1])
			if headerMap, ok := requests["headers"].(map[string][]string); ok {
				headerMap[line[0]] = []string{line[1]}
			} else {
				fmt.Printf("map did not have expected type. Instead it had type %T\n", requests["headers"])
			}
		} else {
			fmt.Printf("length of line: %s was not >1\n", fileScanner.Text())
			if !isTag {
				fmt.Printf("%s was judged to be not a tag\n", line[0])
				if headerMap, ok := requests["headers"].(map[string][]string); ok {
					fmt.Printf("length of requests['headers'] before is %d\n", len(headerMap[currTag]))
					headerMap[currTag] = append(headerMap[currTag], line[0])
					fmt.Printf("length of requests['headers'] after is %d\n", len(headerMap[currTag]))
					fmt.Printf("currTag is %s and its value is: %v\n", currTag, headerMap[currTag])
				} else {
					fmt.Printf("map did not have expected type. Instead it had type %T\n", requests["headers"])
				}

			} else {
				fmt.Printf("%s was judged to be a tag\n", line[0])
				currTag = line[0]
				if headerMap, ok := requests["headers"].(map[string][]string); ok {
					headerMap[currTag] = make([]string, 0)
				} else {
					fmt.Printf("map did not have expected type. Instead it had type %T\n", requests["headers"])
				}

			}
		}
		if !fileScanner.Scan() {
			fmt.Printf("\nExiting headerParser\n\n")
			return nil
		}
	}
}
