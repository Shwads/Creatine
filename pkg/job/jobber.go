package job

import "fmt"

func Jobber(requests map[string]interface{}) error {

	jobList, jobListErr := ConstructJob(requests)
	if jobListErr != nil {
		return jobListErr
	}

	PrintJobList(jobList)

    count := 0

	for _, job := range jobList {
        fmt.Printf("Sending request: %d\n", count)
        count += 1
        job.SendRequest()
                    
        fmt.Printf("Returned after sending request\n")
	}

	return nil
}
