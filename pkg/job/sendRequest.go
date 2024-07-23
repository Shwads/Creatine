package job

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

	//data, dataReadErr := io.ReadAll(res.Body)
	//if dataReadErr != nil {
		//log.Printf("Encountered error: %s. In function 'SendRequest'.", dataReadErr)
		//return dataReadErr
	//}

    if dirCreateErr := os.Mkdir("responses", os.ModePerm); dirCreateErr != nil && dirCreateErr.Error() != "mkdir responses: file exists" {
        log.Printf("Encountered error: %s. Upon attempted directory creation\n", dirCreateErr)
        return dirCreateErr
    }

	file, fileCreateErr := os.Create(fmt.Sprintf("responses/Request-%d:%s.txt", job.RequestNum, job.Method))
	if fileCreateErr != nil {
		log.Printf("Encountered error: %s\n", fileCreateErr)
		return fileCreateErr
	}
	defer file.Close()

    statusString := fmt.Sprintf("Status: %s\n\n", job.Res.Status)
    _, fileWriteErr := file.Write([]byte(statusString))
    if fileWriteErr != nil {
        log.Printf("Encountered error: %s. When attempting to write status to file.\n", fileWriteErr)
        return fileWriteErr
    }

	if res.Header != nil {
		_, fileWriteErr := file.Write([]byte("Headers:\n"))
		if fileWriteErr != nil {
			log.Printf("Encountered error: %s. When writing 'Headers:' to file.\n", fileWriteErr)
		}

		for header, headerList := range res.Header {
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

    reader := bufio.NewReader(res.Body)

    file.Write([]byte("Body:\n"))

    for {
        bytesRead, readerErr := reader.Read(buffer);

        file.Write(buffer[:bytesRead])

        if readerErr != nil {
            if readerErr != io.EOF {
                log.Printf("Encountered error: %s. When reading http response body", readerErr)
                return readerErr
            }
            break
        }
    }

	//_, fileWriteErr = file.Write([]byte("hi there\n"))
	//if fileWriteErr != nil {
		//fmt.Printf("Encountered error: %s. When trying to write to the file", fileWriteErr)
	//}

	return nil
}
