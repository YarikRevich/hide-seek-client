package trackmanager

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
)

//Manager for player
//Tracks the sequnce of played songs
type TrackManager struct {
	trackCollection []struct {
		track, prev *sources.Track
	}
}

//Returns a track which it connected to the passed trackPath
func (tm *TrackManager) NextToPlayAfter(trackPath string) *sources.Track {
	for _, v := range tm.trackCollection {
		if v.track.TrackPath == trackPath {
			return v.prev
		}
	}
	return nil
}

//Pushed a new track to the sequence connecting
//it to already present ones
func (tm *TrackManager) Push(track *sources.Track) {
	if len(tm.trackCollection) != 0 {
		tm.trackCollection = append(tm.trackCollection, struct {
			track *sources.Track
			prev  *sources.Track
		}{
			track: track,
			prev:  tm.trackCollection[len(tm.trackCollection)-1].track,
		})
	} else {
		tm.trackCollection = append(tm.trackCollection, struct {
			track *sources.Track
			prev  *sources.Track
		}{track: track})
	}
}

//Removes passed track from the player sequence
func (tm *TrackManager) Remove(track *sources.Track) {
	for i := 0; i < len(tm.trackCollection); i++ {
		if tm.trackCollection[i].track == track {
			if i != len(tm.trackCollection)-1 && i != 0 {
				tm.trackCollection[i+1].prev = tm.trackCollection[i-1].track
			}
			tm.trackCollection = append(tm.trackCollection[:i], tm.trackCollection[i+1:]...)
		}
	}
}

func (tm *TrackManager) RemoveAll() {
	tm.trackCollection = tm.trackCollection[:0]
}

func (tm *TrackManager) Find(trackPath string) *sources.Track {
	for _, v := range tm.trackCollection {
		if v.track.TrackPath == trackPath {
			return v.track
		}
	}
	return nil
}

//Returns track which is not paused and is being currently played
func (tm *TrackManager) TopCurrentTrack() *sources.Track {
	for i := len(tm.trackCollection) - 1; i != 0; i-- {
		if v := tm.trackCollection[i]; v.track.Ctrl.Paused {
			return v.track
		}
	}
	return nil
}

func New() *TrackManager {
	return new(TrackManager)
}
