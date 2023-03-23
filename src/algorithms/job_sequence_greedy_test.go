package algorithms

import "testing"

func TestJobSequence(t *testing.T) {
	var jobs []Job
	jobs = append(jobs, Job{id: "J1", priority: 20, deadline: 2})
	jobs = append(jobs, Job{id: "J2", priority: 15, deadline: 2})
	jobs = append(jobs, Job{id: "J3", priority: 10, deadline: 1})
	jobs = append(jobs, Job{id: "J4", priority: 5, deadline: 3})
	jobs = append(jobs, Job{id: "J5", priority: 1, deadline: 3})

	result, err := jobSequence(jobs)
	expected := 40
	if err != nil {
		t.Errorf("error should be nil, but %v", err)
	}
	if result != expected {
		t.Errorf("total profit is mismatch. actual %d expected %d", result, expected)
	}
}

func TestJobSequence_Another(t *testing.T) {
	var jobs []Job
	jobs = append(jobs, Job{id: "J1", priority: 35, deadline: 3})
	jobs = append(jobs, Job{id: "J2", priority: 30, deadline: 4})
	jobs = append(jobs, Job{id: "J3", priority: 25, deadline: 4})
	jobs = append(jobs, Job{id: "J4", priority: 20, deadline: 2})
	jobs = append(jobs, Job{id: "J5", priority: 15, deadline: 3})
	jobs = append(jobs, Job{id: "J5", priority: 12, deadline: 1})
	jobs = append(jobs, Job{id: "J5", priority: 5, deadline: 2})

	result, err := jobSequence(jobs)
	expected := 110

	if err != nil {
		t.Errorf("error should be nil, but %v", err)
	}

	if result != expected {
		t.Errorf("total profit is mismatch. actual %d expected %d", result, expected)
	}
}

func TestJobSequence_NegativePriority(t *testing.T) {
	var jobs []Job
	jobs = append(jobs, Job{id: "J1", priority: 20, deadline: 2})
	jobs = append(jobs, Job{id: "J2", priority: -5, deadline: 2})
	jobs = append(jobs, Job{id: "J3", priority: 10, deadline: 1})
	jobs = append(jobs, Job{id: "J4", priority: 5, deadline: 3})
	jobs = append(jobs, Job{id: "J5", priority: 1, deadline: 3})

	_, err := jobSequence(jobs)
	if err == nil {
		t.Errorf("error is expected, but got nil")
	}
}

func TestJobSequence_NegativeDeadline(t *testing.T) {
	var jobs []Job
	jobs = append(jobs, Job{id: "J1", priority: 20, deadline: 2})
	jobs = append(jobs, Job{id: "J2", priority: -5, deadline: 2})
	jobs = append(jobs, Job{id: "J3", priority: 10, deadline: 1})
	jobs = append(jobs, Job{id: "J4", priority: 5, deadline: -3})
	jobs = append(jobs, Job{id: "J5", priority: 1, deadline: 3})

	_, err := jobSequence(jobs)
	if err == nil {
		t.Errorf("error is expected, but got nil")
	}
}
