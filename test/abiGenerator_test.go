package test

import (
	"github.com/stretchr/testify/assert"
	"go-made/abi"
	"testing"
)

func TestSolcAbi(t *testing.T) {
	err := abi.SolcAbi("MoodNft.sol")
	assert.Nil(t, err)
	err = abi.SolcAbi("BasicNft.sol")
	assert.Nil(t, err)
}

func TestSolcBin(t *testing.T) {
	err := abi.SolcBin("BasicNft.sol")
	assert.Nil(t, err)
	err = abi.SolcBin("MoodNft.sol")
	assert.Nil(t, err)
}

func TestAbiGen(t *testing.T) {
	err := abi.AbiGen("BasicNft.sol")
	assert.Nil(t, err)
	err = abi.AbiGen("MoodNft.sol")
	assert.Nil(t, err)
}
