package collection

import (
	"fmt"

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