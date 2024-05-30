package main

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gookit/slog"
	"github.com/joho/godotenv"
	"go-made/script"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Fatal("Error loading .env file")
		return
	}
	var moodNft *abi.MoodNft
	client, err := ethclient.Dial(os.Getenv("SEPOLIA_URL"))
	privateKey, err := crypto.HexToECDSA(os.Getenv("SEPOLIA_PRIVATE_KEY"))
	happySvg, err := os.ReadFile("../img/happy.svg")
	happySvgUri, err := script.SvgToUri(string(happySvg))
	sadSvg, err := os.ReadFile("../img/sad.svg")
	sadSvgUri, err := script.SvgToUri(string(sadSvg))
	moodNft, err = script.DeployMoodNft(privateKey, client, sadSvgUri, happySvgUri)
	if err != nil {
		slog.Fatal("Get name failed:", err)
		return
	}
	name, err := moodNft.Name(nil)
	if err != nil {
		slog.Fatal("Get name failed:", err)
		return
	}
	slog.Info("Get contract name:", name)
}
