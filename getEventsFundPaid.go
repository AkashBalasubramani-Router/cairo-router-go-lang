// package main

// import (
// 	"context"
// 	"fmt"
// 	"math/big"

// 	"github.com/NethermindEth/juno/core/felt"
// 	"github.com/NethermindEth/starknet.go/rpc"
// 	"github.com/NethermindEth/starknet.go/types"
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
// 	// FundsDeposited, _ := utils.HexToFelt("0x338f51ed1b54a9f739b53c6835296497c012ceac5a613b7219505273b71f507")
// 	FundsPaid, _ := utils.HexToFelt("0xb018b34bf524c9807d3cddb007f4d300a18f111bd2c3cf3234e9de181a4cef")
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
// 				FundsPaid,
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

// 	// contractAddress, _ := utils.HexToFelt("0x04eef2db94e7f0e2a71ce38630b6ecef4e3c00d1646c888f7c669ad2965baf7f")

// 	tx := rpc.FunctionCall{
// 		ContractAddress:    contractAddress,
// 		EntryPointSelector: types.GetSelectorFromNameFelt("getChainId"),
// 	}

// 	fmt.Println("Making Call() request")
// 	callResp, err := clientv02.Call(context.Background(), tx, rpc.BlockID{Tag: "latest"})
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	lowPart := new(big.Int)
// 	lowPart.SetString(callResp[0], 0)

// 	highPart := new(big.Int)
// 	highPart.SetString(callResp[1], 0)

// 	combinedBigInt := new(big.Int).Mul(highPart, new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil))
// 	combinedBigInt.Add(combinedBigInt, lowPart)

// 	fmt.Println("Resulttt ", combinedBigInt.String())

// 	for _, emittedEvent := range events.Events {
// 		dataCount := len(emittedEvent.Event.Data)
// 		fmt.Println("Data count :", dataCount)
// 		var nonce *big.Int
// 		// var partnerId, amount *big.Int

// 		var forwarder, messageHash, forwarderRouterAddress string
// 		// var destChainIdBytes string
// 		// , destChainId,
// 		if dataCount >= 6 {
// 			fmt.Println("Event Data:")

// 			lowPart := new(big.Int)
// 			lowPart.SetString(emittedEvent.Event.Data[0].String(), 0)

// 			highPart := new(big.Int)
// 			highPart.SetString(emittedEvent.Event.Data[1].String(), 0)

// 			combinedBigInt := new(big.Int).Mul(highPart, new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil))
// 			combinedBigInt.Add(combinedBigInt, lowPart)

// 			// messageHash = combinedBigInt.String()
// 			res2, _ := utils.BigIntToFelt(combinedBigInt)
// 			messageHash = res2.String()
// 			forwarder = emittedEvent.Event.Data[2].String()

// 			lowPart1 := new(big.Int)
// 			lowPart1.SetString(emittedEvent.Event.Data[3].String(), 0)

// 			highPart1 := new(big.Int)
// 			highPart1.SetString(emittedEvent.Event.Data[4].String(), 0)

// 			combinedBigInt1 := new(big.Int).Mul(highPart1, new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil))
// 			combinedBigInt1.Add(combinedBigInt1, lowPart1)

// 			nonce = combinedBigInt1

// 			forwarderRouterAddress = emittedEvent.Event.Data[5].String()

// 			// blocknumber := emittedEvent.BlockNumber

// 			// transactionkHash := emittedEvent.TransactionHash

// 			// var example rpc.BlockID
// 			// example = rpc.BlockID{
// 			// 	Number: &blocknumber,
// 			// 	Hash:   transactionkHash,
// 			// 	Tag:    "latest",
// 			// }
// 			// res100, _ := clientv02.
// 			// fmt.Println("Result:", res100)

// 			fmt.Println("messageHash:", messageHash)

// 			fmt.Println("forwarder:", forwarder)
// 			fmt.Println("nonce:", nonce)
// 			fmt.Println("forwarderRouterAddress:", forwarderRouterAddress)
// 			// fmt.Println("depositId:", depositId)
// 			// fmt.Println("srcToken:", srcToken)
// 			// fmt.Println("recipient:", recipient)
// 			// fmt.Println("depositor:", depositor)

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
