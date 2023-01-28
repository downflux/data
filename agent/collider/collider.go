package collider

import (
	"github.com/downflux/bvh/bvh"
	"github.com/downflux/go-bvh/id"
	"github.com/downflux/go-geometry/2d/hyperrectangle"
	"github.com/downflux/go-geometry/2d/vector"
)

type O struct {
	BVH *bvh.BVH

	ID id.ID

	CollisionLayer bvh.Layer

	Radius   float64
	Position vector.V
}

type C struct {
	bvh *bvh.BVH

	id id.ID

	l bvh.Layer

	p    vector.M
	r    float64
	aabb hyperrectangle.M
}

func New(o O) *C {
	r := o.Radius
	x, y := o.Position.X(), o.Position.Y()
	c := &C{
		bvh:  o.BVH,
		id:   o.ID,
		l:    o.CollisionLayer,
		p:    vector.M{x, y},
		r:    r,
		aabb: hyperrectangle.New(vector.V{x - r, y - r}, vector.V{x + r, y + r}).M(),
	}
	c.bvh.Insert(c.ID(), c.l, c.aabb.R())

	return c
}

func (c *C) ID() id.ID              { return c.id }
func (c *C) Radius() float64        { return c.r }
func (c *C) AABB() hyperrectangle.R { return c.aabb.R() }

func (c *C) CollisionLayer() bvh.Layer { return c.l }
func (c *C) SetCollisionLayer(l bvh.Layer) {
	c.l = l

	c.bvh.Remove(c.ID())
	c.bvh.Insert(c.ID(), l, c.aabb.R())
}

// Position returns the current position of the agent.
//
// External callers must not set the position via the position reference object
// returned here.
func (c *C) Position() vector.V { return c.p.V() }
func (c *C) SetPosition(v vector.V) {
	c.p.Copy(v)

	x, y := v.X(), v.Y()
	c.aabb.Min().SetX(x - c.r)
	c.aabb.Min().SetY(y - c.r)
	c.aabb.Max().SetX(x + c.r)
	c.aabb.Max().SetY(y + c.r)

	c.bvh.Update(c.ID(), c.aabb.R())
}

func (c *C) Close() error {
	c.bvh.Remove(c.ID())
	return nil
}
