# airdrop

On chain data analysis to generate airdrop balances for arkeo's upcoming drop.


### adding new contracts
- generate contracts from abis using `make generate-contracts`
- https://geth.ethereum.org/docs/getting-started/installing-geth#install-on-macos-via-homebrew
- https://goethereumbook.org/en/smart-contract-compile/


# OOM
If you experience OOM errors, you can increase the memory limit by setting the environment variable GOMEMLIMIT 
```
export GOMEMLIMIT=16GiB
```

#### TODO:
- Hedgeys
- should we blacklist addresses that are inacessbile (FOXY, LP pool, etc? ) makes things more accurate, but also more complicated.
- fix docker env issues