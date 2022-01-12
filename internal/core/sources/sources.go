package sources

import (
	"embed"
	"sync"

	"github.com/YarikRevich/hide-seek-client/internal/core/latency"
)

var instance SourcesProvider

type provider struct {
	audio     *Audio
	font      *Font
	images    *Images
	metadata  *Metadata
	shaders   *Shaders
	colliders *Colliders
}

type SourcesProvider interface {
	LoadSources(embed.FS)

	Audio() *Audio
	Font() *Font
	Images() *Images
	Metadata() *Metadata
	Shaders() *Shaders
	Colliders() *Colliders
	IsLoadingEnded() bool
}

//Loads all sources asynchronously
func (p *provider) LoadSources(fs embed.FS) {
	var wg sync.WaitGroup
	wg.Add(6)

	go p.audio.Load(fs, "dist/audio", &wg)
	latency.Seq(func() {
		var s sync.WaitGroup
		s.Add(1)
		go p.images.Load(fs, "dist/images", &wg, &s)
		s.Wait()
	}, func() {
		var s sync.WaitGroup
		s.Add(1)
		go p.metadata.Load(fs, "dist/metadata", &wg, &s)
		s.Wait()
	})
	go p.font.Load(fs, "dist/fonts", &wg)
	go p.shaders.Load(fs, "dist/shaders", &wg)
	go p.colliders.Load(fs, "dist/colliders", &wg)

	wg.Wait()
}

func (p *provider) Audio() *Audio {
	return p.audio
}

func (p *provider) Font() *Font {
	return p.font
}

func (p *provider) Images() *Images {
	return p.images
}

func (p *provider) Metadata() *Metadata {
	return p.metadata
}

func (p *provider) Shaders() *Shaders {
	return p.shaders
}

func (p *provider) Colliders() *Colliders {
	return p.colliders
}

//Returns a condition of sources' loading process
func (p *provider) IsLoadingEnded() bool {
	if p.audio.IsAllAudioLoaded() {
		return true
	}
	//TODO: include other load checkers
	return false
}

func UseSources() SourcesProvider {
	if instance == nil {
		instance = &provider{
			font:      NewFont(),
			metadata:  NewMetadata(),
			images:    NewImages(),
			audio:     NewAudio(),
			shaders:   NewShaders(),
			colliders: NewColliders(),
		}
	}
	return instance
}
