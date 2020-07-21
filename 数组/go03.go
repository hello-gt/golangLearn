package main

const (
	WIDTH  = 1930
	HEIGHT = 1080
)

type pixel int

var scree [WIDTH][HEIGHT]pixel

func main() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			scree[x][y] = 0
		}
	}
}
