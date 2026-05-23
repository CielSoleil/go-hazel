package main

func main() {
	jobs, err := exec.Command("atq", "-o", "%s").Output()

	if err != nil {
		slog.Error("failed to fetch jobs list", "err", err)
	}
}
