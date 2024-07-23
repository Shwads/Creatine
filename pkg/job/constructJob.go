package job

import (
	"errors"
	"fmt"
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

			if verbose, ok := requestMap["verbose"]; ok {
				if verbose, ok := verbose.(string); ok {
					nextJob.Verbose = strings.ToLower(verbose) == "true"
				} else {
					nextJob.Verbose = true
				}
			}

			if printToFile, ok := requestMap["file"]; ok {
				if printToFile, ok := printToFile.(string); ok {
					nextJob.PrintToFile = strings.ToLower(printToFile) == "true"
				} else {
					nextJob.PrintToFile = true
				}
			}

			if printToConsole, ok := requestMap["console"]; ok {
				if printToConsole, ok := printToConsole.(string); ok {
					nextJob.PrintToConsole = strings.ToLower(printToConsole) == "true"
				} else {
					nextJob.PrintToConsole = false
				}
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
