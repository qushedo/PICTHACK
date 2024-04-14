package states

import "sync"

type InputMap struct {
	Mx  sync.RWMutex
	Map map[int64]int
}

const (
	WaitingForUniversity = 1
	WaitingForFaculty    = 2
)

func MakeInputMap() *InputMap {
	return &InputMap{Map: make(map[int64]int), Mx: sync.RWMutex{}}
}
