package collider

import (
	"github.com/downflux/bvh/bvh"
	"github.com/downflux/go-bvh/id"
	"github.com/downflux/go-geometry/2d/hyperrectangle"
	"github.com/downflux/go-geometry/2d/vector"
)

type O struct {
	ID id.ID

	CollisionLayer bvh.Layer

	AABB hyperrectangle.R
}

type C struct {
	bvh *bvh.BVH

	id id.ID

	l bvh.Layer

	aabb hyperrectangle.M
}

func New(bvh *bvh.BVH, o O) *C {
	c := &C{
		bvh:  bvh,
		id:   o.ID,
		l:    o.CollisionLayer,
		aabb: hyperrectangle.New(vector.V{0, 0}, vector.V{0, 0}).M(),
	}
	c.aabb.Copy(o.AABB)
	c.bvh.Insert(c.ID(), c.l, c.aabb.R())

	return c
}

func (c *C) ID() id.ID              { return c.id }
func (c *C) AABB() hyperrectangle.R { return c.aabb.R() }

func (c *C) CollisionLayer() bvh.Layer { return c.l }
func (c *C) SetCollisionLayer(l bvh.Layer) {
	c.l = l

	c.bvh.Remove(c.ID())
	c.bvh.Insert(c.ID(), l, c.aabb.R())
}

func (c *C) Close() error {
	c.bvh.Remove(c.ID())
	return nil
}
