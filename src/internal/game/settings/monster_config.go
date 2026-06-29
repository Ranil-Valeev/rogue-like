package settings

import "rog/src/internal/game"

const (
	TypeZombie    = "zombie"
	TypeVampire   = "vampire"
	TypeGhost     = "ghost"
	TypeOgre      = "ogre"
	TypeSnakeMage = "snake_mage"
)

type MonsterStat struct {
	HP        int
	Dex       int
	Str       int
	Hostility int
}

var MonsterSymbols = map[string]rune{
	TypeZombie:    game.CellZombie,
	TypeVampire:   game.CellVampire,
	TypeGhost:     game.CellGhost,
	TypeOgre:      game.CellOgre,
	TypeSnakeMage: game.CellSnakeMage,
}

var MonsterStats = map[string]MonsterStat{
	TypeZombie:    {HP: 50, Dex: 5, Str: 15, Hostility: 3},
	TypeVampire:   {HP: 40, Dex: 20, Str: 10, Hostility: 5},
	TypeGhost:     {HP: 30, Dex: 18, Str: 5, Hostility: 2},
	TypeOgre:      {HP: 80, Dex: 3, Str: 25, Hostility: 3},
	TypeSnakeMage: {HP: 35, Dex: 22, Str: 8, Hostility: 6},
}

const StatScalePercentPerLevel = 2 // как в sample: PERCENTS_UPDATE_DIFFICULTY_MONSTERS

func ScaleStatsForLevel(base MonsterStat, level int) MonsterStat {
	if level < 1 {
		level = 1
	}
	scale := StatScalePercentPerLevel * level

	return MonsterStat{
		HP:        scaleStat(base.HP, scale),
		Dex:       scaleStat(base.Dex, scale),
		Str:       scaleStat(base.Str, scale),
		Hostility: base.Hostility,
	}
}

func scaleStat(base, percent int) int {
	return base + base*percent/100
}
