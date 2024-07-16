package job

import "fmt"

func PrintJobList(jobList []Job) {
    fmt.Println("Printing Job List")

    for _, job := range jobList {
        fmt.Printf("url: %s\n", job.Url)  
        fmt.Printf("method: %s\n", job.Method)
        fmt.Printf("print to file: %t\n", job.PrintToFile)
        fmt.Printf("print to console: %t\n", job.PrintToConsole)
        fmt.Printf("verbose: %t\n", job.Verbose)

        for key, value := range job.ReqHeaders {
            fmt.Printf("%s: ", key)
            for _, val := range value {
                fmt.Printf("%s, ", val)
            }
            fmt.Print("\n")
        }
        fmt.Println()
    }
}
