package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth     = 800
	ScreenHeight    = 600
	TileSize        = 20
	Cols            = ScreenWidth / TileSize
	Rows            = ScreenHeight / TileSize
	InitialSpeed    = 4.0
	SpeedMultiplier = 1.1
	MaxSpeed        = 20.0
)

var (
	ColorDay   = color.RGBA{R: 200, G: 200, B: 200, A: 255}
	ColorNight = color.RGBA{R: 50, G: 50, B: 50, A: 255}
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

func (g *Game) Init() {
	g.Ball = nil

	rand.Seed(time.Now().UnixNano())

	for x := 0; x < Cols; x++ {
		for y := 0; y < Rows; y++ {
			if x < Cols/2 {
				g.Grid[x][y] = ColorDay
			} else {
				g.Grid[x][y] = ColorNight
			}
		}
	}

	randomDir := func() float64 {
		if rand.Intn(2) == 0 {
			return InitialSpeed
		}
		return -InitialSpeed
	}

	g.Ball = append(g.Ball, &Ball{
		X: 200, Y: float64(ScreenHeight) / 2,
		DX: randomDir(), DY: randomDir(),
		Color: ColorNight,
	})

	g.Ball = append(g.Ball, &Ball{
		X: 600, Y: float64(ScreenHeight) / 2,
		DX: randomDir(), DY: randomDir(),
		Color: ColorDay,
	})
}

func NewGame() *Game {
	g := &Game{}
	g.Init()
	return g
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Init()
		return nil
	}

	for _, b := range g.Ball {
		b.X += b.DX
		b.Y += b.DY

		if b.X < 0 || b.X > ScreenWidth {
			b.DX = -b.DX
			if b.X < 0 {
				b.X = 0
			}
			if b.X > ScreenWidth {
				b.X = ScreenWidth
			}
		}

		if b.Y < 0 || b.Y > ScreenHeight {
			b.DY = -b.DY
			if b.Y < 0 {
				b.Y = 0
			}
			if b.Y > ScreenHeight {
				b.Y = ScreenHeight
			}
		}

		gridX := int(b.X / TileSize)
		gridY := int(b.Y / TileSize)

		if gridX >= 0 && gridX < Cols && gridY >= 0 && gridY < Rows {
			if g.Grid[gridX][gridY] == b.Color {
				if b.Color == ColorDay {
					g.Grid[gridX][gridY] = ColorNight
				} else {
					g.Grid[gridX][gridY] = ColorDay
				}

				b.DX *= SpeedMultiplier
				b.DY *= SpeedMultiplier

				if b.DX > MaxSpeed {
					b.DX = MaxSpeed
				}
				if b.DX < -MaxSpeed {
					b.DX = -MaxSpeed
				}
				if b.DY > MaxSpeed {
					b.DY = MaxSpeed
				}
				if b.DY < -MaxSpeed {
					b.DY = -MaxSpeed
				}

				b.DX = -b.DX
				b.DY = -b.DY
			}
		}
	}

	dayCount, nightCount := 0, 0
	for x := 0; x < Cols; x++ {
		for y := 0; y < Rows; y++ {
			if g.Grid[x][y] == ColorDay {
				dayCount++
			} else {
				nightCount++
			}
		}
	}
	ebiten.SetWindowTitle(fmt.Sprintf("Pong Wars | Day: %d | Night: %d | [R] Restart", dayCount, nightCount))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for x := 0; x < Cols; x++ {
		for y := 0; y < Rows; y++ {
			vector.FillRect(screen, float32(x*TileSize), float32(y*TileSize), TileSize, TileSize, g.Grid[x][y], false)
		}
	}

	for _, b := range g.Ball {
		vector.FillCircle(screen, float32(b.X), float32(b.Y), 8, b.Color, true)
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
