package eq

type slicesInterface interface {
	LenX() int
	LenY() int
	AreEqual(i int) bool
}

func slicesAreEqual(s slicesInterface) bool {
	if s.LenX() != s.LenY() {
		return false
	}
	for i := 0; i < s.LenX(); i++ {
		if !s.AreEqual(i) {
			return false
		}
	}
	return true
}

type intSlices struct {
	x, y []int
}

func (s *intSlices) LenX() int           { return len(s.x) }
func (s *intSlices) LenY() int           { return len(s.y) }
func (s *intSlices) AreEqual(i int) bool { return s.x[i] == s.y[i] }

type byteSlices struct {
	x, y []byte
}

func (s *byteSlices) LenX() int           { return len(s.x) }
func (s *byteSlices) LenY() int           { return len(s.y) }
func (s *byteSlices) AreEqual(i int) bool { return s.x[i] == s.y[i] }

type stringSlices struct {
	x, y []string
}

func (s *stringSlices) LenX() int           { return len(s.x) }
func (s *stringSlices) LenY() int           { return len(s.y) }
func (s *stringSlices) AreEqual(i int) bool { return s.x[i] == s.y[i] }

// IntSlices checks if two slices contain the same ints in the same order.
func IntSlices(x, y []int) bool {
	b := &intSlices{x, y}
	return slicesAreEqual(b)
}

// ByteSlices checks if two slices contain the same bytes in the same order.
func ByteSlices(x, y []byte) bool {
	b := &byteSlices{x, y}
	return slicesAreEqual(b)
}

// StringSlices checks if two slices contain the same strings in the same order.
func StringSlices(x, y []string) bool {
	b := &stringSlices{x, y}
	return slicesAreEqual(b)
}
