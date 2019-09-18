package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
	"math/rand"
)


func dist(p, q mgl32.Vec2) float32 {
	return float32(math.Sqrt(math.Pow(float64(q[0] - p[0]), 2) + math.Pow(float64(q[1] - p[1]), 2)))
}

func limit(p mgl32.Vec2, lim float32) mgl32.Vec2{
	if p[0] > lim {
		p[0] = lim
	} else if p[0] < -lim {
		p[0] = -lim
	}
	if p[1] > lim {
		p[1] = lim
	} else if p[1] < -lim {
		p[1] = -lim
	}
	return p
}


var alignmentPerception = float32(20.0)
var cohesionPerception = float32(40.0)
var divisionPerception = float32(20.0)
var alignmentCoef = float32(1.0)
var cohesionCoef = float32(1.0)
var divisionCoef = float32(1.0)
var maxSpeed = float32(4)
var maxForce = float32(1)
var avg mgl32.Vec2
var total int


type Boid struct{
	siblings *[]Boid
	position mgl32.Vec2
	velocity mgl32.Vec2
	acceleration mgl32.Vec2
}


func NewBoid(x, y float32, boids *[]Boid) Boid {
	b := Boid{
		siblings: boids,
		position:mgl32.Vec2{x, y},
		velocity:mgl32.Vec2{(rand.Float32() * 3) - 1.5, (rand.Float32() * 3) - 1.5},
		acceleration:mgl32.Vec2{0,0},
	}
	return b
}


func (boid* Boid) alignment() mgl32.Vec2 {
	avg = mgl32.Vec2{0,0}
	total = 0

	for _, sibling := range *boid.siblings {
		if (sibling != *boid) && dist(boid.position, sibling.position) < alignmentPerception {
			avg = avg.Add(sibling.velocity)
			total++
		}
	}
	if total > 0 {
		avg = avg.Mul(1.0 / float32(total) * alignmentCoef)
		avg = avg.Normalize().Mul(maxSpeed)
		avg = avg.Sub(boid.velocity)
		avg = limit(avg, maxForce)
	}
	return avg
}


func (boid* Boid) cohesion() mgl32.Vec2 {
	avg = mgl32.Vec2{0,0}
	total = 0

	for _, sibling := range *boid.siblings {
		if (sibling != *boid) && dist(boid.position, sibling.position) < cohesionPerception {
			avg = avg.Add(sibling.position)
			total++
		}
	}
	if total > 0 {
		avg = avg.Mul(1.0 / float32(total) * cohesionCoef)
		avg = avg.Sub(boid.position)
		avg = avg.Normalize().Mul(maxSpeed)
		avg = avg.Sub(boid.velocity)
		avg = limit(avg, maxForce)
	}
	return avg
}


func (boid* Boid) divison() mgl32.Vec2 {
	avg = mgl32.Vec2{0,0}
	total = 0

	for _, sibling := range *boid.siblings {
		d := dist(boid.position, sibling.position)
		if (sibling != *boid) && d < divisionPerception {
			diff := boid.position.Sub(sibling.position).Mul(1/(d*d))
			avg = avg.Add(diff)
			total++
		}
	}
	if total > 0 {
		avg = avg.Mul(1.0 / float32(total) * divisionCoef)
		avg = avg.Normalize().Mul(maxSpeed)
		avg = avg.Sub(boid.velocity)
		avg = limit(avg, maxForce)
	}
	return avg
}


func (boid* Boid) Tick() {
	boid.acceleration = boid.acceleration.Add(boid.alignment().Mul(1.5))
	boid.acceleration = boid.acceleration.Add(boid.cohesion().Mul(1.0))
	boid.acceleration = boid.acceleration.Add(boid.divison().Mul(2.0))

	boid.position = boid.position.Add(boid.velocity)
	boid.velocity = boid.velocity.Add(boid.acceleration.Mul(0.5))
	boid.velocity = limit(boid.velocity, maxSpeed)
	boid.acceleration = boid.acceleration.Mul(0)

	if float64(boid.position.X()) > width {
		boid.position[0] = 0
	} else if float64(boid.position.X()) < 0 {
		boid.position[0] = float32(width)
	}
	if float64(boid.position.Y()) > height {
		boid.position[1] = 0
	} else if float64(boid.position.Y()) < 0 {
		boid.position[1] = float32(height)
	}


}