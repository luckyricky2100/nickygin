package app

import (
	"errors"
	"time"
)

var durations = [8](time.Duration){
	5 * time.Minute,
	30 * time.Minute,
	12 * time.Hour,
	24 * time.Hour,
	2 * 24 * time.Hour,
	4 * 24 * time.Hour,
	7 * 24 * time.Hour,
	15 * 24 * time.Hour}

const (
	Finished = "Finished"
	NoNext   = "NoNext"
)

func GetNextStep(currentDuration time.Duration) (*time.Duration, error) {
	if currentDuration == durations[len(durations)-1] {
		return nil, errors.New(Finished)
	}
	for i, v := range durations {
		if v == currentDuration {
			return &(durations[i+1]), nil
		}
	}

	return nil, errors.New(NoNext)
}
