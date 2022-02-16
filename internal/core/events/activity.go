package events

import "time"

type ActivityManager struct {
	lastActivity time.Time
}

func (am *ActivityManager) SetLastActivity() {
	am.lastActivity = time.Now()
}

func NewActivityManager() *ActivityManager {
	return new(ActivityManager)
}
