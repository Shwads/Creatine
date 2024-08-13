package main

import (
	"Creatine/pkg/yamlParser"
	"Creatine/pkg/job"
	"Creatine/pkg/scriptParser"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

const SCRIPTNAME = "requestScript.txt"

func main() {
	// FLAGS FOR CONSOLE PRINTING
	verbose := flag.Bool("v", false, "print all response content\n")
	printToFile := flag.Bool("pf", false, "write the output to a file\n")
	printToConsole := flag.Bool("pc", true, "Print the response to the console. Default true.\n")

	var method *string

	method = flag.String("m", "GET", "Used to specify the http request method.\n")

	// FLAGS FOR PARSING REQUEST SCRIPTS
	executeNonIdempotent := flag.Bool("ni", false, "exectute all requests including non-idempotent\n")

	parseScript := flag.Bool("s", false, "Read batches from requestScript.txt")

	// FLAGS FOR PARSING REQUEST FILES
	var fileName *string

	fileName = flag.String("f", "", "Provide a file to construct a request from.\n")

	flag.Parse()

    // Check for -s flag and read from the requestScript.txt
	if *parseScript {
		scriptParseErr := scriptParser.ParseScript(*executeNonIdempotent, SCRIPTNAME)
		if scriptParseErr != nil {
			os.Exit(1)
		}
	}

	if len(*fileName) > 0 {
		requests, parseFileErr := fileParser.ParseFile(*fileName)
		if parseFileErr != nil {
			log.Printf("%s\n", parseFileErr)
			return
		}

		jobberErr := job.Jobber(requests)
		if jobberErr != nil {
			log.Printf("Encountered Error: %s\n", jobberErr)
		}

		return
	}

    // Default request functionality
	*method = strings.ToUpper(*method)

	suppliedArgs := flag.Args()

	if len(suppliedArgs) < 1 {
		fmt.Println("Please prove a URL to which to send the request.")
        os.Exit(1)
	}

	url, urlParseErr := url.Parse(suppliedArgs[0])
	if urlParseErr != nil {
		fmt.Printf("Encountered error: %s\n", urlParseErr)
		fmt.Println("Please provide a valid URL")
        os.Exit(1)
	}

	job := job.Job{
		Url:            url,
		Verbose:        *verbose,
		PrintToFile:    *printToFile,
		PrintToConsole: *printToConsole,
		Method:         *method,
	}

    sendRequestErr := job.SendRequest()
    if sendRequestErr != nil {
        os.Exit(1)
    }
}
