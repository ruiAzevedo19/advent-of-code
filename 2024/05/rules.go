package solution

import (
	"slices"
)

type Rule struct {
	X int
	Y int
}

func NewRule(x int, y int) (rule *Rule) {
	return &Rule{
		X: x,
		Y: y,
	}
}

type Rules struct {
	edges map[int][]int
}

func NewRules(rules []*Rule) (r *Rules) {
	r = &Rules{
		edges: map[int][]int{},
	}

LOOP:
	for _, rule := range rules {
		nodes, ok := r.edges[rule.X]
		if !ok {
			r.edges[rule.X] = []int{rule.Y}

			continue
		}

		// Check if there is already an edge with the same destination node.
		for _, node := range nodes {
			if node == rule.Y {
				continue LOOP
			}
		}

		r.edges[rule.X] = append(nodes, rule.Y)
	}

	return r
}

func (r *Rules) AddRule(x int, y int) {
	if _, ok := r.edges[x]; !ok {
		r.edges[x] = []int{y}
	} else {
		r.edges[x] = append(r.edges[x], y)
	}
}

func (r *Rules) UpdateIsValid(update []int) (ok bool) {
	for i, value := range update {
		nodes := r.edges[value]
		for j := 0; j < i; j++ {
			if slices.Contains(nodes, update[j]) {
				return false
			}
		}
	}

	return true
}

func (r *Rules) FixUpdate(update []int) (fixedUpdate []int) {
	fixedUpdate = slices.Clone(update)

	for !r.UpdateIsValid(fixedUpdate) {
		for i, value := range fixedUpdate {
			after := r.edges[value]
			for _, v := range after {
				iv := slices.Index(fixedUpdate[:i], v)
				if iv < 0 {
					continue
				}
				fixedUpdate[i], fixedUpdate[iv] = fixedUpdate[iv], fixedUpdate[i]
			}
		}
	}

	return fixedUpdate
}
