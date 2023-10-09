package node

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/ghostsecurity/reaper/backend/workflow/transmission"
)

type StopNode struct {
	*base
	noInjections
}

func NewStop() *StopNode {
	return &StopNode{
		base: newBase(
			"Stop",
			TypeStop,
			false,
			NewVarStorage(
				Connectors{
					NewConnector("input", transmission.TypeAnyIn, true),
				},
				nil,
				nil,
			),
		),
	}
}

var ErrStopped = fmt.Errorf("stopped")

func (n *StopNode) Start(ctx context.Context, in <-chan Input, _ chan<- OutputInstance, output chan<- Output) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case _, ok := <-in:
		if !ok {
			return nil
		}
		return ErrStopped
	}
}
