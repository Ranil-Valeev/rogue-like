package utils

import (
	"rog/src/internal/game/entities"
)

func AddTreasuresToPlayer(p *entities.Player, m *entities.Monster, w *entities.World) {
	switch m.GetType() {
	case "zombie":
		p.SetTreasures(p.GetTreasures() + w.GetCurrentLevel())
	case "vampire":
		p.SetTreasures(p.GetTreasures() + 2*w.GetCurrentLevel())
	case "ghost":
		p.SetTreasures(p.GetTreasures() + 3*w.GetCurrentLevel())
	case "ogre":
		p.SetTreasures(p.GetTreasures() + 4*w.GetCurrentLevel())
	case "snake_mage":
		p.SetTreasures(p.GetTreasures() + 5*w.GetCurrentLevel())
	}
}
