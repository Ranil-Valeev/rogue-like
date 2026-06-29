package settings

type MonsterSpawnConfig struct {
	MinMonsters        int
	MaxMonsters        int
	BaseMonsters       int
	MonsterPerLevel    int
	MaxMonstersPerRoom int

	TypeDistribution map[int]map[string]int
}

func DefaultMonsterSpawnConfig() *MonsterSpawnConfig {
	return &MonsterSpawnConfig{
		MinMonsters:        2,
		MaxMonsters:        18,
		BaseMonsters:       3,
		MonsterPerLevel:    1,
		MaxMonstersPerRoom: 5,

		TypeDistribution: map[int]map[string]int{
			// Уровни 1-2: только зомби
			1: {TypeZombie: 100},
			2: {TypeZombie: 100},

			// Уровни 3-4: зомби + вампиры + привидения
			3: {TypeZombie: 55, TypeVampire: 20, TypeGhost: 25},
			4: {TypeZombie: 50, TypeVampire: 25, TypeGhost: 25},

			// Уровни 5-7: + огры
			5: {TypeZombie: 45, TypeVampire: 20, TypeGhost: 25, TypeOgre: 10},
			6: {TypeZombie: 40, TypeVampire: 20, TypeGhost: 25, TypeOgre: 15},
			7: {TypeZombie: 35, TypeVampire: 20, TypeGhost: 20, TypeOgre: 25},

			// Уровни 8-10: + змеи-маги
			8:  {TypeZombie: 30, TypeVampire: 20, TypeGhost: 20, TypeOgre: 20, TypeSnakeMage: 10},
			9:  {TypeZombie: 25, TypeVampire: 20, TypeGhost: 15, TypeOgre: 20, TypeSnakeMage: 20},
			10: {TypeZombie: 20, TypeVampire: 20, TypeGhost: 15, TypeOgre: 20, TypeSnakeMage: 25},

			// Уровни 11-15: змеи-маги чаще
			11: {TypeZombie: 20, TypeVampire: 20, TypeGhost: 15, TypeOgre: 20, TypeSnakeMage: 25},
			12: {TypeZombie: 18, TypeVampire: 20, TypeGhost: 12, TypeOgre: 22, TypeSnakeMage: 28},
			13: {TypeZombie: 15, TypeVampire: 20, TypeGhost: 10, TypeOgre: 25, TypeSnakeMage: 30},
			14: {TypeZombie: 12, TypeVampire: 20, TypeGhost: 10, TypeOgre: 25, TypeSnakeMage: 33},
			15: {TypeZombie: 10, TypeVampire: 20, TypeGhost: 10, TypeOgre: 25, TypeSnakeMage: 35},

			// Уровни 16-21: эндгейм — сильные враги доминируют
			16: {TypeZombie: 8, TypeVampire: 22, TypeGhost: 7, TypeOgre: 28, TypeSnakeMage: 35},
			17: {TypeZombie: 6, TypeVampire: 23, TypeGhost: 6, TypeOgre: 30, TypeSnakeMage: 35},
			18: {TypeZombie: 5, TypeVampire: 23, TypeGhost: 5, TypeOgre: 32, TypeSnakeMage: 35},
			19: {TypeZombie: 3, TypeVampire: 25, TypeGhost: 5, TypeOgre: 32, TypeSnakeMage: 35},
			20: {TypeZombie: 3, TypeVampire: 25, TypeGhost: 3, TypeOgre: 34, TypeSnakeMage: 35},
			21: {TypeZombie: 2, TypeVampire: 28, TypeGhost: 2, TypeOgre: 33, TypeSnakeMage: 35},
		},
	}
}

func (c *MonsterSpawnConfig) GetConfigForLevel(level int) map[string]int {
	if dist, ok := c.TypeDistribution[level]; ok {
		return dist
	}

	for l := level; l >= 1; l-- {
		if dist, ok := c.TypeDistribution[l]; ok {
			return dist
		}
	}

	return map[string]int{TypeZombie: 100}
}
