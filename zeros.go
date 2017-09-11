package randomtext

// ZeroGenerator generates zeros
type ZeroGenerator struct {
}

// Generate generates zeros
func (z ZeroGenerator) Generate() string {
	return "0"
}
