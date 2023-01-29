package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(input[0])
	X, _ := strconv.Atoi(input[1])
	Y, _ := strconv.Atoi(input[2])

	for i := 1; i <= N; i++ {
		if i%getAB(X, Y) == 0 {
			fmt.Println("AB")
			continue
		}
		if i%X == 0 {
			fmt.Println("A")
			continue
		}
		if i%Y == 0 {
			fmt.Println("B")
			continue
		}
		fmt.Println("N")
	}
}

func getAB(X int, Y int) int {
	if X%Y == 0 {
		return X
	}
	if Y%X == 0 {
		return Y
	}
	return X * Y
}
