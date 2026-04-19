package algorithm

type SM2Result struct {
	Repetition   int
	Easiness     float64
	IntervalDays int
}

func Calculate(repetition int, easiness float64, intervalDays int, rating int) SM2Result {
	if rating < 3 {
		return SM2Result{
			Repetition:   0,
			Easiness:     easiness,
			IntervalDays: 1,
		}
	}

	newEasiness := easiness + (0.1 - float64(5-rating)*(0.08+float64(5-rating)*0.02))
	if newEasiness < 1.3 {
		newEasiness = 1.3
	}

	var newInterval int
	if repetition == 0 {
		newInterval = 1
	} else if repetition == 1 {
		newInterval = 6
	} else {
		newInterval = int(float64(intervalDays) * newEasiness)
	}

	return SM2Result{
		Repetition:   repetition + 1,
		Easiness:     newEasiness,
		IntervalDays: newInterval,
	}
}
