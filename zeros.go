package randomtext

// ZeroGenerator returns a function that generates zeros
func ZeroGenerator() func() string {
	return func() string {
		return "0"
	}
}
