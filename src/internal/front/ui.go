package front

import (
	"github.com/gdamore/tcell/v2"

	"rog/src/internal/game/entities"
)

// UI - основной интерфейс пользователя
type UI struct {
	screen tcell.Screen
	world  *entities.World
	player *entities.Player
	monsters []*entities.Monster

	camera *Camera

	gameWidth  int
	gameHeight int

	screenWidth  int
	screenHeight int

	// // Смещение для центрирования(потом сделать)
	// offsetX int
	// offsetY int

	// Высота интерфейса
	uiHeight int

	// Флаг для выхода
	running bool

	// Рендерер
	renderer *Renderer

	// Что было под игроком (для восстановления)
	underPlayer rune
}

// NewUI - создает новый UI
// func NewUI(screen tcell.Screen, world *entities.World) *UI {
// 	screenWidth, screenHeight := screen.Size()

// 	// Высота интерфейса - 1 строка (статус)
// 	uiHeight := 1

// 	// Ищем игрока в мире
// 	playerX, playerY := findPlayer(world)

// 	// Создаем игрока
// 	player := entities.NewPlayer(playerX, playerY)

// 	// Размеры игровой области
// 	gameWidth := screenWidth
// 	gameHeight := screenHeight - uiHeight

// 	// Сохраняем что было под игроком
// 	underPlayer := world.GetCell(playerX, playerY)

// 	ui := &UI{
// 		screen:       screen,
// 		world:        world,
// 		player:       player,
// 		screenWidth:  screenWidth,
// 		screenHeight: screenHeight,
// 		gameWidth:    gameWidth,
// 		gameHeight:   gameHeight,
// 		uiHeight:     uiHeight,
// 		running:      true,
// 		underPlayer:  underPlayer,
// 	}

// 	// Создаем камеру
// 	ui.camera = NewCamera(player.GetX(), player.GetY(), gameWidth, gameHeight)

// 	// Создаем рендерер
// 	ui.renderer = NewRenderer(screen, world, ui.camera, gameWidth, gameHeight)

// 	return ui
// }


// NewUI - создает новый UI
func NewUI(screen tcell.Screen, world *entities.World, player *entities.Player, monsters []*entities.Monster) *UI {
	screenWidth, screenHeight := screen.Size()

	// Высота интерфейса - 1 строка (статус)
	uiHeight := 1

	// Размеры игровой области
	gameWidth := screenWidth
	gameHeight := screenHeight - uiHeight

	// Сохраняем что было под игроком
	underPlayer := world.GetCell(player.GetX(), player.GetY())

	ui := &UI{
		screen:       screen,
		world:        world,
		player:       player,
		monsters:     monsters,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		gameWidth:    gameWidth,
		gameHeight:   gameHeight,
		uiHeight:     uiHeight,
		running:      true,
		underPlayer:  underPlayer,
	}

	// Создаем камеру
	ui.camera = NewCamera(player.GetX(), player.GetY(), gameWidth, gameHeight)

	// Создаем рендерер
	ui.renderer = NewRenderer(screen, world, ui.camera, gameWidth, gameHeight)
	ui.renderer.SetMonsters(monsters)

	return ui
}

func (ui *UI) Run() error {
	ui.camera.Follow(ui.player.GetX(), ui.player.GetY())
	ui.renderer.Render(ui.player.GetX(), ui.player.GetY())
	ui.screen.Show()

	for ui.running {
		ev := ui.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			ui.handleKey(ev)
		case *tcell.EventResize:
			ui.handleResize(ev)
		}

		ui.camera.Follow(ui.player.GetX(), ui.player.GetY())
		ui.renderer.Render(ui.player.GetX(), ui.player.GetY())
		ui.screen.Show()
	}

	return nil
}