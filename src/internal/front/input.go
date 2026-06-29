package front

import (
	"github.com/gdamore/tcell/v2"

	"rog/src/internal/game/movement"
)

// handleKey - обработка нажатий клавиш
func (ui *UI) handleKey(ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyEscape, tcell.KeyCtrlC:
		ui.running = false
		return

	case tcell.KeyUp:
		ui.tryMove(0, -1)
	case tcell.KeyDown:
		ui.tryMove(0, 1)
	case tcell.KeyLeft:
		ui.tryMove(-1, 0)
	case tcell.KeyRight:
		ui.tryMove(1, 0)

	case tcell.KeyRune:
		switch ev.Rune() {
		case 'w', 'W':
			ui.tryMove(0, -1)
		case 's', 'S':
			ui.tryMove(0, 1)
		case 'a', 'A':
			ui.tryMove(-1, 0)
		case 'd', 'D':
			ui.tryMove(1, 0)
		case 'q', 'Q':
			ui.running = false
			return
		default:
			return
		}
	default:
		return
	}
	
}

// tryMove - пытается переместить игрока
func (ui *UI) tryMove(dx, dy int) {
	oldX, oldY := ui.player.GetX(), ui.player.GetY()
	
	// Вызываем movement.MoveTo
	newWorld, newPlayer, newMonsters := movement.MoveTo(
		ui.world,
		ui.player,
		ui.monsters,
		dx, dy,
	)
	
	// Проверяем, изменилось ли что-то
	if newWorld != ui.world {
		ui.world = newWorld
		
		ui.player = newPlayer
		
		if newMonsters != nil {
			ui.monsters = newMonsters
		}

		ui.renderer.SetWorld(newWorld)
		ui.renderer.SetMonsters(ui.monsters)
		
		// Обновляем камеру на новую позицию игрока
		newX, newY := ui.player.GetX(), ui.player.GetY()
		ui.camera.Follow(newX, newY)
		
		// Обновляем клетку под игроком
		ui.underPlayer = ui.world.GetCell(newX, newY)
		
		// Перерисовываем
		ui.renderer.Render(newX, newY)
		ui.screen.Show()
		
		return
	}

	if newPlayer != ui.player {
		ui.player = newPlayer
		
		if newMonsters != nil {
			ui.monsters = newMonsters
			ui.renderer.SetMonsters(newMonsters)
		}
		
		newX, newY := ui.player.GetX(), ui.player.GetY()
		if newX != oldX || newY != oldY {
			ui.underPlayer = ui.world.GetCell(newX, newY)
		}
		
		ui.camera.Follow(newX, newY)
	}
}



// handleAction - обработка действий (атака, взаимодействие и т.д.)
func (ui *UI) handleAction() {
	// Здесь логика атаки или взаимодействия

	ui.renderer.Render(ui.player.GetX(), ui.player.GetY())
	ui.screen.Show()
}

// handleResize - обработка изменения размера окна
func (ui *UI) handleResize(ev *tcell.EventResize) {
	// Получаем новые размеры
	ui.screenWidth, ui.screenHeight = ev.Size()

	// Обновляем размеры игровой области
	ui.gameWidth = ui.screenWidth
	ui.gameHeight = ui.screenHeight - ui.uiHeight

	// Обновляем камеру
	ui.camera.SetBounds(ui.gameWidth, ui.gameHeight)
	ui.camera.Follow(ui.player.GetX(), ui.player.GetY())

	// Обновляем рендерер
	ui.renderer.SetSize(ui.gameWidth, ui.gameHeight)

	// Перерисовываем
	ui.renderer.Render(ui.player.GetX(), ui.player.GetY())
	ui.screen.Show()
}
