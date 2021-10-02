package animation

import (
	"image"

	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	metadataloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/metadata_loader"
	"github.com/hajimehoshi/ebiten/v2"
)

//func convertRectToStrList(R pixel.Rect) []string {
//	MinX := fmt.Sprintf("%f", R.Min.X)
//	MinY := fmt.Sprintf("%f", R.Min.Y)
//	MaxX := fmt.Sprintf("%f", R.Max.X)
//	MaxY := fmt.Sprintf("%f", R.Max.Y)
//	return []string{
//		MinX, MinY, MaxX, MaxY,
//	}
//}
//
//
//func convertRectToIntList(R pixel.Rect) []float64 {
//
//	return []float64{
//		R.Min.X, R.Min.Y, R.Max.X, R.Max.Y,
//	}
//}
//
//func convertToStrList(intlist []float64) []string {
//	var strlist []string
//	for _, i := range intlist {
//		num := fmt.Sprintf("%f", i)
//		strlist = append(strlist, num)
//	}
//	return strlist
//}
//
//func compareEqualBeetwenSlices(a []string, b []string) bool {
//	for index, value := range a {
//		if value != b[index] {
//			return false
//		}
//	}
//	return true
//}
//
//func convertStringSliceToRect(floatlist []float64) pixel.Rect {
//	return pixel.R(
//		floatlist[0], floatlist[1], floatlist[2], floatlist[3],
//	)
//}
//
//func getFrame(frameWidth float64, frameHeight float64, xGrid float64, yGrid int) pixel.Rect {
//	return pixel.R(
//		float64(xGrid)*frameWidth,
//		float64(yGrid)*frameHeight,
//		float64(xGrid+1)*frameWidth,
//		float64(yGrid+1)*frameHeight,
//	)
//}

func WithAnimation(src *ebiten.Image, m metadataloader.Metadata, p *pc.Animation) *ebiten.Image {
	p.FrameDelayCounter++
	p.FrameDelayCounter %= p.FrameDelay
	if p.FrameDelayCounter == 0 {
		p.FrameCount++
		p.FrameCount %= uint32(m.Animation.FrameNum)
	}
	sx, sy := int((m.Animation.FrameX+float64(p.FrameCount))*m.Animation.FrameWidth), int(m.Animation.FrameY)
	// fmt.Println(int((m.Animation.FrameX+float64(p.FrameCount))*m.Animation.FrameWidth))

	return src.SubImage(image.Rect(sx, sy, sx+int(m.Animation.FrameWidth), sy+int(m.Animation.FrameHeight))).(*ebiten.Image)
}
