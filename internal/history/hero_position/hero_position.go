package heroposition

var lastHeroPositions []struct {
	X, Y float64
}

func SetLastAudioTrackPath(x, y float64) {
	lastHeroPositions = append(lastHeroPositions, struct{X float64; Y float64}{
		x, y,
	})
}

func GetLastAudioTrackPath() struct{ X, Y float64 } {
	return lastHeroPosition
}
