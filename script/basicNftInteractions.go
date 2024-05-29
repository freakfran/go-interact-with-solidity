package script

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gookit/slog"
	"go-made/interactions"
	"log"
	"math/big"
)

// DeployBasicNft
//
//	@Description: 部署BasicNft合约
//	@param privateKey
//	@param client
//	@return *interactions.BasicNft
//	@return error
func DeployBasicNft(privateKey *ecdsa.PrivateKey, client *ethclient.Client) (*interactions.BasicNft, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

	if err != nil {
		return nil, err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = 0          // in units,0 = estimate
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	gasTipCap, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, err
	}
	//auth.GasPrice = gasPrice // only for legacy transactions
	auth.GasFeeCap = gasPrice
	auth.GasTipCap = gasTipCap

	address, transaction, basicNft, err := interactions.DeployBasicNft(auth, client)
	if err != nil {
		return nil, err
	}
	log.Printf("Contract address: %s\n", address.Hex())
	log.Printf("Transaction hash: %s\n", transaction.Hash().Hex())
	return basicNft, nil
}

// GetBasicNft
//
//	@Description: 通过地址取得BasicNft
//	@param address
//	@param client
//	@return *interactions.BasicNft
//	@return error
func GetBasicNft(address string, client *ethclient.Client) (*interactions.BasicNft, error) {
	basicNft, err := interactions.NewBasicNft(common.HexToAddress(address), client)
	if err != nil {
		return nil, err
	}
	return basicNft, nil
}

// MintBasicNft
//
//	@Description: 铸币
//	@param privateKey
//	@param client
//	@param basicNft
//	@param tokenURI
func MintBasicNft(privateKey *ecdsa.PrivateKey, client *ethclient.Client, basicNft *interactions.BasicNft, tokenURI string) (*types.Transaction, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

	if err != nil {
		return nil, err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = 0          // in units,0 = estimate
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	gasTipCap, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, err
	}
	//auth.GasPrice = gasPrice // only for legacy transactions
	gasFeeCap := new(big.Int).Add(gasPrice, gasTipCap)
	auth.GasFeeCap = gasFeeCap
	slog.Info("mint with gasFeeCap:", gasFeeCap)
	auth.GasTipCap = gasTipCap
	slog.Info("mint with gasTipCap:", gasTipCap)
	return basicNft.MintNft(auth, tokenURI)
}
