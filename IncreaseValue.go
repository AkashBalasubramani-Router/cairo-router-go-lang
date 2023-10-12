// package main

// import (
// 	"context"
// 	"fmt"
// 	"math/big"
// 	"os"

// 	"github.com/NethermindEth/juno/core/felt"
// 	starknetgo "github.com/NethermindEth/starknet.go"
// 	"github.com/NethermindEth/starknet.go/account"
// 	"github.com/NethermindEth/starknet.go/rpc"
// 	"github.com/NethermindEth/starknet.go/types"
// 	"github.com/NethermindEth/starknet.go/utils"
// 	ethrpc "github.com/ethereum/go-ethereum/rpc"
// 	"github.com/joho/godotenv"
// )

// var (
// 	name           string = "testnet"
// 	account_addr   string = "0x06f36e8a0fc06518125bbb1c63553e8a7d8597d437f9d56d891b8c7d3c977716"
// 	privateKey     string = "0x0687bf84896ee63f52d69e6de1b41492abeadc0dc3cb7bd351d0a52116915937"
// 	public_key     string = "0x58b0824ee8480133cad03533c8930eda6888b3c5170db2f6e4f51b519141963"
// 	someContract   string = "0x5044dfb70b9475663e3ddddb11bbbeccc71614b8db86fc3dc0c16b2b9d3151d"
// 	contractMethod string = "increase_value_8"
// )

// func main() {

// 	godotenv.Load(fmt.Sprintf(".env.%s", name))
// 	base := os.Getenv("INTEGRATION_BASE")
// 	fmt.Println("Starting simpeCall example")
// 	c, err := ethrpc.DialContext(context.Background(), base)
// 	if err != nil {
// 		fmt.Println("Failed to connect to the client, did you specify the url in the .env.mainnet?")
// 		panic(err)
// 	}
// 	clientv02 := rpc.NewProvider(c)
// 	account_address, _ := utils.HexToFelt(account_addr)
// 	ks := starknetgo.NewMemKeystore()
// 	fakePrivKeyBI, _ := new(big.Int).SetString(privateKey, 0)
// 	ks.Put(public_key, fakePrivKeyBI)

// 	fmt.Println("Established connection with the client")

// 	maxfee, _ := utils.HexToFelt("0x9184e72a000")
// 	accnt, err := account.NewAccount(clientv02, 1, account_address, public_key, ks)

// 	InvokeTx := rpc.BroadcastedInvokeV1Transaction{
// 		BroadcastedTxnCommonProperties: rpc.BroadcastedTxnCommonProperties{
// 			Nonce:   new(felt.Felt).SetUint64(110),
// 			MaxFee:  maxfee,
// 			Version: rpc.TransactionV1,
// 			Type:    rpc.TransactionType_Invoke,
// 		},
// 		SenderAddress: account_address,
// 	}

// 	contractAddress, _ := utils.HexToFelt(someContract)

// 	// utils.BigIntToFelt()
// 	// size_ethaddress, _ := utils.HexToFelt("0x1")
// 	// ethAddress_low, _ := utils.HexToFelt("0x82cfc121ff244cb3c3733af712f9df02")
// 	// ethAddress_high, _ := utils.HexToFelt("0x5f046934")
// 	ethaddress, _ := utils.HexToFelt("0xaC25b2B9F4ca06179fA0D2522F47Bc86A9DF9316")

// 	FnCall := rpc.FunctionCall{
// 		ContractAddress:    contractAddress,
// 		EntryPointSelector: types.GetSelectorFromNameFelt("increase_value_8"),
// 		Calldata:           []*felt.Felt{ethaddress},
// 	}

// 	accnt.BuildInvokeTx(context.Background(), &InvokeTx, &[]rpc.FunctionCall{FnCall})
// 	txHash, err := accnt.TransactionHashInvoke(InvokeTx.Calldata, InvokeTx.Nonce, InvokeTx.MaxFee, accnt.AccountAddress)
// 	fmt.Printf("TxHash :", txHash)
// 	resp, err := accnt.AddInvokeTransaction(context.Background(), &InvokeTx)
// 	fmt.Printf("Response : ", resp)
// 	// accnt.BuildInvokeTx()

// }

// // func main() {
// // 	ExecuteIncreaseValue()
// // }
