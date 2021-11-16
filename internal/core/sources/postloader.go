package sources

import (
	// "fmt"
	"path/filepath"
	"strings"
)

type PostLoader struct{}

func (p *PostLoader) getBase(path string) string {
	r, _ := filepath.Split(path)
	return r
}

func (p *PostLoader) cleanPrefix(path string) string {
	return strings.SplitN(path, "/", 3)[2]
}

//Connects image size in image collection
//to metadata in metadata collection
func (p *PostLoader) ConnectImageSizeToMetadata() {
	for k, v := range UseSources().Metadata().Collection {
		if len(v.Info.Parent) != 0 {
			base := p.cleanPrefix(p.getBase(k))
			path := filepath.Join(base, v.Info.Parent)
			img := UseSources().Images().GetImage(path)
			
			imageW, imageH := img.Size()
			
			// fmt.Println(imageW, v.Animation.FrameNum - int(v.Animation.FrameWidth), k, v.Animation.FrameWidth)
			if v.Animation.FrameNum != 0 {
				imageW /= int(v.Animation.FrameNum)
			}
	
			// fmt.Println(k, imageW, int(v.Animation.FrameWidth))
			v.Size.Width = float64(imageW)
			v.Size.Height = float64(imageH)	
		}
	}
}

func NewPostLoader() *PostLoader {
	return new(PostLoader)
}
