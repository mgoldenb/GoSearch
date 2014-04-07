package main

import (
	"GoSearch/Domains/Pancake"
	"fmt"
	//"GoSearch/Algorithms"
	"GoSearch/Algorithms/Idastar"
)

func main() {
	var domain pancake.Pancake
	var search IDAStar.Search
	start := domain.MakeTestState()
	fmt.Println(start)
	search.Verbose = false
	search.ID(domain, start)
}
