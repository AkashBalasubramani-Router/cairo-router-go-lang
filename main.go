// package main

// import (
// 	"context"
// 	"encoding/hex"
// 	"fmt"
// 	"math/big"
// 	"regexp"

// 	"github.com/NethermindEth/juno/core/felt"
// 	starknetgo "github.com/NethermindEth/starknet.go"
// 	gateway "github.com/NethermindEth/starknet.go/gateway"
// 	rpc "github.com/NethermindEth/starknet.go/rpc"
// 	"github.com/NethermindEth/starknet.go/types"
// 	"github.com/NethermindEth/starknet.go/utils"
// 	"github.com/ethereum/go-ethereum/common"
// )

// type StarknetMetaData struct {
// 	TransactionHash *felt.Felt
// 	BlockNumber     uint64
// 	Timestamp       uint64
// }

// type TestSetType struct {
// 	eventFilter        rpc.EventFilter
// 	ExpectedEventCount int
// }

// type StarknetGatewayUpgradeableISendEvent struct {
// 	RouteAmount     *big.Int       `json:"route_amount"`
// 	RouteRecipient  *felt.Felt     `json:"route_recipient"`
// 	EventNonce      *big.Int       `json:"event_nonce"`
// 	SrcChainID      *felt.Felt     `json:"src_chain_id"`
// 	DestChainID     *felt.Felt     `json:"dest_chain_id"`
// 	RequestSender   common.Address `json:"request_sender"`
// 	Version         *big.Int       `json:"version"`
// 	RequestMetadata []*felt.Felt   `json:"request_metadata"`
// 	RequestPacket   []*felt.Felt   `json:"request_packet"`
// 	Raw             StarknetMetaData
// }

// // type StarknetGatewayUpgradeableISendEvent struct {
// // 	RouteAmount     *big.Int       `json:"route_amount"`
// // 	RouteRecipient  *felt.Felt     `json:"route_recipient"`
// // 	EventNonce      *big.Int       `json:"event_nonce"`
// // 	SrcChainID      *felt.Felt     `json:"src_chain_id"`
// // 	DestChainID     *felt.Felt     `json:"dest_chain_id"`
// // 	RequestSender   common.Address `json:"request_sender"`
// // 	Version         *big.Int       `json:"version"`
// // 	RequestMetadata []*felt.Felt   `json:"request_metadata"`
// // 	RequestPacket   []*felt.Felt   `json:"request_packet"`
// // 	Raw             StarknetMetaData
// // }

// // type TestSetType struct {
// // 	eventFilter        rpc.EventFilter
// // 	ExpectedEventCount int
// // }

// // type StarknetMetaData struct {
// // 	TransactionHash *felt.Felt
// // 	BlockNumber     uint64
// // 	Timestamp       uint64
// // }

// // type StarknetEventProcessor struct {
// // 	chainSpec         config.ChainSpecs
// // 	starknetClient    *rpc.Provider
// // 	logger            *log.Entry
// // 	ethPrivateKey     string
// // 	routerChainClient routerclient.ChainClient
// // 	messageProducer   types.IMqSender
// // 	dbHandler         *store.DbHandler
// // }

// // func (eventProcessor *StarknetEventProcessor) QueryGatewayiSendEvents(startBlock *uint64, endBlock *uint64) ([]*StarknetGatewayUpgradeableISendEvent, error) {

// // 	// var addr *felt.Felt = felt.NewFelt("0x00322c67c6ef761f88ba49dad0ad5388336caec7cc55b5f2e48578c11e27d102")
// // 	myFelt, err := utils.HexToFelt("0x99cd8bde557814842a3121e8ddfd433a539b8c9f14bf31ebf108d12e6196e9")
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	myKeys := [][]*felt.Felt{[]*felt.Felt{myFelt}}

// // 	caddress, _ := utils.HexToFelt("0x00322c67c6ef761f88ba49dad0ad5388336caec7cc55b5f2e48578c11e27d102")

// // 	testSet := []TestSetType{
// // 		{
// // 			eventFilter: rpc.EventFilter{
// // 				FromBlock: rpc.BlockID{
// // 					Number: startBlock,
// // 				},
// // 				ToBlock: rpc.BlockID{
// // 					Number: endBlock,
// // 				},
// // 				Address: caddress,
// // 				Keys:    myKeys,
// // 			},
// // 			ExpectedEventCount: 100,
// // 		},
// // 	}

// // 	var iSendEvents []*StarknetGatewayUpgradeableISendEvent
// // 	for _, test := range testSet {

// // 		eventInput := rpc.EventsInput{
// // 			EventFilter: test.eventFilter,
// // 		}

// // 		events, err := eventProcessor.starknetClient.Events(context.Background(), eventInput)
// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 		for _, event := range events.Events {
// // 			metaData := StarknetMetaData{
// // 				TransactionHash: event.TransactionHash,
// // 				BlockNumber:     event.BlockNumber,
// // 				Timestamp:       0, // Fill in the appropriate timestamp value if available
// // 			}
// // 			isendEvent := createStarknetGatewayUpgradeableISendEvent(event, metaData)
// // 			iSendEvents = append(iSendEvents, isendEvent)
// // 		}
// // 	}
// // 	return iSendEvents, nil
// // }

// // func createStarknetGatewayUpgradeableISendEvent(event rpc.EmittedEvent, metaData StarknetMetaData) *StarknetGatewayUpgradeableISendEvent {
// // 	isendEvent := &StarknetGatewayUpgradeableISendEvent{
// // 		RouteAmount:     feltToBig(event.Data[1]),
// // 		RouteRecipient:  event.Data[6],
// // 		EventNonce:      feltToBig(event.Data[2]),
// // 		SrcChainID:      event.Data[4],
// // 		DestChainID:     event.Data[5],
// // 		RequestSender:   common.HexToAddress(event.Data[3].String()),
// // 		Version:         feltToBig(event.Data[0]),
// // 		RequestMetadata: event.Data[7:],
// // 		RequestPacket:   event.Data[8:],
// // 		Raw:             metaData,
// // 	}

// // 	return isendEvent
// // }

// func feltToBig(feltNum *felt.Felt) *big.Int {
// 	// Use SetString() to convert the felt.Felt value to *big.Int
// 	bigIntValue, _ := new(big.Int).SetString(feltNum.String(), 0)
// 	return bigIntValue
// }

// func ExecuteTransaction(name string, counterContract string, fnselector string, fnselector1 string) {

// 	// rpc_1 := "https://starknet-goerli.g.alchemy.com/v2/bbDTAS9Qs68Po3ssBBPVm378Pye6oS2A"
// 	// c, err := ethrpc.DialContext(context.Background(), rpc_1)
// 	// if err != nil {
// 	// 	fmt.Println("Failed to connect to the client, did you specify the url in the .env.mainnet?")
// 	// 	panic(err)
// 	// }
// 	// connection := rpc.NewProvider(c)
// 	// fmt.Println("Established connection with the client")

// 	// myFelt, err := utils.HexToFelt("0x99cd8bde557814842a3121e8ddfd433a539b8c9f14bf31ebf108d12e6196e9")
// 	// caddress, err := utils.HexToFelt("0x00322c67c6ef761f88ba49dad0ad5388336caec7cc55b5f2e48578c11e27d102")
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// myKeys := [][]*felt.Felt{[]*felt.Felt{myFelt}}

// 	// testSet := []TestSetType{
// 	// 	{
// 	// 		eventFilter: rpc.EventFilter{
// 	// 			FromBlock: rpc.BlockID{
// 	// 				Number: startBlock,
// 	// 			},
// 	// 			ToBlock: rpc.BlockID{
// 	// 				Number: endBlock,
// 	// 			},
// 	// 			Address: caddress,
// 	// 			Keys:    myKeys,
// 	// 		},
// 	// 		ExpectedEventCount: 100,
// 	// 	},
// 	// }

// 	// var iSendEvents []*StarknetGatewayUpgradeableISendEvent
// 	// for _, test := range testSet {

// 	// 	eventInput := rpc.EventsInput{
// 	// 		EventFilter: test.eventFilter,
// 	// 	}

// 	// 	events, err := connection.Events(context.Background(), eventInput)
// 	// 	if err != nil {
// 	// 		return nil, err
// 	// 	}
// 	// 	for _, event := range events.Events {
// 	// 		metaData := StarknetMetaData{
// 	// 			TransactionHash: event.TransactionHash,
// 	// 			BlockNumber:     event.BlockNumber,
// 	// 			Timestamp:       0, // Fill in the appropriate timestamp value if available
// 	// 		}
// 	// 		isendEvent := createStarknetGatewayUpgradeableISendEvent(event, metaData)
// 	// 		iSendEvents = append(iSendEvents, isendEvent)
// 	// 	}
// 	// }
// 	// return iSendEvents, nil

// 	gw := gateway.NewProvider(gateway.WithChain(name))
// 	caddress, _ := utils.HexToFelt(counterContract)
// 	selector, _ := utils.HexToFelt(fnselector)

// 	// get count before tx
// 	callResp, err := gw.Call(context.Background(), types.FunctionCall{
// 		ContractAddress:    caddress,
// 		EntryPointSelector: selector,
// 	}, "")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("Counter is currently at: ", callResp[0])

// 	// init account handler
// 	ks := starknetgo.NewMemKeystore()
// 	fakeSenderAddress := "0x03dd748c5bd6c16bbaf6fa68a5a6c2c62012caa2055d3447f7c6408f0d9a3155"
// 	ks.Put(fakeSenderAddress, types.SNValToBN(fakeSenderAddress))
// 	fakeSenderAddress1, _ := utils.HexToFelt("0x03dd748c5bd6c16bbaf6fa68a5a6c2c62012caa2055d3447f7c6408f0d9a3155")
// 	address, _ := utils.HexToFelt("0x07f80089FB3F69B796b0445A67F47A680a71F241a009A5dA4546353dB90AA125")
// 	account, err := starknetgo.NewGatewayAccount(fakeSenderAddress1, address, ks, gw)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	fnselector2, _ := utils.HexToFelt(fnselector1)

// 	increment := []types.FunctionCall{
// 		{
// 			ContractAddress:    caddress,
// 			EntryPointSelector: fnselector2,
// 		},
// 	}

// 	// estimate fee for executing transaction
// 	feeEstimate, err := account.EstimateFee(context.Background(), increment, types.ExecuteDetails{})
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	// fee, _ := big.NewInt(0).SetString(string(feeEstimate.OverallFee), 0)
// 	// expandedFee := big.NewInt(0).Mul(fee, big.NewInt(int64(feeMargin)))
// 	// max := big.NewInt(0).Div(expandedFee, big.NewInt(100))
// 	fmt.Printf("Fee:\n\tEstimate\t\t%v wei\n\tEstimate+Margin\t\t%v wei\n\n", feeEstimate.OverallFee)

// 	// execute transaction
// 	execResp, err := account.Execute(context.Background(), increment, types.ExecuteDetails{MaxFee: big.NewInt(1000000000000)})
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	n, receipt, err := gw.WaitForTransaction(context.Background(), execResp.TransactionHash.String(), 5, 100)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Printf("transaction(%s) receipt: %s\n\n", n, execResp.TransactionHash, receipt.Status)

// 	// AccountAddress := "0x07f80089FB3F69B796b0445A67F47A680a71F241a009A5dA4546353dB90AA125"
// 	// pk := "0x03dd748c5bd6c16bbaf6fa68a5a6c2c62012caa2055d3447f7c6408f0d9a3155"
// 	// fakeSenderAddress := pk
// 	// ks := starknetgo.NewMemKeystore()
// 	// k := types.SNValToBN("0x03dd748c5bd6c16bbaf6fa68a5a6c2c62012caa2055d3447f7c6408f0d9a3155")
// 	// ks.Put(fakeSenderAddress, k)

// 	// fakeSenderAddress_felt, _ := utils.HexToFelt(fakeSenderAddress)
// 	// AccountAddress_felt, _ := utils.HexToFelt(AccountAddress)

// 	// account, err := starknetgo.NewRPCAccount(fakeSenderAddress_felt, AccountAddress_felt, ks, connection)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// var paramters []*felt.Felt
// 	// // args, _ := utils.BigIntToFelt(number)
// 	// // appendToFeltSlice(paramters, args)
// 	// caddress, _ := utils.HexToFelt("0x057b40a7ad58cd274192609d0dfe6ef70a57d9c03ed27ffc1d06b1e7b4078866")
// 	// tx := []types.FunctionCall{
// 	// 	{
// 	// 		ContractAddress:    caddress,
// 	// 		EntryPointSelector: types.GetSelectorFromNameFelt("increase_Value"),
// 	// 		Calldata:           paramters,
// 	// 	},
// 	// }

// 	// Transaction, err := account.Execute(context.Background(), tx, types.ExecuteDetails{MaxFee: big.NewInt(1000000000)})
// 	// if err != nil {
// 	// 	panic(err.Error())
// 	// }

// 	// tx1 := rpc.FunctionCall{
// 	// 	ContractAddress:    caddress,
// 	// 	EntryPointSelector: types.GetSelectorFromNameFelt("display_value"),
// 	// 	Calldata:           paramters,
// 	// }

// 	// Transaction, err := account.Call(context.Background(), tx1, rpc.BlockID{Tag: "latest"})
// 	// if err != nil {
// 	// 	panic(err.Error())
// 	// }

// 	// return Transaction, err
// }

// func createStarknetGatewayUpgradeableISendEvent(event rpc.EmittedEvent, metaData StarknetMetaData) *StarknetGatewayUpgradeableISendEvent {
// 	isendEvent := &StarknetGatewayUpgradeableISendEvent{
// 		RouteAmount:     feltToBig(event.Data[1]),
// 		RouteRecipient:  event.Data[6],
// 		EventNonce:      feltToBig(event.Data[2]),
// 		SrcChainID:      event.Data[4],
// 		DestChainID:     event.Data[5],
// 		RequestSender:   common.HexToAddress(event.Data[3].String()),
// 		Version:         feltToBig(event.Data[0]),
// 		RequestMetadata: event.Data[7:],
// 		RequestPacket:   event.Data[8:],
// 		Raw:             metaData,
// 	}

// 	return isendEvent
// }

// func appendToFeltSlice(slice []*felt.Felt, element *felt.Felt) []*felt.Felt {
// 	return append(slice, element)
// }

// func strToFelt(str string) *felt.Felt {
// 	var f *felt.Felt
// 	asciiRegexp := regexp.MustCompile(`^([[:graph:]]|[[:space:]]){1,31}$`)

// 	if b, ok := new(big.Int).SetString(str, 0); ok {
// 		fBytes := f.Bytes()
// 		b.FillBytes(fBytes[:])
// 		return f
// 	}
// 	// TODO: revisit conversation on seperate 'ShortString' conversion
// 	if asciiRegexp.MatchString(str) {
// 		hexStr := hex.EncodeToString([]byte(str))
// 		if b, ok := new(big.Int).SetString(hexStr, 16); ok {
// 			fBytes := f.Bytes()
// 			b.FillBytes(fBytes[:])
// 			return f
// 		}
// 	}

// 	return f
// }

// func main() {

// 	// startBlock := uint64(844022) // Replace with your actual start block
// 	// endingBlock := uint64(844037)

// 	// // var endingBlock uint64 = 844037

// 	// iSendEvents, err := ExecuteTransaction(&startBlock, &endingBlock)
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// fmt.Println("Received iSendEvents:", iSendEvents)
// 	// ExecuteTransaction(, "0x057b40a7ad58cd274192609d0dfe6ef70a57d9c03ed27ffc1d06b1e7b4078866", "0x646973706c61795f76616c7565", "0x696e6372656173655f56616c7565")
// 	gw := gateway.NewProvider(gateway.WithChain("test"))
// 	caddress, _ := utils.HexToFelt("0x057b40a7ad58cd274192609d0dfe6ef70a57d9c03ed27ffc1d06b1e7b4078866")
// 	selector, _ := utils.HexToFelt("0x646973706c61795f76616c7565")

// 	// get count before tx
// 	callResp, err := gw.Call(context.Background(), types.FunctionCall{
// 		ContractAddress:    caddress,
// 		EntryPointSelector: selector,
// 	}, "")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("Counter is currently at: ", callResp[0])

// 	// init account handler
// 	ks := starknetgo.NewMemKeystore()
// 	fakeSenderAddress := "0x03dd748c5bd6c16bbaf6fa68a5a6c2c62012caa2055d3447f7c6408f0d9a3155"
// 	ks.Put(fakeSenderAddress, types.SNValToBN(fakeSenderAddress))
// 	fakeSenderAddress1, _ := utils.HexToFelt("0x03dd748c5bd6c16bbaf6fa68a5a6c2c62012caa2055d3447f7c6408f0d9a3155")
// 	address, _ := utils.HexToFelt("0x07f80089FB3F69B796b0445A67F47A680a71F241a009A5dA4546353dB90AA125")
// 	account, err := starknetgo.NewGatewayAccount(fakeSenderAddress1, address, ks, gw)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	fnselector2, _ := utils.HexToFelt("0x696e6372656173655f56616c7565")

// 	increment := []types.FunctionCall{
// 		{
// 			ContractAddress:    caddress,
// 			EntryPointSelector: fnselector2,
// 		},
// 	}

// 	// estimate fee for executing transaction
// 	feeEstimate, err := account.EstimateFee(context.Background(), increment, types.ExecuteDetails{})
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	// fee, _ := big.NewInt(0).SetString(string(feeEstimate.OverallFee), 0)
// 	// expandedFee := big.NewInt(0).Mul(fee, big.NewInt(int64(feeMargin)))
// 	// max := big.NewInt(0).Div(expandedFee, big.NewInt(100))
// 	fmt.Printf("Fee:\n\tEstimate\t\t%v wei\n\tEstimate+Margin\t\t%v wei\n\n", feeEstimate.OverallFee)

// 	// execute transaction
// 	execResp, err := account.Execute(context.Background(), increment, types.ExecuteDetails{MaxFee: big.NewInt(1000000000000)})
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	n, receipt, err := gw.WaitForTransaction(context.Background(), execResp.TransactionHash.String(), 5, 100)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Printf("transaction(%s) receipt: %s\n\n", n, execResp.TransactionHash, receipt.Status)

// }
