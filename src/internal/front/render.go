package front

import (
	"fmt"

	"github.com/gdamore/tcell/v2"

	"rog/src/internal/game"
	"rog/src/internal/game/entities"
)

// Renderer - отвечает за отрисовку
type Renderer struct {
	screen   tcell.Screen
	world    *entities.World
	camera   *Camera
	monsters []*entities.Monster

	// Размеры игровой области
	width  int
	height int
}

// NewRenderer - создает новый рендерер
func NewRenderer(screen tcell.Screen, world *entities.World, camera *Camera, width, height int) *Renderer {
	return &Renderer{
		screen: screen,
		world:  world,
		camera: camera,
		width:  width,
		height: height,
	}
}

func (r *Renderer) SetMonsters(monsters []*entities.Monster) {
	r.monsters = monsters
}

func (r *Renderer) SetWorld(world *entities.World) {
	r.world = world
}

func (r *Renderer) SetSize(width, height int) {
	r.width = width
	r.height = height
}

// Render - отрисовывает всё на экране
func (r *Renderer) Render(playerX, playerY int) {
	r.screen.Clear()

	// Рисуем мир
	r.renderWorld()

	// Рисуем монстров поверх мира
	r.renderMonsters()

	// Рисуем игрока поверх всего
	r.renderPlayer(playerX, playerY)

	// Рисуем статусную строку
	r.renderStatus(playerX, playerY)
}

// renderWorld - отрисовывает видимую часть мира
func (r *Renderer) renderWorld() {
	offsetX, offsetY := r.camera.GetOffset()

	for y := 0; y < r.height; y++ {
		for x := 0; x < r.width; x++ {
			worldX := x + offsetX
			worldY := y + offsetY

			// Проверяем, что координаты в пределах мира
			if worldX >= 0 && worldX < game.WorldWidth &&
				worldY >= 0 && worldY < game.WorldHeight {

				ch := r.world.GetCell(worldX, worldY)

				if ch == game.CellPlayer {
					ch = game.CellFloor
				}

				style := r.getCellStyle(ch)
				r.screen.SetContent(x, y, ch, nil, style)
			}
		}
	}
}

// renderMonsters - отрисовывает монстров
func (r *Renderer) renderMonsters() {
	if r.monsters == nil {
		return
	}

	for _, m := range r.monsters {
		if !m.IsAlive() {
			continue
		}

		// Проверяем, виден ли монстр в камере
		if !r.camera.IsVisible(m.GetX(), m.GetY()) {
			continue
		}

		// Получаем экранные координаты
		screenX, screenY := r.camera.WorldToScreen(m.GetX(), m.GetY())

		// Проверяем границы экрана
		if screenX < 0 || screenX >= r.width || screenY < 0 || screenY >= r.height {
			continue
		}

		// Получаем символ монстра
		symbol := m.MapSymbol()

		// Получаем цвет для типа монстра
		color := r.getMonsterColor(m.GetType())

		style := tcell.StyleDefault.
			Foreground(color).
			Background(tcell.ColorBlack)

		r.screen.SetContent(screenX, screenY, symbol, nil, style)
	}
}

// getMonsterColor - возвращает цвет для типа монстра
func (r *Renderer) getMonsterColor(monsterType string) tcell.Color {
	switch monsterType {
	case "zombie":
		return tcell.ColorGreen
	case "vampire":
		return tcell.ColorRed
	case "ghost":
		return tcell.ColorWhite
	case "ogre":
		return tcell.ColorBrown
	case "snake_mage":
		return tcell.ColorPurple
	default:
		return tcell.ColorYellow
	}
}

// renderPlayer - отрисовывает игрока
func (r *Renderer) renderPlayer(playerX, playerY int) {
	screenX, screenY := r.camera.WorldToScreen(playerX, playerY)

	if screenX >= 0 && screenX < r.width &&
		screenY >= 0 && screenY < r.height {

		style := tcell.StyleDefault.
			Foreground(ColorPlayer).
			Background(tcell.ColorBlack)

		r.screen.SetContent(screenX, screenY, game.CellPlayer, nil, style)
	}
}

// renderStatus - отрисовывает статусную строку внизу
func (r *Renderer) renderStatus(playerX, playerY int) {
	// Проверяем, что есть место для статуса
	if r.height < 1 {
		return
	}

	y := r.height

	// Считаем живых монстров
	monsterCount := 0
	if r.monsters != nil {
		for _, m := range r.monsters {
			if m.IsAlive() {
				monsterCount++
			}
		}
	}

	status := fmt.Sprintf(" POS: (%d, %d) | Monsters: %d | Level: %d | Press Q to exit",
		playerX, playerY, monsterCount, r.world.GetCurrentLevel())

	// Обрезаем строку, если она длиннее экрана
	if len(status) > r.width {
		status = status[:r.width]
	}

	style := tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack)

	for x, ch := range status {
		if x < r.width {
			r.screen.SetContent(x, y, ch, nil, style)
		}
	}
}

// getCellStyle - возвращает стиль для ячейки
func (r *Renderer) getCellStyle(ch rune) tcell.Style {
	style := tcell.StyleDefault.Background(tcell.ColorBlack)

	switch ch {
	case game.CellEmpty:
		style = style.Foreground(tcell.ColorBlack)

	case game.CellFloor:
		style = style.Foreground(ColorFloor)

	case game.CellCorridor:
		style = style.Foreground(ColorFloorDark)

	case game.CellRoomWall:
		style = style.Foreground(ColorWall)

	case game.CellRoomDoor:
		style = style.Foreground(ColorDoor)

	case game.CellBorderWallH, game.CellBorderWallV,
		game.CellBorderCornerTL, game.CellBorderCornerTR,
		game.CellBorderCornerBL, game.CellBorderCornerBR:
		style = style.Foreground(ColorBorder)

	case game.CellPlayer:
		style = style.Foreground(ColorPlayer)

	default:
		style = style.Foreground(tcell.ColorWhite)
	}

	return style
}