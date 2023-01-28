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

type Collider struct {
	bvh *bvh.BVH

	id id.ID

	l bvh.Layer

	aabb hyperrectangle.M
}

func New(bvh *bvh.BVH, o O) *Collider {
	f := &Collider{
		bvh:  bvh,
		id:   o.ID,
		l:    o.CollisionLayer,
		aabb: hyperrectangle.New(vector.V{0, 0}, vector.V{0, 0}).M(),
	}
	f.aabb.Copy(o.AABB)
	f.bvh.Insert(f.ID(), f.l, f.aabb.R())

	return f
}

func (f *Collider) ID() id.ID              { return f.id }
func (f *Collider) AABB() hyperrectangle.R { return f.aabb.R() }

func (f *Collider) CollisionLayer() bvh.Layer { return f.l }
func (f *Collider) SetCollisionLayer(l bvh.Layer) {
	f.l = l

	f.bvh.Remove(f.ID())
	f.bvh.Insert(f.ID(), l, f.aabb.R())
}

func (f *Collider) Close() error {
	f.bvh.Remove(f.ID())
	return nil
}
