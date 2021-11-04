package sources

import "path/filepath"

//Connects image size in image collection
//to metadata in metadata collection
func ConnectImageSizeToMetadata() {
	for k, v := range UseSources().Metadata().Collection {
		if len(v.Info.Parent) != 0 {
			dir, _ := filepath.Split(k)
			imageW, imageH := UseSources().Images().GetImage(filepath.Join(dir, v.Info.Parent)).Size()
			if v.Animation.FrameNum != 0 {
				imageW /= int(v.Animation.FrameNum)
			}
			v.Size.Width = float64(imageW)
			v.Size.Height = float64(imageH)
		}
	}
}
