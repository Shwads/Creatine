package job

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func (job Job) SendRequest() {
    fmt.Println("Beginning SendRequest")
	req, newRequestErr := http.NewRequest(job.Method, job.Url.String(), bytes.NewReader(job.ReqBody))
	if newRequestErr != nil {
		fmt.Printf("Encountered error: %s. In function 'SendRequest'.", newRequestErr)
		return
	}

    fmt.Println("Created Request")
	for header, values := range job.ReqHeaders {
		for _, value := range values {
			req.Header.Add(header, value)
		}
	}

    fmt.Println("parsed headers")
	client := &http.Client{}

    fmt.Println("Sending Request")
	res, sendRequestErr := client.Do(req)
	if sendRequestErr != nil {
		fmt.Printf("Encountered error: %s. In function: 'SendRequest'.", sendRequestErr)
		return
	}
	defer res.Body.Close()

    fmt.Println("Sent request")

	data, dataReadErr := io.ReadAll(res.Body)
	if dataReadErr != nil {
		fmt.Printf("Encountered error: %s. In function 'SendRequest'.", dataReadErr)
		return
	}

    fmt.Println("Read data")

    fmt.Printf("%s", string(data))
}
