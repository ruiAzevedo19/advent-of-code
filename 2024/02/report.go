package solution

import (
	"math"
)

type Report struct {
	values []int
}

func NewReport(values []int) (report *Report) {
	return &Report{
		values: values,
	}
}

func (r *Report) IsSafe() (ok bool) {
	if len(r.values) < 2 {
		return true
	}

	// If sign is 1 then the report must be in an asceding order, if -1 then the array must be in descending order.
	var sign int
	if r.values[0] < r.values[1] {
		sign = 1
	} else if r.values[0] > r.values[1] {
		sign = -1
	} else {
		return false
	}

	for i := 0; i < len(r.values)-1; i++ {
		x, y := r.values[i], r.values[i+1]

		// Two adjacent report values must differ at least one, and at most three.
		diff := math.Abs(float64(x) - float64(y))
		if diff < 1 || diff > 3 {
			return false
		}

		// The report must be in ascending or descending order.
		if x > y && sign > 0 || x < y && sign < 0 {
			return false
		}
	}

	return true
}

func (r *Report) CanBeSafe() (ok bool) {
	for i := range r.values {
		newReport := &Report{}
		newReport.values = append(newReport.values, r.values[:i]...)
		newReport.values = append(newReport.values, r.values[i+1:]...)

		if newReport.IsSafe() {
			return true
		}
	}

	return false
}
