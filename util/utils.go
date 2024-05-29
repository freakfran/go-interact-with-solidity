package util

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gookit/slog"
	"math/big"
	"os"
)

type EthConfig struct {
	Client     *ethclient.Client
	PrivateKey *ecdsa.PrivateKey
}

func GetEthConfigFromEnv() (EthConfig, error) {
	prefix := os.Getenv("ENV")
	slog.Info("load " + prefix + " env")
	url := os.Getenv(GetUrlKey(prefix))
	privateKeyStr := os.Getenv(GetPrivateKeyKey(prefix))
	client, err := ethclient.Dial(url)
	if err != nil {
		return EthConfig{}, fmt.Errorf("error dialing eth client: %v", err)
	}
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return EthConfig{}, fmt.Errorf("error parsing private key: %v", err)
	}
	return EthConfig{
		Client:     client,
		PrivateKey: privateKey,
	}, nil
}

func GetUrlKey(prefix string) string {
	return prefix + "_" + "URL"
}

func GetPrivateKeyKey(prefix string) string {
	return prefix + "_" + "PRIVATE_KEY"
}

func GetDeployAddressKey(prefix string) string {
	return prefix + "_" + "LAST_DEPLOY_ADDRESS"
}

func BuildAuth(privateKey *ecdsa.PrivateKey, client *ethclient.Client) (*bind.TransactOpts, error) {
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
	return auth, nil
}
