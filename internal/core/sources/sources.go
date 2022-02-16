package sources

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var tileMapCollection = make(map[string]Tilemap)
var shaderCollection = make(map[string]Shader)
var fontCollection = make(map[string]*Font)
var audioCollection = make(map[string]*Audio)

var ResourceNotFoundError error = errors.New("'%s' with path '%s' not found")

func GetTileMap(path string) Tilemap {
	path = filepath.Join("dist/tilemaps", path)

	tileMap, ok := tileMapCollection[path]
	if !ok {
		newTileMap := NewTilemap()
		if err := newTileMap.load(path); err != nil {
			logrus.Fatalln(err, fmt.Sprintf(ResourceNotFoundError.Error(), "tilemap", path))
		}
		return newTileMap
	}
	return tileMap
}

func GetShader(path string) Shader {
	path = filepath.Join("dist/shaders", path)

	shader, ok := shaderCollection[path]
	if !ok {
		newShader := NewShader()
		if err := newShader.load(path); err != nil {
			logrus.Fatalln(err, fmt.Sprintf(ResourceNotFoundError.Error(), "shader", path))
		}
		return newShader
	}
	return shader
}

func GetFont(path string, size int) *Font {
	path = filepath.Join("dist/fonts", path)

	font, ok := fontCollection[path]
	if !ok {
		newFont := new(Font)
		if err := newFont.load(path, size); err != nil {
			logrus.Fatalln(err, fmt.Sprintf(ResourceNotFoundError.Error(), "font", path))
		}
		return newFont
	}
	return font
}

func GetAudio(path string) *Audio {
	path = filepath.Join("dist/audio", path)

	audio, ok := audioCollection[path]
	if !ok {
		newAudio := new(Audio)
		if err := newAudio.load(path); err != nil {
			logrus.Fatalln(err, fmt.Sprintf(ResourceNotFoundError.Error(), "audio", path))
		}
		return newAudio
	}
	return audio
}

// 	IsLoadingEnded() bool

// //Returns a condition of sources' loading process
// func (p *provider) IsLoadingEnded() bool {
// 	if p.audio.IsAllAudioLoaded() {
// 		return true
// 	}
// 	//TODO: include other load checkers
// 	return false
// }
