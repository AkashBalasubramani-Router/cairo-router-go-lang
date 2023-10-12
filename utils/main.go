package main

import (
	"fmt"
	"math/big"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/utils"
)

func BigIntToHexU128Parts(num *big.Int) (lowHex, highHex string) {
	// Create a big.Int with a value of 2^128
	u128 := new(big.Int)
	u128.Exp(big.NewInt(2), big.NewInt(128), nil)

	// Calculate the low part
	low := new(big.Int).And(num, new(big.Int).Sub(u128, big.NewInt(1)))
	lowHex = fmt.Sprintf("0x%032x", low)

	// Calculate the high part
	high := new(big.Int).Rsh(num, 128)
	highHex = fmt.Sprintf("0x%04x", high)

	return lowHex, highHex

}

func BigIntArrToFelt(val []*big.Int) []*felt.Felt {
	var feltArr []*felt.Felt
	for _, hex := range val {
		felt := utils.BigIntToFelt(hex)
		feltArr = append(feltArr, felt)
	}
	return feltArr
}

func main() {
	// Create a big integer
	bigIntValue := new(big.Int)
	bigIntValue.SetString("12998997974179775878787878787638634793485935040902930582895377529925928350290", 10)

	bigIntArray := []*big.Int{
		big.NewInt(123),
		big.NewInt(456),
		big.NewInt(789),
	}

	// Split the big integer into low and high U128 parts in hexadecimal format
	// lowHex, highHex := BigIntToHexU128Parts(bigIntValue)

	yes := BigIntArrToFelt(bigIntArray)

	fmt.Println("Low:", yes)
	// fmt.Println("High:", highHex)
}
