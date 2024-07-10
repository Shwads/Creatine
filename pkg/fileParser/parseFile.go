/*
This is incredibly hacky. Should've just used and tweaked someone elses yaml parser but I wanted to do it myself. Maybes you can write a better one in a future project
but for now this will have to do.
*/

package fileParser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var multiRequest bool = false

func headerParser(fileScanner *bufio.Scanner, requests map[string]interface{}) {
}

func bodyParser(fileScanner *bufio.Scanner, requests map[string]interface{}) {
}

func ParseFile(fileName string) (map[string]interface{}, error) {
    file, fileOpenErr := os.Open(fileName)
    if fileOpenErr != nil {
        log.Printf("Encountered Error: %s. In function 'Parsefile'.", fileOpenErr)
        return nil, fileOpenErr
    }
    defer file.Close()

    fileScanner := bufio.NewScanner(file)

    var requestTag string
    requestNum := 0

    requests := make(map[string]interface{})

    for fileScanner.Scan() {
        fmt.Printf("%s\n", fileScanner.Text())

        line := strings.Split(fileScanner.Text(), ":")

        line[0] = strings.TrimSpace(line[0])

        if len(line) > 1 {
            line[1] = strings.TrimSpace(line[1])
            if multiRequest {
                if requestMap, ok := requests[requestTag].(map[string]interface{}); ok {
                    requestMap[line[0]] = line[1]
                }
            }
        }

        switch line[0] {
        case "request":
            requestNum += 1
            multiRequest = true
            requestTag := fmt.Sprintf("request-%d", requestNum)
            requests[requestTag] = make(map[string]interface{})
            break

        case "headers":
            if multiRequest {
                if requestMap, ok := requests[requestTag].(map[string]interface{}); ok {
                    headerParser(fileScanner, requestMap)
                }
            } else {
                headerParser(fileScanner, requests)
            }
            break

        case "body":
            if multiRequest {
                if requestMap, ok := requests[requestTag].(map[string]interface{}); ok {
                    bodyParser(fileScanner, requestMap)
                }
            } else {
                bodyParser(fileScanner, requests)
            }
            break
        }
    }

    return requests, nil
}
