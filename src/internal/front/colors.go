package front

import "github.com/gdamore/tcell/v2"

var (
	ColorWall      = tcell.NewRGBColor(128, 128, 128) // Серый
	ColorWallDark  = tcell.NewRGBColor(80, 80, 80)    // Темно-серый
	ColorFloor     = tcell.NewRGBColor(64, 128, 64)   // Темно-зеленый
	ColorFloorDark = tcell.NewRGBColor(40, 80, 40)    // Очень темно-зеленый
	ColorDoor      = tcell.NewRGBColor(200, 150, 50)  // Золотой
	ColorCorridor  = tcell.NewRGBColor(80, 80, 40)    // Темно-желтый
	ColorPlayer    = tcell.NewRGBColor(255, 255, 255) // Белый
	ColorBorder    = tcell.NewRGBColor(150, 150, 150) // Светло-серый
)
