package algorithms

import (
	"errors"
	"sort"
)

type Job struct {
	id       string
	priority int
	deadline int
}

func jobSequence(jobs []Job) (int, error) {

	for _, job := range jobs {
		if job.priority < 0 || job.deadline < 0 {
			return 0, errors.New("priority or deadline cannot be empty")
		}
	}

	max := 0
	sort.Slice(jobs, func(i, j int) bool {
		if max < jobs[i].deadline {
			max = jobs[i].deadline
		}
		return jobs[i].priority > jobs[j].priority
	})

	var result []string
	for i := 0; i < max+1; i++ {
		result = append(result, "0")
	}
	profit := 0
	for i := 0; i < len(jobs); i++ {
		if max < 0 {
			break
		}
		idx := findPlaceForJob(result, jobs[i].deadline)
		if idx != -1 {
			result[idx] = jobs[i].id
			profit += jobs[i].priority
			max--
		}
	}
	return profit, nil
}

func findPlaceForJob(result []string, limit int) int {
	for i := limit; i > 0; i-- {
		if result[i] == "0" {
			return i
		}
	}
	return -1
}
