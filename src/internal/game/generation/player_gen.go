package generation

import (
	"rog/src/internal/game"
	"rog/src/internal/game/entities"
	"rog/src/internal/game/utils"
)

func createPlayer(centralRoom *entities.Room) *entities.Player {
	x := centralRoom.GetCenterX() + utils.ChooseRandom(-2, -1, 0, 1, 2)
	y := centralRoom.GetCenterY() + utils.ChooseRandom(-2, -1, 0, 1, 2)
	player := entities.NewPlayer(x, y)

	return player
}

func carvePlayer(w *entities.World, p *entities.Player) {
	w.SetCell(p.GetX(), p.GetY(), game.CellPlayer)
}
