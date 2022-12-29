# merkle-drop


1. Copy sample.env to .env and update node urls and needed end snapshot.
2. generate contracts from abis `generate-contracts`
 

# generating contract files from abis
https://geth.ethereum.org/docs/getting-started/installing-geth#install-on-macos-via-homebrew
https://goethereumbook.org/en/smart-contract-compile/

```
abigen --abi=./abis/ERC20.abi --pkg=contracts --out=./contracts/ERC20.go
```