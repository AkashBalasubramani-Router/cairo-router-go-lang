// package main

// import (
// 	"encoding/hex"
// 	"fmt"
// 	"math/big"
// 	"strconv"

// 	// "strings"

// 	starknetgo "github.com/NethermindEth/starknet.go"
// 	"github.com/NethermindEth/starknet.go/types"
// )

// //	func stringToHex(inputStr string) string {
// //		strHex := hex.EncodeToString([]byte(inputStr))
// //		return "0x" + strHex
// //	}

// func calculate_packet(packet []string) ([]string, error) {
// 	// var totalLength int
// 	var serializedData []string

// 	if len(packet) == 0 {
// 		// totalLength = len(packet)
// 		// serialized_length, _ := stringToHex(strconv.Itoa(totalLength))
// 		// serializedData = append(serializedData, serialized_length)
// 		serializedData = append(serializedData, "0x00")
// 	} else {

// 		// serialized_length, _ := stringToHex(strconv.Itoa(totalLength))

// 		// // Append the calculated total length to serializedData.
// 		// serializedData = append(serializedData, serialized_length)
// 		serializedData = append(serializedData, fmt.Sprintf("0x%x", len(packet)))

// 		for _, item := range packet {
// 			if num, err := strconv.Atoi(item); err == nil {
// 				// If item is a number, convert to hex and append.
// 				hexNum := fmt.Sprintf("0x%x", num)
// 				serializedData = append(serializedData, hexNum)
// 				// serializedData = append(serializedData, "0x00")
// 			} else {
// 				// If item is not a number, hex encode and append.
// 				serializedData = append(serializedData, item)
// 			}
// 			// fmt.Println("Item:", item)
// 		}
// 	}
// 	return serializedData, nil
// }

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

// // func serializing(IReceive_method_name string, route_amount *big.Int,srcchainId string requestIdentifier *big.Int, requestTimestamp *big.Int, srcchainId string, route_recipient string, destchainId string, asmaddress string, requestSender string, handleraddress string, packet []string, isReadcall bool, MSG_PREFIX string) {
// 	func serializing(route_amount *big.Int,srcchainId string,depositId *big.int,destToken string,recipient string,depositor []string,contract_address string) {
// 	var serializedData []string
// 	// serialized_Ireceive := stringToHex(IReceive_method_name)

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
// 	if asmaddress == "0" {
// 		serialized_asm_address, err := stringToHex(asmaddress)
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return
// 		}
// 		serializedData = append(serializedData, serialized_asm_address)
// 	} else {
// 		serializedData = append(serializedData, asmaddress)
// 	}

// 	if requestSender == "0" {
// 		serialized_requestSender, err := stringToHex(requestSender)
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return
// 		}
// 		serializedData = append(serializedData, serialized_requestSender)
// 	} else {
// 		serializedData = append(serializedData, requestSender)
// 	}

// 	// // serialized_handleraddress,_ := stringToHex(handleraddress)
// 	if handleraddress == "0" {
// 		serialized_handleraddress, err := stringToHex(handleraddress)
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return
// 		}
// 		serializedData = append(serializedData, serialized_handleraddress)
// 	} else {
// 		serializedData = append(serializedData, handleraddress)
// 	}
// 	// serializedData = append(serializedData, handleraddress)

// 	// // if len(packet) == 0 {
// 	// // 	serializedData = append(serializedData, serialized_length)
// 	// // } else {
// 	// // 	serializedData = append(serializedData, serialized_length)

// 	// // 	for _, item := range packet {
// 	// // 		if num, err := strconv.Atoi(item); err == nil {
// 	// // 			// If item is a number, convert it to hex and append to serializedData.
// 	// // 			hexNum := fmt.Sprintf("0x%x", num)
// 	// // 			serializedData = append(serializedData, hexNum)
// 	// // 			serializedData = append(serializedData, "0x00")
// 	// // 		} else {
// 	// // 			// If item is not a number, hex encode it and append to serializedData.
// 	// // 			// hexStr := hex.EncodeToString([]byte(item))
// 	// // 			serializedData = append(serializedData, item)
// 	// // 		}
// 	// // 		fmt.Println("Item:", item)
// 	// // 	}
// 	// // }

// 	// var totalLength int

// 	// if len(packet) == 0 {
// 	// 	totalLength = len(packet)
// 	// } else {
// 	// 	totalLength = len(packet)

// 	// 	for _, item := range packet {
// 	// 		if _, err := strconv.Atoi(item); err == nil {
// 	// 			// If item is a number, increase totalLength by 1.
// 	// 			totalLength++
// 	// 		}
// 	// 	}
// 	// }
// 	// serialized_length, _ := stringToHex(strconv.Itoa(totalLength))

// 	// // Append the calculated total length to serializedData.
// 	// serializedData = append(serializedData, serialized_length)

// 	for _, item := range packet {
// 		// if num, err := strconv.Atoi(item); err == nil {
// 		// If item is a number, convert to hex and append.
// 		// hexNum := fmt.Sprintf("0x%x", num)
// 		// serializedData = append(serializedData, hexNum)
// 		// serializedData = append(serializedData, "0x00")
// 		// } else {
// 		// If item is not a number, hex encode and append.
// 		serializedData = append(serializedData, item)
// 		// }

// 	}

// 	// // // // // serializedData = append(serializedData, serialized_length)

// 	serialized_isReadcall, _ := boolToHex(isReadcall)
// 	serializedData = append(serializedData, serialized_isReadcall)

// 	// return serializedData

// 	// Hash the 0th element outside the loop
// 	// hash, err := starknetgo.Curve.PedersenHash([]*big.Int{big.NewInt(0), types.HexToBN(serializedData[0])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash1, err := starknetgo.Curve.PedersenHash([]*big.Int{hash, types.HexToBN(serializedData[1])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// //zzzzzzzz
// 	// hash2, err := starknetgo.Curve.PedersenHash([]*big.Int{hash1, types.HexToBN(serializedData[2])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash3, err := starknetgo.Curve.PedersenHash([]*big.Int{hash2, types.HexToBN(serializedData[3])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash4, err := starknetgo.Curve.PedersenHash([]*big.Int{hash3, types.HexToBN(serializedData[4])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash5, err := starknetgo.Curve.PedersenHash([]*big.Int{hash4, types.HexToBN(serializedData[5])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash6, err := starknetgo.Curve.PedersenHash([]*big.Int{hash5, types.HexToBN(serializedData[6])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash7, err := starknetgo.Curve.PedersenHash([]*big.Int{hash6, types.HexToBN(serializedData[7])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash8, err := starknetgo.Curve.PedersenHash([]*big.Int{hash7, types.HexToBN(serializedData[8])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash9, err := starknetgo.Curve.PedersenHash([]*big.Int{hash8, types.HexToBN(serializedData[9])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash10, err := starknetgo.Curve.PedersenHash([]*big.Int{hash9, types.HexToBN(serializedData[10])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash11, err := starknetgo.Curve.PedersenHash([]*big.Int{hash10, types.HexToBN(serializedData[11])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash12, err := starknetgo.Curve.PedersenHash([]*big.Int{hash11, types.HexToBN(serializedData[12])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash13, err := starknetgo.Curve.PedersenHash([]*big.Int{hash12, types.HexToBN(serializedData[13])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash14, err := starknetgo.Curve.PedersenHash([]*big.Int{hash13, types.HexToBN(serializedData[14])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash15, err := starknetgo.Curve.PedersenHash([]*big.Int{hash14, types.HexToBN(serializedData[15])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// hash16, err := starknetgo.Curve.PedersenHash([]*big.Int{hash15, types.HexToBN(serializedData[16])})
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }

// 	hash, err := starknetgo.Curve.PedersenHash([]*big.Int{big.NewInt(0), types.HexToBN(serializedData[0])})
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println("messageHash:", types.BigToHex(hash))

// 	// fmt.Println("Hash of the first element:", hash)

// 	for i := 1; i < len(serializedData); i++ {
// 		if i == len(serializedData) {
// 			break
// 		}

// 		hash, err = starknetgo.Curve.PedersenHash([]*big.Int{hash, types.HexToBN(serializedData[i])})
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return
// 		}
// 	}
// 	fmt.Println("messageHash:", types.BigToHex(hash))

// 	hash1, err := starknetgo.Curve.PedersenHash([]*big.Int{hash, big.NewInt(10)})
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println("Hash of all elements:", types.BigToHex(hash1))
// 	//comehere
// 	// fmt.Printf("Final Hash with 10: %s\n", types.BigToHex(hash16))
// 	// fmt.Printf("Final Hash with 10: %s\n", serializedData[10])

// 	// // var serializedData2 []string
// 	hexString := hex.EncodeToString([]byte(MSG_PREFIX))
// 	hexWithPrefix := "0x" + hexString
// 	fmt.Printf("Calculation------- %s\n", hexWithPrefix)
// 	// // serializedData2 = append(serializedData2, hexWithPrefix)
// 	// // // serializedData2 = append(serializedData2, hash1)

// 	hash2, err := starknetgo.Curve.PedersenHash([]*big.Int{big.NewInt(0), types.HexToBN(hexWithPrefix)})
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Printf("Calculation %s\n", types.BigToHex(hash2))

// 	hash3, err := starknetgo.Curve.PedersenHash([]*big.Int{hash2, types.HexToBN(types.BigToHex(hash1))})
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	// parts := strings.SplitN(types.BigToHex(hash3), "x", 2)
// 	// paddedHex := "0x" + fmt.Sprintf("%02s", parts[1])

// 	fmt.Printf("Digest Calcu %s\n", types.BigToHex(hash3))

// }

// func main() {

// 	Packet := []string{"0x3038ae29ffd0258880b34b9ffdd37a02bd1b7a7e15ff183c69a0a1c18d30998", "100"}
// 	res, _ := calculate_packet(Packet)
// 	// joinedRes := strings.Join(res, ", ")
// 	fmt.Println("Result --->", res)

// 	requestTimestamp := big.NewInt(1752503506)
// 	requestIdentifier := big.NewInt(7)
// 	RouteAmount := big.NewInt(0)

// 	// serializing("0x6952656365697665000000000000000000000000000000000000000000000000", &RouteAmount, &RouteAmount, &requestTimestamp)
// 	// stringSlice := []string{"0x02a8a373dba71fa37deb1db033dd66e5aa7013f3baade60d844deaffff61c7d7", "0x64", "0x0"}
// 	serializing("0x695265636569766500000000000000", RouteAmount, requestIdentifier, requestTimestamp, "1", "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266", "1", "0x03c12f1e651adbbb65ed4ea4a6524094ad6a8951740a6f6ed229d28462133ee2", "0", "0x015583814f3374302630072712cb894f98079bdfcba84c788363848da54a2610", res, false, "x19Ethereum Signed Message:\n31")
// 	// serializing("0x695265636569766500000000000000", RouteAmount, RouteAmount, requestTimestamp, "1", "0", "1", "0x03c12f1e651adbbb65ed4ea4a6524094ad6a8951740a6f6ed229d28462133ee2", "0", "0x03c12f1e651adbbb65ed4ea4a6524094ad6a8951740a6f6ed229d28462133ee2", stringSlice, false, "x19Ethereum Signed Message:\n31")
// }
