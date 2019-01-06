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

func calcFrequency(input string) <-chan FreqMap {
	c := make(chan FreqMap)
	go func() {
		c <- Frequency(input)
	}()
	return c
}

// ConcurrentFrequency calculates letter frequency concurrently
func ConcurrentFrequency(input []string) FreqMap {
	res := FreqMap{}
	for _, passage := range input {
		mergeMaps(res, <-calcFrequency(passage))
	}
	return res
}
