package assets

import "embed"

var (
	//go:embed dist
	Assets embed.FS
)
