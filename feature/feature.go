package feature

import (
	"github.com/downflux/data/feature/collider"
)

type O struct {
	Collider collider.O
}

type F struct {
	*collider.C
}

func New(o O) *F {
	return &F{
		C: collider.New(o.Collider),
	}
}
