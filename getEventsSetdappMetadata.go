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
// 	contractAddress, _ := utils.HexToFelt("0x029d74077978f8fd0f680fb17fc439dc38b54b750ad45ca78bb9823e2d5c1c67")
// 	// iAckEvent, _ := utils.HexToFelt("0x4abe8311ae6bce0c81d9a2a4819b3e213ddc510b39641291c8713002303f16")
// 	// valsetUpdateEvent, _ := utils.HexToFelt("0x206930eb6d782c3657f345b7cd4fa1c9f14dd743d26cc26f252875bc7b9d275")
// 	// iSendEvent, _ := utils.HexToFelt("0x20ff74db85f0d9700fcdb580fe4f2dfefd06b4ecde7b8a1753c0f142bda6402")
// 	// iReceiveEvent, _ := utils.HexToFelt("0x3d3a6fa480ecc176ffb470cfbadabc9c568e614a3ff5877bd8c1bcf1ba7c332")
// 	// valsetUpdated, _ := utils.HexToFelt("0x206930eb6d782c3657f345b7cd4fa1c9f14dd743d26cc26f252875bc7b9d275")
// 	setDappMetadataEvent, _ := utils.HexToFelt("0x1dbc9453b52bb94d5026964efc9b92cea7d8d01427ff17a39fd038e44d77520")

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
// 				setDappMetadataEvent,
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
// 	fmt.Println(events)

// 	for _, emittedEvent := range events.Events {
// 		dataCount := len(emittedEvent.Event.Data)
// 		fmt.Println("Datacount", dataCount)
// 		// fmt.Println(emittedEvent.Event.Data)
// 		var eventNonce *big.Int

// 		var dappAddress, chainId, feePayerAddress string

// 		// var execStatus bool
// 		// // , destChainId,
// 		if dataCount >= 1 {
// 			fmt.Println("Event Data:")

// 			lowPart := new(big.Int)
// 			lowPart.SetString(emittedEvent.Event.Data[0].String(), 0)

// 			highPart := new(big.Int)
// 			highPart.SetString(emittedEvent.Event.Data[1].String(), 0)

// 			combinedBigInt := new(big.Int).Mul(highPart, new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil))
// 			combinedBigInt.Add(combinedBigInt, lowPart)

// 			eventNonce = combinedBigInt

// 			fmt.Println("Event Nonce :", eventNonce)

// 			dappAddress = emittedEvent.Event.Data[2].String()
// 			fmt.Println("dappAddress :", dappAddress)

// 			stringVariable3 := emittedEvent.Event.Data[3].String()
// 			stringVariable4Int, err := strconv.Atoi(stringVariable3[2:]) // Removing "0x" prefix
// 			if err != nil {
// 				fmt.Println("Error converting hex for stringVariable7:", err)
// 			}

// 			chainId = strconv.Itoa(stringVariable4Int)
// 			fmt.Println("chainId :", chainId)

// 			feePayerAddress = emittedEvent.Event.Data[4].String()
// 			fmt.Println("feePayerAddress :", feePayerAddress)

// 			transaction := rpc.InvokeTransactionReceipt(rpc.CommonTransactionReceipt{
// 				TransactionHash: emittedEvent.TransactionHash,
// 				Status:          rpc.TransactionAcceptedOnL2,
// 				Events: []rpc.Event{{
// 					FromAddress: contractAddress,
// 				}},
// 			})
// 			// fees := transaction.ActualFee
// 			// fees1, _ := utils.FeltToBigInt(fees)
// 			transactionhash := transaction.TransactionHash
// 			blocknumber100 := emittedEvent.BlockNumber
// 			blocknumbeHash100 := emittedEvent.BlockHash
// 			println("transactionhash --->", transactionhash.String())
// 			println("Blocknumber --->", blocknumber100)

// 			var blockid1 rpc.BlockID = rpc.BlockID{
// 				Number: &blocknumber100,
// 				Hash:   blocknumbeHash100,
// 				Tag:    "",
// 			}

// 			res102, _ := clientv02.BlockWithTxHashes(context.Background(), blockid1)
// 			block := res102.(*rpc.Block)

// 			time := block.Timestamp

// 			fmt.Printf("Timestamp :", time)

// 		}

// 	}
// }

// func main() {
// 	var startblock, endblock uint64
// 	startblock = 851137
// 	endblock =
// 		851144

// 	QueryIsendEvents(startblock, endblock)
// }
