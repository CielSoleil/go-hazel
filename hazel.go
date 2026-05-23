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
}
