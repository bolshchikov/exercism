package letter

func mergeMaps(mapA, mapB FreqMap) FreqMap {
	for k, frB := range mapB {
		frA, ok := mapA[k]
		if ok {
			mapA[k] = frA + frB
		} else {
			mapA[k] = frB
		}
	}
	return mapA
}

// ConcurrentFrequency calculates letter frequency concurrently
func ConcurrentFrequency(input []string) FreqMap {
	res := FreqMap{}
	c := make(chan FreqMap)
	for _, text := range input {
		go func(t string) {
			c <- Frequency(t)
		}(text)
	}
	received := 0
	for received != 3 {
		mergeMaps(res, <-c)
		received++
	}
	return res
}
