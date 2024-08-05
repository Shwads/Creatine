package job

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
)

func ConstructJob(requests map[string]interface{}) ([]Job, error) {
	jobList := make([]Job, 0)
	requestNum := 0

	for request, _ := range requests {
		requestNum += 1
		if requestMap, ok := requests[request].(map[string]interface{}); ok {
			nextJob := Job{}
			nextJob.RequestNum = requestNum

			if title, ok := requestMap["title"]; ok {
				if title, ok := title.(string); ok {
					nextJob.Title = title
				} else {
					errString := fmt.Sprintf("did not find expected type string for title instead found %T\n", title)
					log.Printf(errString)
					return jobList, errors.New(errString)
				}

			} else {
                nextJob.Title = ""
            }

			if verbose, ok := requestMap["verbose"]; ok {
				if verbose, ok := verbose.(string); ok {
					nextJob.Verbose = strings.ToLower(verbose) == "true"
					fmt.Printf("Found val %s for verbosity\n", verbose)
				} else {
					errString := fmt.Sprintf("did not find expected type string for verbosity instead found %T\n", verbose)
					log.Printf(errString)
					return jobList, errors.New(errString)
				}
			} else {
				nextJob.Verbose = true
				fmt.Printf("Set verbose to %t\n", nextJob.Verbose)
			}

			if printToFile, ok := requestMap["file"]; ok {
				if printToFile, ok := printToFile.(string); ok {
					nextJob.PrintToFile = strings.ToLower(printToFile) == "true"
				} else {
					errString := fmt.Sprintf("did not find expected type string for verbosity instead found %T\n", printToFile)
					log.Printf(errString)
					return jobList, errors.New(errString)
				}

			} else {
				nextJob.PrintToFile = true
			}

			if printToConsole, ok := requestMap["console"]; ok {
				if printToConsole, ok := printToConsole.(string); ok {
					nextJob.PrintToConsole = strings.ToLower(printToConsole) == "true"
				} else {
					errString := fmt.Sprintf("did not find expected type string for verbosity instead found %T\n", printToConsole)
					log.Printf(errString)
					return jobList, errors.New(errString)
				}

			} else {
				nextJob.PrintToConsole = false
			}

			if urlString, ok := requestMap["url"]; ok {
				if confirmed, ok := urlString.(string); ok {
					var urlParseErr error

					nextJob.Url, urlParseErr = url.Parse(confirmed)
					if urlParseErr != nil {
						return jobList, urlParseErr
					}
				}
			} else {
				return jobList, errors.New("requests must contain a url tag")
			}

			if method, ok := requestMap["method"]; ok {
				if method, ok := method.(string); ok {
					nextJob.Method = method
				}
			} else {
				return jobList, errors.New(fmt.Sprintf("request %d must contain a method tag", requestNum))
			}

			if headerMap, ok := requestMap["headers"].(map[string][]string); ok {
				nextJob.ReqHeaders = headerMap
			}

			if body, ok := requestMap["body"]; ok {
				if body, ok := body.(string); ok {
					bodyBytes := []byte(body)
					nextJob.ReqBody = bodyBytes
				}
			}

			jobList = append(jobList, nextJob)

		} else {
			return jobList, errors.New(fmt.Sprintf("improperly formatted map object provided - expected type map[string]interface{} - found instead type %T", requests[request]))
		}
	}

	return jobList, nil
}
