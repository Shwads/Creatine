package job

func Jobber(requests map[string]interface{}) error {

	jobList, jobListErr := ConstructJob(requests)
	if jobListErr != nil {
		return jobListErr
	}

	PrintJobList(jobList)

    count := 0

	for _, job := range jobList {
        count += 1
        job.SendRequest()
	}

	return nil
}
