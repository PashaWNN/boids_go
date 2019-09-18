package main

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"github.com/markfarnan/go-canvas/canvas"
	"image/color"
	"math/rand"
)

var boids []Boid

var done chan struct{}

var cvs *canvas.Canvas2d
var width float64
var height float64


func main() {
	cvs, _ = canvas.NewCanvas2d(true)
	height = float64(cvs.Height())
	width = float64(cvs.Width())
	for i := 0; i < 70; i++ {

		boids = append(boids, NewBoid(float32(rand.Intn(500)), float32(rand.Intn(700)), &boids))
	}
	cvs.Start(60, Render)
	<-done
}


func tick() {
	for i, _ := range boids {
		boids[i].Tick()
	}
}

func drawBoid(gc *draw2dimg.GraphicContext, b Boid) {
	gc.BeginPath()
	draw2dkit.Circle(gc, float64(b.position.X()), float64(b.position.Y()), 5)
	gc.FillStroke()
	gc.Close()

	x, y := float64(b.position[0]), float64(b.position[1])
	sec := b.position.Sub(b.velocity.Normalize().Mul(10))
	tX, tY := float64(sec[0]), float64(sec[1])
	gc.BeginPath()
	gc.MoveTo(x, y)
	gc.LineTo(tX, tY)
	gc.Stroke()
	gc.Close()

	//gc.BeginPath()
	//draw2dkit.Circle(gc, float64(b.position.X()), float64(b.position.Y()), 20)
	//gc.Stroke()
	//gc.Close()
}

func Render(gc *draw2dimg.GraphicContext) bool {
	tick()
	gc.SetFillColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	gc.Clear()

	gc.SetFillColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
	gc.SetStrokeColor(color.RGBA{0xff, 0x00, 0x00, 0xff})

	for _, boid := range boids {
		drawBoid(gc, boid)
	}

	return true
}
