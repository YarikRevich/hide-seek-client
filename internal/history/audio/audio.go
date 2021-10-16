package audio

var lastAudioPath string

func SetLastAudioTrackPath(p string) {
	lastAudioPath = p
}

func GetLastAudioTrackPath() string {
	return lastAudioPath
}
