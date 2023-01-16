package db

var (
	sqlFindTokensByChain         = `select address,name,symbol,decimals,chain,genesis_block,height from tokens where chain = $1`
	sqlFindTokenByChainAndSymbol = `select address,name,symbol,decimals,chain,genesis_block,height from tokens where chain = $1 and symbol = $2 limit 1`
	sqlUpdateTokenHeight         = `update tokens set height = $1 where address = $2`
)
