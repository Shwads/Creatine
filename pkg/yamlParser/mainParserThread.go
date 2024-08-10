package fileParser

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"strings"
)

// TODO: add safety checking for too many requests.
// Stagger them and do them in batches.

// Deal with requests in groups? Use a buffer send the requests, write the files 
// and then process the next group of requests in the file?

func mainParserThread(fileScanner *bufio.Scanner, requests map[string]interface{}) error {

    // We're using separate functions for the body and header tags to parse the file with state
    // because of that we need to track whether the current line has been processed
    // this avoids the instance where we exit the header or body parser, reach the end of the 
    // loop and scan a new line discarding the last. Only scan new line if the current has been
    // dealt with.
	var currLineProcessed bool
	var requestTag string

	requestNum := 0

	for {
		if strings.TrimSpace(fileScanner.Text()) == "" {
			fileScanner.Scan()
			continue
		}

		currLineProcessed = false

		line := strings.Split(fileScanner.Text(), ":")
		line[0] = strings.TrimSpace(line[0])

		switch line[0] {

		case "request":
			requestNum += 1
			requestTag = fmt.Sprintf("request-%d", requestNum)
			requests[requestTag] = make(map[string]interface{})
			currLineProcessed = true
			break

		case "headers":
			if requestMap, ok := requests[requestTag].(map[string]interface{}); ok {
				requestMap["headers"] = make(map[string][]string)

				anotherRequest, headerParseErr := headerParser(fileScanner, requestMap)
				if headerParseErr != nil {
					return headerParseErr
				}

				currLineProcessed = !anotherRequest
			}
			break

		case "body":
			if requestMap, ok := requests[requestTag].(map[string]interface{}); ok {

				if len(line) > 1 {
					line[1] = strings.Join(line[1:], ":")
					line[1] = strings.TrimSpace(line[1])

					if line[1] == "|" {
						requestMap["body"] = ""
						anotherRequest, bodyParseErr := bodyParser(true, fileScanner, requestMap)
						if bodyParseErr != nil {
							return bodyParseErr
						}

						// if there's another request (anotherRequest == true) then the current line
						// hasn't been processed
						currLineProcessed = !anotherRequest

					} else if line[1] == ">" {
						requestMap["body"] = ""
						anotherRequest, bodyParseErr := bodyParser(false, fileScanner, requestMap)
						if bodyParseErr != nil {
							return bodyParseErr
						}

						currLineProcessed = !anotherRequest
					} else {
						requestMap["body"] = line[1]
						currLineProcessed = true
					}

				} else {
					log.Println("tag 'body:' should be accompanied with a label")
				}

			}
			break
		}

		line = strings.Split(fileScanner.Text(), ":")
		line[0] = strings.TrimSpace(line[0])

		if len(line) > 1 && len(line) < 3 {
			if line[1] == "" {
				line = line[:1]
			}
		}

		// Deal with case we have a key and value on the same line, return an error if something goes wrong with
		// the map type.
		if len(line) > 1 {
			line[1] = strings.Join(line[1:], ":")
            line[1] = strings.TrimSpace(line[1])

			if requestMap, ok := requests[requestTag].(map[string]interface{}); ok {
				requestMap[line[0]] = line[1]
				currLineProcessed = true
			} else {
				// check in case requestTag has not been set
				if requestTag == "" {
					log.Printf("parser expects requests to begin with a 'request:' tag")
					return errors.New("parser expects requests to begin with a 'request:' tag")
				}

				log.Printf("requests[requestTag] did not have expected type map[string]interface{} instead found type %T.", requests)
				return errors.New(fmt.Sprintf("requests[requestTag] did not have expected type map[string]interface{} instead found type %T.", requests))
			}
		}

		// If the current line has been dealt with fetch the next line, if that's the EOF then return
		if currLineProcessed {
			if !fileScanner.Scan() {
				return nil
			}
		}
	}
}
