package job

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func (job Job) writeToFile() error {
    if dirCreateErr := os.Mkdir("responses", os.ModePerm); dirCreateErr != nil && dirCreateErr.Error() != "mkdir responses: file exists" {
        log.Printf("Encountered error: %s. Upon attempted directory creation\n", dirCreateErr)
        return dirCreateErr
    }

    var requestTitle string

    if job.Title != "" {
        requestTitle = job.Title
    } else {
        requestTitle = fmt.Sprintf("Request-%d:%s", job.RequestNum, job.Method)
    }

	file, fileCreateErr := os.Create(fmt.Sprintf("responses/%s.txt", requestTitle))
	if fileCreateErr != nil {
		log.Printf("Encountered error: %s\n", fileCreateErr)
		return fileCreateErr
	}
	defer file.Close()

    nameString := fmt.Sprintf("\nRequest %d: %s %s\n\n", job.RequestNum, job.Method, job.Url)
    _, fileWriteErr := file.WriteString(nameString)
    if fileWriteErr != nil {
        log.Printf("Encountered error: %s. On atttempted file write\n", fileWriteErr)
        return fileWriteErr
    }

    statusString := fmt.Sprintf("Status: %s\n\n", job.Res.Status)
    _, fileWriteErr = file.Write([]byte(statusString))
    if fileWriteErr != nil {
        log.Printf("Encountered error: %s. When attempting to write status to file.\n", fileWriteErr)
        return fileWriteErr
    }

	if job.Res.Header != nil {
		_, fileWriteErr := file.Write([]byte("Headers:\n"))
		if fileWriteErr != nil {
			log.Printf("Encountered error: %s. When writing 'Headers:' to file.\n", fileWriteErr)
		}

		for header, headerList := range job.Res.Header {
			writeString := fmt.Sprintf("%s:", header)

			for _, val := range headerList {
				writeString = fmt.Sprintf("%s %s", writeString, val)
			}

            writeString = fmt.Sprintf("%s\n", writeString)

			_, fileWriteErr := file.Write([]byte(writeString))
			if fileWriteErr != nil {
				log.Printf("Encountered error: %s. When trying to write headers for header: %s to file\n", fileWriteErr, header)
				return fileWriteErr
			}
		}
        file.Write([]byte("\n"))
	}

    buffer := make([]byte, 256)

    reader := bufio.NewReader(job.Res.Body)

    file.Write([]byte("Body:\n"))

    for {
        bytesRead, readerErr := reader.Read(buffer)

        file.Write(buffer[:bytesRead])

        if readerErr != nil {
            if readerErr != io.EOF {
                log.Printf("Encountered error: %s. When reading http response body", readerErr)
                return readerErr
            }
            break
        }
    }

	return nil
}
