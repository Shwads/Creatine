package main

import (
	"Creatine/pkg/fileParser"
	"Creatine/pkg/job"
	"Creatine/pkg/scriptParser"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
    // FLAGS FOR CONSOLE PRINTING
	verbose := flag.Bool("v", false, "print all response content\n")
	printToFile := flag.Bool("pf", false, "write the output to a file\n")
	printToConsole := flag.Bool("pc", true, "Print the response to the console. Default true.\n")

	var method *string

	method = flag.String("m", "GET", "Used to specify the http request method.\n")

    // FLAGS FOR PARSING REQUEST SCRIPTS
    executeNonIdempotent := flag.Bool("ni", false, "exectute all requests including non-idempotent\n")

    var scriptName *string

    scriptName = flag.String("s", "", "Provide a request script file to construct requests from\n")

    // FLAGS FOR PARSING REQUEST FILES
    var fileName *string

    fileName = flag.String("f", "", "Provide a file to construct a request from.\n")

	flag.Parse()

    if len(*scriptName) > 0 {
        scriptParseErr := scriptParser.ParseScript(*executeNonIdempotent, *scriptName)
        if scriptParseErr != nil {
            os.Exit(1)
        }
        return
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

	*method = strings.ToUpper(*method)

	switch *method {
	case "GET":
		fmt.Println("GET method requested.")
		fmt.Println("There's nothing here yet m8, get to it")
		break
	case "POST":
		fmt.Println("POST method requested.")
		fmt.Println("There's nothing here yet m8, get to it")
		break
	case "PATCH":
		fmt.Println("PATCH method requested.")
		fmt.Println("There's nothing here yet m8, get to it")
		break
	case "PUT":
		fmt.Println("PUT method requested.")
		fmt.Println("There's nothing here yet m8, get to it")
		break
	case "DELETE":
		fmt.Println("DELETE method requested.")
		fmt.Println("There's nothing here yet m8, get to it")
		break
	default:
		fmt.Println("Please use one of the following supported methods:")
		fmt.Print("GET\nPOST\nPATCH\nPUT\nDELETE\n")
		return
	}

	suppliedArgs := flag.Args()

	if len(suppliedArgs) < 1 {
		fmt.Println("Please prove a URL to which to send the request.")
		return
	}

	url, urlParseErr := url.Parse(suppliedArgs[0])
	if urlParseErr != nil {
		fmt.Printf("Encountered error: %s\n", urlParseErr)
		fmt.Println("Please provide a valid URL")
		return
	}

	job := job.Job{
		Url:            url,
		Verbose:        *verbose,
		PrintToFile:    *printToFile,
		PrintToConsole: *printToConsole,
		Method:         *method,
	}

	job.ProcessJob()

}
