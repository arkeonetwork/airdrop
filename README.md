# airdrop
On chain data analysis to generate airdrop balances for arkeo's upcoming drop.

### Setup requirements
- docker 
- [tern](https://github.com/jackc/tern)
- go 1.19

### adding new smart contracts
- generate contracts from abis using `make generate-contracts`
- https://geth.ethereum.org/docs/getting-started/installing-geth#install-on-macos-via-homebrew
- https://goethereumbook.org/en/smart-contract-compile/

### cli ###
a [cobra](https://github.com/spf13/cobra) based cli is being built to faciliate launching the various indexing commands for various chains/tokens etc
- register new commands in [root.go](https://github.com/ArkeoNetwork/airdrop/blob/e1e7f19370852abf344baf74948c15057e361948/cli/root.go#L33)
- See the [ETH Indexer](https://github.com/ArkeoNetwork/airdrop/blob/main/cli/indexer.go) for example
### TODO:
- audit block heights!