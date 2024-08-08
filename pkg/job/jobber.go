package job

func Jobber(requests map[string]interface{}, sendAll bool) error {

	jobList, jobListErr := ConstructJob(requests)
	if jobListErr != nil {
		return jobListErr
	}

	PrintJobList(jobList)

	for _, job := range jobList {
		if !sendAll {
			if job.Method == "POST" || job.Method == "PATCH" {
				continue
			}
		}
		job.SendRequest()
	}

	return nil
}
