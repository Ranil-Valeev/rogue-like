package generation

import (
	"rog/src/internal/game"
	"rog/src/internal/game/entities"
)

func createNewWorld() *entities.World {
	world := entities.NewWorld()

	return world
}

func fillWorldWithEmptyCells(w *entities.World) {
	for y := 0; y < game.WorldHeight; y++ {
		for x := 0; x < game.WorldWidth; x++ {
			w.SetCell(x, y, game.CellEmpty)
		}
	}
}

func fillWorldWithBorders(w *entities.World) {
	for x := 0; x < game.WorldWidth; x++ {
		w.SetCell(x, 0, game.CellBorderWallH)
		w.SetCell(x, game.WorldHeight-1, game.CellBorderWallH)
	}

	for y := 0; y < game.WorldHeight; y++ {
		w.SetCell(0, y, game.CellBorderWallV)
		w.SetCell(game.WorldWidth-1, y, game.CellBorderWallV)
	}

	w.SetCell(0, 0, game.CellBorderCornerTL)
	w.SetCell(game.WorldWidth-1, 0, game.CellBorderCornerTR)
	w.SetCell(0, game.WorldHeight-1, game.CellBorderCornerBL)
	w.SetCell(game.WorldWidth-1, game.WorldHeight-1, game.CellBorderCornerBR)
}
