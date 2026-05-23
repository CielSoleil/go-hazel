package main
func getJobInfo(jobId string) string {
	jobInfo, err := exec.Command("at", "-c", jobId).Output()

	if err != nil {
		slog.Error("failed to fetch job data", "err", err)
	}

	jobInfoSplit := strings.Split(string(jobInfo), "\n")

	jobInfoSplit = slices.DeleteFunc(jobInfoSplit, func(e string) bool {
		return e == ""
	})

	commandIssued := jobInfoSplit[len(jobInfoSplit)-1]

	// fmt.Printf("Last line is \"%s\"\n", commandIssued)

	return commandIssued

}

func main() {
	outFormat := flag.String("format", "cont", "Either \"org\" or \"cont\"")
	flag.Parse()

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

	var parsedJobs [][]any

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

			data := []any{id, cTimestamp, command}
			parsedJobs = append(parsedJobs, data)
		}
	}

	// Sorter kindly given by Gemini
	slices.SortFunc(parsedJobs, func(a, b []any) int {
		valA := a[1].(int64)
		valB := b[1].(int64)

		if valA < valB {
			return -1
		}
		if valA > valB {
			return 1
		}
		return 0
	})

	re_invalid, err := regexp.Compile(`^\d+`)

	if err != nil {
		slog.Error("failed to compile regex", "err", err)
	}
	for _, p := range parsedJobs {
		// Time conversion
		t := time.Unix(p[1].(int64), 0)

		// Load local timezone
		loc, err := time.LoadLocation("America/Mexico_City")

		if err != nil {
			slog.Error("fail to load location", "err", err)
		}

		// Load Tokyo timezone
		tokyo, err := time.LoadLocation("Asia/Tokyo")

		if err != nil {
			slog.Error("fail to load location", "err", err)
		}

		timeLocal := t.In(loc).Format("06 Mon 02 15:04")
		timeTokyo := t.In(tokyo).Format("06 Mon 02 15:04")

		if !re_invalid.MatchString(p[len(p)-1].(string)) {
			fmt.Printf("|%s|%s|%s|%s||\n", p[0], timeTokyo, timeLocal, p[len(p)-1])
		}
	}
}
