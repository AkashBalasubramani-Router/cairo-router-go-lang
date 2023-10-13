package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/account"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/NethermindEth/starknet.go/utils"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
)

var (
	name         string = "testnet"
	account_addr string = "0x06f36e8a0fc06518125bbb1c63553e8a7d8597d437f9d56d891b8c7d3c977716"
	privateKey   string = "0x0687bf84896ee63f52d69e6de1b41492abeadc0dc3cb7bd351d0a52116915937"
	public_key   string = "0x58b0824ee8480133cad03533c8930eda6888b3c5170db2f6e4f51b519141963"
	someContract string = "0x5ff48139a2784ced5ac4a83c6a17913f7c97e407656401503c97e4419a34633" //main contract
	// someContract   string = "0x207d4810c08361e42dda527ae2565f1337d35624c5623df650662360e133bf3"
	contractMethod string = "0x005daf52fdbe5eb8b0f49ba6bfd6b2f34c55c6f539f4c2a3f36e9ea68e0738bd"
	// contractMethod string = "0x004b0981bffc9f82c5b597bc66fac38f60eb3528424c806e4921bfb72720459b"
)

func main() {

	godotenv.Load(fmt.Sprintf(".env.%s", name))
	base := os.Getenv("INTEGRATION_BASE") //please modify the .env.testnet and replace the INTEGRATION_BASE with an starknet goerli RPC.
	fmt.Println("Starting simpleInvoke example")

	//Initialising the connection
	c, err := ethrpc.DialContext(context.Background(), base)
	if err != nil {
		fmt.Println("Failed to connect to the client, did you specify the url in the .env.testnet?")
		panic(err)
	}

	//Initialising the provider
	clientv02 := rpc.NewProvider(c)

	//Here we are converting the account address to felt
	account_address, err := utils.HexToFelt(account_addr)
	if err != nil {
		panic(err.Error())
	}
	//Initializing the account memkeyStore
	ks := account.NewMemKeystore()
	fakePrivKeyBI, ok := new(big.Int).SetString(privateKey, 0)
	if !ok {
		panic(err.Error())
	}
	ks.Put(public_key, fakePrivKeyBI)

	fmt.Println("Established connection with the client")

	maxfee, _ := utils.HexToFelt("0x127156ff7224")

	//Initializing the account
	accnt, err := account.NewAccount(clientv02, account_address, public_key, ks)
	if err != nil {
		panic(err.Error())
	}

	//Getting the nonce from the account, and then converting it into felt
	nonce_string, _ := accnt.Nonce(context.Background(), rpc.BlockID{Tag: "latest"}, accnt.AccountAddress)
	nonce, _ := utils.HexToFelt(*nonce_string)

	//Building the InvokeTx struct
	InvokeTx := rpc.InvokeTxnV1{
		MaxFee:        maxfee,
		Version:       rpc.TransactionV1,
		Nonce:         nonce,
		Type:          rpc.TransactionType_Invoke,
		SenderAddress: account_address,
	}

	contractAddress, _ := utils.HexToFelt(someContract)
	ContractMethod, _ := utils.HexToFelt(contractMethod)

	//Function paramters :

	// fmt.Println(validator_felt[0])

	validators, _ := utils.HexToFelt("0x3D3a12a5B194C30858D6295C68F37D309AfDfa5D")
	size, _ := utils.HexToFelt("0x1")
	powers, _ := utils.HexToFelt("0xffffffff") //4294967295
	valsetNonce_low, _ := utils.HexToFelt("0x01")
	valsetNonce_high, _ := utils.HexToFelt("0x00")

	r1_low, _ := utils.HexToFelt("0x0181523f74528543b9b2aa4fd3621c01")
	r1_high, _ := utils.HexToFelt("0xe82b315f409a002560e6e1b1e51c9ff6")
	// r1_low, _ := utils.HexToFelt("0x0181523f74528543b9b2aa4fd3621c010xe82b315f409a002560e6e1b1e51c9ff6")
	// r1_high, _ := utils.HexToFelt("0x00")
	s1_low, _ := utils.HexToFelt("0x0181523f74528543b9b2aa4fd3621c01")
	s1_high, _ := utils.HexToFelt("0xe82b315f409a002560e6e1b1e51c9ff6")
	v1, _ := utils.HexToFelt("0x1b")

	routeAmount, _ := utils.HexToFelt("0x0")

	requestIdentifier_low, _ := utils.HexToFelt("0x07")
	requestIdentifier_high, _ := utils.HexToFelt("0x00")
	timeStamp_low, _ := utils.HexToFelt("0x687514d2")
	timeStamp_high, _ := utils.HexToFelt("0x00")
	srcChainId, _ := utils.HexToFelt("0x01")
	routeRecipient, _ := utils.HexToFelt("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266")
	// // // routeRecipient, _ := utils.HexToFelt("0x1aad88f6f4ce6ab8827279cfffb922660xf39fd6e5")
	destChainId, _ := utils.HexToFelt("0x01")
	asmAddress, _ := utils.HexToFelt("0x03c12f1e651adbbb65ed4ea4a6524094ad6a8951740a6f6ed229d28462133ee2")
	requestSender, _ := utils.HexToFelt("0x00")
	handlerAddress, _ := utils.HexToFelt("0x015583814f3374302630072712cb894f98079bdfcba84c788363848da54a2610")
	packetsize, _ := utils.HexToFelt("0x02")
	packet1, _ := utils.HexToFelt("0x03038ae29ffd0258880b34b9ffdd37a02bd1b7a7e15ff183c69a0a1c18d30998")
	packet2, _ := utils.HexToFelt("0x64")

	// // // packets := []string{"0x03038ae29ffd0258880b34b9ffdd37a02bd1b7a7e15ff183c69a0a1c18d30998", "0x64"}
	// // // packet,_ := utils.HexArrToFelt(packets)
	isReadCall, _ := utils.HexToFelt("0x00")
	relayerRouterAddress, _ := utils.HexToFelt("456789")

	callDataa := []*felt.Felt{size, validators, size, powers, valsetNonce_low, valsetNonce_high, size, r1_low, r1_high, size, s1_low, s1_high, size, v1, routeAmount, routeAmount,
		requestIdentifier_low, requestIdentifier_high, timeStamp_low, timeStamp_high, srcChainId, routeRecipient, destChainId,
		asmAddress, requestSender, handlerAddress, packetsize, packet1, packet2, isReadCall, relayerRouterAddress}
	fmt.Println(callDataa, "CallData Old")
	fmt.Println("")
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	//Passing it as an array
	validator_array := []string{"0x3D3a12a5B194C30858D6295C68F37D309AfDfa5D"}
	validator_felt, _ := utils.HexArrToFelt(validator_array)

	powers_array := []uint64{4294967295}
	powers_felt := Uint64ArrToFelt(powers_array)

	valsetnonce_single := big.NewInt(1)
	// valsetnonce_felt := B(valsetnonce_array)
	// fmt.Println(valsetnonce_felt, "Valset !!")
	// fmt.Println(valsetnonce_felt, "valsetnonce_felt")

	r1_array := []string{"0xe82b315f409a002560e6e1b1e51c9ff60181523f74528543b9b2aa4fd3621c01"}
	r1_felt := ConvertHexStringsToFelt(r1_array)

	s1_array := []string{"0xe82b315f409a002560e6e1b1e51c9ff60181523f74528543b9b2aa4fd3621c01"}
	s1_felt := ConvertHexStringsToFelt(s1_array)

	v1_string := []string{"27"}
	v1_felt, _ := utils.HexArrToFelt(v1_string)

	routeAmount_single := big.NewInt(0)

	requestIdentifier_single := big.NewInt(7)
	timeStamp_single := big.NewInt(1752503506)

	srcChainId_single, _ := utils.HexToFelt("0x01")

	routeRecipient_single, _ := utils.HexToFelt("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266")

	destChainId_single, _ := utils.HexToFelt("0x01")
	asmAddress_single, _ := utils.HexToFelt("0x03c12f1e651adbbb65ed4ea4a6524094ad6a8951740a6f6ed229d28462133ee2")
	requestSender_single, _ := utils.HexToFelt("0x011293887327823")
	handlerAddress_single, _ := utils.HexToFelt("0x015583814f3374302630072712cb894f98079bdfcba84c788363848da54a2610")
	packet_array := []string{"0x03038ae29ffd0258880b34b9ffdd37a02bd1b7a7e15ff183c69a0a1c18d30998", "0x64"}
	packet_felt, _ := utils.HexArrToFelt(packet_array)
	fmt.Println(packet_felt, "Packet felt")
	isReadCall_single, _ := utils.HexToFelt("0x00")
	relayerRouterAddress_single, _ := utils.HexToFelt("456789")

	var callData1 []*felt.Felt
	//validator
	callData1 = generic_append(validator_felt, callData1)
	// size = utils.Uint64ToFelt((uint64(len(validator_felt))))
	// fmt.Println(size, "Validator size")
	// // Append the size to callData1
	// callData1 = append(callData1, size)
	// for _, element := range validator_felt {
	// 	callData1 = append(callData1, element)
	// }
	// fmt.Println(callData1, "Validator part")

	//powers
	callData1 = generic_append(powers_felt, callData1)

	// size = utils.Uint64ToFelt((uint64(len(powers_felt))))
	// fmt.Println(size, "Powers size")

	// // Append the size to callData1
	// callData1 = append(callData1, size)
	// for _, element := range powers_felt {
	// 	callData1 = append(callData1, element)
	// }

	//valset nonce
	callData1 = BigIntToFeltParts(valsetnonce_single, callData1)
	// fmt.Println(callData1, "Valset")
	// size = utils.Uint64ToFelt((uint64(len(valsetnonce_felt))))
	// valuehex, _ := utils.HexToFelt("0x02")
	// new_size := felt.Zero.Div(size, valuehex)
	// fmt.Println(new_size, "valset size")
	// callData1 = append(callData1, new_size)
	// for _, element := range valsetnonce_felt {
	// 	callData1 = append(callData1, element)
	// }

	//r1
	callData1 = generic_append_2(r1_felt, callData1)

	//s1
	callData1 = generic_append_2(s1_felt, callData1)

	// //v1
	callData1 = generic_append(v1_felt, callData1)

	// //route amount

	callData1 = BigIntToFeltParts(routeAmount_single, callData1)

	// // //request identifier
	callData1 = BigIntToFeltParts(requestIdentifier_single, callData1)

	// // //timestamp
	callData1 = BigIntToFeltParts(timeStamp_single, callData1)

	// // //srcchainId
	callData1 = append(callData1, srcChainId_single)

	// // //routeRecipient
	callData1 = append(callData1, routeRecipient_single)

	// // //destChainId
	callData1 = append(callData1, destChainId_single)

	// // //asmAddress
	callData1 = append(callData1, asmAddress_single)

	// // //requestSender
	callData1 = append(callData1, requestSender_single)

	// // //handlerAddress
	callData1 = append(callData1, handlerAddress_single)

	// // //packet

	callData1 = generic_append(packet_felt, callData1)

	// // //isReadCall
	callData1 = append(callData1, isReadCall_single)

	// // //relayerRouterAddress
	callData1 = append(callData1, relayerRouterAddress_single)
	fmt.Println(callData1, "Calldata New")

	FnCall := rpc.FunctionCall{
		ContractAddress:    contractAddress,
		EntryPointSelector: ContractMethod,
		Calldata:           callData1,
	}

	CairoContractVersion := 2

	InvokeTx.Calldata, err = accnt.FmtCalldata([]rpc.FunctionCall{FnCall}, CairoContractVersion)

	err = accnt.SignInvokeTransaction(context.Background(), &InvokeTx)

	resp, err := accnt.AddInvokeTransaction(context.Background(), InvokeTx)
	fmt.Printf("Response : ", resp)

}

func Uint64ArrToFelt(num []uint64) []*felt.Felt {
	var feltArr []*felt.Felt
	for _, hex := range num {
		felt := utils.Uint64ToFelt(hex)

		feltArr = append(feltArr, felt)
	}
	return feltArr
}

func generic_append(val []*felt.Felt, to_append []*felt.Felt) []*felt.Felt {
	size := utils.Uint64ToFelt(uint64(len(val)))
	to_append = append(to_append, size)
	for _, element := range val {
		to_append = append(to_append, element)
	}
	return to_append
}
func generic_append_2(val []*felt.Felt, to_append []*felt.Felt) []*felt.Felt {
	size := utils.Uint64ToFelt(uint64(len(val) / 2))
	// fmt.Println(size, "Sizee")
	// valuehex, _ := utils.HexToFelt("0x2")
	// new_size := felt.Zero.Div(size, valuehex)
	// fmt.Println(size, "newSizee")
	to_append = append(to_append, size)
	for _, element := range val {
		to_append = append(to_append, element)
	}
	return to_append
}

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
func BigIntArrToHexU128Parts(val []*big.Int) []*felt.Felt {
	var feltArr []*felt.Felt

	for _, hex := range val {
		lowHex, highHex := BigIntToHexU128Parts(hex)
		low_hex, _ := utils.HexToFelt(lowHex)
		high_hex, _ := utils.HexToFelt(highHex)
		feltArr = append(feltArr, low_hex)
		feltArr = append(feltArr, high_hex)
	}

	return feltArr
}

func BigIntToFeltParts(val *big.Int, arrayToAppend []*felt.Felt) []*felt.Felt {
	lowHex, highHex := BigIntToHexU128Parts(val)
	low_hex, _ := utils.HexToFelt(lowHex)
	high_hex, _ := utils.HexToFelt(highHex)
	arrayToAppend = append(arrayToAppend, low_hex)
	arrayToAppend = append(arrayToAppend, high_hex)
	return arrayToAppend
}

func ConvertHexStringsToFelt(hexStrings []string) []*felt.Felt {
	var feltArr []*felt.Felt

	for _, hexString := range hexStrings {
		lowHex, highHex, err := ConvertHexStringToU128Parts(hexString)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		low_hex, _ := utils.HexToFelt(lowHex)
		high_Hex, _ := utils.HexToFelt(highHex)

		feltArr = append(feltArr, low_hex)
		feltArr = append(feltArr, high_Hex)
	}

	return feltArr
}

func ConvertHexStringToU128Parts(hexString string) (lowHex, highHex string, err error) {

	bigIntValue := new(big.Int)

	// Convert the hexadecimal string to a big.Int
	_, success := bigIntValue.SetString(hexString, 0)
	if !success {
		return "", "", fmt.Errorf("Error converting the hex string to a big.Int")
	}

	u128 := new(big.Int)
	u128.Exp(big.NewInt(2), big.NewInt(128), nil)

	low := new(big.Int).And(bigIntValue, new(big.Int).Sub(u128, big.NewInt(1)))
	lowHex = fmt.Sprintf("0x%032x", low)

	high := new(big.Int).Rsh(bigIntValue, 128)
	highHex = fmt.Sprintf("0x%04x", high)

	return lowHex, highHex, nil
}

func BigIntArrToFelt(val []*big.Int) []*felt.Felt {
	var feltArr []*felt.Felt
	for _, hex := range val {
		felt := utils.BigIntToFelt(hex)
		feltArr = append(feltArr, felt)
	}
	return feltArr
}

// func main() {
// 	ExecuteIncreaseValue()
// }
