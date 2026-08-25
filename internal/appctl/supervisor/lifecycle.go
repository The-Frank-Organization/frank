// Package supervisor owns app-side worker and connector lifecycle state.
package supervisor

import "fmt"

type WorkerState string

const (
	WorkerAllocated  WorkerState = "ALLOCATED"
	WorkerSpawning   WorkerState = "SPAWNING"
	WorkerReady      WorkerState = "READY"
	WorkerLeased     WorkerState = "LEASED"
	WorkerRetiring   WorkerState = "RETIRING"
	WorkerFailed     WorkerState = "FAILED"
	WorkerTerminated WorkerState = "TERMINATED"
)

type Machine struct{ state WorkerState }

func NewMachine(initial WorkerState) *Machine { return &Machine{state: initial} }
func (machine *Machine) State() WorkerState   { return machine.state }

func (machine *Machine) Transition(next WorkerState) error {
	allowed := map[WorkerState]map[WorkerState]bool{
		WorkerAllocated: {WorkerSpawning: true, WorkerFailed: true},
		WorkerSpawning:  {WorkerReady: true, WorkerFailed: true},
		WorkerReady:     {WorkerLeased: true, WorkerFailed: true},
		WorkerLeased:    {WorkerRetiring: true, WorkerFailed: true},
		WorkerRetiring:  {WorkerTerminated: true},
		WorkerFailed:    {WorkerRetiring: true, WorkerTerminated: true},
	}
	if !allowed[machine.state][next] {
		return fmt.Errorf("supervisor: illegal worker transition %s -> %s", machine.state, next)
	}
	machine.state = next
	return nil
}
