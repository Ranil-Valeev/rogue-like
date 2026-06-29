package movement

import (
	"rog/src/internal/game"
	"rog/src/internal/game/entities"
	"rog/src/internal/game/generation"
	"rog/src/internal/game/monster"
	"rog/src/internal/game/utils"
)

func MoveTo(w *entities.World, p *entities.Player, m []*entities.Monster, dx, dy int) (*entities.World, *entities.Player, []*entities.Monster) {
	if canMove(w, p, dx, dy) {
		if nextLevel(w, p, dx, dy) {
			newW, newP, newM := generation.CreateWorld()
			newW.SetCurrentLevel(w.GetCurrentLevel() + 1)
			return newW, newP, newM
		}
		move(w, p, dx, dy)
		monster.ProcessTurn(w, p)
		return w, p, m
	}

	if canAttack(w, p, dx, dy) {
		monsterIdx, mon := findMonster(w, p.GetX()+dx, p.GetY()+dy, m)
		attack(p, mon)

		if !mon.IsAlive() {
			utils.AddTreasuresToPlayer(p, mon, w)
			m = append(m[:monsterIdx], m[monsterIdx+1:]...)
			w.SetCell(p.GetX()+dx, p.GetY()+dy, game.CellFloor)
		}
		monster.ProcessTurn(w, p)

	}

	// if canPickUp(w, p, x, y) {}
	return w, p, m
}

func canMove(w *entities.World, p *entities.Player, dx, dy int) bool {
	switch w.GetCell(p.GetX()+dx, p.GetY()+dy) {
	case game.CellFloor:
		return true
	case game.CellRoomDoor:
		return true
	case game.CellCorridor:
		return true
	case game.CellLevelExit:
		return true
	default:
		return false
	}
}

func move(w *entities.World, p *entities.Player, dx, dy int) {
	w.SetCell(p.GetX(), p.GetY(), p.GetOccupiedCell())
	p.SetOccupiedCell(w.GetCell(p.GetX()+dx, p.GetY()+dy))
	w.SetCell(p.GetX()+dx, p.GetY()+dy, game.CellPlayer)

	p.SetX(p.GetX() + dx)
	p.SetY(p.GetY() + dy)
}

func nextLevel(w *entities.World, p *entities.Player, dx, dy int) bool {
	if w.GetCell(p.GetX()+dx, p.GetY()+dy) == game.CellLevelExit {
		return true
	}

	return false
}

func canAttack(w *entities.World, p *entities.Player, dx, dy int) bool {
	switch w.GetCell(p.GetX()+dx, p.GetY()+dy) {
	case game.CellZombie:
		return true
	case game.CellVampire:
		return true
	case game.CellGhost:
		return true
	case game.CellOgre:
		return true
	case game.CellSnakeMage:
		return true
	default:
		return false
	}
}

func findMonster(w *entities.World, x, y int, monsters []*entities.Monster) (int, *entities.Monster) {
	for i, monster := range monsters {
		if monster.GetX() == x && monster.GetY() == y {
			return i, monster
		}
	}

	return -1, nil
}

func attack(p *entities.Player, m *entities.Monster) {
	playerDamage := utils.CalcPlayerDamage(p)

	m.SetHealth(m.GetHealth() - playerDamage)
}
