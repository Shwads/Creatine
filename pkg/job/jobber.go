package job

func Jobber(requests map[string]interface{}) error {

	jobList, jobListErr := ConstructJob(requests)
	if jobListErr != nil {
		return jobListErr
	}

	PrintJobList(jobList)

	for _, job := range jobList {
		job.SendRequest()
	}

	return nil
}
