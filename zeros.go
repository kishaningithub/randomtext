package randomtext

// ZeroGenerator returns a function that generates zeros
func ZeroGenerator() func() string {
	return randomGenerator([]string{"0"})
}
