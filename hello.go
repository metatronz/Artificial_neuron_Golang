package main

import (
	"fmt"
)

var problem = [][]byte{
	[]byte("...?.....?............#"),
	[]byte("......?.......?....####"),
	[]byte(".?........?.....#####?#"),
	[]byte(".......?.....#####?####"),
	[]byte("..?..?....#############"),
	[]byte(".......######?######?##"),
	[]byte("....#####?#######?#####"),
	[]byte(".###################?##"),
}

// var problem = [][]byte{
// []byte(".###################?##"),
// []byte("....#####?#######?#####"),
// []byte(".......######?######?##"),
// []byte("..?..?....#############"),
// []byte(".......?.....#####?####"),
// []byte(".?........?.....#####?#"),
// []byte("......?.......?....####"),
// []byte("...?.....?............#"),
// }

var wx, wy, wb float32

func neuron(x, y, bias int) int {
	f := wx*float32(x) + wy*float32(y) + wb*float32(bias)
	if f < 0 {
		return -1
	}
	return 1
}

func learn(x, y, desired int) {
	const learningConst float32 = 0.1
	err := desired - neuron(x, y, 1)
	wx += float32(err*x) * learningConst
	wy += float32(err*y) * learningConst
	wb += float32(err) * 1 * learningConst
}

func stepOfLearning() {
	y := 0
	for line := range problem {
		x := 0
		for b := range problem[line] {
			if problem[line][b] == byte('.') {
				learn(x, y, -1)
			}
			if problem[line][b] == byte('#') {
				learn(x, y, 1)
			}
			x++
		}
		y++
	}
}

func solution() {
	y := 0

	for line := range problem {
		x := 0
		var cp []byte
		for b := range problem[line] {
			if problem[line][b] == '?' {
				r := neuron(x, y, 1)
				if r < 0 {
					cp = append(cp, '!')
				} else {
					cp = append(cp, '%')
				}

			} else {
				cp = append(cp, problem[line][b])
			}
			x++

		}

		y++
		fmt.Println(string(cp))
	}
}

func main() {

	x := 0
	for x < 100 {
		stepOfLearning()
		solution()
		x++
		fmt.Println(x)
	}

	fmt.Printf("wx = %f,wy = %f,wb = %f\n", wx, wy, wb)

}
