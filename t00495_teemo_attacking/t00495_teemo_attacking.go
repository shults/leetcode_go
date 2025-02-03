package t00495_teemo_attacking

func findPoisonedDuration(timeSeries []int, duration int) int {
	acc := 0
	acc += duration // last poisoned spit

	for i, upto := 0, len(timeSeries)-1; i < upto; i++ {
		diff := timeSeries[i+1] - timeSeries[i]

		if diff < duration {
			acc += diff
		} else {
			acc += duration
		}
	}

	return acc
}
