-include .env

abi:
	solc --abi contracts/$(TAR).sol -o contracts/build/abi --overwrite
	solc --bin contracts/$(TAR).sol -o contracts/build/bin --overwrite
	abigen --bin=contracts/build/bin/$(TAR).bin --abi=contracts/build/abi/$(TAR).abi --pkg=interactions --out=interactions/$(TAR).go --type=$(TAR)


