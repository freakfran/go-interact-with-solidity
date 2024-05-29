package demo

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gookit/slog"
	"github.com/joho/godotenv"
	"go-made/interactions"
	"go-made/script"
	"math/big"
	"os"
)

func example() {
	err := godotenv.Load()
	if err != nil {
		slog.Fatal("Error loading .env file")
		return
	}
	var basicNft *interactions.BasicNft
	client, err := ethclient.Dial(os.Getenv("SEPOLIA_URL"))
	privateKey, err := crypto.HexToECDSA(os.Getenv("SEPOLIA_PRIVATE_KEY"))
	lastDeployAddress := os.Getenv("LAST_DEPLOY_ADDRESS")
	if lastDeployAddress == "" {
		var err error
		basicNft, err = script.DeployBasicNft(privateKey, client)
		if err != nil {
			slog.Fatal("Deploy BasicNft failed:", err)
			return
		}
	} else {
		basicNft, err = script.GetBasicNft(lastDeployAddress, client)
	}
	name, err := basicNft.Name(nil)
	if err != nil {
		slog.Fatal("Get name failed:", err)
		return
	}
	slog.Info("Get contract name:", name)
	balance, err := basicNft.BalanceOf(nil, crypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		slog.Fatal("Get balance failed:", err)
		return
	}
	slog.Info("Get account balance:", balance)
	if balance.Uint64() == 0 {
		_, err = script.MintBasicNft(privateKey, client, basicNft, "ipfs://QmYJouWczWMFXeA4UwQVHrJM4oe9j7wdMcvrkFtEosiFVR?filename=jiangjiang.json")
		if err != nil {
			slog.Fatal("Mint Nft failed:", err)
			return
		}
		slog.Info("jiangjiang minted")
	}
	uri, err := basicNft.TokenURI(nil, big.NewInt(0))
	if err != nil {
		slog.Fatal("Get token uri failed:", err)
		return
	}
	slog.Info("uri:", uri)
}
