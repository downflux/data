package collection

import (
	"fmt"

	"github.com/downflux/go-bvh/id"
)

type E interface {
	ID() id.ID

	Open() error
	Close() error
}

type M[T E] map[id.ID]T

func New[T E]() *M[T] {
	m := make(map[id.ID]T, 1024)
	return (*M[T])(&m)
}

func (m M[T]) Insert(e T) {
	x := e.ID()
	if _, ok := m[x]; ok {
		panic(fmt.Sprintf("cannot insert node: %v", x))
	}
	if err := e.Open(); err != nil {
		panic(fmt.Sprintf("cannot insert node: %v", err))
	}
	m[x] = e
}

func (m M[T]) Iterate() chan T {
	ch := make(chan T, 128)
	go func(ch chan<- T) {
		defer close(ch)
		for _, e := range m {
			ch <- e
		}
	}(ch)
	return ch

}

func (m M[T]) Get(x id.ID) T {
	if e, ok := m[x]; !ok {
		panic(fmt.Sprintf("cannot find node: %v", x))
	} else {
		return e
	}
}

func (m M[T]) Remove(x id.ID) {
	e, ok := m[x]
	if !ok {
		panic(fmt.Sprintf("cannot remove node: %v", x))
	}
	if err := e.Close(); err != nil {
		panic(fmt.Sprintf("cannot remove node: %v", err))
	}
	delete(m, x)
}
