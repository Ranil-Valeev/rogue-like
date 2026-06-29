package front

import (
	"rog/src/internal/game"
)

// Camera - камера, которая следует за игроком
type Camera struct {
	// Смещение камеры (левый верхний угол видимой области)
	offsetX int
	offsetY int

	// Размеры видимой области
	width  int
	height int
}

func NewCamera(playerX, playerY, width, height int) *Camera {
	cam := &Camera{
		width:  width,
		height: height,
	}
	cam.Follow(playerX, playerY)
	return cam
}

func (c *Camera) Follow(playerX, playerY int) {
	c.offsetX = playerX - c.width/2
	c.offsetY = playerY - c.height/2

	c.clamp()
}

// SetBounds - обновляет размеры видимой области
func (c *Camera) SetBounds(width, height int) {
	c.width = width
	c.height = height
	c.clamp()
}

// clamp - ограничивает камеру границами мира
func (c *Camera) clamp() {
	// Если камера шире мира - центрируем мир
	if c.width >= game.WorldWidth {
		c.offsetX = 0
	} else {
		// Не даем камере выйти за границы
		if c.offsetX < 0 {
			c.offsetX = 0
		}
		if c.offsetX+c.width > game.WorldWidth {
			c.offsetX = game.WorldWidth - c.width
		}
	}

	// Если камера выше мира - центрируем мир
	if c.height >= game.WorldHeight {
		c.offsetY = 0
	} else {
		// Не даем камере выйти за границы
		if c.offsetY < 0 {
			c.offsetY = 0
		}
		if c.offsetY+c.height > game.WorldHeight {
			c.offsetY = game.WorldHeight - c.height
		}
	}
}

// GetOffset - возвращает смещение камеры
func (c *Camera) GetOffset() (int, int) {
	return c.offsetX, c.offsetY
}

// WorldToScreen - преобразует мировые координаты в экранные
func (c *Camera) WorldToScreen(worldX, worldY int) (int, int) {
	return worldX - c.offsetX, worldY - c.offsetY
}

// ScreenToWorld - преобразует экранные координаты в мировые
func (c *Camera) ScreenToWorld(screenX, screenY int) (int, int) {
	return screenX + c.offsetX, screenY + c.offsetY
}

// IsVisible - проверяет, видна ли точка в мире
func (c *Camera) IsVisible(worldX, worldY int) bool {
	screenX, screenY := c.WorldToScreen(worldX, worldY)
	return screenX >= 0 && screenX < c.width &&
		screenY >= 0 && screenY < c.height
}
