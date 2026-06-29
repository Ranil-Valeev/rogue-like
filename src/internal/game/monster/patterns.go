package monster

import (
	"rog/src/internal/game/entities"
	"rog/src/internal/game/settings"
	"rog/src/internal/game/utils"
)

func moveByType(ctx TurnContext, m *entities.Monster) {
	switch m.GetType() {
	case settings.TypeZombie, settings.TypeVampire:
		moveRandom(ctx, m)

	case settings.TypeGhost:
		moveGhost(ctx, m)

	case settings.TypeOgre:
		moveTwoSteps(ctx, m)

	case settings.TypeSnakeMage:
		moveDiagonal(ctx, m)

	default:
		moveRandom(ctx, m)
	}
}

func moveRandom(ctx TurnContext, m *entities.Monster) {
	for _, d := range shuffledDirections() {
		if moveMonster(ctx, m, d[0], d[1]) {
			return
		}
	}
}

func moveTwoSteps(ctx TurnContext, m *entities.Monster) {
	for _, d := range shuffledDirections() {
		if moveMonster(ctx, m, d[0], d[1]) {
			moveMonster(ctx, m, d[0], d[1])
			return
		}
	}
}

func moveGhost(ctx TurnContext, m *entities.Monster) {
	playerHere := ctx.Room.ContainsFloor(ctx.Player.GetX(), ctx.Player.GetY())

	if playerHere {
		m.SetVisible(true)
	} else if utils.RandomFrom0ToN(4) == 0 {
		m.SetVisible(!m.IsVisible())
	}

	moveTeleport(ctx, m)
}

func moveTeleport(ctx TurnContext, m *entities.Monster) {
	var freeCells [][2]int

	for y := ctx.Room.GetY1() + 1; y < ctx.Room.GetY2(); y++ {
		for x := ctx.Room.GetX1() + 1; x < ctx.Room.GetX2(); x++ {
			if x == m.GetX() && y == m.GetY() {
				continue
			}
			if !canMonsterMove(ctx, m, x, y) {
				continue
			}
			freeCells = append(freeCells, [2]int{x, y})
		}
	}

	if len(freeCells) == 0 {
		return
	}

	idx := utils.RandomFrom0ToN(len(freeCells))
	nx, ny := freeCells[idx][0], freeCells[idx][1]
	placeMonster(ctx.World, m, nx, ny)
}

var diagonals = [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func moveDiagonal(ctx TurnContext, m *entities.Monster) {
	dx, dy := m.MoveDir()

	if dx == 0 && dy == 0 {
		d := diagonals[utils.RandomFrom0ToN(len(diagonals))]
		dx, dy = d[0], d[1]
	}

	if !moveMonster(ctx, m, dx, dy) {
		dx, dy = -dx, -dy
		moveMonster(ctx, m, dx, dy)
	}

	m.SetMoveDir(dx, dy)
}
