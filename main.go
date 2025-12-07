package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 800
	TileSize     = 20
	Cols         = ScreenWidth / TileSize
	Rows         = ScreenHeight / TileSize
)

var (
	ColorDay   = color.RGBA{R: 200, G: 200, B: 200, A: 255}
	ColorNight = color.RGBA{R: 50, G: 50, B: 50, A: 255}
	ColorBall  = color.RGBA{R: 255, G: 0, B: 0, A: 255}
)

type Ball struct {
	X, Y   float64
	DX, DY float64
	Color  color.Color
}

type Game struct {
	Grid [Cols][Rows]color.Color
	Ball []*Ball
}

func NewGame() *Game {
	g := &Game{}

	for x := 0; x < Cols; x++ {
		for y := 0; y < Rows; y++ {
			if x < Cols/2 {
				g.Grid[x][y] = ColorDay
			} else {
				g.Grid[x][y] = ColorNight
			}
		}
	}

	g.Ball = append(g.Ball, &Ball{X: 200, Y: 400, DX: 4, DY: 4, Color: ColorNight})
	g.Ball = append(g.Ball, &Ball{X: 600, Y: 400, DX: -4, DY: -4, Color: ColorDay})

	return g
}

func (g *Game) Update() error {

	for _, b := range g.Ball {
		b.X += b.DX
		b.Y += b.DY

		if b.X < 0 || b.X > ScreenWidth {
			b.DX = -b.DX
		}

		if b.Y < 0 || b.Y > ScreenHeight {
			b.DY = -b.DY
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for x := 0; x < Cols; x++ {
		for y := 0; y < Rows; y++ {
			vector.FillRect(screen, float32(x*TileSize), float32(y*TileSize), TileSize, TileSize, g.Grid[x][y], false)
		}
	}

	for _, b := range g.Ball {
		vector.FillCircle(screen, float32(b.X), float32(b.Y), 10, b.Color, true)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func main() {

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Go-pongwars")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
