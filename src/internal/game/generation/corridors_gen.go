package generation

import (
	"errors"
	"rog/src/internal/game"
	"rog/src/internal/game/entities"
	"rog/src/internal/game/utils"
)

func carveCorridors(w *entities.World, rooms []*entities.Room) error {
	selectedDoors := selectOneRandomDoorPerRoom(rooms)
	shuffleDoors(selectedDoors)

	for i, j := 0, 1; i < len(selectedDoors)-1; i, j = i+1, j+1 {
		door1 := selectedDoors[i]
		door2 := selectedDoors[j]

		err := drawCorridor(w, door1.GetX(), door1.GetY(), door2.GetX(), door2.GetY())

		if err != nil {
			return err
		}
	}

	return nil
}

func selectOneRandomDoorPerRoom(rooms []*entities.Room) []entities.Door {
	selected := make([]entities.Door, len(rooms))

	for i, room := range rooms {
		doors := room.GetDoors()
		idx := utils.RandomFrom0ToN(len(doors))
		selected[i] = doors[idx]
	}

	return selected
}

func shuffleDoors(doors []entities.Door) {
	for i := range doors {
		j := utils.RandomFrom0ToN(i + 1)
		doors[i], doors[j] = doors[j], doors[i]
	}
}

func drawCorridor(w *entities.World, x1, y1, x2, y2 int) error {
	x1, y1 = drawCorridorStart(w, x1, y1)
	x2, y2 = drawCorridorStart(w, x2, y2)

	for i := 0; i < 3; i++ {
		x1 = findPathXForCorridor(w, x1, y1, x2, y2)
		y1 = findPathYForCorridor(w, x1, y1, x2, y2)

		if x1 == x2 && y1 == y2 {
			break
		}
	}

	if x1 != x2 || y1 != y2 {
		x1, y1 = adjustPath(w, x1, y1, x2, y2)
	}

	if x1 != x2 || y1 != y2 {
		return errors.New("corridors_gen.go: drawCorridor: corridors did not connect")
	}

	return nil
}

func drawCorridorStart(w *entities.World, x, y int) (int, int) {
	if w.GetCell(x, y-1) == game.CellEmpty || w.GetCell(x, y-1) == game.CellCorridor {
		w.SetCell(x, y-1, game.CellCorridor)
		w.SetCell(x, y-2, game.CellCorridor)
		return x, y - 2
	} else if w.GetCell(x+1, y) == game.CellEmpty || w.GetCell(x+1, y) == game.CellCorridor {
		w.SetCell(x+1, y, game.CellCorridor)
		w.SetCell(x+2, y, game.CellCorridor)
		return x + 2, y
	} else if w.GetCell(x, y+1) == game.CellEmpty || w.GetCell(x, y+1) == game.CellCorridor {
		w.SetCell(x, y+1, game.CellCorridor)
		w.SetCell(x, y+2, game.CellCorridor)
		return x, y + 2
	} else {
		w.SetCell(x-1, y, game.CellCorridor)
		w.SetCell(x-2, y, game.CellCorridor)
		return x - 2, y
	}
}

func findPathXForCorridor(w *entities.World, x1, y1, x2, y2 int) int {
	for x1 != x2 {
		dirX, _ := findDirection(x1, y1, x2, y2)

		switch dirX {
		case game.Unknown:
			return x1
		case game.Right:
			if (w.GetCell(x1+1, y1) == game.CellEmpty || w.GetCell(x1+1, y1) == game.CellCorridor) &&
				(w.GetCell(x1+2, y1) == game.CellEmpty || w.GetCell(x1+2, y1) == game.CellCorridor) {
				w.SetCell(x1+1, y1, game.CellCorridor)
				x1++
			} else {
				return x1
			}
		case game.Left:
			if (w.GetCell(x1-1, y1) == game.CellEmpty || w.GetCell(x1-1, y1) == game.CellCorridor) &&
				(w.GetCell(x1-2, y1) == game.CellEmpty || w.GetCell(x1-2, y1) == game.CellCorridor) {
				w.SetCell(x1-1, y1, game.CellCorridor)
				x1--
			} else {
				return x1
			}
		}
	}

	return x1
}

func findPathYForCorridor(w *entities.World, x1, y1, x2, y2 int) int {
	for y1 != y2 {
		_, dirY := findDirection(x1, y1, x2, y2)

		switch dirY {
		case game.Unknown:
			return y1
		case game.Up:
			if (w.GetCell(x1, y1-1) == game.CellEmpty || w.GetCell(x1, y1-1) == game.CellCorridor) &&
				(w.GetCell(x1, y1-2) == game.CellEmpty || w.GetCell(x1, y1-2) == game.CellCorridor) {
				w.SetCell(x1, y1-1, game.CellCorridor)
				y1--
			} else {
				return y1
			}
		case game.Down:
			if (w.GetCell(x1, y1+1) == game.CellEmpty || w.GetCell(x1, y1+1) == game.CellCorridor) &&
				(w.GetCell(x1, y1+2) == game.CellEmpty || w.GetCell(x1, y1+2) == game.CellCorridor) {
				w.SetCell(x1, y1+1, game.CellCorridor)
				y1++
			} else {
				return y1
			}
		}
	}

	return y1
}

func findDirection(x1, y1, x2, y2 int) (int, int) {
	dirX := game.Unknown
	dirY := game.Unknown

	if x1 < x2 {
		dirX = game.Right
	} else if x1 > x2 {
		dirX = game.Left
	}

	if y1 < y2 {
		dirY = game.Down
	} else if y1 > y2 {
		dirY = game.Up
	}

	return dirX, dirY
}

func adjustPath(w *entities.World, x1, y1, x2, y2 int) (int, int) {
	wallX, wallY, wallDir, closeRange := findWall(w, x1, y1)

	if closeRange {
		x1, y1 = stepBackFromWall(w, wallX, wallY, wallDir, x1, y1)
	}

	x1, y1 = passWall(w, wallX, wallY, wallDir, x1, y1, x2, y2)

	return x1, y1
}

func findWall(w *entities.World, x, y int) (int, int, int, bool) {
	checks := []struct {
		dx, dy   int
		dir      int
		distance bool
	}{
		{0, -1, game.Up, true},
		{1, 0, game.Right, true},
		{0, 1, game.Down, true},
		{-1, 0, game.Left, true},
		{0, -2, game.Up, false},
		{2, 0, game.Right, false},
		{0, 2, game.Down, false},
		{-2, 0, game.Left, false},
	}

	for _, check := range checks {
		nx, ny := x+check.dx, y+check.dy
		cell := w.GetCell(nx, ny)

		if cell == game.CellRoomWall || cell == game.CellRoomDoor {
			return nx, ny, check.dir, check.distance
		}
	}

	return -1, -1, 0, false
}

func stepBackFromWall(w *entities.World, wallX, wallY, wallDir, x1, y1 int) (int, int) {
	switch wallDir {
	case game.Up:
		w.SetCell(x1, y1+1, game.CellCorridor)
		return x1, y1 + 1
	case game.Down:
		w.SetCell(x1, y1-1, game.CellCorridor)
		return x1, y1 - 1
	case game.Left:
		w.SetCell(x1+1, y1, game.CellCorridor)
		return x1 + 1, y1
	case game.Right:
		w.SetCell(x1-1, y1, game.CellCorridor)
		return x1 - 1, y1
	default:
		return x1, y1
	}
}

func passWall(w *entities.World, wallX, wallY, wallDir, x1, y1, x2, y2 int) (int, int) {
	dirToCenterX, dirToCenterY := findDirection(wallX, wallY, game.WorldCenterX, game.WorldCenterY)

	if wallDir == game.Up || wallDir == game.Down {
		x1, y1 = passHorizontalWall(w, x1, y1, wallDir, dirToCenterX, x2, y2)
	}

	if wallDir == game.Left || wallDir == game.Right {
		x1, y1 = passVerticalWall(w, x1, y1, wallDir, dirToCenterY, x2, y2)
	}

	return x1, y1
}

func passHorizontalWall(w *entities.World, x1, y1, wallDir, dirToCenterX, x2, y2 int) (int, int) {
	cell := game.CellRoomWall
	chosenDir := game.Unknown

	for cell == game.CellRoomWall || cell == game.CellRoomDoor {
		if wallDir == game.Up {
			cell = w.GetCell(x1, y1-2)
		}

		if wallDir == game.Down {
			cell = w.GetCell(x1, y1+2)
		}

		if cell == game.CellRoomWall || cell == game.CellRoomDoor {
			if dirToCenterX == game.Left {
				x1--
				w.SetCell(x1, y1, game.CellCorridor)
				chosenDir = game.Left
			}

			if dirToCenterX == game.Right {
				x1++
				w.SetCell(x1, y1, game.CellCorridor)
				chosenDir = game.Right
			}

			if dirToCenterX == game.Unknown {
				x1++
				w.SetCell(x1, y1, game.CellCorridor)
				chosenDir = game.Right
			}
		}
	}

	x1, y1 = makeExtraStep(w, x1, y1, chosenDir)

	y1 = findPathYForCorridor(w, x1, y1, x2, y2)
	x1 = findPathXForCorridor(w, x1, y1, x2, y2)

	return x1, y1
}

func passVerticalWall(w *entities.World, x1, y1, wallDir, dirToCenterY, x2, y2 int) (int, int) {
	cell := game.CellRoomWall
	chosenDir := game.Unknown

	for cell == game.CellRoomWall || cell == game.CellRoomDoor {
		if wallDir == game.Left {
			cell = w.GetCell(x1-2, y1)
		}

		if wallDir == game.Right {
			cell = w.GetCell(x1+2, y1)
		}

		if cell == game.CellRoomWall || cell == game.CellRoomDoor {
			if dirToCenterY == game.Up {
				y1--
				w.SetCell(x1, y1, game.CellCorridor)
				chosenDir = game.Up
			}

			if dirToCenterY == game.Down {
				y1++
				w.SetCell(x1, y1, game.CellCorridor)
				chosenDir = game.Down
			}

			if dirToCenterY == game.Unknown {
				y1++
				w.SetCell(x1, y1, game.CellCorridor)
				chosenDir = game.Down
			}
		}
	}

	x1, y1 = makeExtraStep(w, x1, y1, chosenDir)

	x1 = findPathXForCorridor(w, x1, y1, x2, y2)
	y1 = findPathYForCorridor(w, x1, y1, x2, y2)

	return x1, y1
}

func makeExtraStep(w *entities.World, x1, y1 int, chosenDir int) (int, int) {
	if chosenDir == game.Left {
		x1--
		w.SetCell(x1, y1, game.CellCorridor)
	}

	if chosenDir == game.Right {
		x1++
		w.SetCell(x1, y1, game.CellCorridor)
	}

	return x1, y1
}
