// package main

// import (
// 	"encoding/hex"
// 	"fmt"
// 	"math/big"
// 	"strconv"

// 	starknetgo "github.com/NethermindEth/starknet.go"
// 	"github.com/NethermindEth/starknet.go/types"
// )

// //	func stringToHex(inputStr string) string {
// //		strHex := hex.EncodeToString([]byte(inputStr))
// //		return "0x" + strHex
// //	}
// func stringToHex(input string) (string, error) {
// 	// Convert string to integer
// 	i, err := strconv.Atoi(input)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Convert integer to hexadecimal string with leading '0x'
// 	hexStr := fmt.Sprintf("0x%x", i)

// 	return hexStr, nil
// }

// // func stringidToHex(input string) string {

// // 	// Convert string to integer
// // 	i, err := strconv.Atoi(input)
// // 	if err != nil {
// // 		fmt.Println("Error:", err)
// // 		return
// // 	}

// // 	// Convert integer to hexadecimal string
// // 	hexStr := fmt.Sprintf("0x%02x", i)

// // 	fmt.Println(hexStr)
// // }

// func bigIntToHex(bigInt *big.Int) string {
// 	if bigInt.Sign() == 0 {
// 		return "0x0"
// 	}

// 	bigHex := hex.EncodeToString(bigInt.Bytes())
// 	return "0x" + bigHex
// }

// func boolToHex(input bool) (string, error) {
// 	var intValue string
// 	if input {
// 		intValue = "0x01"
// 	} else {
// 		intValue = "0x00"
// 	}

// 	return intValue, nil
// }

// // func serializing(IACK_method_name string, chainId string, requestIdentifier big.Int, ackrequestIdentifier big.Int, destchainId string, requestsender string, execdata *big.Int, execflag bool) {
// // 	var serializedData []string
// // 	serializedData = append(serializedData, IACK_method_name)
// // 	serialized_chainId := stringToHex(chainId)
// // 	serializedData = append(serializedData, serialized_chainId)
// // 	serialized_requestIdentifier := bigIntToHex(&requestIdentifier)
// // 	serializedData = append(serializedData, serialized_requestIdentifier)
// // 	serialized_ackrequestIdentifier := bigIntToHex(&ackrequestIdentifier)
// // 	serializedData = append(serializedData, serialized_ackrequestIdentifier)
// // 	serialized_destchainId := stringToHex(destchainId)
// // 	serializedData = append(serializedData, serialized_destchainId)
// // 	serialized_requestSender := stringToHex(requestsender)
// // 	serializedData = append(serializedData, serialized_requestSender)
// // 	serialized_execdata := bigIntToHex(execdata)
// // 	serializedData = append(serializedData, serialized_execdata)
// // 	serialized_execflag, _ := boolToHex(execflag)
// // 	serializedData = append(serializedData, serialized_execflag)

// // 	fmt.Println(serializedData)

// // }
// // func serializing(IReceive_method_name string, route_amount *big.Int, requestIdentifier *big.Int, requestTimestamp *big.Int) {

// func serializing(IReceive_method_name string, route_amount *big.Int, requestIdentifier *big.Int, requestTimestamp *big.Int, srcchainId string, route_recipient string, destchainId string, asmaddress string, requestSender string, handleraddress string, packet []string, isReadcall bool) {
// 	var serializedData []string
// 	// serialized_Ireceive := stringToHex(IReceive_method_name)
// 	// serializedData = append(serializedData, serialized_Ireceive)
// 	serializedData = append(serializedData, "0x00")
// 	serializedData = append(serializedData, IReceive_method_name)
// 	serialized_route_amount := bigIntToHex(route_amount)
// 	serializedData = append(serializedData, serialized_route_amount)
// 	serializedData = append(serializedData, "0x00")
// 	serialized_requestIdentifier := bigIntToHex(requestIdentifier)
// 	serializedData = append(serializedData, serialized_requestIdentifier)
// 	serializedData = append(serializedData, "0x00")
// 	serialized_requestTimestamp := bigIntToHex(requestTimestamp)
// 	serializedData = append(serializedData, serialized_requestTimestamp)
// 	serializedData = append(serializedData, "0x00")
// 	serialized_srcchainId, _ := stringToHex(srcchainId)
// 	serializedData = append(serializedData, serialized_srcchainId)
// 	if route_recipient == "0" {
// 		serialized_route_recipient, err := stringToHex(route_recipient)
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return
// 		}
// 		serializedData = append(serializedData, serialized_route_recipient)
// 	} else {
// 		serializedData = append(serializedData, route_recipient)
// 	}

// 	serialized_destchainId, _ := stringToHex(destchainId)
// 	serializedData = append(serializedData, serialized_destchainId)
// 	// serialized_asmaddress := stringToHex(asmaddress)
// 	serializedData = append(serializedData, asmaddress)
// 	serialized_requestSender, _ := stringToHex(requestSender)
// 	serializedData = append(serializedData, serialized_requestSender)
// 	// serialized_handleraddress := stringToHex(handleraddress)
// 	serializedData = append(serializedData, asmaddress)

// 	serialized_length, _ := stringToHex(strconv.Itoa((len(packet))))

// 	if len(packet) == 0 {
// 		serializedData = append(serializedData, serialized_length)
// 	} else {
// 		serializedData = append(serializedData, serialized_length)
// 		for _, item := range packet {
// 			serializedData = append(serializedData, item)
// 			fmt.Println("Item:", item)

// 		}
// 	}

// 	// // serializedData = append(serializedData, serialized_length)

// 	serialized_isReadcall, _ := boolToHex(isReadcall)
// 	serializedData = append(serializedData, serialized_isReadcall)

// 	// return serializedData

// 	// Hash the 0th element outside the loop
// 	// hash, err := starknetgo.Curve.PedersenHash([]*big.Int{big.NewInt(0), types.HexToBN("0x00")})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash1, err := starknetgo.Curve.PedersenHash([]*big.Int{hash, types.HexToBN(IReceive_method_name)})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// //zzzzzzzz
// 	// hash2, err := starknetgo.Curve.PedersenHash([]*big.Int{hash1, types.HexToBN(serialized_route_amount)})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash3, err := starknetgo.Curve.PedersenHash([]*big.Int{hash2, types.HexToBN("0x00")})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash4, err := starknetgo.Curve.PedersenHash([]*big.Int{hash3, types.HexToBN(serialized_requestIdentifier)})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash5, err := starknetgo.Curve.PedersenHash([]*big.Int{hash4, types.HexToBN("0x00")})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash6, err := starknetgo.Curve.PedersenHash([]*big.Int{hash5, types.HexToBN(serialized_requestTimestamp)})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash7, err := starknetgo.Curve.PedersenHash([]*big.Int{hash6, types.HexToBN("0x00")})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash8, err := starknetgo.Curve.PedersenHash([]*big.Int{hash7, types.HexToBN(serialized_srcchainId)})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash9, err := starknetgo.Curve.PedersenHash([]*big.Int{hash8, types.HexToBN(serializedData[4])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash10, err := starknetgo.Curve.PedersenHash([]*big.Int{hash9, types.HexToBN(serializedData[5])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash11, err := starknetgo.Curve.PedersenHash([]*big.Int{hash10, types.HexToBN(serializedData[6])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash12, err := starknetgo.Curve.PedersenHash([]*big.Int{hash11, types.HexToBN(serializedData[7])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash13, err := starknetgo.Curve.PedersenHash([]*big.Int{hash12, types.HexToBN(serializedData[8])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash14, err := starknetgo.Curve.PedersenHash([]*big.Int{hash13, types.HexToBN(serializedData[9])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash15, err := starknetgo.Curve.PedersenHash([]*big.Int{hash14, types.HexToBN(serializedData[10])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash16, err := starknetgo.Curve.PedersenHash([]*big.Int{hash15, types.HexToBN(serializedData[11])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// //comehere
// 	// fmt.Printf("Final Hash with 10: %s\n", types.BigToHex(hash16))
// 	// fmt.Printf("Final Hash with 10: %s\n", serializedData[10])

// 	hash, err := starknetgo.Curve.PedersenHash([]*big.Int{big.NewInt(0), types.HexToBN(serializedData[0])})
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// fmt.Println("Hash of the first element:", hash)

// 	for i := 1; i < len(serializedData); i++ {
// 		hash, err = starknetgo.Curve.PedersenHash([]*big.Int{hash, types.HexToBN(serializedData[i])})
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return
// 		}
// 	}

// 	hash1, err := starknetgo.Curve.PedersenHash([]*big.Int{hash, big.NewInt(10)})
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println("Hash of all elements:", types.BigToHex(hash1))

// 	// // // Print the hash result of the 0th element
// 	// fmt.Printf("Hash of %s: %s\n", serializedData[0], hash.String())

// 	// // Hash the remaining elements starting from the first element
// 	// for _, data := range serializedData[1:] {
// 	// 	newHash, err := starknetgo.Curve.PedersenHash([]*big.Int{hash, types.HexToBN(data)})
// 	// 	if err != nil {
// 	// 		fmt.Println("Error:", err)
// 	// 		return
// 	// 	}
// 	// 	hash = newHash

// 	// 	// Print the hash result
// 	// 	fmt.Printf("Hash of %s: %s\n", data, hash.String())
// 	// }

// 	// // Hash the result with 10
// 	// finalHash, err := starknetgo.Curve.PedersenHash([]*big.Int{hash, big.NewInt(10)})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }

// 	// // Print the final hash result
// 	// fmt.Printf("Final Hash with 10: %s\n", finalHash.String())

// }

// func main() {

// 	hash, err := starknetgo.Curve.PedersenHash([]*big.Int{big.NewInt(0), types.HexToBN("0x64")})
// 	fmt.Println(hash)
// 	// Create a big.Int
// 	value := new(big.Int)
// 	value.SetString("100", 10)

// 	hexOutput := bigIntToHex(value)
// 	fmt.Println(hexOutput)

// 	value1 := false

// 	hexOutput1, err := boolToHex(value1)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println(hexOutput1)

// 	inputString := "Akash Balasubramani"
// 	hexOutput2, _ := stringToHex(inputString)
// 	fmt.Println(hexOutput2) // Output: 48656c6c6f2c20576f726c6421
// 	requestTimestamp := big.NewInt(10)
// 	RouteAmount := big.NewInt(0)
// 	// var requestTimestamp big.Int
// 	// requestTimestamp.SetInt64(10)
// 	// var RouteAmount big.Int
// 	// RouteAmount.SetInt64(0)

// 	// serializing("0x6952656365697665000000000000000000000000000000000000000000000000", &RouteAmount, &RouteAmount, &requestTimestamp)
// 	stringSlice := []string{"0x48656c6c6f20576f726c64"}
// 	serializing("0x695265636569766500000000000000", RouteAmount, RouteAmount, requestTimestamp, "1", "0", "1", "0x03c12f1e651adbbb65ed4ea4a6524094ad6a8951740a6f6ed229d28462133ee2", "0", "0x03c12f1e651adbbb65ed4ea4a6524094ad6a8951740a6f6ed229d28462133ee2", stringSlice, false)
// 	// ress, _ := stringToHex("100")
// 	// fmt.Println(ress)

// }
