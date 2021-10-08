package collection

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
	"golang.org/x/image/font"
)

var FontCollection = make(map[string]font.Face)

func GetFont(path string)font.Face{
	i, ok := FontCollection[path]
	if !ok {
		logrus.Fatal(fmt.Sprintf("font with path '%s' not found", path))
	}
	return i
}

func GetFontBySize(rawSize float64)font.Face{
	size := strconv.FormatFloat(rawSize, 'f', 0, 64)
	return GetFont(fmt.Sprintf("assets/fonts/base_%s", size))
}
