package main

func main() {
	jobs, err := exec.Command("atq", "-o", "%s").Output()

	if err != nil {
		slog.Error("failed to fetch jobs list", "err", err)
	}

	jobsSplit := strings.Split(string(jobs), "\n")

	// Each resulting string looks like this
	// 11528   1779418800 a applejack

	re_job, err := regexp.Compile(`(\d+)\s+(\d+)\s+(\w+)\s+(\w+)`)

	if err != nil {
		slog.Error("failed to compile regex", "err", err)
	}
	// Parse each job and get its data
	for _, j := range jobsSplit {
		if re_job.MatchString(j) {
			matches := re_job.FindStringSubmatch(j)
			// matches = matches[1:]

			var id string = matches[1]
			var timestamp string = matches[2]
			// var queue string = matches[3]
			// var username string = matches[4]
			var command string = getJobInfo(id)

			cTimestamp, err := strconv.ParseInt(timestamp, 10, 64)

			if err != nil {
				slog.Error("failed to convert", "err", err)
			}

		}
	}
}
