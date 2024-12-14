package solution

type Result struct {
	ValidUpdates              [][]int
	InvalidUpdates            [][]int
	MiddlePageSumValidUpdates int

	FixedUpdates              [][]int
	MiddlePageSumFixedUpdates int
}

func NewResult() (result *Result) {
	return &Result{}
}
