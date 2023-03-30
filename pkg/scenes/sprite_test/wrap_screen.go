package sprite_test

type Rectangle struct {
	X, Y, Width, Height float32
}

func WrapCoordinates(x, y float32, rect Rectangle) (float32, float32) {
	// Calculate the minimum and maximum values for x and y within the rectangle
	minX := rect.X
	maxX := rect.X + rect.Width
	minY := rect.Y
	maxY := rect.Y + rect.Height

	// Wrap the x coordinate within the rectangle
	for x < minX {
		x += rect.Width
	}
	for x >= maxX {
		x -= rect.Width
	}

	// Wrap the y coordinate within the rectangle
	for y < minY {
		y += rect.Height
	}
	for y >= maxY {
		y -= rect.Height
	}

	return x, y
}
