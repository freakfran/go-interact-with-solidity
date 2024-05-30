package abi

import (
	"go-made/util/cmdUtils"
	"strings"
)

// SolcAbi
//
//	@Description: 执行 solc --abi命令
//	@param contractName
//	@return error
func SolcAbi(contractName string) error {
	folderName := GenFolderName(contractName)
	contractName = strings.ReplaceAll(contractName, ".sol", "")
	cmd := "solc" + " --abi contracts/src/" + contractName + ".sol -o contracts/build/abi/" + folderName + " --overwrite"
	return cmdUtils.Execute(cmd)
}

// SolcBin
//
//	@Description: 执行 solc --bin命令
//	@param contractName
//	@return error
func SolcBin(contractName string) error {
	folderName := GenFolderName(contractName)
	contractName = strings.ReplaceAll(contractName, ".sol", "")
	cmd := "solc" + " --bin contracts/src/" + contractName + ".sol -o contracts/build/bin/" + folderName + " --overwrite"
	return cmdUtils.Execute(cmd)
}

func AbiGen(contractName string) error {
	folderName := GenFolderName(contractName)
	contractName = strings.ReplaceAll(contractName, ".sol", "")
	_ = MkDir("abi/" + folderName)
	cmd := "abigen " +
		"--bin=contracts/build/bin/" + folderName + "/" + contractName + ".bin " +
		"--abi=contracts/build/abi/" + folderName + "/" + contractName + ".abi " +
		"--pkg=" + folderName + " " +
		"--out=abi/" + folderName + "/" + folderName + ".go "
	return cmdUtils.Execute(cmd)
}

// GenFolderName
//
//	@Description: 返回合约名称第一个字母的小写格式 SimpleStorage.sol -> simpleStorage
//	@param contractName
//	@return string
func GenFolderName(contractName string) string {
	folderName := strings.ToLower(contractName[0:1]) + contractName[1:]
	folderName = strings.ReplaceAll(folderName, ".sol", "")
	return folderName
}

func MkDir(dir string) error {
	return cmdUtils.Execute("mkdir " + dir)
}
