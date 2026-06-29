package monster

import (
	"rog/src/internal/game"
	"rog/src/internal/game/entities"
	"rog/src/internal/game/utils"
)

var directions = [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func placeMonster(w *entities.World, m *entities.Monster, nx, ny int) {
	w.SetCell(m.GetX(), m.GetY(), m.GetOccupiedCell())
	m.SetOccupiedCell(w.GetCell(nx, ny))
	w.SetCell(nx, ny, m.MapSymbol())
	m.SetX(nx)
	m.SetY(ny)
}

func moveMonster(ctx TurnContext, m *entities.Monster, dx, dy int) bool {
	nx := m.GetX() + dx
	ny := m.GetY() + dy

	if !canMonsterMove(ctx, m, nx, ny) {
		return false
	}

	placeMonster(ctx.World, m, nx, ny)
	return true
}

func canMonsterMove(ctx TurnContext, m *entities.Monster, nx, ny int) bool {
	if ctx.World.GetCell(nx, ny) != game.CellFloor {
		return false
	}
	if !ctx.Room.ContainsFloor(nx, ny) {
		return false
	}
	if nx == ctx.Player.GetX() && ny == ctx.Player.GetY() {
		return false
	}
	for _, other := range ctx.World.Monsters {
		if other != m && other.IsAlive() && other.GetX() == nx && other.GetY() == ny {
			return false
		}
	}
	return true
}

func shuffledDirections() [][2]int {
	dirs := make([][2]int, len(directions))
	copy(dirs, directions)

	for i := len(dirs) - 1; i > 0; i-- {
		j := utils.RandomFrom0ToN(i + 1)
		dirs[i], dirs[j] = dirs[j], dirs[i]
	}
	return dirs
}
