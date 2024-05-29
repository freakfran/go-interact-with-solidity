package test

import (
	"github.com/gookit/slog"
	"github.com/stretchr/testify/assert"
	"go-made/initializers"
	"go-made/script"
	"go-made/util"
	"math/big"
	"os"
	"testing"
)

func init() {
	initializers.LoadEnvVariables()
}

// moodNft sepolia:0x4e99C239EE5b7dDD078F01060027Fb5F67D78a03
func TestDeployMoodNft(t *testing.T) {
	var err error

	ethConfig, err := util.GetEthConfigFromEnv()
	assert.Nil(t, err)

	happySvg, err := os.ReadFile("../img/happy.svg")
	assert.Nil(t, err)
	happySvgUri, err := script.SvgToUri(string(happySvg))
	assert.Nil(t, err)
	sadSvg, err := os.ReadFile("../img/sad.svg")
	assert.Nil(t, err)
	sadSvgUri, err := script.SvgToUri(string(sadSvg))
	assert.Nil(t, err)

	_, err = script.DeployMoodNft(ethConfig.PrivateKey, ethConfig.Client, sadSvgUri, happySvgUri)
	assert.Nil(t, err)
}

func TestNameAndSymbolIsCorrect(t *testing.T) {
	ethConfig, err := util.GetEthConfigFromEnv()
	assert.Nil(t, err)
	nft, err := script.GetMoodNft("0x4e99C239EE5b7dDD078F01060027Fb5F67D78a03", ethConfig.Client)
	assert.Nil(t, err)
	name, err := nft.Name(nil)
	assert.Nil(t, err)
	assert.Equal(t, "Mood NFT", name)
	symbol, err := nft.Symbol(nil)
	assert.Nil(t, err)
	assert.Equal(t, "MN", symbol)
}

func TestMintNft(t *testing.T) {
	ethConfig, err := util.GetEthConfigFromEnv()
	assert.Nil(t, err)
	nft, err := script.GetMoodNft("0x4e99C239EE5b7dDD078F01060027Fb5F67D78a03", ethConfig.Client)
	assert.Nil(t, err)
	_, err = script.MintMoodNft(ethConfig.PrivateKey, ethConfig.Client, nft)
	assert.Nil(t, err)
}

func TestTokenUri(t *testing.T) {
	ethConfig, err := util.GetEthConfigFromEnv()
	assert.Nil(t, err)
	nft, err := script.GetMoodNft("0x4e99C239EE5b7dDD078F01060027Fb5F67D78a03", ethConfig.Client)
	assert.Nil(t, err)
	uri, err := nft.TokenURI(nil, big.NewInt(0))
	assert.Nil(t, err)
	slog.Info(uri)
}

func TestFlipMood(t *testing.T) {
	ethConfig, err := util.GetEthConfigFromEnv()
	assert.Nil(t, err)
	nft, err := script.GetMoodNft("0x4e99C239EE5b7dDD078F01060027Fb5F67D78a03", ethConfig.Client)
	assert.Nil(t, err)
	_, err = script.FlippMood(ethConfig.PrivateKey, ethConfig.Client, nft, big.NewInt(0))
	assert.Nil(t, err)
}

func TestSvgToUri(t *testing.T) {
	svg, err := os.ReadFile("../img/happy.svg")
	assert.Nil(t, err)
	uri, err := script.SvgToUri(string(svg))
	assert.Nil(t, err)
	slog.Info(uri)
	svg, err = os.ReadFile("../img/sad.svg")
	assert.Nil(t, err)
	uri, err = script.SvgToUri(string(svg))
	assert.Nil(t, err)
	slog.Info(uri)
}
