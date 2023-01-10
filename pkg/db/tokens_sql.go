package db

var (
	sqlFindTokensByChain = `select address,name,symbol,decimals,chain,genesis_block,height from tokens where chain = $1`
	sqlFindAllChains     = `select distinct(chain) from tokens`
	sqlUpdateTokenHeight = `update tokens set height = $1 where address = $2`
)
