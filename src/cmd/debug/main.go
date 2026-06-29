package main

import (
	"fmt"
	"os/exec"
	"rog/src/internal/game"
	"rog/src/internal/game/entities"
	"rog/src/internal/game/generation"
	"rog/src/internal/game/movement"
)

func main() {
	world, player, monsters := generation.CreateWorld()

	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	clearScreen()
	printWorld(world)

	for {
		var ch byte
		fmt.Scanf("%c", &ch)

		switch ch {
		case 'w', 'W':
			world, player, monsters = movement.MoveTo(world, player, monsters, 0, -1)
		case 's', 'S':
			world, player, monsters = movement.MoveTo(world, player, monsters, 0, 1)
		case 'a', 'A':
			world, player, monsters = movement.MoveTo(world, player, monsters, -1, 0)
		case 'd', 'D':
			world, player, monsters = movement.MoveTo(world, player, monsters, 1, 0)
		case 'q', 'Q':
			return
		default:
			continue
		}

		clearScreen()
		printWorld(world)
	}
}

func printWorld(world *entities.World) {
	for y := 0; y < game.WorldHeight; y++ {
		for x := 0; x < game.WorldWidth; x++ {
			fmt.Print(string(world.GetCell(x, y)))
		}
		fmt.Println()
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
