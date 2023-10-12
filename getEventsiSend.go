// package main

// import (
// 	"context"
// 	"fmt"
// 	"math/big"
// 	"strconv"

// 	"github.com/NethermindEth/juno/core/felt"
// 	"github.com/NethermindEth/starknet.go/rpc"
// 	"github.com/NethermindEth/starknet.go/utils"
// 	ethrpc "github.com/ethereum/go-ethereum/rpc"
// )

// func QueryIsendEvents(startblock uint64, endblock uint64) {
// 	base := "https://starknet-goerli.g.alchemy.com/v2/bbDTAS9Qs68Po3ssBBPVm378Pye6oS2A"
// 	fmt.Println("Starting simpeCall example")
// 	c, err := ethrpc.DialContext(context.Background(), base)
// 	if err != nil {
// 		fmt.Println("Failed to connect to the client, did you specify the url in the .env.mainnet?")
// 		panic(err)
// 	}
// 	clientv02 := rpc.NewProvider(c)
// 	fmt.Println("Established connection with the client")

// 	// startblock_felt := utils.Uint64ToFelt(startblock)
// 	// endblock_felt := utils.Uint64ToFelt(endblock)

// 	// blockhash1, _ := utils.HexToFelt("0x3f56b1eda98fe5f1baf347a4dd62d9870133a9b417812f104724bcda3decfd9")
// 	// blockhash2, _ := utils.HexToFelt("0x1f91401f309e87f5e8113a5d022ce84e45ddcc434f6c7bd21e66d6715b7204c")
// 	// blockhash3, _ := utils.HexToFelt("0x2b388197c31599e3a774c6d8a62a498884b5ec211b10827bfd484bd27aa2a40")
// 	contractAddress, _ := utils.HexToFelt("0x29d74077978f8fd0f680fb17fc439dc38b54b750ad45ca78bb9823e2d5c1c67")
// 	// iSendEvent, _ := utils.HexToFelt("0x20ff74db85f0d9700fcdb580fe4f2dfefd06b4ecde7b8a1753c0f142bda6402")
// 	iReceiveEvent, _ := utils.HexToFelt("0x20ff74db85f0d9700fcdb580fe4f2dfefd06b4ecde7b8a1753c0f142bda6402")
// 	// valsetUpdated, _ := utils.HexToFelt("0x206930eb6d782c3657f345b7cd4fa1c9f14dd743d26cc26f252875bc7b9d275")

// 	eventInput := rpc.EventsInput{
// 		EventFilter: rpc.EventFilter{
// 			FromBlock: rpc.BlockID{
// 				Number: &startblock,
// 			},
// 			ToBlock: rpc.BlockID{
// 				Number: &endblock,
// 			},
// 			Address: contractAddress,
// 			Keys: [][]*felt.Felt{[]*felt.Felt{
// 				iReceiveEvent,
// 			}},
// 		},
// 		ResultPageRequest: rpc.ResultPageRequest{
// 			ChunkSize: 5,
// 		},
// 	}

// 	events, err := clientv02.Events(context.Background(), eventInput)
// 	if err != nil {
// 		fmt.Println("Unsuccessful")
// 		panic(err)
// 	}

// 	for _, emittedEvent := range events.Events {
// 		dataCount := len(emittedEvent.Event.Data)
// 		println("Dtaa count :", dataCount)
// 		if dataCount >= 14 {
// 			// Process low and high pairs of data elements
// 			var version, routeAmount, eventNonce *big.Int
// 			var requestSender, srcChainId, destChainId, routeRecipient string

// 			for i := 0; i < 5; i += 2 {
// 				lowPart := new(big.Int)
// 				lowPart.SetString(emittedEvent.Event.Data[i].String(), 0)

// 				highPart := new(big.Int)
// 				highPart.SetString(emittedEvent.Event.Data[i+1].String(), 0)

// 				combinedBigInt := new(big.Int).Mul(highPart, new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil))
// 				combinedBigInt.Add(combinedBigInt, lowPart)

// 				switch i {
// 				case 0:
// 					version = combinedBigInt
// 				case 2:
// 					routeAmount = combinedBigInt
// 				case 4:
// 					eventNonce = combinedBigInt
// 				}
// 			}

// 			fmt.Println("Version:", version)
// 			fmt.Println("Route Amount:", routeAmount)
// 			fmt.Println("Event Nonce:", eventNonce)
// 			requestSender = emittedEvent.Event.Data[6].String()
// 			stringVariable7 := emittedEvent.Event.Data[7].String()
// 			// stringVariable8 := emittedEvent.Event.Data[8].String()

// 			// Process individual string data elements
// 			stringVariable7Int, err := strconv.Atoi(stringVariable7[2:]) // Removing "0x" prefix
// 			if err != nil {
// 				fmt.Println("Error converting hex for stringVariable7:", err)
// 			}
// 			srcChainId = strconv.Itoa(stringVariable7Int)

// 			stringVariable8Int, err := strconv.Atoi(emittedEvent.Event.Data[8].String()[2:]) // Removing "0x" prefix
// 			if err != nil {
// 				fmt.Println("Error converting hex for stringVariable7:", err)
// 			}
// 			destChainId = strconv.Itoa(stringVariable8Int)

// 			// stringVariable8 := emittedEvent.Event.Data[8].String()
// 			routeRecipient = emittedEvent.Event.Data[9].String()

// 			fmt.Println("requestSender:", requestSender)
// 			fmt.Println("srcChainId:", srcChainId)
// 			fmt.Println("destChainId:", destChainId)
// 			fmt.Println("routeRecipient:", routeRecipient)

// 			// Process new arrays based on element 10 and 12
// 			arraySize1 := new(big.Int)
// 			arraySize1.SetString(emittedEvent.Event.Data[10].String(), 0)

// 			// Construct arrayValue1 based on arraySize1
// 			arrayValue1 := []byte{}
// 			for j := 0; j < int(arraySize1.Int64()); j++ {
// 				value := emittedEvent.Event.Data[11+j].String()
// 				arrayValue1 = append(arrayValue1, value...)
// 			}

// 			// Process new arrays based on element 12
// 			arraySize2Offset := 10 + int(arraySize1.Int64()) + 1
// 			arraySize2 := new(big.Int)
// 			arraySize2.SetString(emittedEvent.Event.Data[arraySize2Offset].String(), 0)

// 			if arraySize2.Int64() > 0 {
// 				// Construct arrayValue2 based on arraySize2
// 				arrayValue2 := []byte{}
// 				for j := 0; j < int(arraySize2.Int64()); j++ {
// 					value := emittedEvent.Event.Data[arraySize2Offset+1+j].String()
// 					arrayValue2 = append(arrayValue2, value...)
// 				}

// 				fmt.Println("New Array 2:", arrayValue2)
// 			} else {
// 				fmt.Println("New Array 2 is empty")
// 			}
// 		}
// 	}
// }

// func main() {
// 	var startblock, endblock uint64
// 	startblock = 851129
// 	endblock = 851131

// 	QueryIsendEvents(startblock, endblock)
// }

// // for _, emittedEvent := range events.Events {
// // 	dataCount := len(emittedEvent.Event.Data)
// // 	if dataCount >= 14 {
// // 		// Process low and high pairs of data elements
// // 		for i := 0; i < 5; i += 2 {
// // 			lowPart := new(big.Int)
// // 			lowPart.SetString(emittedEvent.Event.Data[i].String(), 0)

// // 			highPart := new(big.Int)
// // 			highPart.SetString(emittedEvent.Event.Data[i+1].String(), 0)

// // 			combinedBigInt := new(big.Int).Mul(highPart, new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil))
// // 			combinedBigInt.Add(combinedBigInt, lowPart)

// // 			fmt.Println("Combined BigInt:", combinedBigInt)
// // 		}

// // 		// Process individual string data elements
// // 		stringVariable6 := emittedEvent.Event.Data[6].String()
// // 		stringVariable7 := emittedEvent.Event.Data[7].String()
// // 		stringVariable8 := emittedEvent.Event.Data[8].String()
// // 		stringVariable9 := emittedEvent.Event.Data[9].String()

// // 		fmt.Println("String Variable 6:", stringVariable6)
// // 		fmt.Println("String Variable 7:", stringVariable7)
// // 		fmt.Println("String Variable 8:", stringVariable8)
// // 		fmt.Println("String Variable 9:", stringVariable9)

// // 		// Process new arrays based on element 10 and 12
// // 		arraySize1 := new(big.Int)
// // 		arraySize1.SetString(emittedEvent.Event.Data[10].String(), 0)

// // 		// Construct arrayValue1 based on arraySize1
// // 		arrayValue1 := []byte{}
// // 		for j := 0; j < int(arraySize1.Int64()); j++ {
// // 			value, _ := hex.DecodeString(emittedEvent.Event.Data[11+j].String())
// // 			arrayValue1 = append(arrayValue1, value...)
// // 		}

// // 		fmt.Println("New Array 1:", arrayValue1)

// // 		// Process new arrays based on element 12
// // 		arraySize2Offset := 10 + int(arraySize1.Int64()) + 1
// // 		arraySize2 := new(big.Int)
// // 		arraySize2.SetString(emittedEvent.Event.Data[arraySize2Offset].String(), 0)

// // 		// Construct arrayValue2 based on arraySize2
// // 		arrayValue2 := []byte{}
// // 		for j := 0; j < int(arraySize2.Int64()); j++ {
// // 			value, _ := hex.DecodeString(emittedEvent.Event.Data[arraySize2Offset+1+j].String())
// // 			arrayValue2 = append(arrayValue2, value...)
// // 		}

// // 		fmt.Println("New Array 2:", arrayValue2)
// // 	}
// // }

// // for _, emittedEvent := range events.Events {
// // 	dataCount := len(emittedEvent.Event.Data)
// // 	// fmt.Printf("Datacoun%s", dataCount)
// // 	if dataCount >= 14 {
// // 		// Process low and high pairs of data elements
// // 		for i := 0; i < 5; i += 2 {
// // 			lowPart := new(big.Int)
// // 			lowPart.SetString(emittedEvent.Event.Data[i].String(), 0)

// // 			highPart := new(big.Int)
// // 			highPart.SetString(emittedEvent.Event.Data[i+1].String(), 0)

// // 			combinedBigInt := new(big.Int).Mul(highPart, new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil))
// // 			combinedBigInt.Add(combinedBigInt, lowPart)

// // 			fmt.Println("Combined BigInt:", combinedBigInt)
// // 		}

// // 		// Process individual string data elements
// // 		stringVariable6 := emittedEvent.Event.Data[6].String()
// // 		stringVariable7 := emittedEvent.Event.Data[7]
// // 		stringVariable8 := emittedEvent.Event.Data[8].String()
// // 		stringVariable9 := emittedEvent.Event.Data[9].String()

// // 		fmt.Println("String Variable 6:", stringVariable6)
// // 		fmt.Println("String Variable 7:", stringVariable7)
// // 		fmt.Println("String Variable 8:", stringVariable8)
// // 		fmt.Println("String Variable 9:", stringVariable9)

// // 		// Process new arrays based on element 10 and 12
// // 		arraySize1 := new(big.Int)
// // 		arraySize1.SetString(emittedEvent.Event.Data[10].String(), 0)

// // 		arrayValue1 := emittedEvent.Event.Data[11].String() + emittedEvent.Event.Data[12].String() +
// // 			emittedEvent.Event.Data[13].String()

// // 		newArray1 := make([]byte, arraySize1.Int64())
// // 		decodedValue1, _ := hex.DecodeString(arrayValue1)
// // 		copy(newArray1, decodedValue1)
// // 		fmt.Println("New Array 1:", newArray1)

// // 		arraySize2 := new(big.Int)
// // 		arraySize2.SetString(emittedEvent.Event.Data[13].String(), 0)

// // 		// arrayValue2 := emittedEvent.Event.Data[15].String() + emittedEvent.Event.Data[16].String() +
// // 		// 	emittedEvent.Event.Data[17].String()

// // 		// newArray2 := make([]byte, arraySize2.Int64())
// // 		// decodedValue2, _ := hex.DecodeString(arrayValue2)
// // 		// copy(newArray2, decodedValue2)
// // 		// fmt.Println("New Array 2:", newArray2)
// // 	}

// // }
