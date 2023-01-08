# merkle-drop


1. Copy sample.env to .env and update node urls and needed end snapshot.
2. generate contracts from abis `make generate-contracts`
 

# generating contract files from abis
https://geth.ethereum.org/docs/getting-started/installing-geth#install-on-macos-via-homebrew
https://goethereumbook.org/en/smart-contract-compile/


# OOM
GOMEMLIMIT may need to be set higher to run this depending on the lenght of the snapshot. 
```
export GOMEMLIMIT=16GiB
```


#### TODO:
- Hedgeys
- should we blacklist addresses that are inacessbile (FOXY, LP pool, etc? ) makes things more accurate, but also more complicated.