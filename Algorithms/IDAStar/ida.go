package IDAStar

import (
	"GoSearch/Algorithms"
	"fmt"
	"strings"
)

type Search struct {
	Solution  []algorithms.State
	Threshold int
	Verbose   bool
	// all parameters are also here
}

func (alg Search) DF(domain algorithms.Domain, state algorithms.State, depth int) bool {
	if depth == 0 {
		fmt.Printf("Threshold %d\n", alg.Threshold)
	}
	if domain.Heuristic(state) == 0 {
		alg.Solution[depth] = state.CopyState()
		return true
	}
	if depth == alg.Threshold {
		return false
	}
	if alg.Verbose {
		fmt.Println(strings.Repeat(" ", depth*4), state)
	}
	for i := 0; i < domain.BF(state); i++ {
		move := domain.GetMove(i)
		domain.MakeMove(state, move)
		if alg.DF(domain, state, depth+1) {
			domain.UnmakeMove(state, move)
			alg.Solution[depth] = state.CopyState()
			return true
		}
		domain.UnmakeMove(state, move)
	}
	return false
}

func (alg Search) ID(domain algorithms.Domain, start algorithms.State) {
	for alg.Threshold = 0; true; alg.Threshold++ {
		alg.Solution = make([]algorithms.State, alg.Threshold+1, alg.Threshold+1)
		if alg.DF(domain, start, 0) {
			break
		}
	}

	for i := 0; i < len(alg.Solution); i++ {
		fmt.Println(alg.Solution[i])
	}
}
