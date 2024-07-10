/*
This is incredibly hacky. Should've just used and tweaked someone elses yaml parser but I wanted to do it myself. Maybes you can write a better one in a future project
but for now this will have to do.
*/

package fileParser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Mode int8

const (
	Normal Mode = iota
	Headers
	List
	MultiLineVal
)

var parserMode Mode

func ParseRequestFile(fileName string) (map[string]interface{}, error) {
    file, fileReadErr := os.Open(fileName)
    if fileReadErr != nil {
        fmt.Printf("Encountered error: %s. In function ParseRequestFile", fileReadErr)
        return nil, fileReadErr
    }
    defer file.Close()

    return nil, nil
}

func ParseFile(fileName string) {
	file, fileReadErr := os.Open(fileName)
	defer file.Close()

	if fileReadErr != nil {
		fmt.Printf("Encountered Error: %s.\n", fileReadErr)
		return
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var currTag string
	var requestTag string
    var isTag bool

	parserMode = Normal
	requestNum := 0
	requests := make(map[string]interface{})

	for fileScanner.Scan() {
        isTag = false
        for _, rn := range []rune(fileScanner.Text()) {
            if rn == ':' {
                isTag = true
            }
        }

		line := strings.Split(fileScanner.Text(), ":")
		switch parserMode {
		case Normal:
			if line[0] == "request" {
				requestNum += 1
				requestTag = fmt.Sprintf("request-%d", requestNum)
				requests[requestTag] = make(map[string]interface{})
				break
			}

			if line[0] == "headers" {
				if requestMap, ok := requests[requestTag].(map[string]interface{}); ok {
					requestMap["headers"] = make(map[string][]string)
				}
				parserMode = Headers
				break
			}

			if len(line) > 1 {
				line[0] = strings.TrimSpace(line[0])
				line[1] = strings.TrimSpace(line[1])

				if requestMap, ok := requests[requestTag].(map[string]interface{}); ok {
					if line[1] == "|" {
						requestMap[line[0]] = ""
						parserMode = MultiLineVal
					}
					requestMap[line[0]] = line[1]
				}
			}
			break

		case Headers:
			line[0] = strings.TrimSpace(line[0])
			asRunes := []rune(line[0])

			if requestMap, ok := requests[requestTag].(map[string]interface{}); ok {
				if asRunes[0] == '-' {
					line[0] = string(asRunes[1:])
					line[0] = strings.TrimSpace(line[0])

					if headerMap, ok := requestMap["Headers"].(map[string][]string); ok {
						if len(line) > 1 {
							headerMap[line[0]] = []string{strings.TrimSpace(line[1])}
							break
						}
						parserMode = List
						currTag = line[0]
						headerMap[line[0]] = make([]string, 0)
						break
					}
				}

				if len(line) > 1 {
					line[1] = strings.TrimSpace(line[1])
					if line[1] == "|" {
						parserMode = MultiLineVal
						break
					}
					requestMap[line[0]] = strings.TrimSpace(line[1])
				}
			}
			break
		case List:
			if requestMap, ok := requests[requestTag].(map[string]interface{}); ok {
				line[0] = strings.TrimSpace(line[0])
				asRunes := []rune(line[0])

				if asRunes[0] != '-' {
                    line[0] = string(asRunes[1:])
                    if len(line) > 1 {
                        if strings.TrimSpace(line[1]) == "|" {
                            currTag = strings.TrimSpace(line[0])
                            requestMap[currTag] = make([]byte, 0)
                            parserMode = MultiLineVal
                            break
                        }

                        requestMap[line[0]] = line[1]
                        parserMode = Normal
                        break
                    }
				}

                if len(line) > 1 {
                    if headerMap, ok := requestMap["Headers"].(map[string][]string); ok {
                        line[0] = string(asRunes[1:])
                        line[0] = strings.TrimSpace(line[0])

                        headerMap[line[0]] = []string{line[1]}
                        parserMode = Normal
                        break
                    }
                }

                if isTag {
                    if headerMap, ok := requestMap["Headers"].(map[string][]string); ok {
                        line[0] = string(asRunes[1:])
                        line[0] = strings.TrimSpace(line[0])

                        headerMap[line[0]] = make([]string, 0)
                        break
                    }
                }

                if headerMap, ok := requestMap["Headers"].(map[string][]string); ok {
                    asRunes = asRunes[1:]
                    val := strings.TrimSpace(string(asRunes))
                    headerMap[currTag] = append(headerMap[currTag], val)
                }
			}
			break
		case MultiLineVal:
			break
		}
	}
}
