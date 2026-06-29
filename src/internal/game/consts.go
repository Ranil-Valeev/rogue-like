package game

// Размеры мира
const (
	WorldWidth  = 120
	WorldHeight = 40

	WorldCenterX = WorldWidth / 2
	WorldCenterY = WorldHeight / 2
)

// Типы статичных клеток
const (
	CellEmpty     = ' '
	CellFloor     = '⋅'
	CellRoomWall  = '#'
	CellRoomDoor  = '+'
	CellCorridor  = '░'
	CellLevelExit = 'X'

	CellBorderWallH    = '═'
	CellBorderWallV    = '║'
	CellBorderCornerTL = '╔'
	CellBorderCornerTR = '╗'
	CellBorderCornerBL = '╚'
	CellBorderCornerBR = '╝'
)

// Типы динамических клеток
const (
	CellPlayer    = '☺'
	CellZombie    = 'z'
	CellVampire   = 'v'
	CellGhost     = 'g'
	CellOgre      = 'O'
	CellSnakeMage = 's'
)

// Направления
const (
	Unknown = iota
	Up
	Right
	Down
	Left
)

// Комнаты
const (
	TopLeft = iota
	Top
	TopRight
	MidLeft
	Mid
	MidRight
	BotLeft
	Bot
	BotRight
)

// Типы статичных клеток
/*
const (
	CellEmpty    = ' '
	CellFloor    = '⋅'
	CellRoomWall = '#'
	CellRoomDoor = '+'
	CellCorridor = '░'

	CellBorderWallH    = '─'
	CellBorderWallV    = '|'
	CellBorderCornerTL = '┌'
	CellBorderCornerTR = '┐'
	CellBorderCornerBL = '└'
	CellBorderCornerBR = '┘'
) */
