package main

const pi = 3.1415

func main() {
	printCircleArea()
	printCircleArea()
}

func printCircleArea(radius int) {
	panic("unimplemented")
}

func calculateCircleArea(radius int) float32 {
	return float32(radius) * float32(radius) * pi
}
