package archievements

type TenKillsArchievement struct{}

func (tka *TenKillsArchievement) IsUnlocked() bool {
	return false
}

func NewTenKillsArchievement() *TenKillsArchievement {
	return new(TenKillsArchievement)
}
