package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

func flip_a_coin() int64 {
	res, _ := rand.Int(rand.Reader, big.NewInt(2))
	res = big.NewInt(0).Add(res, big.NewInt(2))
	return res.Int64()
}

func make_a_yao() int64 {
	var res int64 = 0
	for i := 1; i <= 3; i++ {
		res = res + flip_a_coin()
	}
	return res
}

func make_a_hexagram() [6]int64 {
	var a [6]int64
	for i := 0; i < 6; i++ {
		a[i] = make_a_yao()
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter question: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	var six = "--- x ---"
	var seven = "---------"
	var eight = "---   ---"
	var nine = "----0----"

	var a = make_a_hexagram()

	for i := 5; i >= 0; i-- {
		switch a[i] {
		case 6:
			fmt.Println(six)
		case 7:
			fmt.Println(seven)
		case 8:
			fmt.Println(eight)
		case 9:
			fmt.Println(nine)
		default:
			fmt.Println("oops")
		}
	}

	fmt.Println(a)
}
