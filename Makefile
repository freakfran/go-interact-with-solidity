-include .env


abi:
	solc --abi contracts/$(TAR).sol -o contracts/build/$(TAR)/abi --overwrite
	solc --bin contracts/$(TAR).sol -o contracts/build/$(TAR)/bin --overwrite
	abigen --bin=contracts/build/bin/$(TAR).bin --abi=contracts/build/abi/$(TAR).abi --pkg=interactions --out=interactions/$(TAR).go


