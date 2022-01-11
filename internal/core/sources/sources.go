package sources

import (
	"embed"
	"sync"
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
	wg.Add(5)

	go p.audio.Load(fs, "dist/audio", &wg)
	go p.metadata.Load(fs, "dist/metadata", &wg)
	go p.images.Load(fs, "dist/images", &wg)
	go p.font.Load(fs, "dist/fonts", &wg)
	go p.shaders.Load(fs, "dist/shaders", &wg)
	go p.colliders.Load(fs, "dist/colliders", &wg)

	wg.Wait()

	NewPostLoader().ConnectImageSizeToMetadata()
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
