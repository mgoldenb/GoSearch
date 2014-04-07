package pancake

import (
	//"fmt"
	"GoSearch/Algorithms"
	"GoSearch/Helpful"
)

type PancakeState []int

func (state PancakeState) CopyState() algorithms.State {
	return append(PancakeState(nil), state...)
}

type PancakeMove int

type Pancake struct{}

func (Pancake) MakeTestState() algorithms.State {
	return PancakeState{0, 5, 4, 7, 2, 6, 1, 3}
}

func (Pancake) Heuristic(state algorithms.State) (res int) {
	myState := state.(PancakeState)
	for i := 1; i < len(myState); i++ {
		if helpful.Abs(myState[i]-myState[i-1]) > 1 {
			res++
		}
	}
	if myState[0] != 0 {
		res++
	}
	return
}

func (Pancake) BF(state algorithms.State) int {
	return len(state.(PancakeState)) - 1
}

func (Pancake) GetMove(i int) (res algorithms.Move) {
	return (PancakeMove)(i + 2)
}

func ReverseSlice(slice []int) {
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-i-1] = slice[len(slice)-i-1], slice[i]
	}
}

func (Pancake) MakeMove(state algorithms.State, move algorithms.Move) {
	ReverseSlice(state.(PancakeState)[:move.(PancakeMove)])
}

func (domain Pancake) UnmakeMove(state algorithms.State, move algorithms.Move) {
	domain.MakeMove(state, move)
}
