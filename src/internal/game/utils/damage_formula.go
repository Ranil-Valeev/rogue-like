package utils

import (
	"rog/src/internal/game/entities"
)

func CalcPlayerDamage(p *entities.Player) int {
	damage := p.GetDexterity() + p.GetStrenght()

	return damage
}
