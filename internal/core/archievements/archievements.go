package archievements

var instance *Archievements

type Archievements struct {
	tenkillsarchievements *TenKillsArchievement
}

func (a *Archievements) TenKillsArchievement() *TenKillsArchievement {
	return a.tenkillsarchievements
}

func (a *Archievements) Update() {

}

func UseArchievements() *Archievements {
	if instance == nil {
		instance = &Archievements{
			tenkillsarchievements: NewTenKillsArchievement(),
		}
	}
	return instance
}
