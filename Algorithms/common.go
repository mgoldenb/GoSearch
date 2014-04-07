package algorithms

type State interface {
	CopyState() State
}

type Move interface{}

type Domain interface {
	MakeTestState() State
	Heuristic(State) int
	BF(State) int
	GetMove(int) Move
	MakeMove(State, Move)
	UnmakeMove(State, Move)
}
