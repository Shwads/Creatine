package job

import (
	"fmt"
	"net/url"
	"testing"
)

func TestSendRequest(t *testing.T) {
	url, parseURLErr := url.Parse("https://blog.boot.dev/index.xml")
	if parseURLErr != nil {
		fmt.Printf("Encountered error: %s. In function 'TestSendRequest'.", parseURLErr)
		t.Fatal()
	}

	headers := make(map[string][]string)

	testJob := Job{
		Url:        url,
		Method:     "GET",
		ReqHeaders: headers,
	}

	testJob.SendRequest()
}
