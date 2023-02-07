package main

import (
	//importerar nödvändiga bibliotek
	"fmt"
	"math/big"
	"strconv"
)

func fac(n int64) *big.Int {
	var i int64
	if n == 0 {
		return big.NewInt(1)
	}
	ret := big.NewInt(1)
	for i = n; i > 1; i-- {
		ret = ret.Mul(ret, big.NewInt(i))
	}
	return ret
}

//Jämför värdena från funktionen fac med de faktiska värdena och slår larm om något går fel
func tfac() {
	var i int64
	actual := []int64{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
	for i = 0; i < 10; i++ {
		facint := fac(i).Int64()
		if facint != actual[i] {
			panic("Fel i fakultetsfunktionen!")
		}
	}
}

func main() {
	//indata
	var input string
	fmt.Println("Vilket tal?")
	fmt.Scanln(&input)
	//omvandlar indata från string till int64
	input2, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		panic("Fel: indata måste vara positiva heltal.")
	}
	ret := fac(int64(input2))
	fmt.Println(ret)

	//Test
	tfac()
}
