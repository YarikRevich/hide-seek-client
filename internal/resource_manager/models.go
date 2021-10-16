package resource_manager

import (
	"embed"
	"strings"
	"sync"
)

type Loader func(embed.FS, string, string, *sync.WaitGroup)

type Component struct {
	Embed embed.FS
	Path  string
}

func (c *Component) SeparatePath() []string {
	return strings.Split(c.Path, ":")
}
