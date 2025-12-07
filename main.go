package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 800
}

func main(){
	
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Go-pongwars")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}