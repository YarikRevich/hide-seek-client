package sources

import (
	"embed"
	"fmt"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

type Shaders struct {
	sync.Mutex

	Collection map[string]*ebiten.Shader
}

func (i *Shaders) loadFile(fs embed.FS, path string) {
	file, err := fs.ReadFile(path)
	if err != nil {
		logrus.Fatal("error happened opening shader file from embedded fs", err)
	}

	reg := regexp.MustCompile(`\.[a-z0-9]*$`)
	if reg.MatchString(path) {
		i.Lock()
		shader, err := ebiten.NewShader(file)
		if err != nil {
			logrus.Fatal(err)
		}
		i.Collection[reg.Split(path, -1)[0]] = shader
		i.Unlock()
	}
}

func (i *Shaders) Load(fs embed.FS, path string, wg *sync.WaitGroup) {
	NewParser(fs, path, i.loadFile).Parse()
	wg.Done()
}

func (i *Shaders) GetShader(path string) *ebiten.Shader {
	path = filepath.Join("assets/shaders", path)

	shader, ok := i.Collection[path]
	if !ok {
		logrus.Fatal(fmt.Sprintf("shader with path '%s' not found", path))
	}

	return shader
}

func NewShaders() *Shaders {
	return &Shaders{Collection: make(map[string]*ebiten.Shader)}
}
