package job

import (
	"bytes"
	"log"
	"net/http"
)

func (job Job) SendRequest() error {
	req, newRequestErr := http.NewRequest(job.Method, job.Url.String(), bytes.NewReader(job.ReqBody))
	if newRequestErr != nil {
		log.Printf("Encountered error: %s. In function 'SendRequest'.", newRequestErr)
		return newRequestErr
	}

	for header, values := range job.ReqHeaders {
		for _, value := range values {
			req.Header.Add(header, value)
		}
	}

	client := &http.Client{}

	res, sendRequestErr := client.Do(req)
	if sendRequestErr != nil {
		log.Printf("Encountered error: %s. In function: 'SendRequest'.", sendRequestErr)
		return sendRequestErr
	}
	defer res.Body.Close()

	job.Res = res

	if job.PrintToFile {
		writeToFileErr := job.writeToFile()
		if writeToFileErr != nil {
			return writeToFileErr
		}
	}

	if job.PrintToConsole {
		printToConsoleErr := job.printToConsole()
		if printToConsoleErr != nil {
			return printToConsoleErr
		}
	}

	return nil
}
