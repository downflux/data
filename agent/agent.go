package agent

import (
	"github.com/downflux/bvh/bvh"
	"github.com/downflux/go-bvh/id"
	"github.com/downflux/go-geometry/2d/hyperrectangle"
	"github.com/downflux/go-geometry/2d/vector"
)

type O struct {
	ID id.ID

	CollisionLayer bvh.Layer

	Radius   float64
	Position vector.V
}

type AgentCollider struct {
	bvh *bvh.BVH

	id id.ID

	l bvh.Layer

	p    vector.M
	r    float64
	aabb hyperrectangle.M
}

func New(bvh *bvh.BVH, o O) *AgentCollider {
	r := o.Radius
	x, y := o.Position.X(), o.Position.Y()
	a := &AgentCollider{
		id:   o.ID,
		l:    o.CollisionLayer,
		p:    vector.M{x, y},
		r:    r,
		aabb: hyperrectangle.New(vector.V{x - r, y - r}, vector.V{x + r, y + r}).M(),
	}
	a.bvh.Insert(a.ID(), a.l, a.aabb.R())

	return a
}

func (a *AgentCollider) ID() id.ID { return a.id }

func (a *AgentCollider) CollisionLayer() bvh.Layer { return a.l }
func (a *AgentCollider) SetCollisionLayer(l bvh.Layer) {
	a.l = l

	a.bvh.Remove(a.ID())
	a.bvh.Insert(a.ID(), l, a.aabb.R())
}

func (a *AgentCollider) Radius() float64    { return a.r }
func (a *AgentCollider) Position() vector.V { return a.p.V() }
func (a *AgentCollider) SetPosition(v vector.V) {
	a.p.Copy(v)

	x, y := v.X(), v.Y()
	a.aabb.Min().SetX(x - a.r)
	a.aabb.Min().SetY(y - a.r)
	a.aabb.Max().SetX(x + a.r)
	a.aabb.Max().SetY(y + a.r)

	a.bvh.Update(a.ID(), a.aabb.R())
}

func (a *AgentCollider) Close() error {
	a.bvh.Remove(a.ID())
	return nil
}
