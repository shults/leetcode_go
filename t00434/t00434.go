package t00434

func countSegments(s string) int {
	var segments = 0
	var segmentCounter sCounter = startCounter

	for _, val := range []byte(s) {
		segmentCounter = segmentCounter(val, &segments)
	}

	return segments
}

type sCounter func(ch byte, segments *int) sCounter

func startCounter(ch byte, segments *int) sCounter {
	if ch == ' ' {
		return spaceCounter
	} else {
		*segments++
		return charCounter
	}
}

func spaceCounter(ch byte, segments *int) sCounter {
	if ch == ' ' {
		return spaceCounter
	}

	*segments++
	return charCounter
}

func charCounter(ch byte, segments *int) sCounter {
	if ch == ' ' {
		return spaceCounter
	}

	return charCounter
}
