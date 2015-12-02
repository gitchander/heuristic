package hexm

type Size struct {
	Dx, Dy, Dz int
}

func (s Size) getError() error {

	if (s.Dx <= 0) || (s.Dy <= 0) || (s.Dz <= 0) {
		return ErrorSizeZeroParameter
	}

	return nil
}

func (s Size) IsValid() bool {
	return (s.getError() == nil)
}

func (s Size) IsEmpty() bool {
	return (s.Dx == 0) || (s.Dy == 0) || (s.Dz == 0)
}

func (s Size) Сontained(c Coord) bool {

	if (c.X < 0) || (c.X >= s.Dx) {
		return false
	}
	if (c.Y < 0) || (c.Y >= s.Dy) {
		return false
	}
	if (c.Z < 0) || (c.Z >= s.Dz) {
		return false
	}

	return true
}
