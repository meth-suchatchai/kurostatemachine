package kurostatemachine

import "testing"

var transitions = map[State][]State{
	"TODO":      {"DOING"},
	"DOING":     {"FAILED", "COMPLETED"},
	"FAILED":    {"DOING"},
	"COMPLETED": {},
}

var sm = NewStateMachine(transitions)

func TestStateMachine_ValidTransitions(t *testing.T) {
	item := &Item{TaskID: 1, State: "TODO"}
	// TODO → DOING
	if err := sm.Transition(item, "DOING"); err != nil {
		t.Errorf("expected valid transition, got error: %v", err)
	}
	if item.State != "DOING" {
		t.Errorf("expected state DOING, got %s", item.State)
	}

	// DOING → FAILED
	if err := sm.Transition(item, "FAILED"); err != nil {
		t.Errorf("expected valid transition, got error: %v", err)
	}
	if item.State != "FAILED" {
		t.Errorf("expected state FAILED, got %s", item.State)
	}

	// FAILED → DOING
	if err := sm.Transition(item, "DOING"); err != nil {
		t.Errorf("expected valid transition, got error: %v", err)
	}
	if item.State != "DOING" {
		t.Errorf("expected state DOING, got %s", item.State)
	}

	// DOING → COMPLETED
	if err := sm.Transition(item, "COMPLETED"); err != nil {
		t.Errorf("expected valid transition, got error: %v", err)
	}
	if item.State != "COMPLETED" {
		t.Errorf("expected state COMPLETED, got %s", item.State)
	}
}

func TestSuccessNewStateMachine(t *testing.T) {
	item := &Item{TaskID: 1, State: "TODO"}
	err := sm.Transition(item, "COMPLETED")
	if err == nil {
		t.Errorf("expected error on invalid transition, got none")
	}
	if item.State != "TODO" {
		t.Errorf("expected state to remain TODO, got %s", item.State)
	}
}
