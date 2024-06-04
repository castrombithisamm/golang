package main

// import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Allocate memory for the image
	image := make([][]uint8, dy)

	// Loop through rows
	for y := 0; y < dy; y++ {
		// Allocate memory for each row
		image[y] = make([]uint8, dx)
		// Loop through columns
		for x := 0; x < dx; x++ {
			// Generate pixel intensity using (x+y)/2 function
			image[y][x] = uint8((x + y) / 2)
		}
	}
	return image
}

func main() {
	// Display the image
	// pic.Show(Pic)
}
