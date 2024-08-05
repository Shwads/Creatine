package job

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func (job Job) printToConsole() error {

    fmt.Println("=======================================================================")
    fmt.Println("=======================================================================")

	if job.PrintToFile {
		file, fileOpenErr := os.Open(fmt.Sprintf("responses/Request-%d:%s.txt", job.RequestNum, job.Method))
		if fileOpenErr != nil {
			log.Printf("Encountered error: %s. On attempted file open.\n", fileOpenErr)
			return fileOpenErr
		}
		defer file.Close()

		buffer := make([]byte, 256)
		reader := bufio.NewReader(file)

		for {
			numBytesRead, readErr := reader.Read(buffer)
			fmt.Print(string(buffer[:numBytesRead]))

			if readErr != nil {
				if readErr != io.EOF {
					log.Printf("Encountered error: %s. On attempting file read.", readErr)
					return readErr
				}
                break
			}
		}

        fmt.Println()
		return nil
	}

	fmt.Printf("\nRequest %d: %s %s\n\n", job.RequestNum, job.Method, job.Url)
	fmt.Println("\tStatus:")
	fmt.Printf("\t\t%s\n", job.Res.Status)

	if job.Verbose {
		if job.Res.Header != nil {
			fmt.Print("\n\tHeaders:\n")
			for header, headerList := range job.Res.Header {
				fmt.Printf("\t\t%s:", header)
				for _, val := range headerList {
					fmt.Printf(" %s;", val)
				}
				fmt.Println()
			}
			fmt.Println()
		}

		buffer := make([]byte, 256)

		reader := bufio.NewReader(job.Res.Body)

		fmt.Print("\n\tBody:\n\n")

		for {
			bytesRead, readerErr := reader.Read(buffer)

			fmt.Print(string(buffer[:bytesRead]))

			if readerErr != nil {
				if readerErr != io.EOF {
					log.Printf("Encountered error: %s. On attempted read of response body.\n", readerErr)
					return readerErr
				}
				break
			}
		}

		fmt.Println()
	}

    fmt.Println()
    fmt.Println("=======================================================================")
    fmt.Println("=======================================================================")
    fmt.Println()


	return nil
}
