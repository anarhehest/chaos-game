package main

import (
	"flag"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 800
)

var (
	N              = flag.Int("n", 3, "number of polygon vertices (N)")
	rFactor        = flag.Float64("r", 0.5, "move fraction towards chosen vertex (0..1)")
	pointsPerFrame = flag.Int("ppf", 1000, "points drawn per frame")
)

type Point struct {
	X float32
	Y float32
}

type Game struct {
	verts         []Point
	pos           Point
	img           *ebiten.Image
	rng           *rand.Rand
	init          bool
	colorByVertex bool
}

func (g *Game) Init() {
	flag.Parse()
	if *N < 3 {
		*N = 3
	}
	cx, cy := float64(screenWidth)/2, float64(screenHeight)/2
	radius := math.Min(cx, cy) - 40

	g.verts = make([]Point, *N)
	for i := 0; i < *N; i++ {
		ang := 2*math.Pi*float64(i)/float64(*N) - math.Pi/2
		g.verts[i] = Point{
			X: float32(cx + radius*math.Cos(ang)),
			Y: float32(cy + radius*math.Sin(ang)),
		}
	}

	g.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	g.pos = Point{
		X: float32(g.rng.Float64() * float64(screenWidth)),
		Y: float32(g.rng.Float64() * float64(screenHeight)),
	}
	g.img = ebiten.NewImage(screenWidth, screenHeight)
	g.img.Fill(color.RGBA{10, 10, 10, 255})
	g.init = true
	g.colorByVertex = true
}

func (g *Game) Update() error {
	if !g.init {
		g.Init()
	}
	r := float32(*rFactor)
	if r <= 0 {
		r = 0.5
	}
	for i := 0; i < *pointsPerFrame; i++ {
		idx := g.rng.Intn(len(g.verts))
		v := g.verts[idx]
		g.pos.X = g.pos.X + (v.X-g.pos.X)*r
		g.pos.Y = g.pos.Y + (v.Y-g.pos.Y)*r

		x := int(g.pos.X + 0.5)
		y := int(g.pos.Y + 0.5)
		if x < 0 || x >= screenWidth || y < 0 || y >= screenHeight {
			continue
		}

		col := color.RGBA{220, 220, 220, 255}
		if g.colorByVertex {
			switch idx % 6 {
			case 0:
				col = color.RGBA{255, 100, 100, 255}
			case 1:
				col = color.RGBA{100, 255, 100, 255}
			case 2:
				col = color.RGBA{100, 100, 255, 255}
			case 3:
				col = color.RGBA{255, 200, 50, 255}
			case 4:
				col = color.RGBA{200, 100, 255, 255}
			case 5:
				col = color.RGBA{100, 255, 255, 255}
			}
		}
		g.img.Set(x, y, col)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.img, &ebiten.DrawImageOptions{})
	// draw vertex markers
	for _, v := range g.verts {
		drawCircle(screen, int(v.X), int(v.Y), 3, color.RGBA{255, 255, 255, 255})
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func drawCircle(img *ebiten.Image, cx, cy, r int, col color.RGBA) {
	minX := cx - r
	maxX := cx + r
	minY := cy - r
	maxY := cy + r
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			dx := x - cx
			dy := y - cy
			if dx*dx+dy*dy <= r*r {
				if x >= 0 && x < screenWidth && y >= 0 && y < screenHeight {
					img.Set(x, y, col)
				}
			}
		}
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Chaos Game — N-gon")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
