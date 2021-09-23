package screen

const (
	SMALLSCREEN screenType = iota
	FULLSCREEN
)

var (
	screenState *ScreenState
)

type screenType int

type ScreenState struct {
	screenState screenType
}

func (ss *ScreenState) SetState(s screenType) {
	ss.screenState = s
}

func (ss *ScreenState) GetState() screenType {
	return ss.screenState
}

func GetInstance() *ScreenState {
	if screenState == nil {
		return new(ScreenState)
	}
	return screenState
}
