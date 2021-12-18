package statistics

import "github.com/google/uuid"

type Statistics struct {
	//Decimal stats
	Kills, Shoots map[uuid.UUID]int
}

func NewStatistics() *Statistics {
	return new(Statistics)
}
