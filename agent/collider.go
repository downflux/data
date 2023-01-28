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

type Collider struct {
	bvh *bvh.BVH

	id id.ID

	l bvh.Layer

	p    vector.M
	r    float64
	aabb hyperrectangle.M
}

func New(bvh *bvh.BVH, o O) *Collider {
	r := o.Radius
	x, y := o.Position.X(), o.Position.Y()
	a := &Collider{
		id:   o.ID,
		l:    o.CollisionLayer,
		p:    vector.M{x, y},
		r:    r,
		aabb: hyperrectangle.New(vector.V{x - r, y - r}, vector.V{x + r, y + r}).M(),
	}
	a.bvh.Insert(a.ID(), a.l, a.aabb.R())

	return a
}

func (a *Collider) ID() id.ID { return a.id }

func (a *Collider) CollisionLayer() bvh.Layer { return a.l }
func (a *Collider) SetCollisionLayer(l bvh.Layer) {
	a.l = l

	a.bvh.Remove(a.ID())
	a.bvh.Insert(a.ID(), l, a.aabb.R())
}

func (a *Collider) Radius() float64    { return a.r }
func (a *Collider) Position() vector.V { return a.p.V() }
func (a *Collider) SetPosition(v vector.V) {
	a.p.Copy(v)

	x, y := v.X(), v.Y()
	a.aabb.Min().SetX(x - a.r)
	a.aabb.Min().SetY(y - a.r)
	a.aabb.Max().SetX(x + a.r)
	a.aabb.Max().SetY(y + a.r)

	a.bvh.Update(a.ID(), a.aabb.R())
}

func (a *Collider) Close() error {
	a.bvh.Remove(a.ID())
	return nil
}
