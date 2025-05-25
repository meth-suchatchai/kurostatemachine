package kurostm

import "fmt"

type State string

type Item struct {
	TaskID int
	State  State
}

type StateMachine struct {
	validTransitions map[State][]State
}

func NewStateMachine(transitions map[State][]State) *StateMachine {
	return &StateMachine{
		validTransitions: transitions,
	}
}

func (sm *StateMachine) CanTransition(item *Item, to State) bool {
	for _, allowed := range sm.validTransitions[item.State] {
		if allowed == to {
			return true
		}
	}
	return false
}

func (sm *StateMachine) Transition(item *Item, to State) error {
	if sm.CanTransition(item, to) {
		item.State = to
		return nil
	}
	return fmt.Errorf("invalid transition from %s to %s failed", item.State, to)
}
