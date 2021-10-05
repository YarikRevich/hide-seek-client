package fonts

import (
	// "golang.org/x/image/font"
	// "golang.org/x/image/font/opentype"

	// "github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/text"
	// "golang.org/x/image/font/basicfont"
	"fmt"
	"strconv"

	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/font_loader/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
	"golang.org/x/image/font"
)

func GetFont(m models.Metadata)font.Face{
	size := strconv.FormatFloat(m.Button.Font, 'f', 0, 64)
	return collection.GetFont(fmt.Sprintf("assets/fonts/base_%s", size))
}


// func {
	

// 	text.Draw(screen, "", &basicfont.Face{}, 0, 0, nil)
// }