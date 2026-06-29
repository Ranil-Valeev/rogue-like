package generation

import (
	"rog/src/internal/game"
	"rog/src/internal/game/entities"
)

func CreateWorld() (*entities.World, *entities.Player, []*entities.Monster) {
	world := createNewWorld()
	fillWorldWithEmptyCells(world)
	fillWorldWithBorders(world)

	rooms := createRooms()
	world.SetRooms(rooms)
	carveRoom(world, rooms)
	createDoors(world, rooms)

	for {
		err := carveCorridors(world, rooms)
		if err == nil {
			break
		} else {
			continue
		}
	}

	carveLevelExit(world, rooms)

	player := createPlayer(rooms[game.Mid])
	carvePlayer(world, player)
	level := world.GetCurrentLevel()
	monsters := CreateMonstersForLevel(rooms, level)
	for _, m := range monsters {
		world.AddMonster(m)
	}
	carveMonsters(world, monsters)

	return world, player, monsters
}
