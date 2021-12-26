package transition

// import "github.com/YarikRevich/hide-seek-client/internal/core/sources"

var instance *TransitionPool

type TransitionPool struct {
	//
}

func (tp *TransitionPool) Push() {

}

func (tp *TransitionPool) Process() {

}

func UseTransitionPool() *TransitionPool {
	if instance == nil {
		instance = new(TransitionPool)
	}
	return instance
}

// func Transite(m sources.Model){
// 	tsx := m.Transition.Scale.X
// 	tsy := m.Transition.Scale.Y
// }
