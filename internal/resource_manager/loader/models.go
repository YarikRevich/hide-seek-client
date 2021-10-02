package loader

import (
	"embed"
	"sync"
)

type Loader func(embed.FS, string, string, *sync.WaitGroup)

type Component struct {
	Embed embed.FS
	Path  string
}
