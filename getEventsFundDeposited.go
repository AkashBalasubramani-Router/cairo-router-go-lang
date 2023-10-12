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
// 	contractAddress, _ := utils.HexToFelt("0x04eef2db94e7f0e2a71ce38630b6ecef4e3c00d1646c888f7c669ad2965baf7f")
// 	// iSendEvent, _ := utils.HexToFelt("0x20ff74db85f0d9700fcdb580fe4f2dfefd06b4ecde7b8a1753c0f142bda6402")
// 	// iReceiveEvent, _ := utils.HexToFelt("0x3d3a6fa480ecc176ffb470cfbadabc9c568e614a3ff5877bd8c1bcf1ba7c332")
// 	// valsetUpdated, _ := utils.HexToFelt("0x206930eb6d782c3657f345b7cd4fa1c9f14dd743d26cc26f252875bc7b9d275")
// 	FundsDeposited, _ := utils.HexToFelt("0x338f51ed1b54a9f739b53c6835296497c012ceac5a613b7219505273b71f507")

// 	eventInput := rpc.EventsInput{
// 		EventFilter: rpc.EventFilter{
// 			FromBlock: rpc.BlockID{
// 				Number: &startblock,
// 			},
// 			ToBlock: rpc.BlockID{
// 				Number: &endblock,
// 			},
// 			Address: contractAddress,
// 			Keys: [][]*felt.Felt{{
// 				FundsDeposited,
// 			}},
// 		},
// 		ResultPageRequest: rpc.ResultPageRequest{
// 			ChunkSize: 1,
// 		},
// 	}

// 	events, err := clientv02.Events(context.Background(), eventInput)
// 	if err != nil {
// 		fmt.Println("Unsuccessful")
// 		panic(err)
// 	}

// 	for _, emittedEvent := range events.Events {
// 		dataCount := len(emittedEvent.Event.Data)
// 		fmt.Println("Data count :", dataCount)
// 		var partnerId, amount, destAmount, depositId *big.Int
// 		// var partnerId, amount *big.Int
// 		var recipient []string
// 		var destChainIdBytes, srcToken, depositor string
// 		// var destChainIdBytes string
// 		// , destChainId,
// 		if dataCount >= 14 {
// 			fmt.Println("Event Data:")
// 			for i := 0; i < 3; i += 2 {
// 				lowPart := new(big.Int)
// 				lowPart.SetString(emittedEvent.Event.Data[i].String(), 0)

// 				highPart := new(big.Int)
// 				highPart.SetString(emittedEvent.Event.Data[i+1].String(), 0)

// 				combinedBigInt := new(big.Int).Mul(highPart, new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil))
// 				combinedBigInt.Add(combinedBigInt, lowPart)

// 				switch i {
// 				case 0:
// 					partnerId = combinedBigInt
// 				case 2:
// 					amount = combinedBigInt
// 				}

// 			}

// 			stringVariable4 := emittedEvent.Event.Data[4].String()

// 			stringVariable4Int, err := strconv.Atoi(stringVariable4[2:]) // Removing "0x" prefix
// 			if err != nil {
// 				fmt.Println("Error converting hex for stringVariable7:", err)
// 			}

// 			destChainIdBytes = strconv.Itoa(stringVariable4Int)

// 			for i := 5; i < 8; i += 2 {
// 				lowPart := new(big.Int)
// 				lowPart.SetString(emittedEvent.Event.Data[i].String(), 0)

// 				highPart := new(big.Int)
// 				highPart.SetString(emittedEvent.Event.Data[i+1].String(), 0)

// 				combinedBigInt := new(big.Int).Mul(highPart, new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil))
// 				combinedBigInt.Add(combinedBigInt, lowPart)

// 				switch i {
// 				case 5:
// 					destAmount = combinedBigInt
// 				case 7:
// 					depositId = combinedBigInt
// 				}

// 			}

// 			srcToken = emittedEvent.Event.Data[9].String()

// 			arraySize1 := new(big.Int)
// 			arraySize1.SetString(emittedEvent.Event.Data[10].String(), 0)

// 			// Construct arrayValue1 based on arraySize1
// 			arrayValue1 := []string{}
// 			for j := 0; j < int(arraySize1.Int64()); j++ {
// 				value := emittedEvent.Event.Data[11+j].String()
// 				arrayValue1 = append(arrayValue1, value)
// 			}

// 			recipient = arrayValue1

// 			depositor = emittedEvent.Event.Data[10+int(arraySize1.Int64())+1].String()

// 			fmt.Println("partnerId:", partnerId)

// 			fmt.Println("amount:", amount)
// 			fmt.Println("destChainIdBytes:", destChainIdBytes)
// 			fmt.Println("destAmount:", destAmount)
// 			fmt.Println("depositId:", depositId)
// 			fmt.Println("srcToken:", srcToken)
// 			fmt.Println("recipient:", recipient)
// 			fmt.Println("depositor:", depositor)

// 		}
// 	}
// }

// func main() {
// 	var startblock, endblock uint64
// 	startblock =
// 		848220
// 	endblock = 850313

// 	QueryIsendEvents(startblock, endblock)
// }
