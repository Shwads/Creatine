package job

import (
	"net/http"
	"net/url"
)

type Job struct {
	Url            *url.URL
	Verbose        bool
	PrintToFile    bool
	PrintToConsole bool
	Method         string
	ReqHeaders     map[string][]string
	ReqBody        []byte
	Res            *http.Response
}
