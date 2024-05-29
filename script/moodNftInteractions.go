package script

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gookit/slog"
	"go-made/interactions"
	"go-made/util"
	"log"
	"math/big"
)

// DeployMoodNft
//
//	@Description: 部署MoodNft合约
//	@param privateKey
//	@param client
//	@return *interactions.MoodNft
//	@return error
func DeployMoodNft(privateKey *ecdsa.PrivateKey, client *ethclient.Client, sadSvgImgUri string, happySvgImgUri string) (*interactions.MoodNft, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	slog.Info("fromAddress:", fromAddress.Hex())
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

	if err != nil {
		return nil, fmt.Errorf("error getting nonce:%v", err)
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

	address, transaction, moodNft, err := interactions.DeployMoodNft(auth, client, sadSvgImgUri, happySvgImgUri)
	if err != nil {
		return nil, err
	}
	log.Printf("Contract address: %s\n", address.Hex())
	log.Printf("Transaction hash: %s\n", transaction.Hash().Hex())
	return moodNft, nil
}

// GetMoodNft
//
//	@Description: 通过地址取得MoodNft
//	@param address
//	@param client
//	@return *interactions.MoodNft
//	@return error
func GetMoodNft(address string, client *ethclient.Client) (*interactions.MoodNft, error) {
	moodNft, err := interactions.NewMoodNft(common.HexToAddress(address), client)
	if err != nil {
		return nil, err
	}
	return moodNft, nil
}

// MintMoodNft
//
//	@Description: 铸币
//	@param privateKey
//	@param client
//	@param basicNft
//	@param tokenURI
func MintMoodNft(privateKey *ecdsa.PrivateKey, client *ethclient.Client, moodNft *interactions.MoodNft) (*types.Transaction, error) {
	auth, err := util.BuildAuth(privateKey, client)
	if err != nil {
		return nil, err
	}
	return moodNft.MintNft(auth)
}

func FlippMood(privateKey *ecdsa.PrivateKey, client *ethclient.Client, moodNft *interactions.MoodNft, tokenId *big.Int) (*types.Transaction, error) {
	auth, err := util.BuildAuth(privateKey, client)
	if err != nil {
		return nil, err
	}
	return moodNft.FlipMood(auth, tokenId)
}

func SvgToUri(svg string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(svg))
	baseURL := "data:image/svg+xml;base64,"
	return baseURL + encoded, nil
}
