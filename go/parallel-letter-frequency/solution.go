package letter

// ConcurrentFrequency calculates letter frequency concurrently
func ConcurrentFrequency(input []string) FreqMap {
	res := FreqMap{}
	c := make(chan FreqMap)
	for _, text := range input {
		go func(t string) {
			c <- Frequency(t)
		}(text)
	}
	for range input {
		for k, v := range <-c {
			res[k] += v
		}
	}
	return res
}
