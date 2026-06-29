package generation

import (
	"rog/src/internal/game"
	"rog/src/internal/game/entities"
	"rog/src/internal/game/utils"
)

func carveLevelExit(w *entities.World, rooms []*entities.Room) {
	chosenRoom := utils.ChooseRandom(0, 1, 2, 3, 5, 6, 7, 8)
	dx := utils.ChooseRandom(-2, -1, 0, 1, 2)
	dy := utils.ChooseRandom(-2, -1, 0, 1, 2)

	w.SetCell(rooms[chosenRoom].GetCenterX()+dx, rooms[chosenRoom].GetCenterY()+dy, game.CellLevelExit)
}
