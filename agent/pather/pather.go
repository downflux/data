package pather

import (
	"github.com/downflux/data/agent/collider"
	"github.com/downflux/go-geometry/2d/vector"
	"github.com/downflux/go-geometry/2d/vector/polar"
)

type O struct {
	Collider collider.O

	Heading        polar.V
	TargetPosition vector.V
	Velocity       vector.V
	TargetVelocity vector.V

	MaxVelocity        float64
	MaxAngularVelocity float64
	MaxAcceleration    float64
}

type P struct {
	*collider.C

	heading        polar.M
	targetPosition vector.M
	velocity       vector.M
	targetVelocity vector.M

	maxVelocity        float64
	maxAngularVelocity float64
	maxAcceleration    float64
}

func New(o O) *P {
	p := &P{
		C: collider.New(o.Collider),

		heading:        polar.M{0, 0},
		targetPosition: vector.M{0, 0},
		velocity:       vector.M{0, 0},
		targetVelocity: vector.M{0, 0},

		maxVelocity:        o.MaxVelocity,
		maxAngularVelocity: o.MaxAngularVelocity,
		maxAcceleration:    o.MaxAcceleration,
	}

	p.heading.Copy(o.Heading)
	p.targetPosition.Copy(o.TargetPosition)
	p.velocity.Copy(o.Velocity)
	p.targetVelocity.Copy(o.TargetVelocity)

	return p
}

func (p *P) MaxVelocity() float64         { return p.maxVelocity }
func (p *P) MaxAngularVelocity() float64  { return p.maxAngularVelocity }
func (p *P) MaxAcceleration() float64     { return p.maxAcceleration }
func (p *P) TargetPosition() vector.V     { return p.targetPosition.V() }
func (p *P) Velocity() vector.V           { return p.velocity.V() }
func (p *P) TargetVelocity() vector.V     { return p.targetVelocity.V() }
func (p *P) Heading() polar.V             { return p.heading.V() }
func (p *P) SetTargetPosition(v vector.V) { p.targetPosition.Copy(v) }
func (p *P) SetVelocity(v vector.V)       { p.velocity.Copy(v) }
func (p *P) SetTargetVelocity(v vector.V) { p.targetVelocity.Copy(v) }
func (p *P) SetHeading(v polar.V)         { p.heading.Copy(v) }
