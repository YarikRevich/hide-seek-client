package printer

import (
	"fmt"
	"strings"

	"github.com/mbndr/figlet4go"
)

func PrintCliMessage(m string) {
	renderer := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorGreen,
	}

	for _, v := range strings.Split(m, "\n") {
		text, err := renderer.RenderOpts(v, options)
		if err != nil {
			panic(err)
		}

		fmt.Print(text)

	}
}
