package main

import (
	"log"

	"github.com/gdamore/tcell/v2"

	"rog/src/internal/front"
	"rog/src/internal/game/entities"
	"rog/src/internal/game/generation"
)

func main() {
	world, player, monsters := generation.CreateWorld()

	StartGame(world, player, monsters)
}

func StartGame(world *entities.World, player *entities.Player, monsters []*entities.Monster) {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err := screen.Init(); err != nil {
		log.Fatal(err)
	}
	defer screen.Fini()

	// Создаем UI и передаем все нужные данные
	ui := front.NewUI(screen, world, player, monsters)
	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}