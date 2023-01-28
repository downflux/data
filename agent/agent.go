package agent

import (
	"github.com/downflux/data/agent/pather"
)

type O struct {
	Pather pather.O
}

type A struct {
	*pather.P
}

func New(o O) *A {
	return &A{
		P: pather.New(o.Pather),
	}
}
